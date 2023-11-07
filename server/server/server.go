package server

import (
	"bufio"
	"chatty/utils"
	"fmt"
	"log"
	"net"
	"strings"
)

type Client struct {
	Conn net.Conn
}

type Server struct {
	port    int
	clients []*Client
}

// create a new server with given port and empty client list
func NewServer(port int) *Server {
	return &Server{port: port, clients: []*Client{}}
}

// start the server
func (s *Server) Start() {
	srv, err := net.Listen("tcp", fmt.Sprint(":", s.port))
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

		// add the new client to the server
		client := &Client{Conn: conn}
		s.addClient(client)

		go s.handleConnection(client)
	}

}

// function to handle a new client connection
func (s *Server) handleConnection(client *Client) {
	defer client.Conn.Close()

	log.Println("Received connection from", client.Conn.RemoteAddr().String())

	for {
		data, err := bufio.NewReader(client.Conn).ReadString('\n')
		if err != nil {
			log.Println("Error reading data:", err)
			s.removeClient(client)
			break
		}

		log.Print(data)

		message := utils.NewMessage(data)
		if message == nil {
			msg := utils.Message{Username: "Server", Content: "Invalid message!"}
			client.Conn.Write(msg.Bytes())
			continue
		}

		// if the client sends "quit", close the connection
		if strings.TrimSpace(message.Content) == "quit" {
			log.Println("Closing connection to", client.Conn.RemoteAddr().String())
			s.removeClient(client)
			break
		}

		// send the data back to all the clients
		s.BroadcastMessage(*message)
	}

}

// broadcast a message to all clients
func (s *Server) BroadcastMessage(message utils.Message) {
	for _, client := range s.clients {
		_, err := client.Conn.Write(message.Bytes())
		if err != nil {
			log.Println("Error sending message to", client.Conn.RemoteAddr().String())
			s.removeClient(client)
		}
	}
}

// add a new client to the server
func (s *Server) addClient(client *Client) {
	s.clients = append(s.clients, client)
}

// remove a client from the server
func (s *Server) removeClient(client *Client) {
	for i, c := range s.clients {
		if c == client {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
			break
		}
	}
}
