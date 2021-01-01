package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Connection Timeout")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(conn, "Bitch, you sayin: %s\n", line)
	}
	defer conn.Close()

	// now we get here
	// the connection will time out
	// that breaks the scanner loop
	fmt.Println("Code got here!")
}
