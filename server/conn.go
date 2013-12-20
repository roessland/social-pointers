package main

import (
    "code.google.com/p/go.net/websocket"
    "fmt"
    "log"
)

type connection struct {
    // The websocket connection
    ws *websocket.Conn

    // Buffered channel of outbound messages
    send chan string

    // Unique id for this connection
    id int
}

func (c *connection) reader() {
    for {
        // Get mouse position
        var message string
        err := websocket.Message.Receive(c.ws, &message)
        if err != nil {
            break
        }
        // Broadcast move command
        h.broadcast <- fmt.Sprintf("move mid%v ", c.id) + message
    }
    c.ws.Close()
}

func (c *connection) writer() {
    for message := range c.send {
        err := websocket.Message.Send(c.ws, message)
        if err != nil {
            break
        }
    }
    c.ws.Close()
}

func wsHandler(ws *websocket.Conn) {
    log.Print("Client connected")
    c := &connection{send: make(chan string, 256), ws: ws, id: <-idGen}
    h.register <- c
    defer func() { h.unregister <- c }()
    go c.writer()
    c.reader()
}
