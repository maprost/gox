package gxcfg

const (
	golangImage               = "golang:latest"
	FileInsideDockerContainer = "FileInsideDockerContainer.gox"
	PostgresDefaultPost       = "5432"
	postgresDriver            = "postgres"
	unknownDriver             = "UnkownDriver"
)

func dbPortFactory(driver string) string {
	if driver == postgresDriver {
		return PostgresDefaultPost
	}

	return unknownDriver
}
