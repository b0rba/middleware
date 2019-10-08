package proxies

import (
	"github.com/b0rba/middleware/my-middleware/client/distribution/requestor"
	"github.com/b0rba/middleware/my-middleware/common/distribution/clientproxy"
	"github.com/b0rba/middleware/my-middleware/common/utils"
)

// CalculatorProxy is a struct that holds the data need to contact the server
//
// Members:
//  Host - Holds an ip address.
//  Port - Stores the used port.
//  ID   - Identifies the process.
//
type CalculatorProxy struct {
	Host string
	Port int
	ID   int
}

// Mul is a function to multiply a number by 2.
//
// Parameters:
//  p1 - a number.
//
// Returns:
//  the result.
//
func (proxy CalculatorProxy) Mul(p1 int) int {
	param := make([]interface{}, 1)
	param[0] = p1
	rq := utils.Request{Op: "Mul", Params: param}
	inv := utils.Invocation{Host: proxy.Host, Port: proxy.Port, Request: rq}
	reqtor := requestor.Requestor{}
	// getting reply
	reply := reqtor.Invoke(inv).([]interface{})
	result := int(reply[0].(float64))
	return result
}

// NewCalculatorProxy is a function to instantiate a new calculator based on clientproxy.
//
// Parameters:
//  cp - the clientproxy.
//
// Returns:
//  a CalculatorProxy.
//
func NewCalculatorProxy(cp clientproxy.ClientProxy) CalculatorProxy {
	return CalculatorProxy{Host: cp.Host, Port: cp.Port, ID: cp.ID}
}
