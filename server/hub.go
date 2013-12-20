package main

import (
    "fmt"
)

type hub struct {
    // Registered connections
    connections map[*connection]bool

    // For broadcasting to all connected clients
    broadcast chan string

    // Register requests from connections
    register chan *connection

    // Unregister requests from connections
    unregister chan *connection
}

var h = hub {
    connections: make(map[*connection]bool),
    broadcast: make(chan string, 1),
    register: make(chan *connection),
    unregister: make(chan *connection),
}

func (h *hub) run() {
    for {
        select {
        case c := <-h.register:
            h.connections[c] = true
        case c := <-h.unregister:
            h.broadcast <- fmt.Sprintf("del mid%v", c.id)
            delete(h.connections, c)
            close(c.send)
        case m := <-h.broadcast:
            for c := range h.connections {
                select {
                case c.send <-m:
                default:
                    delete(h.connections, c)
                    close(c.send)
                    go c.ws.Close()
                }
            }
        }
    }
}
