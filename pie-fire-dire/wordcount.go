package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func WordCount(url string) (map[string]int32, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error %v", err)
	}

	text := string(body)
	words := strings.FieldsFunc(text, func(r rune) bool {
		return r == ' ' || r == '.' || r == ',' || r == '\n'
	})

	wordCount := make(map[string]int32)
	for _, word := range words {
		word = strings.ToLower(word)
		wordCount[word]++
	}

	return wordCount, nil
}
