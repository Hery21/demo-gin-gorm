package config

import "os"

//type AppConfig struct {
//	AppName            string
//	JWTSecret          []byte
//	JWTExpireInMinutes int64
//}
//
//// todo: add database config
//var Config = AppConfig{
//	AppName:            "test",
//	JWTSecret:          []byte("very-secret"),
//	JWTExpireInMinutes: 15,
//}

type dbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

type AppConfig struct {
	ENV                string
	AppName            string
	JWTSecret          []byte
	JWTExpireInMinutes int64
	DBConfig           dbConfig
}

func getENV(key, defaultVal string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}

var Config = AppConfig{
	ENV:                getENV("ENV", "testing"),
	AppName:            "sea-labs-library",
	JWTSecret:          []byte("very-secret"),
	JWTExpireInMinutes: 15,
	DBConfig: dbConfig{
		Host:     getENV("DB_HOST", ""),
		User:     getENV("DB_USER", ""),
		Password: getENV("DB_PASSWORD", ""),
		DBName:   getENV("DB_NAME", ""),
		Port:     getENV("DB_PORT", ""),
	},
}
