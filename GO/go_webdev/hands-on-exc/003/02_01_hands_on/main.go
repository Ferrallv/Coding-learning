package main

import (
	"bufio"
	"net"
	"log"
	"io"
	"os"
	
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		io.WriteString(os.Stdout, "I see you connected")
	}
	defer conn.Close()

}