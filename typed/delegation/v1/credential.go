package v1

import (
	"github.com/rdsutbbp/coresdk/rest"
)

type CredentialGetter interface {
	Credential() CredentialInterface
}

type CredentialInterface interface {
	Init()
	Update()
	UpdateStatus()
	Query()
	CredentialExpansion
}

type credential struct {
	client rest.Interface
}

func newCredential(c *DelegationV1Client) *credential {
	return &credential{
		client: c.RESTClient(),
	}
}

func (c *credential) Init() {}

func (c *credential) Update() {}

func (c *credential) Query() {}

func (c *credential) UpdateStatus() {}
