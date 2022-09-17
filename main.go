package main

import (
    "net/http"
    "fmt"
    "github.com/gorilla/websocket"
    "github.com/Tch1b0/datastream/pkg/chunks"
)

var upgrader = websocket.Upgrader{}

func SendSecureChunk(conn *websocket.Conn, chunk *chunks.Chunk) error {
    for {
        err := conn.WriteMessage(websocket.BinaryMessage, chunk.Data)
        if err != nil {
            return err
        }

        err = conn.WriteMessage(websocket.TextMessage, []byte(chunk.Checksum))
        if err != nil {
            return err
        }

        msgT, msg, err := conn.ReadMessage()
        if err != nil {
            return err
        } else if msgT != websocket.TextMessage || string(msg) != "ok" {
            continue
        } else {
            break
        }
    }
    return nil
}

func main() {
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            return
        }
        defer conn.Close()
        data := chunks.SplitData([]byte("This is a cool test message that is being sent by a secure chunk sender"), 10)
        for _,c := range *data {
            conn.WriteMessage(websocket.TextMessage, []byte("CHUNK"))
            SendSecureChunk(conn, c)
        }
        conn.WriteMessage(websocket.TextMessage, []byte("EOF"))
    })

    fmt.Println("Listening on port :8080")
    fmt.Println(http.ListenAndServe("localhost:8080", nil))
}

