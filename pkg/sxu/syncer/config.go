package syncer

import "github.com/yindaheng98/dion/config"

type Config config.Common

func (c *Config) Load(file string) error {
	return config.Load(file, c)
}
