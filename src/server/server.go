package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func run() {
	var socket = makeListener()
	defer socket.Close()
	log.Println("Listening on port 8080")
	for {
		var conn, err = socket.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func makeListener() net.Listener {
	var socket, err = net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	return socket
}

func handleConnection(connection net.Conn) {
	var buffer = bufio.NewReader(connection)
	var writer = bufio.NewWriter(connection)
	var request, err = recieveMessage(buffer)
	if err != nil {
		log.Println(err)
		return
	}
	var loginData = parse(request)
	var username, err2 = checkValidLogin(loginData)
	if err2 != nil {
		log.Println(err)
		return
	}
	sendMessage(writer, createSuccessfulLogin())
	for {
		var request_data, err = recieveMessage(buffer)
		if err != nil {
			log.Println(err)
			continue
		}
		var request = parse(request_data)
		switch request[command] {
		case commandRead:
			_ = read(request, writer)
		case commandWrite:
			_ = write(request, writer, username)
		case commandLogout:
			err = logout(request, writer)
			if err == nil {
				return
			}
		default:
			continue
		}
	}
}

func checkValidLogin(loginData map[string]string) (string, error) {
	var username = loginData["username"]
	var password = loginData["password"]
	if username == "" || password == "" {
		return "", fmt.Errorf("missing username or password")
	}
	return username, nil
}

func read(request map[string]string, writer *bufio.Writer) error {
	if len(request) > 1 {
		var _, _ = writer.WriteString(createError(3003))
		return nil
	}
	var authors = getAuthros()
	var texts = getTexts()
	var response, err = createSuccessfulRead(authors, texts)
	if err != nil {
		return err
	}
	sendMessage(writer, response)
	return nil
}
func write(request map[string]string, writer *bufio.Writer, username string) error {
	var text = request[attrText]
	if text == "" {
		var _, err = writer.WriteString(createError(2003))
		return err
	}
	addMessage(username, text)
	return nil
}

func logout(request map[string]string, writer *bufio.Writer) error {
	if len(request) > 1 {
		var _, err = writer.WriteString(createError(4003))
		return err
	}
	var err = sendMessage(writer, createSuccessfulLogout())
	return err
}
