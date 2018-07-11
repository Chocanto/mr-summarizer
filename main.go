package main

import (
	"log"
	"os"
)

//go:generate gotext -srclang=en update -out=catalog_gen.go -lang=en,fr

func main() {
	configPath := "config.yml"
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	err := loadConfig(configPath)
	if err != nil {
		log.Fatalln(err.Error())
	}

	initPrinter(Config.Language)
	table := gatherMRs()
	sendToMattermost(table)
}

