package main

import (
	"net/http"

	"golang.org/x/net/websocket"
	"github.com/sgs1370/useless"
)

func main() {
	http.Handle("/useless", websocket.Handler(func(ws *websocket.Conn) {
		ws.Write([]byte(useless.Foobar()))
	}))
	http.ListenAndServe(":3000", nil)
}
