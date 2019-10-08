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
	// registering the calculator
	var cp clientproxy.ClientProxy
	cp = clientproxy.InitClientProxy("localhost", 8080, 2030, "Calculator")
	nclient := namingClient.InitServer(cp.Host)
	nclient.Bind("Calculator", cp)
	fmt.Println("Calculator registered!")
	// control loop passed to middleware
	fmt.Println("Multiplicator Server running!!")
	calcInvoker := invoker.CalculatorInvoker{}

	go calcInvoker.Invoke()
	c := make(chan int)
	<-c
}
