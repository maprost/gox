[![GoDoc](https://godoc.org/github.com/maprost/gox/gxcfg?status.svg)](https://godoc.org/github.com/maprost/gox/gxcfg)

# GoX (0.1 alpha)

GoX is a go tool, to build and run your go server application in a docker container.

## Actions
- init all dependencies in docker for your test environment (IntelliJ) (`gox init`)
- compile and test (init therefore all dependencies) your go server application in a docker container and create a docker image (`gox build`)
- run your docker image + dependencies (`gox run`)
- see all states of your docker container/images (`gox stat`)

## Supported Databases
- Postgres

### Planned Databases
- MySql
- Google Cloud Datastore
- MongoDB

## Dependencies
- docker
- godep

## Usage
1. put the *.gox file into your project root folder
    1. e.g.: ~/go/src/github.com/maprost/gox
1. in this folder will take place the most action
    1. godep save -> creating the vendor folder (in root folder)
    1. Compiling the code (inside docker)
    1. Test the system (inside docker)
    1. Create the docker image (in root folder)
    
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