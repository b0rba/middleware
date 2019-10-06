package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"strconv"

	"github.com/b0rba/middleware/utils"
	"github.com/b0rba/middleware/gRPC/echo/impl"
)

func main() {

	// create new instance of multiplicator
	echo := new(impl.EchoRPC)

	// create new rpc server
	server := rpc.NewServer()
	server.RegisterName("Echo", echo)

	// associate a http handler to servidor
	server.HandleHTTP("/", "/debug")

	// create tcp listen
	l, err := net.Listen("tcp", ":"+strconv.Itoa(8080))
	utils.PrintError(err, "Servidor não inicializado")

	// wait for calls
	fmt.Println("Servidor pronto (RPC-HTTP) ...\n")
	http.Serve(l, nil)
}
