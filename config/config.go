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
	} `yaml:"rpc"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
	Tinkoff struct {
		Endpoint string `yaml:"endpoint"`
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
	fmt.Printf(`Telegram Config:
	offset = %v
	limit = %v
	timeout = %v
	allowed_updates = %v
`,
		Conf.Telegram.GetUpdates.Offset,
		Conf.Telegram.GetUpdates.Limit,
		Conf.Telegram.GetUpdates.Timeout,
		Conf.Telegram.GetUpdates.AllowedUpdates,
	)

	Conf.Database = Database(cy.Database)
	Conf.Rpc = Rpc(cy.Rpc)
	Conf.Tinkoff.Endpoint = cy.Tinkoff.Endpoint

	return nil
}
