package client

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func connect(hostname string, port uint16) (net.Conn, error) {
	var address string = get_address(hostname, port)
	return net.Dial("tcp4", address)
}

func send_message(connection net.Conn, message string) (string, error) {
	var buffer = bufio.NewReadWriter(bufio.NewReader(connection), bufio.NewWriter(connection))
	message = strings.ReplaceAll(message, " ", "_")
	log.Println("Sending message: " + message)
	{
		//var _, err = buffer.WriteString(message + "\n")
		var _, err = connection.Write([]byte(message + "\n"))
		if err != nil {
			log_error(err)
		}
	}
	log.Println("sent")
	var result string
	{
		var err error
		result, err = buffer.ReadString(byte('\n'))
		if err != nil {
			log_error(err)
		}
	}
	log.Println("recieved")
	return string(result), nil
}

func drop(connection net.Conn) {
	connection.Close()
}
