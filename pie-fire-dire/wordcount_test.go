package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWordCount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bacon ipsum dolor amet biltong. Biltong pork belly porchetta."))
	}))
	defer server.Close()

	wordCount, err := WordCount(server.URL)
	if err != nil {
		t.Fatalf("e %v", err)
	}

	expected := map[string]int32{
		"bacon":     1,
		"ipsum":     1,
		"dolor":     1,
		"amet":      1,
		"biltong":   2,
		"pork":      1,
		"belly":     1,
		"porchetta": 1,
	}

	// checkkkkkkkkkkkkk
	for word, count := range expected {
		if wordCount[word] != count {
			t.Errorf("e %v word %q, %v", count, word, wordCount[word])
		}
	}
}
