package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func clientMain() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go copyIo(os.Stdout, conn)
	copyIo(conn, os.Stdin)
}

func copyIo(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println(err)
	}
}