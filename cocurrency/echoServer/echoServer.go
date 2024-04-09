package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Starting tcp server....")
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		go echo(conn, input.Text(), time.Second*5)
	}
	defer conn.Close()
}

func echo(conn net.Conn, respond string, delay time.Duration) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(respond))
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", respond)
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", strings.ToLower(respond))
}