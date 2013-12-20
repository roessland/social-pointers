package main

import (
    "code.google.com/p/go.net/websocket"
    "log"
    "net/http"
)

func main() {
	log.Print("Running...")
    go h.run()
    http.Handle("/", websocket.Handler(wsHandler))
    if err := http.ListenAndServe(":54567", nil); err != nil {

        log.Fatal("ListenAndServe:", err)
    }
}
