package syncer

import (
	log "github.com/pion/ion-log"
	pbion "github.com/pion/ion/proto/ion"
	"github.com/yindaheng98/dion/pkg/isglb"
	"github.com/yindaheng98/dion/pkg/islb"
	pb "github.com/yindaheng98/dion/proto"
	"github.com/yindaheng98/dion/util"
	"google.golang.org/protobuf/proto"
)

// ISGLBSyncer is a Client to sync SFUStatus
type ISGLBSyncer struct {
	client  *isglb.Client
	node    *islb.Node
	descSFU *pbion.Node

	router   trackRouter
	reporter *qualityReporter
	session  SessionTracker

	clientSet       *util.DisorderSet[util.ClientNeededSessionItem]
	forwardTrackSet *util.DisorderSet[util.ForwardTrackItem]
	proceedTrackSet *util.DisorderSet[util.ProceedTrackItem]

	// Just recv and send latest status
	statusRecvCh   chan *pb.SFUStatus
	statusSendCh   chan bool
	sessionEventCh chan *SessionEvent
}

func NewSFUStatusSyncer(node *islb.Node, peerID string, descSFU *pbion.Node, toolbox ToolBox) *ISGLBSyncer {
	isglbClient := isglb.NewClient(node, peerID, map[string]interface{}{})
	if isglbClient == nil {
		return nil
	}
	forwarder, processor, session := toolbox.TrackForwarder, toolbox.TrackProcessor, toolbox.SessionTracker
	if forwarder == nil {
		forwarder = StupidTrackForwarder{}
	}
	if processor == nil {
		processor = StupidTrackProcesser{}
	}
	if session == nil {
		session = StupidSessionTracker{}
	}
	tr, cr := toolbox.TransmissionReporter, toolbox.ComputationReporter
	if tr == nil {
		tr = &StupidTransmissionReporter{}
	}
	if cr == nil {
		cr = &StupidComputationReporter{}
	}
	s := &ISGLBSyncer{
		client:  isglbClient,
		node:    node,
		descSFU: descSFU,

		router: trackRouter{
			TrackForwarder: forwarder,
			TrackProcessor: processor,
		},
		reporter: newQualityReporter(tr, cr),
		session:  session,

		clientSet:       util.NewDisorderSet[util.ClientNeededSessionItem](),
		forwardTrackSet: util.NewDisorderSet[util.ForwardTrackItem](),
		proceedTrackSet: util.NewDisorderSet[util.ProceedTrackItem](),

		statusRecvCh:   make(chan *pb.SFUStatus, 1),
		statusSendCh:   make(chan bool, 1),
		sessionEventCh: make(chan *SessionEvent, 1024),
	}
	s.statusSendCh <- true
	isglbClient.OnSFUStatusRecv = func(st *pb.SFUStatus) {
		select {
		case _, ok := <-s.statusRecvCh:
			if !ok {
				return
			}
		default:
		}
		select {
		case s.statusRecvCh <- st:
		default:
		}
	}
	return s
}

func (s *ISGLBSyncer) NotifySFUStatus() {
	// Only send latest status
	select {
	case s.statusSendCh <- true:
	default:
	}
}

// ↓↓↓↓↓ should access Index, so keep single thread ↓↓↓↓↓

// syncStatus sync the current SFUStatus with the expected SFUStatus
// MUST be single threaded
func (s *ISGLBSyncer) syncStatus(expectedStatus *pb.SFUStatus) {
	if expectedStatus.SFU.String() != s.descSFU.String() { // Check if the SFU status is mine
		// If not
		log.Warnf("Received SFU status is not mine, drop it: %s", expectedStatus.SFU)
		s.NotifySFUStatus() // The server must re-consider the status for our SFU
		return              // And we should wait for the right SFU status to come
	}

	// Check if the client needed Client is same
	sessionIndexDataList := util.ClientNeededSessions(expectedStatus.Clients).ToDisorderSetItemList()
	if !s.clientSet.IsSame(sessionIndexDataList) { // Check if the ClientNeededSessions is same
		// If not
		log.Warnf("Received SFU status have different Client list, drop it: %s", expectedStatus.Clients)
		s.NotifySFUStatus() // The server must re-consider the status for our SFU
		return              // And we should wait for the right SFU status to come
	}

	// Perform track forward change
	forwardIndexDataList := util.ForwardTracks(expectedStatus.ForwardTracks).ToDisorderSetItemList()
	forwardAdd, forwardDel, forwardReplace := s.forwardTrackSet.Update(forwardIndexDataList)
	for _, track := range forwardDel {
		s.router.StopForwardTrack(track.Track)
	}
	for _, track := range forwardReplace {
		s.router.ReplaceForwardTrack(
			track.Old.Track,
			track.New.Track,
		)
	}
	for _, track := range forwardAdd {
		s.router.StartForwardTrack(track.Track)
	}

	//Perform track proceed change
	proceedIndexDataList := util.ProceedTracks(expectedStatus.ProceedTracks).ToDisorderSetItemList()
	proceedAdd, proceedDel, proceedReplace := s.proceedTrackSet.Update(proceedIndexDataList)
	for _, track := range proceedDel {
		s.router.StopProceedTrack(track.Track)
	}
	for _, track := range proceedReplace {
		s.router.ReplaceProceedTrack(
			track.Old.Track,
			track.New.Track,
		)
	}
	for _, track := range proceedAdd {
		s.router.StartProceedTrack(track.Track)
	}
}

// handleSessionEvent handle the SessionEvent
// MUST be single threaded
func (s *ISGLBSyncer) handleSessionEvent(event *SessionEvent) {
	// Just add or remove it, and sand latest status
	switch event.State {
	case SessionEvent_ADD:
		s.clientSet.Add(util.ClientNeededSessionItem{Client: event.Session})
		s.NotifySFUStatus()
	case SessionEvent_REMOVE:
		s.clientSet.Del(util.ClientNeededSessionItem{Client: event.Session})
		s.NotifySFUStatus()
	}
}

// main is the "main function" goroutine of the NewSFUStatusSyncer
// All the methods about Index should be here, to ensure the assess is single-threaded
func (s *ISGLBSyncer) main() {
	for {
		select {
		case event, ok := <-s.sessionEventCh: // handle an event
			if !ok {
				return
			}
			s.handleSessionEvent(event) // should access Index, so keep single thread
		case st, ok := <-s.statusRecvCh: // handle a received SFU status
			if !ok {
				return
			}
			s.syncStatus(st) // should access Index, so keep single thread
		case _, ok := <-s.statusSendCh: // handle SFU status send event
			if !ok {
				return
			}
			st := &pb.SFUStatus{
				SFU:           proto.Clone(s.descSFU).(*pbion.Node),
				ForwardTracks: util.ForwardTrackItemList(s.forwardTrackSet.Sort()).ToForwardTracks(),
				ProceedTracks: util.ProceedTrackItemList(s.proceedTrackSet.Sort()).ToProceedTracks(),
				Clients:       util.ClientSessionItemList(s.clientSet.Sort()).ToClientSessions(),
			} // should access Index, so keep single thread
			go s.client.SendSFUStatus(st)
		}
	}
}

// ↑↑↑↑↑ should access Index, so keep single thread ↑↑↑↑↑

func (s *ISGLBSyncer) sessionFetcher() {
	defer func() {
		if err := recover(); err != nil {
			log.Debugf("error on close: %+v", err)
		}
	}()
	for {
		event := s.session.FetchSessionEvent()
		if event == nil {
			return
		}
		s.sessionEventCh <- event.Clone()
	}
}

func (s *ISGLBSyncer) reportFetcher() {
	for {
		report := s.reporter.FetchReport()
		if report == nil {
			return
		}
		go s.client.SendQualityReport(report)
	}
}

func (s *ISGLBSyncer) Start() {
	go s.main()
	go s.reportFetcher()
	go s.sessionFetcher()
	s.client.Connect()
}

func (s *ISGLBSyncer) Stop() {
	s.client.Close()
	close(s.statusRecvCh)
	close(s.statusSendCh)
	close(s.sessionEventCh)
}
