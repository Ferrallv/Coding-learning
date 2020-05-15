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
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	defer conn.Close()

	io.WriteString(os.Stdout, "I see you connected")
	fmt.Println("Code got here.")

}