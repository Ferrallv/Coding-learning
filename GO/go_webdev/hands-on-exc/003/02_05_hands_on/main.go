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
	io.WriteString(conn, "This shoudl show something, but it don't?")
}