[![GoDoc](https://godoc.org/github.com/maprost/gox/gxcfg?status.svg)](https://godoc.org/github.com/maprost/gox/gxcfg)

# GoX (0.3 alpha)

GoX is a go tool, to build and run your go server application in a docker container.

## Install 

`go get -u github.com/maprost/gox`

## Actions
### `gox init` / `gox`
- init all dependencies in docker container for your test environment 
    - it's mainly to run your tests in IntelliJ or at the terminal
- default behavior of `gox`. 

### `gox build`
1. run `godep [save|update] ./...` in a docker container, if activated (`-godep`)
1. check style with `go vet` and `golint` in a docker container
1. compile your project in a docker container
1. test your project in a docker container (therefor init all dependencies)
1. build docker image
1. build a run script, if activated (`-script`)

### `gox binrun`
1. build a binary of your project
1. init all dependencies, if not deactivated (`-fast`)
1. run the binary

### `gox stat`
- see all states of your docker container/images
- pull (`-pull`)
- clean (`-clean`)

## Supported Databases
- Postgres

### Planned Databases
- MySql
- Google Cloud Datastore
- MongoDB

## Dependencies
- docker (has to be installed on the system)
- godep (will be downloaded into a docker container and execute)
- golint (will be downloaded into a docker container and execute)
    
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
	Clients map[string]string // optional
}
```