package main

import (
	"bufio"
	"net"
	"log"
	"io"
	"os"
	"fmt"
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

	go serve(conn)

	defer io.WriteString(os.Stdout, "I see you connected")
	defer fmt.Println("Code got here.")

}

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln != "" {
			fmt.Println(ln)
		} else {
			break
		}
	}
	defer conn.Close()
}