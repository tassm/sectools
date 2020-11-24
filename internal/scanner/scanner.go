package scanner

import (
	"errors"
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
)

func worker(work, res chan int, addr string) {
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

func getPorts(portsRange string) ([]int, error) {
	o := make([]int, 2)
	s := strings.Split(portsRange, "-")
	if len(s) > 2 {
		return o, errors.New("Invalid Port Range")
	}
	for v := range s {
		p, e := strconv.Atoi(s[v])
		if e != nil {
			return o, errors.New("Invalid Port Range")
		}
		if p < 0 || p > 65535 {
			return o, errors.New("Port Numbers Invalid")
		}
		o[v] = p
	}
	if o[0] > o[1] {
		return o, errors.New("Port Range TO is greater than FROM")
	}
	return o, nil
}

/*
*	Main function
 */
func Main(address, portRange string) error {
	ports := make(chan int, 100)
	results := make(chan int)
	var open []int

	ranges, e := getPorts(portRange)
	if e != nil {
		return e
	}
	rs, re := ranges[0], ranges[1]

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, address)
	}

	go func() {
		for i := rs; i <= re; i++ {
			ports <- i
		}
	}()

	for i := rs; i <= re; i++ {
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
	return nil
}
