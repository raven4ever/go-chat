package main

import (
	"chatty-cli/client"
	"chatty-cli/utils"
	"flag"
)

func main() {
	host := flag.String("host", "127.0.0.1", "server host")
	port := flag.Int("port", 8080, "server port number")
	username := flag.String("username", utils.RandomString(13), "username")

	flag.Parse()

	client := client.NewClient(*host, *port, *username)

	client.HandleConnection()

}
