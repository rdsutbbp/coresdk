package v1

import (
	"context"
	"fmt"

	"github.com/rdsutbbp/coresdk/rest"
)

type DelegationGetter interface {
	Delegation() DelegationInterface
}

type DelegationInterface interface {
	List(ctx context.Context, param *PageParam) (*DelegationListReply, error)
}

type delegation struct {
	client rest.Interface
}

func newDelegation(c *CoreV1Client) *delegation {
	return &delegation{
		client: c.RESTClient(),
	}
}

type DelegationListReply struct {
	Total       int32         `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	Delegations []*Delegation `protobuf:"bytes,2,rep,name=Delegations,proto3" json:"Delegations,omitempty"`
}

type Delegation struct {
	TimeCreate string   `protobuf:"bytes,1,opt,name=TimeCreate,proto3" json:"TimeCreate,omitempty"`
	UUID       string   `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Logo       string   `protobuf:"bytes,3,opt,name=Logo,proto3" json:"Logo,omitempty"`
	Readme     string   `protobuf:"bytes,4,opt,name=Readme,proto3" json:"Readme,omitempty"`
	Action     []string `protobuf:"bytes,5,rep,name=Action,proto3" json:"Action,omitempty"`
}

func (c *delegation) List(ctx context.Context, pageParam *PageParam) (*DelegationListReply, error) {
	var delegations DelegationListReply
	err := c.client.Post().
		SubPath("/gateway/delegation/api/v1/list").
		Params(fmt.Sprintf("?Page=%d&Size=%d&All=%t", pageParam.Page, pageParam.Size, pageParam.All)).
		Do(ctx).
		Into(&delegations)

	if err != nil {
		return nil, err
	}
	return &delegations, nil
}
