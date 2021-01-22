package config

import "github.com/spf13/viper"

type RootConfig struct {
	Server   Server
	Database Database
	Log      Log
}

type Server struct {
	Port  string
	DbLog bool
}

type Database struct {
	Endpoint string
	Port     int
	Username string
	Password string
	DbName   string
}

type Log struct {
	Level  string
	Format string
}

var Config RootConfig

func init() {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}
