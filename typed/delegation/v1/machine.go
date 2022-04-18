package v1

import (
	"context"

	"github.com/rdsutbbp/coresdk/rest"
)

type MachineGetter interface {
	Machine() MachineInterface
}

type MachineInterface interface {
	Init(ctx context.Context, machine *CoreMachine) (*CoreMachine, error)
	Update(ctx context.Context, machine *CoreMachine) error
	Query(ctx context.Context, id int32) (*CoreMachine, error)
	MachineExpansion
}

type machine struct {
	client rest.Interface
}

// CoreMachine core model machine
type CoreMachine struct {
	ID           int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TimeUpdate   string `protobuf:"bytes,3,opt,name=TimeUpdate,proto3" json:"TimeUpdate,omitempty"`
	NickName     string `protobuf:"bytes,5,opt,name=NickName,proto3" json:"NickName,omitempty"`
	HostIP       string `protobuf:"bytes,6,opt,name=HostIP,proto3" json:"HostIP,omitempty"`
	VirtualIP    string `protobuf:"bytes,7,opt,name=VirtualIP,proto3" json:"VirtualIP,omitempty"`
	CPU          string `protobuf:"bytes,8,opt,name=CPU,proto3" json:"CPU,omitempty"`
	Memory       string `protobuf:"bytes,9,opt,name=Memory,proto3" json:"Memory,omitempty"`
	Disk         string `protobuf:"bytes,10,opt,name=Disk,proto3" json:"Disk,omitempty"`
	Bandwidth    string `protobuf:"bytes,11,opt,name=Bandwidth,proto3" json:"Bandwidth,omitempty"`
	State        int32  `protobuf:"varint,12,opt,name=State,proto3" json:"State,omitempty"`
	Message      string `protobuf:"bytes,13,opt,name=Message,proto3" json:"Message,omitempty"`
	Args         string `protobuf:"bytes,14,opt,name=Args,proto3" json:"Args,omitempty"`
	SystemInfo   string `protobuf:"bytes,15,opt,name=SystemInfo,proto3" json:"SystemInfo,omitempty"`
	FullData     string `protobuf:"bytes,16,opt,name=FullData,proto3" json:"FullData,omitempty"`
	HostagentID  int32  `protobuf:"varint,17,opt,name=HostagentID,proto3" json:"HostagentID,omitempty"`
	CredentialID int32  `protobuf:"varint,18,opt,name=CredentialID,proto3" json:"CredentialID,omitempty"`
	HeartBeat    string `protobuf:"bytes,19,opt,name=HeartBeat,proto3" json:"HeartBeat,omitempty"`
	Tag          string `protobuf:"bytes,20,opt,name=Tag,proto3" json:"Tag,omitempty"`
	InstallPath  string `protobuf:"bytes,21,opt,name=InstallPath,proto3" json:"InstallPath,omitempty"`
}

func newMachine(c *DelegationV1Client) *machine {
	return &machine{
		client: c.RESTClient(),
	}
}

func (c *machine) Init(ctx context.Context, machine *CoreMachine) (*CoreMachine, error) {
	var resp CoreMachine
	err := c.client.Post().
		SubPath("/gateway/delegation/api/v1/machine/init").
		Body(machine).
		Do(ctx).
		Into(&resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *machine) Update(ctx context.Context, machine *CoreMachine) error {
	var resp CoreMachine
	err := c.client.Post().
		SubPath("/gateway/delegation/api/v1/machine/update").
		Body(machine).
		Do(ctx).
		Into(&resp)

	if err != nil {
		return err
	}

	return nil
}

func (c *machine) Query(ctx context.Context, id int32) (*CoreMachine, error) {
	var resp CoreMachine
	err := c.client.Post().
		SubPath("/gateway/delegation/api/v1/machine/query").
		Body(struct{ ID int32 }{ID: id}).
		Do(ctx).
		Into(&resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}
