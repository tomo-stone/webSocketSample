package main

import (
	"net/http"
	"golang.org/x/net/websocket"
	"log"
	"time"
	"html/template"
)

func main() {
	http.HandleFunc("/", pageHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/ws", websocket.Handler(wsHandler))
	err := http.ListenAndServe(":9000", nil)
	if err != nil{
		log.Fatal("ListenAndServe:"+err.Error())
	}
}

func wsHandler (ws *websocket.Conn){
	defer func(){
		ws.Close()
		log.Println("Connection is closed")
	}()
	for {
		t := time.Now()
		ws.Write([]byte(t.Format("2006/01/02 15:04:05.000")))
		time.Sleep(time.Second/60)
	}
}

func pageHandler (w http.ResponseWriter, _ *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil{
		log.Fatal("pageHandler:"+err.Error())
	}
	err = t.Execute(w,nil)
	if err != nil{
		log.Fatal("pageHandler:"+err.Error())
	}
	log.Println("pageHandler: OK")
}