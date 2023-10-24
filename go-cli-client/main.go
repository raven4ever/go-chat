package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

func main() {
	host := flag.String("host", "127.0.0.1", "server host")
	port := flag.Int("port", 8080, "server port number")
	flag.Parse()

	connectToServer(*host, *port)
}

func connectToServer(host string, port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Printf("Connected to server %s:%d...", host, port)

	// read input from stdin in a loop and send it to the server
	for {
		var input string
		fmt.Scanln(&input)
		fmt.Fprintf(conn, input+"\n")

		// read the response from the server
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading data:", err)
			return
		}
		fmt.Print("Server response: ", string(data))
	}
}
