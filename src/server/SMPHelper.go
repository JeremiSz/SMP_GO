package server

import (
	"fmt"
	"strings"
)

const command = "command"
const loginUsername = "username"
const attrText = "text"
const commandLogin = "login"
const commandRead = "read"
const commandWrite = "write"
const commandLogout = "logout"

var TEMPLATE_ERROR = "command:%s,code:%d,meaning:$s"

func parse(message string) map[string]string {
	var data = make(map[string]string)
	var tokens = strings.Split(message, ",")
	for _, pair := range tokens {
		var index = strings.Index(pair, ":")
		var key = pair[:index]
		var value = pair[index+1:]
		data[key] = value
	}
	return data
}

func createError(code uint) string {
	switch code {
	case 1002:
		return formatError(commandLogin, 1002, "Other login error")
	case 1003:
		return formatError(commandLogin, 1003, "Login not first command")
	case 1004:
		return formatError(commandLogin, 1004, "Unkown username")
	case 1005:
		return formatError(commandLogin, 1005, "Incorrect password")
	case 2002:
		return formatError(commandWrite, 2002, "Other write error")
	case 2003:
		return formatError(commandWrite, 2003, "Missing text attribute")
	case 3002:
		return formatError(commandRead, 3002, "Other read error")
	case 3003:
		return formatError(commandRead, 3003, "Invalid read attribute")
	case 4002:
		return formatError(commandLogout, 4002, "Other logout error")
	case 4003:
		return formatError(commandLogout, 4003, "Had invalid attributes")
	}
	return ""
}
func formatError(command string, code uint, meaning string) string {
	return fmt.Sprintf(TEMPLATE_ERROR, command, code, meaning)
}
func createSuccessfulLogin() string {
	return fmt.Sprintf(TEMPLATE_ERROR, commandLogin, 1001, "Successful login")
}
func createSuccessfulLogout() string {
	return fmt.Sprintf(TEMPLATE_ERROR, commandLogout, 4001, "Successful logout")
}
func createSuccessfulWrite() string {
	return fmt.Sprintf(TEMPLATE_ERROR, commandWrite, 2001, "Successful write")
}
func createSuccessfulRead(authors []string, texts []string) (string, error) {
	var sb strings.Builder
	var opening = fmt.Sprintf("command:read,code:3001,authors")
	var _, err = sb.WriteString(opening)
	if err != nil {
		return "", err
	}
	for _, author := range authors {
		var _, err = sb.WriteString(":" + author)
		if err != nil {
			return "", err
		}
	}
	var _, err2 = sb.WriteString(",texts")
	if err2 != nil {
		return "", err2
	}
	for _, text := range texts {
		var _, err = sb.WriteString(":" + text)
		if err != nil {
			return "", err
		}
	}
	return sb.String(), nil
}
