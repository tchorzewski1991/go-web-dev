package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"strings"
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

func handle(conn net.Conn) {
	io.WriteString(conn, "\nIN-MEMORY STORAGE\n\n" +
		"USE:          \n" +
		"SET key value \n" +
		"GET value     \n" +
		"DEL value     \n" +
		"EXAMPLE:      \n" +
		"SET name Joe  \n" +
		"GET name    \n\n" +
		"")

	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line   := scanner.Text()
		fields := strings.Fields(line)

		command := fields[0]
		key     := fields[1]

		logger := func(conn net.Conn, format string, data string) {
			fmt.Fprintf(conn, format, data)
		}

		switch command {
		case "GET":
			if value, ok := data[key]; ok {
				logger(conn, "%s \n", value)
			} else {
				logger(conn, "KEY NOT FOUND: %s \n", key)
			}
		case "SET":
			if len(fields) != 3 {
				fmt.Fprintln(conn, "INVALID FORMAT. USE: SET key value")
				continue
			}

			value := fields[2]
			data[key] = value
		case "DELETE":
			if _, ok := data[key]; ok {
				delete(data, key)
				logger(conn, "DELETED KEY: %s \n", key)
			} else {
				logger(conn, "KEY NOT FOUND: %s \n", key)
			}
		default:
			logger(conn, "INVALID COMMAND: %s \n", command)
		}
	}

	defer conn.Close()
}
