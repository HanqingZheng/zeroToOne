package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "one")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
	}
}
