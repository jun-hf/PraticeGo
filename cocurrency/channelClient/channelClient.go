package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan struct{})
	ch1 := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		fmt.Println("Done")
		ch <- struct{}{}
	} ()
	ioCopy(conn, os.Stdin)
	go func() {
		time.Sleep(10*time.Second)
		conn.Close()
		ch1 <- struct{}{}
	}()
	<- ch1
	<- ch
}

func ioCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println(err)
	}
}