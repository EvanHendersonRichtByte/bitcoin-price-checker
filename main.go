package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		url := "https://api.coinbase.com/v2/prices/spot?currency=IDR"
		res, err := http.Get(url)
		if err != nil {
			return
		}
		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			return
		}

		w.Write([]byte(body))
	}
	r.Body.Close()
}

func main() {
	http.HandleFunc("/test", test)
	fmt.Println("Server started at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
