package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server start...")
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/user", userHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello handler")
	fmt.Fprint(w, "Hello World")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	mail := r.Header.Get("X-Goog-Authenticated-User-Email")
	mail2 := r.Header.Get("x-goog-authenticated-user-email")
	id := r.Header.Get("X-Goog-Authenticated-User-ID")
	id2 := r.Header.Get("x-goog-authenticated-user-id")
	fmt.Fprint(w, "ID: "+id+"/"+id2+"\nmail: "+mail+"/"+mail2)
}
