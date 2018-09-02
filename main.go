package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	http.HandleFunc("/", pageHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/ws", websocket.Handler(wsHandler))
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:" + err.Error())
	}
}

func wsHandler(ws *websocket.Conn) {
	var buffer []byte
	var count,loop,timeout uint8

	defer func() {
		ws.Close()
		log.Println("Connection is closed")
	}()
	go func() {
		for {
			websocket.Message.Receive(ws, &buffer)
			if buffer != nil {
				count++
			}
		}
	}()

	for timeout < 5 {
		t := time.Now()
		ws.Write([]byte(t.Format("15:04:05.000")))
		buffer = nil
		loop++
		if loop >= 30 { // 1sec.ごとに確認
			if count == 0 {
				timeout++
			} else {
				timeout = 0
			}
			count = 0
			loop = 0
		}
		time.Sleep(time.Second/30)
	}
}

func pageHandler(w http.ResponseWriter, _ *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal("pageHandler:" + err.Error())
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal("pageHandler:" + err.Error())
	}
	log.Println("pageHandler: OK")
}
