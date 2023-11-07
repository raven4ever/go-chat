package client

import (
	"bufio"
	"chatty-cli/utils"
	"fmt"
	"log"
	"net"
	"os"
)

type Client struct {
	Host     string
	Port     int
	Username string
}

func NewClient(host string, port int, username string) *Client {
	return &Client{
		Host:     host,
		Port:     port,
		Username: username,
	}
}

func (c *Client) HandleConnection() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		log.Fatalln("Error connecting to server:", err)
	}
	defer conn.Close()

	log.Printf("Connected to server %s:%d...\n", c.Host, c.Port)
	log.Printf("Hello %s!\n", c.Username)
	log.Println("You can now start typing your messages.")
	log.Println("Type 'quit' to exit.")

	go ReceiveFromServer(conn)

	for {
		var input string
		fmt.Print("> ")
		fmt.Scanln(&input)

		// create a new message if the input is not empty
		if input == "" {
			continue
		}

		msg := utils.Message{Username: c.Username, Content: input}

		// send the message
		conn.Write(msg.Bytes())

		if input == "quit" {
			log.Println("Closing connection...")
			os.Exit(0)
		}
	}

}

func ReceiveFromServer(conn net.Conn) {
	for {
		// read the response from the server
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal("Error reading data:", err)
		}

		message := utils.NewMessage(data)
		if message == nil {
			log.Println("Received invalid message from server: ", data)
		}

		fmt.Print("\n")
		fmt.Print(message.String())
		fmt.Print("> ")
	}
}
