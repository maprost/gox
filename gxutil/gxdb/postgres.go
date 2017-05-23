package gxdb

import (
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/gxutil/gxdocker"
	"github.com/maprost/gox/gxutil/gxlog"
	"time"
)

type postgres struct {
	cfg gxcfg.Database
}

func newPostgres(cfg gxcfg.Database) *postgres {
	return &postgres{cfg: cfg}
}

func (p *postgres) Run(hdd bool) (err error) {
	err = p.Remove()
	if err != nil {
		return
	}

	docker := gxdocker.NewRunBuilder(p.cfg.Docker.Container, p.cfg.Docker.Image)
	docker.Detach()
	docker.Port(p.cfg.Port(), "5432")
	docker.EnvironmentVariable("POSTGRES_USER=" + p.cfg.Username())

	if hdd {
		docker.Value(p.cfg.DiscSpace, "/var/lib/postgresql/data/")
	}

	_, err = docker.Run()
	if err != nil {
		return
	}

	return p.createDatabase()
}

func (p *postgres) Remove() error {
	return gxdocker.StopAndRemove(p.cfg.Docker.Container)
}

func (p *postgres) Pull() error {
	return gxdocker.Pull(p.cfg.Docker.Image)
}

func (p *postgres) createDatabase() (err error) {
	err = p.checkAvailability()
	if err != nil {
		return
	}

	gxlog.Info("Postgres is running...")

	out, err := gxdocker.Execute(gxlog.LevelDebug, p.cfg.Docker.Container,
		"su postgres --command 'psql -l | grep "+p.cfg.Database()+"'")
	if err != nil {
		return
	}

	// contains no database?
	if out == "" {
		gxlog.Info("Create database '", p.cfg.Database(), "'...")

		_, err = gxdocker.Execute(gxlog.LevelDebug, p.cfg.Docker.Container,
			"su postgres --command 'createdb -O postgres "+p.cfg.Database()+"'")
	}

	return
}

func (p *postgres) checkAvailability() error {
	gxlog.Info("Checking Postgres availability...")

	for {
		out, err := gxdocker.Execute(gxlog.LevelDebug, p.cfg.Docker.Container, "ps aux | grep 'postgres' | grep 'docker-entrypoint.sh' | grep -v 'grep'")
		if err != nil {
			return err
		}

		if out == "" {
			break
		}

		gxlog.Info("waiting for postgres...")
		time.Sleep(3 * time.Second)
	}
	return nil
}
