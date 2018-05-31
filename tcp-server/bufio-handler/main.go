package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil { log.Fatalln(err) }

	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Fatalln(err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn)  {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}

	defer conn.Close()

	fmt.Println("Code got here!")
}
