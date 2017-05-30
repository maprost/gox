package db

import (
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/docker"
	"github.com/maprost/gox/internal/log"
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

	docker := docker.NewRunBuilder(p.cfg.Docker.Container, p.cfg.Docker.Image)
	docker.Detach()
	docker.Port(p.cfg.Port(), "5432")
	docker.EnvironmentVariable("POSTGRES_USER=" + p.cfg.Username())

	if hdd {
		docker.Value(p.cfg.Docker.DiscSpace, "/var/lib/postgresql/data/")
	}

	_, err = docker.Run()
	if err != nil {
		return
	}

	return p.createDatabase()
}

func (p *postgres) Remove() error {
	return docker.StopAndRemove(p.cfg.Docker.Container)
}

func (p *postgres) PullDockerImage() error {
	return docker.Pull(p.cfg.Docker.Image)
}

func (p *postgres) createDatabase() (err error) {
	err = p.checkAvailability()
	if err != nil {
		return
	}

	log.Info("Postgres is running...")

	out, err := docker.Execute(log.LevelDebug, p.cfg.Docker.Container,
		"su postgres --shell 'psql -l | grep "+p.cfg.Database()+"'")
	if err != nil {
		return
	}

	// contains no database?
	if out == "" {
		log.Info("Create database '", p.cfg.Database(), "'...")

		_, err = docker.Execute(log.LevelDebug, p.cfg.Docker.Container,
			"su postgres --shell 'createdb -O postgres "+p.cfg.Database()+"'")
	}

	return
}

func (p *postgres) checkAvailability() error {
	log.Info("Checking Postgres availability...")

	for {
		out, err := docker.Execute(log.LevelDebug, p.cfg.Docker.Container, "ps aux | grep 'postgres' | grep 'docker-entrypoint.sh' | grep -v 'grep'")
		if err != nil {
			return err
		}

		if out == "" {
			break
		}

		log.Info("Waiting for Postgres...")
		time.Sleep(3 * time.Second)
	}
	return nil
}
