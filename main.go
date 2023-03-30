package main

import (
  "fmt"
  "log"
  "net/http"
  
  "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    
    defer conn.Close()
    
    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Prinln(err)
            return
        }
        
        fmt.Printf("Received message: %s\n", string(p))
        
        err = conn.WriteMessage(messageType, p)
        if err != nil {
            log.Println(err)
            return
        }
    }
}

func main() {
    http.HandleFunc("/ws", handleWebSocket)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
