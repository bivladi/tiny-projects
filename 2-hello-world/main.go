package main

import (
	"flag"
	"fmt"
)

type language string

var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",
	"en": "Hello world",
	"fr": "Bonjour le monde",
	"he": "שלום עולם",
	"ur": "ہیلو دنیا",
	"vi": "Xin chào Thế Giới",
	"ru": "Привет, мир",
}

func main() {
	var lang string
	flag.StringVar(
		&lang,
		"lang",
		"en",
		"The required language, e.g. ur...",
	)
	flag.Parse()
	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// greet returns a greeting message to the world
func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
