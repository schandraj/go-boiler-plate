package config

import "github.com/sahalazain/go-common/config"

var Map = map[string]interface{}{
	"APP_PORT":                   "3030",
	"ELASTIC_APM_SERVER_URL":     "",
	"ELASTIC_APM_SERVICE_NAME":   "",
	"APP_PREFIX":                 "",
	"APP_SECRET":                 "",
	"MONGO_URL":                  "",
	"REDIS_HOST":                 "",
	"REDIS_PASSWORD":             "",
	"REDIS_DB":                   0,
	"KAFKA_BROKERS":              "",
	"MYSQL_USER":                 "",
	"MYSQL_PASS":                 "",
	"MYSQL_HOST":                 "",
	"MYSQL_PORT":                 "",
	"MYSQL_DBNAME":               "",
	"POSTGRES_HOST":              "",
	"POSTGRES_USER":              "",
	"POSTGRES_PASS":              "",
	"POSTGRES_DBNAME":            "",
	"POSTGRES_PORT":              "",
	"CB_TIMEOUT":                 5000,
	"CB_MAX_CONCURRENT":          1000,
	"CB_ERROR_PERCENT_THRESHOLD": 30,
}

var Cfg config.Getter
var ConfigUrl string

func Load() error {
	cfgClient, err := config.Load(Map, ConfigUrl)
	if err != nil {
		return err
	}

	Cfg = cfgClient

	return nil
}
