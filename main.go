package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	TCP       = "tcp"
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
)

func main() {
	serverAddr := CONN_HOST + ":" + CONN_PORT
	tcpAddr, err := net.ResolveTCPAddr(TCP, serverAddr)
	if err != nil {
		fmt.Printf("Resolve address failed on %s\n", serverAddr)
		os.Exit(1)
	}

	conn, err := net.DialTCP(TCP, nil, tcpAddr)
	if err != nil {
		fmt.Printf("Dial fialed: %s\n", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	if err := scanner.Err(); err != nil {
		fmt.Printf("Read failed: %s", err.Error())
		os.Exit(1)
	}
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
