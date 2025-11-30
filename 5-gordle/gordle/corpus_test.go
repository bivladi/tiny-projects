package gordle_test

import (
	"errors"
	"testing"
	"learngo/gordle/gordle"
)


func TestReadCorpus(t *testing.T) {
	tt := map[string]struct{
		fileName string
		length int
		err error
	} {
		"empty file": {
			fileName: "../corpus/empty.txt",
			length: 0,
			err: gordle.ErrorCorpusIsEmpty,
		},
		"english words": {
			fileName: "../corpus/english_words.txt",
			length: 34,
			err: nil,
		},
	}
	for name,tc := range tt {
		t.Run(name, func (t *testing.T) {
			result,err := gordle.ReadCorpus(tc.fileName)
			if len(result) != tc.length {
				t.Fatalf("expected length %d, got %d", tc.length, len(result))
			}
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected %v, got %v", tc.err, err)
			}
		})
	}
}