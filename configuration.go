package main

import "github.com/jinzhu/configor"

var Config = struct {
	Mattermost struct {
		Webhook  string `required:"true"`
		Channel  string `required:"true"`
		Username string `default:"Merge Requests"`
		IconUrl  string
	}

	Gitlab struct {
		Url   string `required:"true"`
		Token string `required:"true"`
		Group string `required:"true"`
	}

	Threshold int    `default:"7"`
	Title     string `default:"MERGE REQUESTS WAITING FOR APPROVAL"`
}{}

func loadConfig(filepath string) error {
	return configor.Load(&Config, filepath)
}
