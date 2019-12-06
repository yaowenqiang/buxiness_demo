package main

import (
	"fmt"
	"time"
)

func worker(done chan bool, id int) {
	fmt.Println("working...", id)
	time.Sleep(time.Second * 10000)
	fmt.Println("done", id)
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done, 1)
	<-done
	go worker(done, 2)
	<-done
	go worker(done, 3)
	<-done
	go worker(done, 4)
	<-done
	go worker(done, 5)
	<-done
	go worker(done, 6)
	<-done
	go worker(done, 7)

	<-done
}
