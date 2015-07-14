package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			lnBs := []byte(ln)
			var rotatedBs = make([]byte, len(lnBs))
			for i := 0; i < len(lnBs); i++ {
				rotatedBs[(len(lnBs) - 1 - i)] = lnBs[i]
				// i = 0
				// rotatedBs[10 - 1 - 0] = lnBs[0]
				// rotatedBs[9] = lnBs[0]
				// i = 1
				// rotatedBs[10 - 1 - 1] = lnBs[1]
				// rotatedBs[8] = lnBs[1]
				// i = 2
				// rotatedBs[10 - 1 - 2] = lnBs[2]
				// rotatedBs[7] = lnBs[2]
			}
			rotated := string(rotatedBs)
			fmt.Fprintf(conn, "%s\n", strings.ToUpper(ln))
			fmt.Fprintf(conn, "%s\n", strings.ToUpper(ln[0:1]))
			fmt.Fprintf(conn, "%s\n", rotated)
			}

	}
}



