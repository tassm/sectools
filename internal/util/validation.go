package util

import (
	"log"
	"strconv"
)

func ValidatePort(arg string) int {
	port, err := strconv.Atoi(arg)
	if err != nil || port > 65535 || port < 0 {
		log.Fatalln("Invalid local port specified for proxy")
	}
	return port
}
