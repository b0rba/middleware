package lookup

import (
	"errors"

	"github.com/b0rba/middleware/MeuMiddleware/utils/clientProxy"
)

// NamingService is a structure for holding all the names.
//
// Members:
//  Repository - a map with the name as key and the client proxy as value.
//
type NamingService struct {
	Repository map[string]clientProxy.ClientProxy
}

// add a name to the repository.
func (naming *NamingService) Bind(name string, proxy clientProxy.ClientProxy) error {
	_, present := naming.Repository[name]
	if present {
		return errors.New("Unable to bind: " + name + ". Already in the naming service.")
	}
	naming.Repository[name] = proxy
	return nil
}

// get a ClientProxy from the repository.
func (naming *NamingService) Lookup(name string) (clientProxy.ClientProxy, error) {
	cp, present := naming.Repository[name]
	if !present {
		var nilClientProxy clientProxy.ClientProxy // cannot return nil for struct
		return nilClientProxy, errors.New(name + " not found.")
	}
	return cp, nil
}

//return all data in the naming service.
func (naming *NamingService) List() map[string]clientProxy.ClientProxy {
	return naming.Repository
}
