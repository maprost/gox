package goxcfg

type Database struct {
	driver   string
	dbname   string
	user     string
	password string
	host     string
	port     string
	Docker   string
}

func (d Database) Driver() string {
	return d.driver
}

func (d Database) Database() string {
	return d.dbname
}

func (d Database) Host() string {
	return d.host
}

func (d Database) Port() string {
	return d.port
}

func (d Database) Username() string {
	return d.user
}

func (d Database) Password() string {
	return d.password
}

type Config struct {
	Name         string
	Port         string
	RelativeRoot string
	Database     []Database
	Docker       string
	ClientUrl    map[string]string
}

var singleton *Config = nil

func InitConfig(file string, databaseAccess DatabaseAccess) error {
	cfg, err := loadConfig(file, databaseAccess)
	if err != nil {
		return err
	}
	singleton = &cfg
	return nil
}

func GetConfig() *Config {
	return singleton
}
