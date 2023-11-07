package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	s := newServer()

	fmt.Println("Server is now running at http://localhost:7001")

	http.Handle("/ws", websocket.Handler(s.handleWS))
	http.ListenAndServe(":7001", nil)
}
