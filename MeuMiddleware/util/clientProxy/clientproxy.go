package clientproxy

import (
	"github.com/b0rba/middleware/MeuMiddleware/util/requestor"
	"github.com/b0rba/middleware/MeuMiddleware/util/aux"
	"github.com/b0rba/middleware/MeuMiddleware/identification/aObjectReference"
)
// ClientProxy holds the data needed to contact the server
type ClientProxy struct {
	Host     string
	Port     int
	ID       int
	TypeName string
	AOR      aObjectReference.AbsoluteObjectReference
}

// Echo receives a string and returns the same string
func (clientproxy ClientProxy) Echo (s1 string) string {

	// Sets up the necessary structs for the requestor
	params := make([]interface{},1)
	params[0] = s1
	request := aux.Request{"Echo", params}
	inv := aux.Invocation {Host:clientproxy.Host, Port:clientproxy.Port, Request:request}

	// Invokes requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})

	return int(ter[0].(float64))
}