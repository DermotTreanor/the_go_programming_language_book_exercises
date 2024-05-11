package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"flag"
)

var port *int = flag.Int("p", 8000, "specifies the port to dial")

func main() {
	address := fmt.Sprintf("localhost:%d", *port)

	//Connect to a server
	connection, _ := net.Dial("tcp", address)

	io.Copy(os.Stdout, connection)
}
