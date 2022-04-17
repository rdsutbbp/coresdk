package v1

import (
	"context"

	"github.com/rdsutbbp/coresdk/rest"
)

type HostagentGetter interface {
	Hostagent() HostagentInterface
}

type HostagentInterface interface {
	Init(ctx context.Context, hostagent *CoreHostagent) (*CoreHostagent, error)
	Update(ctx context.Context, hostagent *CoreHostagent) error
	Query(ctx context.Context, id int32) (*CoreHostagent, error)
	HostagentExpansion
}

type hostagent struct {
	client rest.Interface
}

type CoreHostagent struct {
	ID   int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UUID string `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Addr string `protobuf:"bytes,3,opt,name=Addr,proto3" json:"Addr,omitempty"`
}

func newHostagent(c *DelegationV1Client) *hostagent {
	return &hostagent{
		client: c.RESTClient(),
	}
}

func (c *hostagent) Init(ctx context.Context, hostagent *CoreHostagent) (*CoreHostagent, error) {
	var resp CoreHostagent
	err := c.client.Post().
		SubPath("/gateway/delegation/api/v1/hostagent/init").
		Body(hostagent).
		Do(ctx).
		Into(&resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *hostagent) Update(ctx context.Context, hostagent *CoreHostagent) error {
	var resp CoreHostagent
	err := c.client.Post().
		SubPath("/gateway/delegation/api/v1/hostagent/update").
		Body(hostagent).
		Do(ctx).
		Into(&resp)

	if err != nil {
		return err
	}

	return nil
}

func (c *hostagent) Query(ctx context.Context, id int32) (*CoreHostagent, error) {
	var resp CoreHostagent

	err := c.client.Post().
		SubPath("/gateway/delegation/api/v1/hostagent/query").
		Body(struct{ ID int32 }{ID: id}).
		Do(ctx).
		Into(&resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}
