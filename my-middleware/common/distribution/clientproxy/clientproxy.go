package clientproxy

// ClientProxy holds the data need to contact the server.
type ClientProxy struct {
	Host     string
	Port     int
	ID       int
	TypeName string
}

// InitClientProxy initializes a client proxy.
func InitClientProxy(host string, port, id int, nameType string) ClientProxy {
	var clientProxy ClientProxy
	clientProxy.Host = host
	clientProxy.Port = port
	clientProxy.ID = id
	clientProxy.TypeName = nameType
	return clientProxy
}
