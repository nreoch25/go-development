package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// any type with the method ServeHTTP with an http.ResponseWrite
// and a pointer to a http.Request is a type of Handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {

}
