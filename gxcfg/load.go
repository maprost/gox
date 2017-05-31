package gxcfg

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type config struct {
	Port   string
	Docker struct {
		Container string
		Image     string
		Volume    []string
	}
	Databases []struct {
		Driver   string
		Dbname   string
		User     string
		Password string
		Host     map[string]string
		Port     map[string]string
		Docker   struct {
			Container string
			Image     string
			Discspace string
		}
	}
	Clients map[string]string
}

func loadConfig(file string, databaseAccess DatabaseAccess, configSearch bool) (conf Config, err error) {
	deep := 0
	if configSearch {
		deep = 8
	}

	var cfg config
	// TODO: nur gox-init hat die suche, alle anderen müssen die config angeben
	deep, err = searchConfig(file, deep, &cfg)
	if err != nil {
		return
	}

	// build config
	conf.Port = cfg.Port
	conf.Clients = cfg.Clients
	conf.Docker = CfgDocker{
		Image:     cfg.Docker.Image,
		Container: cfg.Docker.Container,
		Volumes:   cfg.Docker.Volume,
	}

	conf.FullProjectPath, err = getFullProjectPath(deep)
	if err != nil {
		return
	}

	conf.ProjectPath, err = getProjectPath(deep)
	if err != nil {
		return
	}
	conf.Docker.ProjectPath = "/go/" + conf.ProjectPath

	nameIndex := strings.LastIndex(conf.ProjectPath, "/")
	conf.Name = conf.ProjectPath[nameIndex+1:]

	// build database list
	conf.Database = make([]Database, len(cfg.Databases))
	for i, db := range cfg.Databases {
		var host string
		host, err = getValue(db.Host, databaseAccess)
		if err != nil {
			return
		}

		var port string
		port, err = getValue(db.Port, databaseAccess)
		if err != nil {
			return
		}

		conf.Database[i] = Database{
			driver:   db.Driver,
			dbname:   db.Dbname,
			user:     db.User,
			password: db.Password,
			host:     host,
			port:     port,
			Docker: DBDocker{
				Image:     db.Docker.Image,
				Container: db.Docker.Container,
				DiscSpace: db.Docker.Discspace,
			},
		}
	}

	return
}

func searchConfig(filename string, levelsDown int, properties interface{}) (int, error) {
	var file []byte
	var err error

	relativeRoot := ""
	index := 0
	for index <= levelsDown {
		file, err = ioutil.ReadFile(relativeRoot + filename)
		if err != nil {
			index++
			relativeRoot += "../"
		} else {
			break
		}
	}
	// nothing found?
	if err != nil {
		return index, err
	}

	// something found? -> convert
	err = json.Unmarshal(file, &properties)
	return index, err
}

func getValue(mapToCheck map[string]string, access DatabaseAccess) (string, error) {
	value, ok := mapToCheck[access.String()]
	if !ok {
		return value, errors.New("Can't find database access " + access.String())
	}
	return value, nil
}

func getFullProjectPath(deep int) (string, error) {
	path, err := getPath()
	if err != nil {
		return "", err
	}

	return trimLast(path, deep), nil
}

func getProjectPath(deep int) (string, error) {
	path, err := getFullProjectPath(deep)
	if err != nil {
		return "", err
	}

	return trimSrc(path), nil
}

func getPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", errors.New("Can't get folder information. " + err.Error())
	}
	return dir, nil
}

func trimLast(path string, deep int) string {
	if deep == 0 {
		return path
	}

	index := strings.LastIndexFunc(path, func(c rune) bool {
		if c == '/' {
			deep--
		}
		return deep == 0
	})
	return path[:index]
}

func trimSrc(path string) string {
	index := strings.Index(path, "/src/") // look for golang root
	return path[index+1:]
}
