package main

import (
	"fmt"

	namingClient "github.com/b0rba/middleware/my-middleware/client/distribution/proxies"
	"github.com/b0rba/middleware/my-middleware/common/distribution/clientproxy"
	"github.com/b0rba/middleware/my-middleware/server/distribution/invoker"
	"github.com/b0rba/middleware/my-middleware/server/distribution/proxies"
)

func main() {
	// setting the naming server on
	namingServer := proxies.InitServer()
	go namingServer.Run()
	// registering the echoer
	var clientProxy clientproxy.ClientProxy
	clientProxy = clientproxy.InitClientProxy("localhost", 8080, 2030, "Echo")
	serverNamingClient := namingClient.InitServer(clientProxy.Host)
	serverNamingClient.Bind("Echo", clientProxy)
	fmt.Println("Echo registered ((:")
	// control loop passed to middleware
	fmt.Println("Running Echoer Server ((: ")
	echoInvoker := invoker.EchoInvoker{}

	go echoInvoker.Invoke()
	aux := make(chan int)
	<-aux
}
