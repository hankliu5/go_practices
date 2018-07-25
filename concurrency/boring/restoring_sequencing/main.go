package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func main() {
	c := make(chan Message)
	go boring("Joe", c)
	go boring("Ann", c)
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're both boring. I'm leaving")
}

func boring(name string, c chan Message) {
	waitForIt := make(chan bool)
	for i := 0; ; i++ {
		c <- Message{fmt.Sprintf("%s %d", name, i), waitForIt}
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		<-waitForIt
	}
}
