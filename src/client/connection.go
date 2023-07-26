package client

import (
	"net"
	"strings"
)

func connect(hostname string, port uint16) (net.Conn, error) {
	var address string = get_address(hostname, port)
	return net.Dial("tcp", address)
}

func send_message(connection net.Conn, message string) (string, error) {
	message = strings.ReplaceAll(message, " ", "_")
	{
		var _, err = connection.Write([]byte(message))
		if err != nil {
			log_error(err)
		}
	}

	var buffer = make([]byte, 1024)
	{
		var _, err = connection.Read(buffer)
		if err != nil {
			log_error(err)
		}
	}
	return string(buffer), nil
}

func drop(connection net.Conn) {
	connection.Close()
}
