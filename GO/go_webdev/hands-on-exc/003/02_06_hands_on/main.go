package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"io"
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
			log.Println(err)
			continue
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE HEND OF THE HTTP REQUEST HEADERS")
			break
		}
	}
	body := "Response body payload"
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, "This should show something, but it don't?")
}