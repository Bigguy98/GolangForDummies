package main

import (
	"time"
)

func getSleep(c chan string) {
	time.Sleep(time.Second * 5)
	c <- "Hello world, i wake up!"
}

func ping(c chan string) {
	// if you only send 5 message but the receiver waiting forever ==> deadlock
	// for i := 0; i < 5; i++ {
	// 	c <- "ping"
	// }
	for {
		c <- "ping"
	}
}

func pingOnChannelAfterSeconds(c chan string, t int) {
	time.Sleep(time.Second * time.Duration(t))
	c <- "ping to channel"
}

func channelDemo() {
	println("*********** CHANNEL DEMO *****************")
	c := make(chan string)
	go getSleep(c) // concurren function

	// this will wait until it get the message from channel c
	message := <-c
	println(message)
}

func bufferChannelDemo() {
	println("*********** BUFFER CHANNEL DEMO *****************")
	messages := make(chan string, 2)
	messages <- "Hello bugsmaker"
	messages <- "Nice to meet you!"
	close(messages)
	println("Pushed 2 messages to channel!")
	time.Sleep(time.Second * 2)
	for msg := range messages {
		println(msg)
	}
}

func channelAndFlowControl() {

	messages := make(chan string)
	go ping(messages)
	// listening forever
	for {
		msg := <-messages
		println(msg)
	}
}

func readOnlyChannel(message <-chan string) {
	msg := <-message
	println(msg)
	// "invalid operation: cannot send to receive-only channel message" if you uncomment the code below
	// message <- "abc"
}

func writeOnlyChannel(message chan<- string) {
	message <- "Hello world, i'm bugsmaker"
	// "invalid operation: cannot receive from send-only channel message" if you uncomment the code below
	// msg := <- message
}

func readWriteChannel(message chan string) {
	message <- "Hello world, i'm bugsmaker"
	msg := <-message
	println(msg)
}

func declarePermissionToChannel() {
	channel := make(chan string)
	go writeOnlyChannel(channel)
	go readOnlyChannel(channel)
	go readWriteChannel(channel)
	time.Sleep(time.Second * 5)
}

func usingSelectKeyword() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go pingOnChannelAfterSeconds(channel1, 3)
	go pingOnChannelAfterSeconds(channel2, 5)
	// select statement creates a series of receivers for Channels and executes whichever one receives a message first
	// if there is no message coming, it will wait forever
	// so you could add a time condition to stop blocks
	select {
	case msg1 := <-channel1:
		println("Get message from channel1 first, content: ", msg1)
	case msg2 := <-channel2:
		println("Get message from channel2 first, content: ", msg2)
	case <-time.After(time.Second * 2):
		println("There's no message")
	}
}

func main() {

	// bufferChannelDemo()
	// channelAndFlowControl()
	// declarePermissionToChannel()
	usingSelectKeyword()
}
