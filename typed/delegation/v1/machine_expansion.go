package v1

type MachineExpansion interface {
	CheckPortsConflict()
	ExportPorts()
	ExecCommand()
}

type CheckPortsConflictRequest struct {
	MachineID int32       `protobuf:"varint,1,opt,name=MachineID,proto3" json:"MachineID,omitempty"`
	Ports     []*PortItem `protobuf:"bytes,2,rep,name=Ports,proto3" json:"Ports,omitempty"`
}

type CheckPortsConflictReply struct {
	Ports []*PortItem `protobuf:"bytes,1,rep,name=Ports,proto3" json:"Ports,omitempty"`
}

type PortItem struct {
	Name   string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Port   int32  `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
	Result bool   `protobuf:"varint,3,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (c *machine) CheckPortsConflict() {
	//TODO implement me
	panic("implement me")
}

func (c *machine) ExportPorts() {
	//TODO implement me
	panic("implement me")
}

type ExecCommandRequest struct {
	MachineID int32    `protobuf:"varint,1,opt,name=MachineID,proto3" json:"MachineID,omitempty"`
	Command   string   `protobuf:"bytes,2,opt,name=Command,proto3" json:"Command,omitempty"`
	Timeout   int32    `protobuf:"varint,3,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	Args      []string `protobuf:"bytes,3,rep,name=Args,proto3" json:"Args,omitempty"`
	Envs      []string `protobuf:"bytes,4,rep,name=Envs,proto3" json:"Envs,omitempty"`
	Files     []*File  `protobuf:"bytes,5,rep,name=Files,proto3" json:"Files,omitempty"`
}

type File struct {
	Name    string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
}

type ExecCommandReply struct {
	Result bool   `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	Stdout string `protobuf:"bytes,2,opt,name=Stdout,proto3" json:"Stdout,omitempty"`
	StdErr string `protobuf:"bytes,3,opt,name=StdErr,proto3" json:"StdErr,omitempty"`
}

func (c *machine) ExecCommand() {
	//TODO implement me
	panic("implement me")
}
