[![GoDoc](https://godoc.org/github.com/maprost/gox/gxcfg?status.svg)](https://godoc.org/github.com/maprost/gox/gxcfg)
[![Go Report Card](https://goreportcard.com/badge/github.com/maprost/gox)](https://goreportcard.com/report/github.com/maprost/gox)

# GoX (0.3 alpha)

GoX is a go tool, to build and run your go server application in a docker container.

## Install 

`go get github.com/maprost/gox`

## Actions
### `gox init` / `gox`
- init all dependencies in docker container for your test environment 
    - it's mainly to run your tests in IntelliJ or at the terminal
- default behavior of `gox`. 

### `gox build`
1. run `godep [save|update] ./...`, if activated (`-godep`)
1. check style in a docker container, build failed if `-style` is activated and there are some check style warnings
    1. `go vet`
    1. `golint`
    1. `gocyclo`, check cyclomatic complexities of functions in Go source code (shouldn't be higher than 10)
1. compile your project in a docker container
1. test your project in a docker container (therefor init all dependencies)
1. build docker image
1. build a shell script for the server, if activated (`-shell`)
1. build a docker compose script for the server, if activated (`-compose`)

### `gox binrun`
- this mode exists mostly for local testing
1. build a binary of your project
1. init all dependencies, if not deactivated (`-fast`)
1. run the binary

### `gox tools`
- see all states of your docker container/images
- pull all needed docker images (`-pull`)
- clean your docker images (`-clean`)
- build a travis ci script (`-travis`)

## Supported Databases
- Postgres

### Planned Databases
- MySql
- Google Cloud Datastore
- MongoDB

## Dependencies
- docker (has to be installed on the system)
- godep (has to be installed `go get github.com/tools/godep`)
- golint (will be downloaded into a docker container and execute)
- gocyclo (will be downloaded into a docker container and execute)
    
## Config
```go
type config struct {
	Port   string   // mandatory, port of server
	Docker struct { // optional,
		Container string   // optional, default: name of project
		Image     string   // optional, default: golang:latest
		Volume    []string // optional, public resource folder like: html, css, images...
	}
	Databases []struct { // optional,
		Driver   string   // mandatory, 'postgres'
		Dbname   string   // mandatory, name of the used database
		User     string   // mandatory
		Password string   // optional, default: ''
		Host     string   // optional, default: localhost
		Port     string   // optional, default: std port of db
		Docker   struct { // optional, if not set -> database is not in a docker container
			Container string // optional, default: Driver+Port
			Image     string // optional, default: Driver:latest
			Discspace string // optional, for mode hdd mandatory
		}
	}
	Property map[string]string // optional, some input like client urls or keys
}
```