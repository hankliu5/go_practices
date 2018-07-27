package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const listenAddr = "localhost:4000"

var partner = make(chan io.ReadWriteCloser)

func main() {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// use goroutine here to prevent from block by the previous connection.
		go match(c)
	}
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
