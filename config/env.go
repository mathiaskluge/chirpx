package config

type Config struct {
	Port   string
	DBPath string
}

var Env = initConfig()

func initConfig() Config {

	return Config{
		Port:   ":8080",
		DBPath: "db/db.json",
	}
}
