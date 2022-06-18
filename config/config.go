package config

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type ConfigYaml struct {
	Telegram struct {
		Token      string `yaml:"token"`
		GetUpdates struct {
			Offset         int64    `yaml:"offset"`
			Limit          int64    `yaml:"limit"`
			Timeout        int      `yaml:"timeout"`
			AllowedUpdates []string `yaml:"allowed_updates"`
		} `yaml:"getUpdates"`
	} `yaml:"telegram"`
	Rpc struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"rpc"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
	Tinkoff struct {
		Endpoint     string `yaml:"endpoint"`
		DefaultToken string `yaml:"defaultToken"`
		UpdateTime   int32  `yaml:"updateTime"`
	} `yaml:"tinkoff"`
}

var Conf Config

func ParseConfig(file []byte) error {
	var cy ConfigYaml
	err := yaml.Unmarshal(file, &cy)
	if err != nil {
		return fmt.Errorf("config.go %v", err)
	}
	Conf.Telegram = Telegram(cy.Telegram)

	Conf.Database = Database(cy.Database)
	Conf.Rpc = Rpc(cy.Rpc)

	Conf.Tinkoff.Endpoint = cy.Tinkoff.Endpoint
	Conf.Tinkoff.DefaultToken = cy.Tinkoff.DefaultToken
	Conf.Tinkoff.UpdateTime = cy.Tinkoff.UpdateTime

	return nil
}
