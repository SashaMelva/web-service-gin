package config

import (
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	HttpServer *ConfigHttpServer
	Logger     *ConfigLogger
	DataBase   *ConfigDB
	ConfigRMQ  *ConfigRMQ
}

type ConfigRMQ struct {
	User     string
	Password string
	Host     string
	Port     string
}

type ConfigHttpServer struct {
	Host    string
	Port    string
	Timeout time.Duration
}

type ConfigLogger struct {
	Level       zapcore.Level
	LogEncoding string `required:"true"`
}

type ConfigDB struct {
	NameDB   string
	Host     string
	Port     string
	User     string
	Password string
}

func New(pahToFile string) Config {
	viper.AddConfigPath(pahToFile)
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	configLog := ConfigLogger{}
	configDB := ConfigDB{
		NameDB:   viper.Get("nameDB").(string),
		Host:     viper.Get("hostDB").(string),
		Port:     viper.Get("portDB").(string),
		User:     viper.Get("usesrDB").(string),
		Password: viper.Get("passwordDB").(string),
	}

	configHttpServer := ConfigHttpServer{
		Host: viper.Get("hostHttpServer").(string),
		Port: viper.Get("portHttpServe").(string),
	}

	level, err := zapcore.ParseLevel(viper.Get("level").(string))
	if err != nil {
		configLog = ConfigLogger{zapcore.DebugLevel, viper.Get("logEncoding").(string)}
	} else {
		configLog = ConfigLogger{level, viper.Get("logEncoding").(string)}
	}

	return Config{
		HttpServer: &configHttpServer,
		Logger:     &configLog,
		DataBase:   &configDB,
	}
}
