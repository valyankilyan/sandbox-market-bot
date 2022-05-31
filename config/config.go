package config

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// type configYaml struct {
// 	Telegram struct {
// 		Token string `yaml:"token"`
// 	} `yaml:"telegram"`

// 	Users []struct {
// 		Name         string `yaml:"name"`
// 		TgUserName   string `yaml:"TgUserName"`
// 		TgID         int    `yaml:"TgId"`
// 		TinkoffToken string `yaml:"TinkoffToken"`
// 	} `yaml:"users"`
// }

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
	Users []struct {
		Name         string `yaml:"name"`
		TgUserName   string `yaml:"TgUserName"`
		TgId         int64  `yaml:"TgId"`
		TinkoffToken string `yaml:"TinkoffToken"`
	} `yaml:"users"`
}

var Conf Config

func ParseConfig(file []byte) error {
	var cy ConfigYaml
	err := yaml.Unmarshal(file, &cy)
	if err != nil {
		return fmt.Errorf("config.go %v", err)
	}

	Conf.Telegram = Telegram(cy.Telegram)
	Conf.Users = make(map[int64]User)

	for _, user := range cy.Users {
		Conf.Users[user.TgId] = User{
			user.Name,
			user.TgUserName,
			user.TgId,
			user.TinkoffToken,
		}
	}

	return nil
}
