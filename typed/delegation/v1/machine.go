package v1

import (
	"context"
	"github.com/rdsutbbp/coresdk/rest"
)

type MachineGetter interface {
	Machine() MachineInterface
}

type MachineInterface interface {
	Init(ctx context.Context, machine *InitMachineBody) (*ReportInitMachineReply, error)
	Update()
	UpdateStatus()
	Query()
	MachineExpansion
}

type machine struct {
	client rest.Interface
}

func newMachine(c *DelegationV1Client) *machine {
	return &machine{
		client: c.RESTClient(),
	}
}

// InitMachineBody init machine body
type InitMachineBody struct {
	NickName     string `protobuf:"bytes,1,opt,name=NickName,proto3" json:"NickName,omitempty"`
	HostIP       string `protobuf:"bytes,2,opt,name=HostIP,proto3" json:"HostIP,omitempty"`
	VirtualIP    string `protobuf:"bytes,3,opt,name=VirtualIP,proto3" json:"VirtualIP,omitempty"`
	CPU          string `protobuf:"bytes,4,opt,name=CPU,proto3" json:"CPU,omitempty"`
	Memory       string `protobuf:"bytes,5,opt,name=Memory,proto3" json:"Memory,omitempty"`
	Disk         string `protobuf:"bytes,6,opt,name=Disk,proto3" json:"Disk,omitempty"`
	Bandwidth    string `protobuf:"bytes,7,opt,name=Bandwidth,proto3" json:"Bandwidth,omitempty"`
	Args         string `protobuf:"bytes,10,opt,name=Args,proto3" json:"Args,omitempty"`
	SystemInfo   string `protobuf:"bytes,11,opt,name=SystemInfo,proto3" json:"SystemInfo,omitempty"`
	FullData     string `protobuf:"bytes,12,opt,name=FullData,proto3" json:"FullData,omitempty"`
	CredentialID int64  `protobuf:"varint,13,opt,name=CredentialID,proto3" json:"CredentialID,omitempty"`
	Tag          string `protobuf:"bytes,14,opt,name=Tag,proto3" json:"Tag,omitempty"`
	InstallPath  string `protobuf:"bytes,15,opt,name=InstallPath,proto3" json:"InstallPath,omitempty"`
}

// ReportInitMachineReply init machine reply
type ReportInitMachineReply struct {
	ID           int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TimeCreate   string `protobuf:"bytes,2,opt,name=TimeCreate,proto3" json:"TimeCreate,omitempty"`
	TimeUpdate   string `protobuf:"bytes,3,opt,name=TimeUpdate,proto3" json:"TimeUpdate,omitempty"`
	MachineUUID  string `protobuf:"bytes,4,opt,name=MachineUUID,proto3" json:"MachineUUID,omitempty"`
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
	DelegationID int64  `protobuf:"varint,17,opt,name=DelegationID,proto3" json:"DelegationID,omitempty"`
	CredentialID int64  `protobuf:"varint,18,opt,name=CredentialID,proto3" json:"CredentialID,omitempty"`
	HeartBeat    string `protobuf:"bytes,19,opt,name=HeartBeat,proto3" json:"HeartBeat,omitempty"`
	Tag          string `protobuf:"bytes,20,opt,name=Tag,proto3" json:"Tag,omitempty"`
	InstallPath  string `protobuf:"bytes,21,opt,name=InstallPath,proto3" json:"InstallPath,omitempty"`
}

func (c *machine) Init(ctx context.Context, machine *InitMachineBody) (*ReportInitMachineReply, error) {
	var resp ReportInitMachineReply
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

func (c *machine) Update() {}

func (c *machine) Query() {}

func (c *machine) UpdateStatus() {}
