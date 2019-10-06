package clientproxy

import (
	"github.com/b0rba/middleware/MeuMiddleware/utils/requestor"
	"github.com/b0rba/middleware/MeuMiddleware/utils/aux"
	"github.com/b0rba/middleware/MeuMiddleware/identification/aObjectReference"
)
// ClientProxy is a struct that holds the data need to contact the server
//
// Members:
//  Host     - Holds a ip address.
//  Port     - Stores the used port.
//  ID       - Identifies the client.
//  TypeName - Declares the type used.
//
type ClientProxy struct {
	Host     string
	Port     int
	ID       int
	TypeName string
	AOR      aObjectReference.AbsoluteObjectReference
}

// Echo is a funcion that receives a string and returns the same string
//
// Parameters:
// s1 - String to get echoed
//
// Returns:
// Same string
//
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