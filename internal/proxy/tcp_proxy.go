package proxy

import (
	"fmt"
	"io"
	"log"
	"net"
)

func proxy(src net.Conn, targ string) {
	log.Println("Connection Handled")
	dst, err := net.Dial("tcp", targ)
	if err != nil {
		log.Fatalln("Target Host Unreachable")
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

/*
*	TCP proxy entry point
 */
func Main(target string, port int) {
	svr, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Unable to bind to port %d\n", port)
	}
	log.Printf("Proxying TCP connections %v -> %v\n", svr.Addr(), target)
	for {
		conn, err := svr.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go proxy(conn, target)
	}
}
