package v1

import (
	"fmt"

	"github.com/rdsutbbp/coresdk/rest"
	"golang.org/x/net/context"
)

type CredentialGetter interface {
	Credential() CredentialInterface
}

type CredentialInterface interface {
	Init(ctx context.Context, credential *CoreCredential)
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

type CoreCredential struct {
	ID       int
	Name     string
	Type     string
	Version  string
	FullData string
	Args     string
}

func (c *credential) Init(ctx context.Context, credential *CoreCredential) {
	err := c.client.Post().
		SubPath("/delegation/v1/credential/init").
		Body(credential).
		Do(ctx).
		Into(credential)
	fmt.Println(err)
}

func (c *credential) Update() {}

func (c *credential) Query() {}

func (c *credential) UpdateStatus() {}
