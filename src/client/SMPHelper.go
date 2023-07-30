package client

import (
	"fmt"
	"strings"
)

const login_message string = "command:login,username:%s,password:%s"
const write_message string = "command:write,text:%s"
const read_message string = "command:read"
const logout_message string = "command:logout"
const ATTR_CODE string = "code"
const ATTR_MEANING string = "meaning"
const ATTR_AUTHORS string = "authors"
const ATTR_TEXTS string = "texts"
const command string = "command"

var GENERIC_LOGIN_ERROR map[string]string = map[string]string{
	ATTR_CODE:    "1002",
	ATTR_MEANING: "Other login error",
	command:      "login",
}

var GENERIC_WRITE_ERROR map[string]string = map[string]string{
	ATTR_CODE:    "2002",
	ATTR_MEANING: "Other write error",
	command:      "write",
}

var GENERIC_READ_ERROR map[string]string = map[string]string{
	ATTR_CODE:    "3002",
	ATTR_MEANING: "Other read error",
	command:      "read",
}
var GENERIC_LOGOUT_ERROR map[string]string = map[string]string{
	ATTR_CODE:    "4002",
	ATTR_MEANING: "Other logout error",
	command:      "logout",
}

func createWrite(message string) string {
	return fmt.Sprintf(write_message, message)
}
func createLogin(username string, password string) string {
	return fmt.Sprintf(login_message, username, password)
}
func createRead() string {
	return read_message
}
func createLogout() string {
	return logout_message
}
func parse(message string) map[string]string {
	var result = make(map[string]string)
	var pairs = strings.Split(message, ",")
	for _, pair := range pairs {
		var index = strings.Index(pair, ":")
		if index == -1 {
			continue
		}
		var key = pair[:index]
		var value = pair[index+1:]
		result[key] = value
	}
	return result
}

func extractArray(value string) []string {
	return strings.Split(value, ":")
}
