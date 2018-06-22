package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// any type with the method ServeHTTP with an http.ResponseWrite
// and a pointer to a http.Request is a type of Handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var d hotdog
	// ServeHTTP will handle requests to this port
	// Under the scenes this uses the net.Listen and Accept
	// Functions to start a TCP connection which will
	// listen to incomping HTTP requests at port 8080
	http.ListenAndServe(":8080", d)
}
