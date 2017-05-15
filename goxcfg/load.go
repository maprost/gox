package goxcfg

import (
	"encoding/json"
	"io/ioutil"
)

type database struct {
	Driver   string
	Dbname   string
	User     string
	Password string
	Host     map[string]string
	Port     map[string]string
	Docker   string
}

type config struct {
	Name      string
	Port      string
	Docker    string
	Databases []database
	Clients   map[string]string
}

func loadConfig(file string, databaseAccess DatabaseAccess) (Config, error) {
	var conf Config
	var err error

	var cfg config
	conf.RelativeRoot, err = searchConfig(file, 8, &cfg)
	if err != nil {
		return conf, err
	}

	// build config
	conf.Name = cfg.Name
	conf.Docker = cfg.Docker
	conf.Port = cfg.Port
	conf.ClientUrl = cfg.Clients

	// build database list
	conf.Database = make([]Database, len(cfg.Databases))
	for i, db := range cfg.Databases {
		conf.Database[i] = Database{
			driver:   db.Driver,
			dbname:   db.Dbname,
			user:     db.User,
			password: db.Password,
			host:     db.Host[databaseAccess.String()],
			port:     db.Port[databaseAccess.String()],
			Docker:   db.Docker,
		}
	}

	return conf, nil
}

func searchConfig(filename string, levelsDown int, properties interface{}) (string, error) {
	var file []byte
	var err error

	relativeRoot := ""
	for levelsDown >= 0 {
		file, err = ioutil.ReadFile(relativeRoot + filename)
		if err != nil {
			levelsDown--
			relativeRoot += "../"
		} else {
			break
		}
	}
	// nothing found?
	if err != nil {
		return relativeRoot, err
	}

	// something found? -> convert
	err = json.Unmarshal(file, &properties)
	return relativeRoot, err
}
