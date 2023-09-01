package main

import (
	"fmt"
	"net/http"

	"gocode.ethan/ethereum-dev/handleRequest"
)

func main() {
	http.HandleFunc("/postjson1", handleRequest.HandlePostJSON)
	http.HandleFunc("/postjson2", handleRequest.HandlePostJSON_Map)

	fmt.Println("Listening on :8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
