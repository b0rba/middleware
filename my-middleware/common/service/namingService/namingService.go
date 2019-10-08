package namingService

import (
	"errors"

	"github.com/b0rba/middleware/my-middleware/common/distribution/clientproxy"
)

// NamingService is a structure for holding all the names.
//
// Members:
//  Repository - a map with the name as key and the client proxy as value.
//
type NamingService struct {
	Repository map[string]clientproxy.ClientProxy
}

// Bind is a function to add a name to the repository.
//
// Parameters:
//  name  - the name to be added.
//  proxy - the ClientProxy to with the name key.
//
// Returns:
//  an error if was unable to add the ClientProxy.
//
func (naming *NamingService) Bind(name string, proxy clientproxy.ClientProxy) error {
	_, present := naming.Repository[name]
	if present {
		return errors.New("Unable to bind " + name + ". This name is already on the naming service.")
	}
	naming.Repository[name] = proxy
	return nil
}

// Lookup is a function to get a ClientProxy from the repository.
//
// Parameters:
//  name - the name key to the ClientProxy.
//
// Returns:
//  the ClientProxy.
//  an error if was unable to find the name.
//
func (naming *NamingService) Lookup(name string) (clientproxy.ClientProxy, error) {
	cp, present := naming.Repository[name]
	if !present {
		var nilClientProxy clientproxy.ClientProxy // cannot return nil for struct
		return nilClientProxy, errors.New(name + " not found on the naming service.")
	}
	return cp, nil
}

// List is a function to return all data in the naming service.
//
// Parameters:
//  none
//
// Returns:
//  all data in the naming service.
//
func (naming *NamingService) List() map[string]clientproxy.ClientProxy {
	return naming.Repository
}
