package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	_ = "breakpoint"
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
func homepage(w http.ResponseWriter, r *http.Request) {
	r.Cookie("GoTestCookie")
	fmt.Fprintf(w, "Welcome to the homepage")
	fmt.Println("EndPoint hit: Homepage")
}
