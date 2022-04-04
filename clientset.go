package coresdk

import (
	"github.com/rdsutbbp/coresdk/rest"
	delegationv1 "github.com/rdsutbbp/coresdk/typed/delegation/v1"
)

type Clientset struct {
	delegationV1 *delegationv1.DelegationV1Client

	// add other client
	// e.g. envClient CoreEnvV1Client
}

func (c *Clientset) DelegationV1() delegationv1.DelegationV1Interface {
	return c.delegationV1
}

func NewClientWithOptions(ops ...rest.Opt) (*Clientset, error) {
	c := &rest.RESTClient{}
	for _, op := range ops {
		if err := op(c); err != nil {
			return nil, err
		}
	}
	configShallowCopy := *c
	var cs Clientset
	var err error
	cs.delegationV1, err = delegationv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}
