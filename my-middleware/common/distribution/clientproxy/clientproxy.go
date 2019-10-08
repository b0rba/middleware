package clientproxy

// ClientProxy is a struct that holds the data need to contact the server.
//
// Members:
//  Host     - Holds an ip address.
//  Port     - Stores the used port.
//  ID       - Identifies the process.
//  TypeName - Declares the type used.
//
type ClientProxy struct {
	Host     string
	Port     int
	ID       int
	TypeName string
}

// InitClientProxy is a function to initialize a client proxy.
//
// Parameters:
//  host     - Holds an ip address.
//  port     - Stores the used port.
//  id       - Identifies the process.
//  typename - Declares the type used.
//
// Returns:
//  a client proxy
//
func InitClientProxy(host string, port, id int, typename string) ClientProxy {
	var cp ClientProxy
	cp.Host = host
	cp.Port = port
	cp.ID = id
	cp.TypeName = typename
	return cp
}
