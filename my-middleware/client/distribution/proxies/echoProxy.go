package proxies

import (
	"github.com/b0rba/middleware/my-middleware/client/distribution/requestor"
	"github.com/b0rba/middleware/my-middleware/common/distribution/clientproxy"
	"github.com/b0rba/middleware/my-middleware/common/utils"
)

// EchoProxy holds the data need to contact the server
type EchoProxy struct {
	Host string
	Port int
	ID   int
}

// Ech makes a echo
func (echoProxy EchoProxy) Ech(input string) string {
	params := make([]interface{}, 1)
	params[0] = input
	request := utils.Request{Op: "Ech", Params: params}
	invocation := utils.Invocation{Host: echoProxy.Host, Port: echoProxy.Port, Request: request}
	_requestor := requestor.Requestor{}
	// getting reply
	reply := _requestor.Invoke(invocation).([]interface{})
	result := string(reply[0].(string))
	return result
}

// NewEchoProxy is a function to instantiate a new echoer based on clientproxy.
//
// Parameters:
//  cp - the clientproxy.
//
// Returns:
//  a EchoProxy.
//
func NewEchoProxy(clientProxy clientproxy.ClientProxy) EchoProxy {
	return EchoProxy{Host: clientProxy.Host, Port: clientProxy.Port, ID: clientProxy.ID}
}
