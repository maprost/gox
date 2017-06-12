package gxcfg

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	Port   string   // mandatory, port of server
	Docker struct { // optional,
		Container string   // optional, default: name of project
		Image     string   // optional, default: golang:latest
		Volume    []string // optional, public resource folder like: html, css, images...
	}
	Databases []struct { // optional,
		Driver   string    // mandatory, 'postgres'
		Dbname   string    // mandatory, name of the used database
		User     string    // mandatory
		Password string    // mandatory
		Host     string    // mandatory
		Port     string    // optional, default: std port of db
		Docker   *struct { // optional, if not set -> database is not in a docker container
			Container string // optional, default: Driver+Port
			Image     string // optional, default: Driver:latest
			Discspace string // optional, for mode hdd mandatory
		}
	}
	Property map[string]string // optional
}

func loadConfig(filename string, properties interface{}, configSearch bool) (int, error) {
	file, index, err := searchFile(filename, configSearch)

	// nothing found?
	if err != nil {
		return index, err
	}

	// something found? -> convert
	err = json.Unmarshal(file, &properties)
	return index, err
}

func checkFileInsideDockerContainer(configSearch bool) bool {
	_, _, err := searchFile(FileInsideDockerContainer, configSearch)
	return err == nil
}

func searchFile(filename string, configSearch bool) ([]byte, int, error) {
	levelsDown := 0
	if configSearch {
		levelsDown = 8
	}

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

	return file, index, err
}
