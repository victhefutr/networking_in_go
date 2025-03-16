package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Make connection to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	// Write request to the server (using proper newline format)
	request := "GET /index.html\r\n"
	_, err = fmt.Fprintf(conn, request)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}

	// Read response from server
	reader := bufio.NewReader(conn)

	fmt.Println("Response from server:")
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break // Stop reading if the server closes the connection
		}
		fmt.Print(line) // Print each line of the response
	}
}
