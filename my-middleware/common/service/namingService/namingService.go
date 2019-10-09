package namingService

import (
	"errors"

	"github.com/b0rba/middleware/my-middleware/common/distribution/clientproxy"
)

// NamingService hold all the names.
type NamingService struct {
	Repository map[string]clientproxy.ClientProxy
}

// Bind adds a name to the repository.
func (naming *NamingService) Bind(name string, proxy clientproxy.ClientProxy) error {
	_, present := naming.Repository[name]
	if present {
		return errors.New("Unable to bind " + name + ". Name already exists.")
	}
	naming.Repository[name] = proxy
	return nil
}

// Lookup gets a ClientProxy from the repository.
func (naming *NamingService) Lookup(name string) (clientproxy.ClientProxy, error) {
	clientProxy, present := naming.Repository[name]
	if !present {
		var nilClientProxy clientproxy.ClientProxy // cannot return nil for struct
		return nilClientProxy, errors.New(name + " not found.")
	}
	return clientProxy, nil
}

// List returns all data in the naming service.
func (naming *NamingService) List() map[string]clientproxy.ClientProxy {
	return naming.Repository
}
