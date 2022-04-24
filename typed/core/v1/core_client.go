package v1

import "github.com/rdsutbbp/coresdk/rest"

type CoreV1Interface interface {
	RESTClient() rest.Interface

	CredentialGetter
	MachineGetter
	DelegationGetter
}

type CoreV1Client struct {
	restClient rest.Interface
}

func (c *CoreV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

func (c *CoreV1Client) Credential() CredentialInterface {
	return newCredential(c)
}

func (c *CoreV1Client) Machine() MachineInterface {
	return newMachine(c)
}

func (c *CoreV1Client) Delegation() DelegationInterface {
	return newDelegation(c)
}

// NewForConfig creates a new CoreV1Client for the given config.
func NewForConfig(c *rest.RESTClient) (*CoreV1Client, error) {
	config := *c
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CoreV1Client{client}, nil
}
