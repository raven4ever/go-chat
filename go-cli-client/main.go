package main

import (
	"bufio"
	"chatty-cli/utils"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	host := flag.String("host", "127.0.0.1", "server host")
	port := flag.Int("port", 8080, "server port number")
	username := flag.String("username", utils.RandomString(13), "username")

	flag.Parse()

	connectToServer(*host, *port, *username)
}

func connectToServer(host string, port int, user string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Connected to server %s:%d...\n", host, port)
	fmt.Printf("Hello %s!\n", user)
	fmt.Println("You can now start typing your messages.")
	fmt.Println("Type 'quit' to exit.")

	// read input from stdin in a loop and send it to the server
	for {
		var input string
		fmt.Scanln(&input)

		msg := utils.Message{Username: user, Content: input}

		// fmt.Fprint(conn, msg.String())

		// send the message to the server
		conn.Write(msg.Bytes())

		// read the response from the server
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal("Error reading data:", err)
		}

		fmt.Print(string(data))

		if input == "quit" {
			log.Println("Closing connection...")
			os.Exit(0)
		}
	}
}
