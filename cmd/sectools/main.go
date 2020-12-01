package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TasSM/sectools/internal/shell"

	"github.com/TasSM/sectools/internal/util"

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
		port := util.ValidatePort(argv[2])
		proxy.Main(argv[1], port)
	case "shell":
		if len(argv) != 2 {
			fmt.Printf("Invalid arguments for remote shell - refer to usage")
		}
		port := util.ValidatePort(argv[1])
		shell.Main(port)
	}
	os.Exit(0)
}
