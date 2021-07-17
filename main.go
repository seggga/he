package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "hello")
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)
}
