package rest

import (
	"net/http"
	"time"
)

type Interface interface {
	Verb(verb string) *Request
	Post() *Request
	Get() *Request
}

type Opt func(client *RESTClient) error

type RESTClient struct {
	protocol string
	addr     string
	port     string

	timeout time.Duration

	lifeCycleUUID string

	// Set specific behavior of the client.  If not set http.DefaultClient will be used.
	Client *http.Client
}

func (r *RESTClient) Verb(verb string) *Request {
	return NewRequest(r).Verb(verb)
}

func (r *RESTClient) Post() *Request {
	return r.Verb("POST")
}

func (r *RESTClient) Get() *Request {
	return r.Verb("GET")
}

func RESTClientFor(config *RESTClient) (*RESTClient, error) {
	httpClient := &http.Client{}

	rest := &RESTClient{
		protocol:      config.protocol,
		addr:          config.addr,
		port:          config.port,
		timeout:       config.timeout,
		lifeCycleUUID: config.lifeCycleUUID,
		Client:        httpClient,
	}
	return rest, nil
}

func WithProtocol(protocol string) Opt {
	return func(c *RESTClient) error {
		c.protocol = protocol
		return nil
	}
}

func WithCoreAddr(addr string) Opt {
	return func(c *RESTClient) error {
		c.addr = addr
		return nil
	}
}

func WithCoreListenPort(port string) Opt {
	return func(c *RESTClient) error {
		c.port = port
		return nil
	}
}

func WithLifeCycleUUID(lifeCycleUUID string) Opt {
	return func(c *RESTClient) error {
		c.lifeCycleUUID = lifeCycleUUID
		return nil
	}
}
