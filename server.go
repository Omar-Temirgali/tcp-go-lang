package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server is running...")

	server, err := net.Listen("tcp", "127.0.0.1:6066")
	if err != nil {
		panic(err)
	}
	defer server.Close()
	fmt.Println("Listening on 127.0.0.1:6066")
	fmt.Println("Waiting for client...")
	for {
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Client established connection")
		go clientInteraction(conn)
	}
}

func clientInteraction(conn net.Conn) {

	mes, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message received: ", string(mes))
	res := "Hello, " + mes
	conn.Write([]byte(res + "\n"))

}
