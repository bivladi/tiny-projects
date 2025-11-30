package gordle

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const ErrorCorpusIsEmpty = corpusError("corpus is empty")

func ReadCorpus(path string) ([]string, error) {
	data,err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("couldn't read the file %s, err: %w", path, err)
	}
	if len(data) == 0 {
		return nil,ErrorCorpusIsEmpty
	}
	return strings.Fields(string(data)), nil
}

func pickWord(corpus []string) string {
	index := rand.Intn(len(corpus))
	return corpus[index]
}