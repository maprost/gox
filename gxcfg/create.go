package gxcfg

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func CreateConfig(filename string, configSearch bool) (conf Config, err error) {
	var cfg config
	_, err = loadConfig(filename, &cfg, configSearch)
	if err != nil {
		return
	}

	conf.ConfigProfile = filename
	if index := strings.LastIndex(filename, "."); index > 0 {
		conf.ConfigProfile = filename[:index]
	}

	// path config
	conf.FullProjectPath, err = getFullProjectPath()
	if err != nil {
		return
	}
	conf.ProjectPath = trimSrc(conf.FullProjectPath)

	// server config
	conf.Port = cfg.Port

	nameIndex := strings.LastIndex(conf.ProjectPath, "/")
	conf.Name = conf.ProjectPath[nameIndex+1:]

	conf.Property = cfg.Property

	// server docker config
	conf.Docker.Image = cfg.Docker.Image
	if conf.Docker.Image == "" {
		conf.Docker.Image = golangImage
	}

	conf.Docker.Container = cfg.Docker.Container
	if conf.Docker.Container == "" {
		conf.Docker.Container = conf.Name
	}

	conf.Docker.Volumes = cfg.Docker.Volume
	conf.Docker.ProjectPath = "/go/" + conf.ProjectPath

	// build database list
	insideDockerContainer := checkFileInsideDockerContainer(configSearch)

	conf.Database = make([]Database, len(cfg.Databases))
	for i, db := range cfg.Databases {
		useLink := db.Docker != nil && insideDockerContainer
		entry := Database{}

		if dbPortFactory(db.Driver) == unknownDriver {
			err = errors.New("Unknown driver found, can't proceed, driver is mandatory.")
			return
		}
		entry.driver = db.Driver

		if db.Dbname == "" {
			err = errors.New("Empty db name found, can't proceed, db name is mandatory.")
			return
		}
		entry.dbname = db.Dbname

		if db.User == "" {
			err = errors.New("Empty user name found, can't proceed, user is mandatory.")
			return
		}
		entry.user = db.User
		entry.password = db.Password

		// set port first
		if db.Port == "" {
			db.Port = dbPortFactory(entry.driver)
		}

		// set docker config
		if db.Docker != nil {
			entry.Docker.Image = db.Docker.Image
			if entry.Docker.Image == "" {
				entry.Docker.Image = entry.driver + ":latest"
			}

			entry.Docker.Container = db.Docker.Container
			if entry.Docker.Container == "" {
				entry.Docker.Container = entry.driver + db.Port
			}

			entry.Docker.DiscSpace = db.Docker.Discspace
		}

		// binary is inside the docker container and tries to reach the database
		if useLink {
			entry.port = dbPortFactory(entry.driver)
			entry.host = entry.Docker.Container

		} else {
			entry.port = db.Port
			entry.host = db.Host
			if entry.host == "" {
				entry.host = "localhost"
			}
		}
		conf.Database[i] = entry
	}

	return
}

func getFullProjectPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", errors.New("Can't get folder information. " + err.Error())
	}
	return dir, nil
}

func trimSrc(path string) string {
	index := strings.Index(path, "/src/") // look for golang root
	return path[index+1:]
}
