package scanner

import (
	"fmt"
	"net"
)

func Worker(work, res chan int, addr string) {
	for p := range work {
		addr = fmt.Sprintf("%v:%d", addr, p)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			res <- 0
			continue
		}
		conn.Close()
		res <- p
	}
}
