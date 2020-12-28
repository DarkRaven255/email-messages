package config

import "github.com/spf13/viper"

type Config struct {
	Port       string
	DbUsername string
	DbPassword string
}

var Cfg Config

func LoadConfig() error {
	viper.AutomaticEnv()
	Cfg.Port = viper.GetString("PORT")
	Cfg.DbUsername = viper.GetString("DBUSERNAME")
	Cfg.DbPassword = viper.GetString("DBPASSWORD")

	return nil
}
