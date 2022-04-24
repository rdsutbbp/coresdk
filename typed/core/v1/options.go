package v1

type PageParam struct {
	Page int32 `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	Size int32 `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	All  bool  `protobuf:"varint,3,opt,name=All,proto3" json:"All,omitempty"`
}
