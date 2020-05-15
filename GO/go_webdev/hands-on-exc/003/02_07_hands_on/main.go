package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"io"
	"strings"
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
		
	var i int
	var resp_Method, resp_URI string

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			the_response := strings.Fields(ln)
			resp_Method = the_response[0]
			resp_URI = the_response[1]
			fmt.Println("METHOD:", resp_Method)
			fmt.Println("URI:", resp_URI)
		}
		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE HEND OF THE HTTP REQUEST HEADERS")
			break
		}
		i++
	}
	body := "Response body payload"
	body += "\nMETHOD:"
	body += resp_Method
	body += "\nURI:"
	body += resp_URI

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}