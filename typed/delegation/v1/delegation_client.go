package v1

import "github.com/rdsutbbp/coresdk/rest"

type DelegationV1Interface interface {
	RESTClient() rest.Interface

	CredentialGetter
	MachineGetter
	HostagentGetter
}

type DelegationV1Client struct {
	restClient rest.Interface
}

func (c *DelegationV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

func (c *DelegationV1Client) Credential() CredentialInterface {
	return newCredential(c)
}

func (c *DelegationV1Client) Machine() MachineInterface {
	return newMachine(c)
}

func (c *DelegationV1Client) Hostagent() HostagentInterface {
	return newHostagent(c)
}

// NewForConfig creates a new CoreV1Client for the given config.
func NewForConfig(c *rest.RESTClient) (*DelegationV1Client, error) {
	config := *c
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &DelegationV1Client{client}, nil
}
