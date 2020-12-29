package config

import "github.com/spf13/viper"

type Config struct {
	Port          string
	DbCluster1    string
	DbUsername    string
	DbPassword    string
	DbKeyspace    string
	EmailHost     string
	EmailPort     int
	EmailUsername string
	EmailPassword string
}

var Cfg Config

func LoadConfig() error {
	viper.AutomaticEnv()

	Cfg.Port = viper.GetString("PORT")
	Cfg.DbCluster1 = viper.GetString("DB_CLUSTER1")
	Cfg.DbUsername = viper.GetString("DB_USERNAME")
	Cfg.DbPassword = viper.GetString("DB_PASSWORD")
	Cfg.DbKeyspace = viper.GetString("DB_KEYSPACE")
	Cfg.EmailHost = viper.GetString("EMAIL_HOST")
	Cfg.EmailPort = viper.GetInt("EMAIL_PORT")
	Cfg.EmailUsername = viper.GetString("EMAIL_LOGIN")
	Cfg.EmailPassword = viper.GetString("EMAIL_PASSWORD")

	return nil
}
