package client

import (
	"net"
	"strings"
)

func login(connection net.Conn, name string, password string) (
	map[string]string, error) {
	var message = createLogin(name, password)
	var response, err = send_message(connection, message)
	if err != nil {
		log_error(err)
	}
	return parse(response), nil
}
func close(connection net.Conn) (map[string]string, error) {
	var message = createLogout()
	var response, err = send_message(connection, message)
	if err != nil {
		log_error(err)
	}
	drop(connection)
	return parse(response), nil
}

func write(connection net.Conn, message string) (
	map[string]string, error) {
	var text = strings.ReplaceAll(message, ":", "")
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, "\n", "")
	var response, err = send_message(connection, createWrite(text))
	if err != nil {
		log_error(err)
	}
	return parse(response), nil
}
func read(connection net.Conn) (map[string]string, error) {
	var message = createRead()
	var response, err = send_message(connection, message)
	if err != nil {
		log_error(err)
	}
	return parse(response), nil
}
