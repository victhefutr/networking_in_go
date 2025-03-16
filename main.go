package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer ln.Close()

	log.Println("Server started on port 8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue // Skip to the next iteration to avoid handling a nil connection
		}

		// Goroutine to handle multiple connections
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Read command line from the client
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading from client:", err)
		fmt.Fprintf(conn, "Error reading command: %v\n", err)
		return
	}

	// Trim newline character and split the line into command and resource
	parts := strings.SplitN(strings.TrimSpace(line), " ", 2)
	if len(parts) < 1 {
		fmt.Fprintln(conn, "Invalid command format. Expected format: COMMAND RESOURCE")
		return
	}

	command := parts[0]
	resource := ""
	if len(parts) == 2 {
		resource = parts[1]
	}

	log.Printf("Received command: %s %s\n", command, resource)

	// Handle the command
	switch command {
	case "GET":
		handleGet(conn, resource)
	default:
		fmt.Fprintf(conn, "Unknown command: %s\n", command)
	}
}

func handleGet(conn net.Conn, resource string) {
	if resource == "" {
		fmt.Fprintln(conn, "GET command requires a resource")
		return
	}

	// Placeholder response
	fmt.Fprintf(conn, "GET command received for resource: %s\n", resource)
}
