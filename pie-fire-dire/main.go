package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func beefMain(w http.ResponseWriter, r *http.Request) {
	wordCount, err := WordCount("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		http.Error(w, fmt.Sprintf("error %v", err), http.StatusInternalServerError)
		return
	}

	result := map[string]interface{}{
		"beef": wordCount,
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	go startGRPCServer()

	r := mux.NewRouter()
	r.HandleFunc("/beef/summary", beefMain).Methods("GET")

	fmt.Println("main 8080")
	http.ListenAndServe(":8080", r)
}
