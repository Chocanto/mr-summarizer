package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
)

var Printer *message.Printer

func initPrinter(selectedLanguage string) {
	matcher := language.NewMatcher(message.DefaultCatalog.Languages())
	tag, _, c := matcher.Match(language.MustParse(selectedLanguage))
	log.Printf("Based on the configured language {%v} we matched {%v} with a confidence of {%v}", Config.Language, tag, c)
	Printer = message.NewPrinter(tag)
}
