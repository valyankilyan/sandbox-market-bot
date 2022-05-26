package config

import (
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

type configYaml struct {
	TgToken TgToken `yaml:"tgtoken"`
	Users   []struct {
		Name         string `yaml:"name"`
		TgUserName   string `yaml:"TgUserName"`
		TgId         int64  `yaml:"TgId"`
		TinkoffToken string `yaml:"TinkoffToken"`
	} `yaml:"users"`
}

func ParseConfig(file []byte) (*Config, error) {
	cy := configYaml{}
	err := yaml.Unmarshal(file, &cy)
	if err != nil {
		return nil, err
	}

	c := Config{}
	c.TgToken = cy.TgToken
	c.Users = make(map[int64]User)

	for _, user := range cy.Users {
		c.Users[user.TgId] = User{
			user.Name,
			user.TgUserName,
			user.TgId,
			user.TinkoffToken,
		}
	}

	return &c, nil
}
