package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func IndexHandler2(w http.ResponseWriter, r *http.Request) {
	e, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(e))
	fmt.Fprintln(w, "two")
}

func main() {
	http.HandleFunc("/", IndexHandler2)
	err := http.ListenAndServe(":8002", nil)
	if err != nil {
		fmt.Println(err)
	}
}
