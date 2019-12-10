package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("server start...")
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/auth", userHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello handler")
	fmt.Fprint(w, "Hello World")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	mail := r.Header.Get("X-Goog-Authenticated-User-Email")
	id := r.Header.Get("X-Goog-Authenticated-User-ID")
	fmt.Fprint(w, "ID: "+id+"\nmail: "+mail)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	auth := false
	fmt.Fprint(w, "auth: "+strconv.FormatBool(auth))
}
