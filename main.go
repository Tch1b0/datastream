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

func testSplitter() {
    x := chunks.SplitData([]byte("Hello, my name is Johannes and I made this downloader!"), 10)
    fmt.Println(x)
}

func testChunks() {
    c := chunks.NewChunk([]byte("This is a test"))
    b := chunks.NewChunk([]byte("Another Test"))
    g := chunks.NewChunk([]byte("This is a test"))
    fmt.Println("Checksum of c: ", c.Checksum)
    fmt.Println("Checksum of b: ", b.Checksum)
    fmt.Println("c == b = ", c.Checksum == b.Checksum)
    fmt.Println("c == g = ", c.Checksum == g.Checksum)
}

