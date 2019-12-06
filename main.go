package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server start...")
    http.HandleFunc("/hello", helloHandler)
    http.ListenAndServe(":8080", nil)
}

// リクエストを処理する関数
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello handler")
    fmt.Fprint(w, "Hello World")
}

