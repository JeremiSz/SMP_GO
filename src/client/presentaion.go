package client

import (
	"fmt"
	"log"
)

func get_input(question string, defualt string) string {
	fmt.Println(question)
	var answer string
	fmt.Scanln(&answer)
	if answer == "" {
		return defualt
	} else {
		return answer
	}
}

func log_message(message string) {
	fmt.Println(message)
}
func log_error(message error) {
	log.Fatal(message)
}
func get_address(hostname string, port uint16) string {
	return fmt.Sprintf("%s:%d", hostname, port)
}
