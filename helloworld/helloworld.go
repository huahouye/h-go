package main

import (
	"fmt"
	"net/http"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World %s!</h1>", r.URL.Path[1:])
}

func main() {
	
	http.HandleFunc("/", defaultHandler)
	fmt.Println("启动服务器localhost:8080")
	http.ListenAndServe(":8080", nil)
}
