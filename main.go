package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "net/http"
    "github.com/gorilla/websocket"
    "github.com/Tch1b0/datastream/pkg/chunks"
)

const PORT = 8080

var upgrader = websocket.Upgrader{}

// send a chunk and make sure the client receives the data without corruption
func SendSecureChunk(conn *websocket.Conn, chunk *chunks.Chunk) error {
    for {
        // send the data chunk
        err := conn.WriteMessage(websocket.BinaryMessage, chunk.Data)
        if err != nil {
            return err
        }

        // send checksum of chunk
        err = conn.WriteMessage(websocket.TextMessage, []byte(chunk.Checksum))
        if err != nil {
            return err
        }

        // wait for message from client, "ok" will return from this function,
        // and "mismatch" and other responses will trigger the data to be resent
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

// ensure file is loaded
func LoadFile(path string) []byte {
    content, err := ioutil.ReadFile(path)
    if err != nil {
        panic(err)
    }
    return content
}

func main() {
    if (len(os.Args) < 2) {
        fmt.Println("Please pass a filepath as the first argument")
        return
    }

    fileData := LoadFile(os.Args[1])

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        // upgrade connection to websocket
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            return
        }
        // close connection on function return
        defer conn.Close()
        // generate data chunks
        data := chunks.SplitData(fileData, 128)
        for _,c := range *data {
            conn.WriteMessage(websocket.TextMessage, []byte("CHUNK"))
            SendSecureChunk(conn, c)
        }
        conn.WriteMessage(websocket.TextMessage, []byte("EOF"))
    })

    fmt.Printf("Listening on port :%d\n", PORT)
    fmt.Printf("File %s can be fetched via ws://localhost:%d/ws\n", os.Args[1], PORT)
    fmt.Println(http.ListenAndServe(fmt.Sprintf("localhost:%d", PORT), nil))
}

