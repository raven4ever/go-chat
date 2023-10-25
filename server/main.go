package main

import (
	"bufio"
	"chatty/utils"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// read the port number from the command line argument
	// if no argument is provided, use 8080 as the default port number
	portNumber := flag.Int("port", 8080, "server port number")
	flag.Parse()

	if utils.IsValidPortNumber(*portNumber) {
		log.Println("Using port number:", *portNumber)
		startServer(*portNumber)
	} else {
		log.Println("Invalid port number:", *portNumber)
	}

}

// startServer starts a TCP server on the specified port and listens for incoming connections.
// It accepts incoming connections and spawns a new goroutine to handle each connection.
func startServer(port int) {
	srv, err := net.Listen("tcp", fmt.Sprint(":", port))
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer srv.Close()

	log.Println("Server started. Waiting for connections...")

	for {
		conn, err := srv.Accept()
		if err != nil {
			log.Fatal("Error accepting connection:", err)
		}

		go handleConnection(conn)
	}
}

// handleConnection handles the incoming connection from a client.
// It reads data from the connection and sends it back to the client.
// If the client sends "quit", the connection is closed.
func handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Println("Received connection from ", conn.RemoteAddr().String())

	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println("Error reading data:", err)
			break
		}

		log.Print(data)

		message := utils.NewMessage(data)
		if message == nil {
			msg := utils.Message{Username: "Server", Content: "Invalid message!"}
			conn.Write(msg.Bytes())
			continue
		}

		// if the client sends "quit", close the connection
		if strings.TrimSpace(message.Content) == "quit" {
			log.Println("Closing connection to ", conn.RemoteAddr().String())
			msg := utils.Message{Username: "Server", Content: "See ya!"}
			conn.Write(msg.Bytes())
			break
		}

		// send the data back to the client
		conn.Write(message.Bytes())
	}
}
