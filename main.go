package main

import (
	"log"
	"os"
)

func main() {
	configPath := "config.yml"
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	err := loadConfig(configPath)
	if err != nil {
		log.Fatalln(err.Error())
	}

	table := gatherMRs()
	sendToMattermost(table)
}

