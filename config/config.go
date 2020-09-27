package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func GetConfig() Configuration {
	conf := Configuration{}

	gonfig.GetConf("config/config.json", &conf)

	return conf
}
