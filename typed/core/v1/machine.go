package v1

import (
	"github.com/rdsutbbp/coresdk/rest"
)

type MachineGetter interface {
	Machine() MachineInterface
}

type MachineInterface interface {
	List()
}

type machine struct {
	client rest.Interface
}

func newMachine(c *CoreV1Client) *machine {
	return &machine{
		client: c.RESTClient(),
	}
}

func (c *machine) List() {}
