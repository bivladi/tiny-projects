package main

import (
	"fmt"
	"learngo/gordle/gordle"
	"os"
)

func main() {
	corpus,err := gordle.ReadCorpus("corpus/english_words.txt")
	if err != nil {
		_,_ = fmt.Fprintf(os.Stderr, "unable to read a corpus file, err: %s", err)
		return
	}
	g,err := gordle.New(os.Stdin, corpus, 6)
	if err != nil {
		_,_ = fmt.Fprintf(os.Stderr, "unabl to create the game, err: %s", err)
		return
	}
	g.Play()
}
