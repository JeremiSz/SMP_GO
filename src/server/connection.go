package server

import (
	"bufio"
)

func recieveMessage(buffer *bufio.Reader) (string, error) {
	return buffer.ReadString('\n')
}
func sendMessage(buffer *bufio.Writer, message string) error {
	_, err := buffer.WriteString(message + "\n")
	if err != nil {
		return err
	}
	return buffer.Flush()
}
