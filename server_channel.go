package main

import (
	"fmt"
	"time"
	"net"
	"bufio"
)

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

type ClientJob struct {
	name string
	conn net.Conn
}

func generateResponses(clientJobs chan ClientJob) {
	for {
		clientJob := <-clientJobs
		for start := time.Now(); time.Now().Sub(start) < time.Second; {
		}
		clientJob.conn.Write([]byte("Hello, " + clientJob.name))
	}
}

func main() {
	clientJobs := make(chan ClientJob)
	go generateResponses(clientJobs)
	ln, err := net.Listen("tcp", ":8082")
	check(err, "SERVER RUNNING@PORT:8082")
	for {
		conn, err := ln.Accept()
		check(err, "Accepted connection.")
		go func() {
			buf := bufio.NewReader(conn)

			for {
				name, err := buf.ReadString('\n')

				if err != nil {
					fmt.Printf("Client disconnected.\n")
					break
				}
				clientJobs <- ClientJob{name, conn}
			}
		}()
	}
}
