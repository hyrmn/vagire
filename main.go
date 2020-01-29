package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var server string
	var port string

	flag.StringVar(&server, "r", "127.0.0.1", "The default remote IP to dial")
	flag.StringVar(&port, "p", "7000", "The default port to dial")
	flag.Parse()

	address := fmt.Sprintf("%v:%v", server, port)
	fmt.Printf("Dialing %v\n", address)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Connected. Type in text. Hit <enter> to send. ctrl+c to exit.")
	reader := bufio.NewReader(os.Stdin)

	for {
		// Read in a line from stdin
		line, err := reader.ReadBytes('\n')

		// send it to the echo server
		_, err = conn.Write(line)

		// read in the line of response from the echo server
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("> %v", response)
	}
}
