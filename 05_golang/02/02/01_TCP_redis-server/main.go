package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

var data = make(map[string]string)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln("listen messed up", err.Error())
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln("conn messed up", err.Error())
			panic(err)
		}

		for {
			var bs = make([]byte, 1024)
			n, err := conn.Read(bs)
			if err != nil {
				break
			}

			// check bs to see if first word is GET, SET, DEL
			// respond appropriately
			fmt.Println(bs[:n])
			fmt.Println(string(bs[:n]))

			userString := string(bs[:n])
			words := strings.Split(userString, " ")
			command := words[0]
			switch command {
			case "GET":
				//code here
				key := strings.TrimSpace(words[1])
				fmt.Println("GET key", key)
				value := data[key]
				_, err = conn.Write([]byte(value))
				if err != nil {
					break
				}
			case "SET":
				//code here
				key := words[1]
				value := words[2]
				data[key] = value
				for k, v := range data {
					_, err = conn.Write([]byte(k + " - " + v + "\n"))
					if err != nil {
						break
					}
				}
			case "DEL":
				//code here
				key := strings.TrimSpace(words[1])
				delete(data, key)
				for k, v := range data {
					_, err = conn.Write([]byte(k + " - " + v + "\n"))
					if err != nil {
						break
					}
				}
			default:
				_, err = conn.Write([]byte("Please enter a command: <GET | SET | DEL> <KEY> [VALUE]\n" + command))
				if err != nil {
					break
				}
			}
		}
		conn.Close()
	}

}
