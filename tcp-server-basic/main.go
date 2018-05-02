package main

import (
	"net"
	"log"
	"io"
)

func main() {
	// Steps:
	// 1. We create a listener for future connections.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil { log.Fatalln(err) }

	// 2. We are responsible for closing listener at the end.
	defer listener.Close()

	// 3. We create an infinite loop for accepting incoming HTTP requests
	//    through connection object.
	for {
		connection, err := listener.Accept()
		if err != nil { log.Fatalln(err) }

		// 4. connection object implements Write() function, so we can
		//    basically use any other function that expects Writer interface.
		//    Polymorphism in its purest form.
		io.WriteString(connection, "Hello from simple TCP server! \n")
		connection.Close()
	}

}
