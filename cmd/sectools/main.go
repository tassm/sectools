package main

import (
	"fmt"
	"os"

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
			fmt.Printf("Invalid arguments for scan refer to usage")
		}
		if e := scanner.Main(argv[1], argv[2]); e != nil {
			fmt.Printf("Error TCP Scan failed: %v\n", e)
			os.Exit(1)
		}
	}
	os.Exit(0)
}
