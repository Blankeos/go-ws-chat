package main

import (
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

// an server object that contains the connections.
type Server struct {
	conns map[*websocket.Conn]bool
}

func newServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

// ===========================================================================
// Methods
// ===========================================================================

// For printing output to all connected clients.
func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("Write error:", err)
			}
		}(ws)
	}
}

// For reading the input.
func (s *Server) readLoop(ws *websocket.Conn) {
	buffer := make([]byte, 1024)

	for {
		n, err := ws.Read(buffer)
		if err != nil {
			// Websocket has closed when we reached end-of-file.
			if err != io.EOF {
				break
			}

			fmt.Println("Read error: ", err)
			continue
		}

		msg := buffer[:n]

		// Print to all connected clients.
		s.broadcast(msg)

		// Print to server
		fmt.Println(msg)

		// Print to this client.
		ws.Write([]byte("Thank you for the message."))
	}
}

// ===========================================================================
// Methods (Handlers)
// ===========================================================================

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client.", ws.RemoteAddr())

	s.conns[ws] = true

	s.readLoop(ws)
}
