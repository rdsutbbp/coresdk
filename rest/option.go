package rest

import (
	"errors"
	"net/http"
	"os"
	"time"
)

func WithDefaultCoreRESTMode(c *RESTClient) error {
	if len(os.Args) < 5 {
		return errors.New("args count error")
	}
	// set default headers
	headers := make(map[string][]string)
	headers["Content-Type"] = []string{"application/json"}
	c.headers = headers

	// set default protocol
	c.protocol = "http"

	// set default addr
	c.addr = "127.0.0.1"

	// set default port
	c.port = os.Args[2]

	// set default driver UUID
	c.delegationUUID = os.Args[3]

	// set default lifeCycleUUID
	c.lifeCycleUUID = os.Args[4]

	// set default retry times
	c.retryTimes = 3

	// set default retry delay time
	c.retryDelay = time.Duration(1) * time.Second

	// set default client
	client := &http.Client{}
	c.client = client
	return nil
}

func WithClient(client *http.Client) Opt {
	return func(c *RESTClient) error {
		c.client = client
		return nil
	}
}

func WithHeaders(headers map[string][]string) Opt {
	return func(c *RESTClient) error {
		c.headers = headers
		return nil
	}
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

func WithDelegationUUID(delegationUUID string) Opt {
	return func(c *RESTClient) error {
		c.delegationUUID = delegationUUID
		return nil
	}
}

func WithRetryTimes(times int) Opt {
	return func(c *RESTClient) error {
		c.retryTimes = times
		return nil
	}
}

func WithRetryDelay(time time.Duration) Opt {
	return func(c *RESTClient) error {
		c.retryDelay = time
		return nil
	}
}

func WithXForwardedAuthUser(auth *XForwardedAuthUser) Opt {
	return func(c *RESTClient) error {
		c.auth = auth
		return nil
	}
}
