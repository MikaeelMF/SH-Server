package server

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func echo(ws *websocket.Conn) {
	fmt.Println("Running Websocket Echo")
	var err error
	for {
		var replay string
		if err = websocket.Message.Receive(ws, &replay); err != nil {
			fmt.Println("Cannot receive message!")
			break
		}
		fmt.Println("Message received: " + replay)
		resp := "I received " + replay
		if err = websocket.Message.Send(ws, &resp); err != nil {
			fmt.Println("Cannot send message!")
		}
	}
}

func Server() {
	http.Handle("/", websocket.Handler(echo))
	fmt.Println("Running Websocket")
	log.Fatal(http.ListenAndServe(":1908", nil))
}
