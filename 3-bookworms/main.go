package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var fileName string
	flag.StringVar(
		&fileName,
		"fileName",
		"testdata/bookworms.json",
		"path to file with bookworms data",
	)
	flag.Parse()
	bookworms, err := loadBookworms(fileName)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}
	commonBooks := findCommonBooks(bookworms)
	fmt.Println("Here are the books in common:")
	displayBooks(commonBooks)
}
