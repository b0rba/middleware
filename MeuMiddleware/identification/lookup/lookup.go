package lookup

import (
	"errors"

	"github.com/b0rba/middleware/MeuMiddleware/utils/clientProxy"
)

// NamingService holds all the names.
type NamingService struct {
	Repository map[string]clientProxy.ClientProxy
}

// Bind adds a name to the repository.
func (naming *NamingService) Bind(name string, proxy clientProxy.ClientProxy) error {
	_, present := naming.Repository[name]
	if present {
		return errors.New("Unable to bind: " + name + ". Already in the naming service.")
	}
	naming.Repository[name] = proxy
	return nil
}

// Lookup gets a ClientProxy from the repository.
func (naming *NamingService) Lookup(name string) (clientProxy.ClientProxy, error) {
	cp, present := naming.Repository[name]
	if !present {
		var nilClientProxy clientProxy.ClientProxy // cannot return nil for struct
		return nilClientProxy, errors.New(name + " not found.")
	}
	return cp, nil
}

//List return all data in the naming service.
func (naming *NamingService) List() map[string]clientProxy.ClientProxy {
	return naming.Repository
}
