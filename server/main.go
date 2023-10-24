package main

import (
	"bufio"
	"chatty/utils"
	"flag"
	"fmt"
	"net"
	"strings"
)

func main() {
	// read the port number from the command line argument
	// if no argument is provided, use 8080 as the default port number
	portNumber := flag.Int("port", 8080, "server port number")
	flag.Parse()

	if utils.IsValidPortNumber(*portNumber) {
		fmt.Println("Using port number:", *portNumber)
		startServer(*portNumber)
	} else {
		fmt.Println("Invalid port number:", *portNumber)
	}

}

// startServer starts a TCP server on the specified port and listens for incoming connections.
// It accepts incoming connections and spawns a new goroutine to handle each connection.
func startServer(port int) {
	srv, err := net.Listen("tcp", fmt.Sprint(":", port))
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer srv.Close()

	fmt.Println("Server started. Waiting for connections...")
	for {
		conn, err := srv.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}
		go handleConnection(conn)
	}
}

// handleConnection handles the incoming connection from a client.
// It reads data from the connection and sends it back to the client.
// If the client sends "quit", the connection is closed.
func handleConnection(conn net.Conn) {
	fmt.Printf("Received connection from %s\n", conn.RemoteAddr().String())
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading data:", err)
			return
		}

		formattedData := strings.TrimSpace(data)

		// if the client sends "quit", close the connection
		if formattedData == "quit" {
			fmt.Printf("Closing connection to %s\n", conn.RemoteAddr().String())
			conn.Write([]byte("See ya!\n"))
			break
		}

		// send the data back to the client
		conn.Write([]byte(data))
	}
	conn.Close()
}
