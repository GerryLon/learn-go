package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/xenolf/lego/log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	var (
		conn    *websocket.Conn
		err     error
		msgType int
		data    []byte
	)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err = upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Fprintln(w, "upgrade failed")
			return
		}

		for {
			msgType, data, err = conn.ReadMessage()
			if err != nil {
				log.Println(err)
				goto End
			}
			fmt.Printf("got %s from client\n", string(data))

			err = conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err)
				goto End
			}
		}
	End:
		conn.Close()
	})

	http.ListenAndServe(":8899", nil)
}
