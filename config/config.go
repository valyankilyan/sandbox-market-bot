package config

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type configYaml struct {
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

var Telegram telegram
var Database database
var Rpc rpc
var Tinkoff tinkoff

func ParseConfig(file []byte) error {
	var cy configYaml
	err := yaml.Unmarshal(file, &cy)
	if err != nil {
		return fmt.Errorf("config.go %v", err)
	}
	Telegram = telegram(cy.Telegram)

	Database = database(cy.Database)
	Rpc = rpc(cy.Rpc)

	Tinkoff.Endpoint = cy.Tinkoff.Endpoint
	Tinkoff.DefaultToken = cy.Tinkoff.DefaultToken
	Tinkoff.UpdateTime = cy.Tinkoff.UpdateTime

	return nil
}
