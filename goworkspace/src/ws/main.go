package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	mensaje := "Docker (\U0001f40b) est&aacute; escrito en Go"
	fmt.Fprint(w, "<h1>"+mensaje+"</h1>")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":7800", nil)
}
