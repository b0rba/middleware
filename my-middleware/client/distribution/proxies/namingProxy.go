package proxies

import (
	"errors"

	"github.com/b0rba/middleware/my-middleware/client/distribution/requestor"
	"github.com/b0rba/middleware/my-middleware/common/distribution/clientproxy"
	"github.com/b0rba/middleware/my-middleware/common/utils"
)

// Server is a structure for managing a naming service.
//
// Members:
//  NS   - the naming service.
//  IP   - the ip of the server.
//  Port - port to the service.
//
type Server struct {
	IP   string
	Port int
}

// Lookup is a function to find the server of an object.
//
// Parameters:
//  name - the name of the object.
//
// Returns:
//  the proxy of the object.
//
func (sv Server) Lookup(name string) interface{} {
	param := make([]interface{}, 1)
	param[0] = name
	rq := utils.Request{Op: "Lookup", Params: param}
	inv := utils.Invocation{Host: sv.IP, Port: sv.Port, Request: rq}
	reqtor := requestor.Requestor{}
	// getting the reply
	reply := reqtor.Invoke(inv).([]interface{})
	if reply[1] != nil {
		err := reply[1].(error)
		utils.PrintError(err, "unable to lookup on naming proxy")
	}
	rpMap := reply[0].(map[string]interface{})
	cp := clientproxy.InitClientProxy(rpMap["Host"].(string), int(rpMap["Port"].(float64)), int(rpMap["ID"].(float64)), rpMap["TypeName"].(string))
	// getting the result
	var result interface{}
	switch cp.TypeName {
	case "Calculator":
		result = CalculatorProxy{Host: cp.Host, Port: cp.Port, ID: cp.ID}
	default:
		utils.PrintError(errors.New("unrecognized clientproxy type"), "type of the clientproxy: "+cp.TypeName)
	}
	return result
}

// Bind is a function to register an object on the naming service.
//
// Parameters:
//  name - the name of the object.
//
// Returns:
//  none
//
func (sv Server) Bind(name string, cp clientproxy.ClientProxy) {
	param := make([]interface{}, 2)
	param[0] = name
	param[1] = cp
	rq := utils.Request{Op: "Bind", Params: param}
	inv := utils.Invocation{Host: sv.IP, Port: sv.Port, Request: rq}
	reqtor := requestor.Requestor{}
	// getting the result
	reply := reqtor.Invoke(inv).([]interface{})
	if reply[0] != nil {
		err := reply[0].(error)
		utils.PrintError(err, "unable to bind on naming proxy")
	}
}

// List is a function to get all clientproxies on the server.
//
// Parameters:
//  none
//
// Returns:
//  the map with the clientproxies
//
func (sv Server) List() map[string]clientproxy.ClientProxy {
	param := make([]interface{}, 0)
	rq := utils.Request{Op: "List", Params: param}
	inv := utils.Invocation{Host: sv.IP, Port: sv.Port, Request: rq}
	reqtor := requestor.Requestor{}
	// getting the result
	reply := reqtor.Invoke(inv).([]interface{})
	result := reply[0].(map[string]clientproxy.ClientProxy)
	return result
}

// InitServer is a function to locate a server.
//
// parameters:
//  none.
//
// Returns:
//  the location of the server.
//
func InitServer(ip string) Server {
	sv := Server{IP: ip, Port: 8090}
	return sv
}
