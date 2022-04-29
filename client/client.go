package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6066")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name: ")
	text, _ := reader.ReadString('\n')

	fmt.Fprintf(conn, text+"\n")
	mes, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Server response: " + mes)
}
