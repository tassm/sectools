package main

import (
	"fmt"
	"sort"

	"github.com/TasSM/sectools/internal/scanner"
)

func main() {
	ports := make(chan int, 50)
	results := make(chan int)
	address := "localhost"

	var open []int

	for i := 0; i < cap(ports); i++ {
		go scanner.Worker(ports, results, address)
	}

	// begin writing to the ports
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			open = append(open, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(open)
	for _, p := range open {
		fmt.Printf("%d Open\n", p)
	}
}
