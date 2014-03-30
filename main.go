package main

import (
	"fmt"
	"handlers"
	"net/http"
)

func main() {
	port := 7765

	http.HandleFunc("/", handlers.HttpHandler)
	http.HandleFunc("/websocket", handlers.WebSocketHandler)

	fmt.Printf("HTTP server listening on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
