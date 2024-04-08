package clock 

import (
	"io"
	"log"
	"net"
	"time"
)

func main1() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:01:01\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}