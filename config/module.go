package config

import (
	"go.uber.org/fx"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Module = fx.Options(
		factories,
	)

var factories = fx.Provide(
	LoadConfig,
)

func LoadConfig() *Config {
	conf := &Config{}
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Println("erro ao ler arquivo", err)
		panic("Erro ao ler arquivo configuração")
	}

	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		panic("Erro no parser do arquivo de configuração")
	}

	return conf
}
