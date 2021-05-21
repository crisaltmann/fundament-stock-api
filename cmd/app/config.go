package app

import (
	"github.com/crisaltmann/fundament-stock-api/config"
	"go.uber.org/fx"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Config = fx.Options(
	configfactories,
)

var configfactories = fx.Provide(
	loadConfig,
)

func loadConfig() *config.Config {
	conf := &config.Config{}
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
