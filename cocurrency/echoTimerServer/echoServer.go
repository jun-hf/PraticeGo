package main

import ( 
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main1() {
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

func scan(c io.Reader, line chan<- string) {
	s := bufio.NewScanner(c)
	for s.Scan() {
		line <- s.Text()
	}
	if s.Err() != nil {
		log.Print("scan: ", s.Err())
	}
}

func handleConnection(conn net.Conn) {
	wg := &sync.WaitGroup{}
	timeOut := time.Second * 5
	timer := time.NewTimer(timeOut)

	line := make(chan string)
	go scan(conn, line)
	defer func() {
		wg.Wait()
		fmt.Printf("Closing connection %v", conn)
		conn.Close()
		timer.Stop()
	}()
	for {	
		select {
		case s := <- line:
			timer.Reset(timeOut)
			wg.Add(1)
			go echo(conn, s, time.Second*3, wg)
		case <-timer.C:
			return
	}}
}

func echo(conn net.Conn, respond string, delay time.Duration, w *sync.WaitGroup) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(respond))
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", respond)
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", strings.ToLower(respond))
}
