package main

import (
	"fmt"
	"handlers"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// General purpose random
	rand.Seed(time.Now().UnixNano())
	handlers.LoadPacketHandlers()

	port := 7765

	http.HandleFunc("/", handlers.HttpHandler)
	http.HandleFunc("/websocket", handlers.WebSocketHandler)

	fmt.Printf("HTTP server listening on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
