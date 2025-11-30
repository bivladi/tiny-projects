package gordle

import "testing"

func TestPickWord(t *testing.T) {
	corpus := []string{"APPLE", "COMMA", "VALUE", "DONAT"}
	word := pickWord(corpus)
	if !inCorpus(corpus, word) {
		t.Fatalf("expected %v is in corpus %v", word, corpus)
	}
}

func inCorpus(corpus [] string, word string) bool {
	for _,value := range corpus {
		if value == word {
			return true
		}
	}
	return false
}