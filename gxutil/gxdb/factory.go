package gxdb

import "github.com/maprost/gox/gxcfg"

type Database interface {
	Run(hdd bool) error
	Remove() error
	Pull() error
}

func New(cfg gxcfg.Database) Database {
	if cfg.Driver() == "postgres" {
		return newPostgres(cfg)
	}

	return nil
}
