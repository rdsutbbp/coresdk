package v1

import (
	"github.com/rdsutbbp/coresdk/rest"
)

type CredentialGetter interface {
	Credential() CredentialInterface
}

type CredentialInterface interface {
	List()
}

type credential struct {
	client rest.Interface
}

func newCredential(c *CoreV1Client) *credential {
	return &credential{
		client: c.RESTClient(),
	}
}

func (c *credential) List() {

}
