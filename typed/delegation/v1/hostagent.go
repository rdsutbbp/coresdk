package v1

import "github.com/rdsutbbp/coresdk/rest"

type HostagentGetter interface {
	Hostagent() HostagentInterface
}

type HostagentInterface interface {
	Init()
	Update()
	Query()
	HostagentExpansion
}

type hostagent struct {
	client rest.Interface
}

func newHostagent(c *DelegationV1Client) *hostagent {
	return &hostagent{
		client: c.RESTClient(),
	}
}

func (c *hostagent) Init() {}

func (c *hostagent) Update() {}

func (c *hostagent) Query() {}
