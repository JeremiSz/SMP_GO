package main

import (
	"SimpleMessagingProtocol/src/client"
	"SimpleMessagingProtocol/src/server"
	"os"
)

func main() {
	var args = os.Args[1:]
	switch args[0] {
	case "client":
		client.Run()
	case "server":
		server.Run()

	}
}
