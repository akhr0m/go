package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalf("%s", err)
			continue
		}
		go serve(conn)
	}
}

func serve(conn net.Conn)  {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		line := scanner.Text()
		fmt.Println(line)
		if line != "" {
			fmt.Println("this is end of the http request headers")
			break
		}
	}
	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code got here.")
	io.WriteString(conn, "in here responses written")
}