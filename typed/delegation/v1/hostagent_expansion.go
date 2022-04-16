package v1

import "context"

type HostagentExpansion interface {
	QueryEnv(ctx context.Context) (*HostagentEnvReply, error)
	QueryPkg(ctx context.Context) (*HostagentPkgReply, error)
	ShutdownService() // shutdown hostagent service
}

type HostagentEnvReply struct {
	CoreAddr string `protobuf:"bytes,1,opt,name=CoreAddr,proto3" json:"CoreAddr,omitempty"`
}

// QueryEnv query hostagent env
func (c *hostagent) QueryEnv(ctx context.Context) (*HostagentEnvReply, error) {
	var resp HostagentEnvReply
	err := c.client.Get().
		SubPath("/gateway/delegation/api/v1/hostagent/env").
		Do(ctx).
		Into(&resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

type HostagentPkgReply struct {
	AMD64 *AMD64 `protobuf:"bytes,1,opt,name=AMD64,proto3" json:"AMD64,omitempty"`
	ARM64 *ARM64 `protobuf:"bytes,2,opt,name=ARM64,proto3" json:"ARM64,omitempty"`
}

type ARM64 struct {
	Binary string `protobuf:"bytes,1,opt,name=Binary,proto3" json:"Binary,omitempty"`
	Image  string `protobuf:"bytes,2,opt,name=Image,proto3" json:"Image,omitempty"`
}

type AMD64 struct {
	Binary string `protobuf:"bytes,1,opt,name=Binary,proto3" json:"Binary,omitempty"`
	Image  string `protobuf:"bytes,2,opt,name=Image,proto3" json:"Image,omitempty"`
}

// QueryPkg query hostagent pkg
// which contains arm64 and amd64 binary url and image tag
func (c *hostagent) QueryPkg(ctx context.Context) (*HostagentPkgReply, error) {
	var resp HostagentPkgReply
	err := c.client.Get().
		SubPath("/gateway/delegation/api/v1/hostagent/pkg").
		Do(ctx).
		Into(&resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

//ShutdownService shutdown hostagent service
func (c *hostagent) ShutdownService() {
	//TODO implement me
	panic("implement me")
}
