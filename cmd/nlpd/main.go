package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	nlp "github.com/parikxxit/ardanlab-nlp"
)

func main() {
	//routing
	http.HandleFunc("/health", HealthHandler)
	http.HandleFunc("/tokenize", TokenizeHandler)
	//run a server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error while starting server: %s", err)
	}
}

func TokenizeHandler(w http.ResponseWriter, r *http.Request) {
	var text string
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error while reading body: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	text = string(res)
	tok := nlp.Tokenize(text)
	resp := map[string]any{
		"tokens": tok,
	}
	result, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error while marsheling: %s", err)
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	_, err = w.Write(result)
	if err != nil {
		fmt.Printf("Error while writing %s", err)
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}
