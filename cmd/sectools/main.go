package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/TasSM/sectools/internal/proxy"
	"github.com/TasSM/sectools/internal/scanner"
)

/*
*	Entrypoint - Select operation
 */
func main() {
	argv := os.Args[1:]

	switch op := argv[0]; op {
	case "scan":
		if len(argv) != 3 {
			fmt.Printf("Invalid arguments for proxy - refer to usage")
		}
		if e := scanner.Main(argv[1], argv[2]); e != nil {
			log.Fatalf("Error TCP Scan failed: %v\n", e)
		}
	case "proxy":
		if len(argv) != 3 {
			fmt.Printf("Invalid arguments for proxy - refer to usage")
		}
		port, err := strconv.Atoi(argv[2])
		if err != nil || port > 65535 || port < 0 {
			log.Fatalln("Invalid local port specified for proxy")
		}
		proxy.Main(argv[1], port)
	}
	os.Exit(0)
}
