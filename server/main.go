package main

import (
	"chatty/server"
	"chatty/utils"
	"flag"
	"log"
)

func main() {
	// read the port number as a command line argument
	// if no argument is provided, use 8080 as the default port number
	portNumber := flag.Int("port", 8080, "server port number")
	flag.Parse()

	if utils.IsValidPortNumber(*portNumber) {
		log.Println("Using port number:", *portNumber)

		server := server.NewServer(*portNumber)

		server.Start()
	} else {
		log.Fatalln("Invalid port number:", *portNumber)
	}

}
