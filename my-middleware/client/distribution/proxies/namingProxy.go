package proxies

import (
	"errors"

	"github.com/b0rba/middleware/my-middleware/client/distribution/requestor"
	"github.com/b0rba/middleware/my-middleware/common/distribution/clientproxy"
	"github.com/b0rba/middleware/my-middleware/common/utils"
)

// Server managing a naming service.
type Server struct {
	IP   string
	Port int
}

// Lookup find the server of an object.
func (server Server) Lookup(name string) interface{} {
	params := make([]interface{}, 1)
	params[0] = name
	request := utils.Request{Op: "Lookup", Params: params}
	invocation := utils.Invocation{Host: server.IP, Port: server.Port, Request: request}
	reqstor := requestor.Requestor{}
	// getting the reply
	reply := reqstor.Invoke(invocation).([]interface{})
	if reply[1] != nil {
		err := reply[1].(error)
		utils.PrintError(err, "unable to lookup on naming proxy")
	}
	replyMap := reply[0].(map[string]interface{})
	clientProxy := clientproxy.InitClientProxy(replyMap["Host"].(string), int(replyMap["Port"].(float64)), int(replyMap["ID"].(float64)), replyMap["nameType"].(string))

	// get the result
	var result interface{}
	switch clientProxy.TypeName {
	case "Echo":
		result = EchoProxy{Host: clientProxy.Host, Port: clientProxy.Port, ID: clientProxy.ID}
	default:
		utils.PrintError(errors.New("type unrecognized"), "type of clientproxy: " + clientProxy.TypeName)
	}
	return result
}

// Bind register an object on the naming service.
func (server Server) Bind(name string, cp clientproxy.ClientProxy) {
	params := make([]interface{}, 2)
	params[0] = name
	params[1] = cp
	request := utils.Request{Op: "Bind", Params: params}
	invocation := utils.Invocation{Host: server.IP, Port: server.Port, Request: request}
	reqtor := requestor.Requestor{}
	// getting the result
	reply := reqtor.Invoke(invocation).([]interface{})
	if reply[0] != nil {
		err := reply[0].(error)
		utils.PrintError(err, "unable to bind on naming proxy")
	}
}

// List get all clientProxies on the server.
func (server Server) List() map[string]clientproxy.ClientProxy {
	params := make([]interface{}, 0)
	request := utils.Request{Op: "List", Params: params}
	invocation := utils.Invocation{Host: server.IP, Port: server.Port, Request: request}
	reqtor := requestor.Requestor{}
	// getting the result
	reply := reqtor.Invoke(invocation).([]interface{})
	result := reply[0].(map[string]clientproxy.ClientProxy)
	return result
}

// InitServer locate a server.
func InitServer(ip string) Server {
	server := Server{IP: ip, Port: 8090}
	return server
}
