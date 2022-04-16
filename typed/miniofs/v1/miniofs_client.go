package v1

import "github.com/rdsutbbp/coresdk/rest"

type MiniofsV1Interface interface {
	RESTClient() rest.Interface

	ObjectGetter
}

type MiniofsV1Client struct {
	restClient rest.Interface
}

func (c *MiniofsV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

func (c *MiniofsV1Client) Object() ObjectInterface {
	return newObject(c)
}

// NewForConfig creates a new CoreV1Client for the given config.
func NewForConfig(c *rest.RESTClient) (*MiniofsV1Client, error) {
	config := *c
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MiniofsV1Client{client}, nil
}
