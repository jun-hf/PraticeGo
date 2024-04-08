package clockwall81

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	portNumber := flag.Int("port", 8080, "port number")
	timeZone := flag.String("timezone", "US/Eastern", "current time of the time zone ")
	flag.Parse()
	loc, err := time.LoadLocation(*timeZone)
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf("localhost:%v", *portNumber)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn, loc)
	}
}

func handleConnection(con net.Conn, loc *time.Location) {
	defer con.Close()
	for {
		_, err := io.WriteString(con, time.Now().In(loc).Format("15:01:02\n"))
		if err != nil {
			return
		}
		time.Sleep(1*time.Second)
	}
}