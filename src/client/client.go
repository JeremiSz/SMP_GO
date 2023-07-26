package client

import (
	"net"
	"strconv"
	"strings"
)

func Run() {
	var connection = make_conneciton()
	login_visual(connection)
	for {
		log_message("Commands\nread\nwrite\nlogout\n")
		var command = get_input("Enter command", "")
		command = strings.Trim(command, " ")
		switch command {
		case "read":
			read_visual(connection)
		case "write":
			write_visual(connection)
		case "logout":
			if logout_visual(connection) {
				return
			}
		}
	}
}

func login_visual(connection net.Conn) {
	var name = get_input("Enter name", "user")
	var password = get_input("Enter password", "password")
	login(connection, name, password)
}

func logout_visual(connection net.Conn) bool {
	var response, err = close(connection)
	if err != nil {
		log_error(err)
	}
	var is_success = response[ATTR_CODE] == "4001"
	if is_success {
		return true
	}
	log_message(response[ATTR_MEANING])
	return false
}

func write_visual(connection net.Conn) {
	var message = get_input("Enter message", "")
	if message == "" {
		log_message("Message cannot be empty")
		return
	}
	var result, err = write(connection, message)
	if err != nil {
		log_error(err)
	}
	if result[ATTR_CODE] == "2001" {
		log_message("Message sent")
		return
	}
	log_message(result[ATTR_MEANING])
}

func read_visual(connection net.Conn) {
	var result, err = read(connection)
	if err != nil {
		log_error(err)
	}
	if result[ATTR_CODE] != "3001" {
		log_message(result[ATTR_MEANING])
		return
	}
	var authors_array = result[ATTR_AUTHORS]
	var texts_array = result[ATTR_TEXTS]
	if len(authors_array) < 1 || len(texts_array) < 1 {
		log_message("No messages")
		return
	}
	var authors = extractArray(authors_array)
	var texts = extractArray(texts_array)
	var size uint
	if len(authors) < len(texts) {
		size = uint(len(authors))
	} else {
		size = uint(len(texts))
	}
	for i := uint(0); i < size; i++ {
		log_message(authors[i] + ": " + texts[i])
	}
}

func make_conneciton() net.Conn {
	var hostname string = get_input("Enter hostname", "localhost")
	var port uint16
	{
		var value, err = strconv.Atoi(get_input("Enter port", "8080"))
		if err != nil {
			log_error(err)
		}
		port = uint16(value)
	}
	var connection, err = connect(hostname, port)
	if err != nil {
		log_error(err)
	}
	return connection
}
