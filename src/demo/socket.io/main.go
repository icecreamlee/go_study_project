package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/googollee/go-socket.io"
)

var server *socketio.Server

func main() {
    var err error
    server, err = socketio.NewServer(nil)
    if err != nil {
        log.Fatal(err)
    }
    server.OnConnect("/", func(s socketio.Conn) error {
        s.SetContext("")
        fmt.Println("connected:", s.ID())
        server.JoinRoom(s.ID(), s)
        //s.Join(s.ID())
        return nil
    })
    server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
        fmt.Println("notice:", msg)
        s.Emit("reply", "have "+msg)
    })
    server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
        s.SetContext(msg)
        return "recv " + msg
    })
    server.OnEvent("/", "bye", func(s socketio.Conn) string {
        last := s.Context().(string)
        s.Emit("bye", last)
        err := s.Close()
        if err != nil {
            log.Fatal("close err: ", err)
        }
        return last
    })
    server.OnError("/", func(e error) {
        fmt.Println("meet error:", e)
    })
    server.OnDisconnect("/", func(s socketio.Conn, msg string) {
        fmt.Println("closed", msg)
    })
    go server.Serve()
    defer server.Close()

    http.Handle("/socket.io/", server)
    http.Handle("/", http.FileServer(http.Dir("./asset")))
    log.Println("Serving at localhost:8000...")
    http.HandleFunc("/p", sayhelloName)
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    //server.BroadcastToRoom(r.Form.Get("id"), "reply", "Broadcast")
    server.BroadcastToRoom(r.Form.Get("id"), "reply", "Broadcast", func (so socketio.Server, data string) {
        log.Println("Client ACK with data: ", data)
    })
    fmt.Fprintf(w, "BroadcastToRoom") //这个写入到w的是输出到客户端的
}
