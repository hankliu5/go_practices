package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

const listenAddr = "localhost:4000"

var partner = make(chan io.ReadWriteCloser)
var rootTemplate = template.Must(template.New("root").Parse(
	`
	<!DOCTYPE html>
	<html>
	<head>
	<meta charset="utf-8" />
	<script>
		websocket = new WebSocket("ws://{{.}}/socket");
		websocket.onmessage = function(m) { console.log("Received:", m.data); };
	</script>
	</head>
	</html>
	`))

type socket struct {
	io.ReadWriter
	done chan bool
}

func (s socket) Close() error {
	s.done <- true
	return nil
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.Handle("/socket", websocket.Handler(socketHandler))
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, listenAddr)
}

func socketHandler(ws *websocket.Conn) {
	s := socket{ws, make(chan bool)}
	go match(s)
	<-s.done
}

func match(c io.ReadWriteCloser) {
	fmt.Fprintln(c, "Waiting for a partner...")
	select {
	case partner <- c:
		// handled by the other goroutine
	case p := <-partner:
		chat(p, c)
	}
}

func chat(a, b io.ReadWriteCloser) {
	fmt.Fprintln(a, "Found one! Say hi.")
	fmt.Fprintln(b, "Found one! Say hi.")
	errc := make(chan error, 1)
	go cp(a, b, errc)
	go cp(b, a, errc)
	if err := <-errc; err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(a, "Leaving the chatroom")
	fmt.Fprintln(b, "Leaving the chatroom")
	a.Close()
	b.Close()
}

func cp(w io.Writer, r io.Reader, errc chan<- error) {
	_, err := io.Copy(w, r)
	errc <- err
}
