// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: netdev.proto

package netdev

import (
	macvlan "github.com/checkpoint-restore/go-criu/v6/crit/images/macvlan"
	_ "github.com/checkpoint-restore/go-criu/v6/crit/images/opts"
	sit "github.com/checkpoint-restore/go-criu/v6/crit/images/sit"
	sysctl "github.com/checkpoint-restore/go-criu/v6/crit/images/sysctl"
	tun "github.com/checkpoint-restore/go-criu/v6/crit/images/tun"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NdType int32

const (
	NdType_LOOPBACK NdType = 1
	NdType_VETH     NdType = 2
	NdType_TUN      NdType = 3
	// External link -- for those CRIU only dumps and restores
	// link parameters such as flags, address, MTU, etc. The
	// existence of the link on restore should be provided
	// by the setup-namespaces script.
	NdType_EXTLINK NdType = 4
	NdType_VENET   NdType = 5 // OpenVZ device
	NdType_BRIDGE  NdType = 6
	NdType_MACVLAN NdType = 7
	NdType_SIT     NdType = 8
)

// Enum value maps for NdType.
var (
	NdType_name = map[int32]string{
		1: "LOOPBACK",
		2: "VETH",
		3: "TUN",
		4: "EXTLINK",
		5: "VENET",
		6: "BRIDGE",
		7: "MACVLAN",
		8: "SIT",
	}
	NdType_value = map[string]int32{
		"LOOPBACK": 1,
		"VETH":     2,
		"TUN":      3,
		"EXTLINK":  4,
		"VENET":    5,
		"BRIDGE":   6,
		"MACVLAN":  7,
		"SIT":      8,
	}
)

func (x NdType) Enum() *NdType {
	p := new(NdType)
	*p = x
	return p
}

func (x NdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NdType) Descriptor() protoreflect.EnumDescriptor {
	return file_netdev_proto_enumTypes[0].Descriptor()
}

func (NdType) Type() protoreflect.EnumType {
	return &file_netdev_proto_enumTypes[0]
}

func (x NdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *NdType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = NdType(num)
	return nil
}

// Deprecated: Use NdType.Descriptor instead.
func (NdType) EnumDescriptor() ([]byte, []int) {
	return file_netdev_proto_rawDescGZIP(), []int{0}
}

type NetDeviceEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        *NdType                   `protobuf:"varint,1,req,name=type,enum=NdType" json:"type,omitempty"`
	Ifindex     *uint32                   `protobuf:"varint,2,req,name=ifindex" json:"ifindex,omitempty"`
	Mtu         *uint32                   `protobuf:"varint,3,req,name=mtu" json:"mtu,omitempty"`
	Flags       *uint32                   `protobuf:"varint,4,req,name=flags" json:"flags,omitempty"`
	Name        *string                   `protobuf:"bytes,5,req,name=name" json:"name,omitempty"`
	Tun         *tun.TunLinkEntry         `protobuf:"bytes,6,opt,name=tun" json:"tun,omitempty"`
	Address     []byte                    `protobuf:"bytes,7,opt,name=address" json:"address,omitempty"`
	Conf        []int32                   `protobuf:"varint,8,rep,name=conf" json:"conf,omitempty"`
	Conf4       []*sysctl.SysctlEntry     `protobuf:"bytes,9,rep,name=conf4" json:"conf4,omitempty"`
	Conf6       []*sysctl.SysctlEntry     `protobuf:"bytes,10,rep,name=conf6" json:"conf6,omitempty"`
	Macvlan     *macvlan.MacvlanLinkEntry `protobuf:"bytes,11,opt,name=macvlan" json:"macvlan,omitempty"`
	PeerIfindex *uint32                   `protobuf:"varint,12,opt,name=peer_ifindex,json=peerIfindex" json:"peer_ifindex,omitempty"`
	PeerNsid    *uint32                   `protobuf:"varint,13,opt,name=peer_nsid,json=peerNsid" json:"peer_nsid,omitempty"`
	Master      *uint32                   `protobuf:"varint,14,opt,name=master" json:"master,omitempty"`
	Sit         *sit.SitEntry             `protobuf:"bytes,15,opt,name=sit" json:"sit,omitempty"`
}

func (x *NetDeviceEntry) Reset() {
	*x = NetDeviceEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netdev_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetDeviceEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetDeviceEntry) ProtoMessage() {}

func (x *NetDeviceEntry) ProtoReflect() protoreflect.Message {
	mi := &file_netdev_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetDeviceEntry.ProtoReflect.Descriptor instead.
func (*NetDeviceEntry) Descriptor() ([]byte, []int) {
	return file_netdev_proto_rawDescGZIP(), []int{0}
}

func (x *NetDeviceEntry) GetType() NdType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return NdType_LOOPBACK
}

func (x *NetDeviceEntry) GetIfindex() uint32 {
	if x != nil && x.Ifindex != nil {
		return *x.Ifindex
	}
	return 0
}

func (x *NetDeviceEntry) GetMtu() uint32 {
	if x != nil && x.Mtu != nil {
		return *x.Mtu
	}
	return 0
}

func (x *NetDeviceEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *NetDeviceEntry) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *NetDeviceEntry) GetTun() *tun.TunLinkEntry {
	if x != nil {
		return x.Tun
	}
	return nil
}

func (x *NetDeviceEntry) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *NetDeviceEntry) GetConf() []int32 {
	if x != nil {
		return x.Conf
	}
	return nil
}

func (x *NetDeviceEntry) GetConf4() []*sysctl.SysctlEntry {
	if x != nil {
		return x.Conf4
	}
	return nil
}

func (x *NetDeviceEntry) GetConf6() []*sysctl.SysctlEntry {
	if x != nil {
		return x.Conf6
	}
	return nil
}

func (x *NetDeviceEntry) GetMacvlan() *macvlan.MacvlanLinkEntry {
	if x != nil {
		return x.Macvlan
	}
	return nil
}

func (x *NetDeviceEntry) GetPeerIfindex() uint32 {
	if x != nil && x.PeerIfindex != nil {
		return *x.PeerIfindex
	}
	return 0
}

func (x *NetDeviceEntry) GetPeerNsid() uint32 {
	if x != nil && x.PeerNsid != nil {
		return *x.PeerNsid
	}
	return 0
}

func (x *NetDeviceEntry) GetMaster() uint32 {
	if x != nil && x.Master != nil {
		return *x.Master
	}
	return 0
}

func (x *NetDeviceEntry) GetSit() *sit.SitEntry {
	if x != nil {
		return x.Sit
	}
	return nil
}

type NetnsId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is CRIU's id which is allocated for each namespace
	TargetNsId *uint32 `protobuf:"varint,1,req,name=target_ns_id,json=targetNsId" json:"target_ns_id,omitempty"`
	// This is an id which can be used to address this namespace
	// from another network namespace. Each network namespace has
	// one set of id-s for other namespaces.
	NetnsidValue *int32 `protobuf:"varint,2,req,name=netnsid_value,json=netnsidValue" json:"netnsid_value,omitempty"`
}

func (x *NetnsId) Reset() {
	*x = NetnsId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netdev_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetnsId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetnsId) ProtoMessage() {}

func (x *NetnsId) ProtoReflect() protoreflect.Message {
	mi := &file_netdev_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetnsId.ProtoReflect.Descriptor instead.
func (*NetnsId) Descriptor() ([]byte, []int) {
	return file_netdev_proto_rawDescGZIP(), []int{1}
}

func (x *NetnsId) GetTargetNsId() uint32 {
	if x != nil && x.TargetNsId != nil {
		return *x.TargetNsId
	}
	return 0
}

func (x *NetnsId) GetNetnsidValue() int32 {
	if x != nil && x.NetnsidValue != nil {
		return *x.NetnsidValue
	}
	return 0
}

type NetnsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefConf  []int32               `protobuf:"varint,1,rep,name=def_conf,json=defConf" json:"def_conf,omitempty"`
	AllConf  []int32               `protobuf:"varint,2,rep,name=all_conf,json=allConf" json:"all_conf,omitempty"`
	DefConf4 []*sysctl.SysctlEntry `protobuf:"bytes,3,rep,name=def_conf4,json=defConf4" json:"def_conf4,omitempty"`
	AllConf4 []*sysctl.SysctlEntry `protobuf:"bytes,4,rep,name=all_conf4,json=allConf4" json:"all_conf4,omitempty"`
	DefConf6 []*sysctl.SysctlEntry `protobuf:"bytes,5,rep,name=def_conf6,json=defConf6" json:"def_conf6,omitempty"`
	AllConf6 []*sysctl.SysctlEntry `protobuf:"bytes,6,rep,name=all_conf6,json=allConf6" json:"all_conf6,omitempty"`
	Nsids    []*NetnsId            `protobuf:"bytes,7,rep,name=nsids" json:"nsids,omitempty"`
	ExtKey   *string               `protobuf:"bytes,8,opt,name=ext_key,json=extKey" json:"ext_key,omitempty"`
	UnixConf []*sysctl.SysctlEntry `protobuf:"bytes,9,rep,name=unix_conf,json=unixConf" json:"unix_conf,omitempty"`
}

func (x *NetnsEntry) Reset() {
	*x = NetnsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netdev_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetnsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetnsEntry) ProtoMessage() {}

func (x *NetnsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_netdev_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetnsEntry.ProtoReflect.Descriptor instead.
func (*NetnsEntry) Descriptor() ([]byte, []int) {
	return file_netdev_proto_rawDescGZIP(), []int{2}
}

func (x *NetnsEntry) GetDefConf() []int32 {
	if x != nil {
		return x.DefConf
	}
	return nil
}

func (x *NetnsEntry) GetAllConf() []int32 {
	if x != nil {
		return x.AllConf
	}
	return nil
}

func (x *NetnsEntry) GetDefConf4() []*sysctl.SysctlEntry {
	if x != nil {
		return x.DefConf4
	}
	return nil
}

func (x *NetnsEntry) GetAllConf4() []*sysctl.SysctlEntry {
	if x != nil {
		return x.AllConf4
	}
	return nil
}

func (x *NetnsEntry) GetDefConf6() []*sysctl.SysctlEntry {
	if x != nil {
		return x.DefConf6
	}
	return nil
}

func (x *NetnsEntry) GetAllConf6() []*sysctl.SysctlEntry {
	if x != nil {
		return x.AllConf6
	}
	return nil
}

func (x *NetnsEntry) GetNsids() []*NetnsId {
	if x != nil {
		return x.Nsids
	}
	return nil
}

func (x *NetnsEntry) GetExtKey() string {
	if x != nil && x.ExtKey != nil {
		return *x.ExtKey
	}
	return ""
}

func (x *NetnsEntry) GetUnixConf() []*sysctl.SysctlEntry {
	if x != nil {
		return x.UnixConf
	}
	return nil
}

var File_netdev_proto protoreflect.FileDescriptor

var file_netdev_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x64, 0x65, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x6d, 0x61, 0x63, 0x76, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6f,
	0x70, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x74, 0x75, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x73, 0x79, 0x73, 0x63, 0x74, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x09, 0x73, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x03,
	0x0a, 0x10, 0x6e, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e,
	0x32, 0x08, 0x2e, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x69, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x07, 0x69, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x74,
	0x75, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x6d, 0x74, 0x75, 0x12, 0x1b, 0x0a, 0x05,
	0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0d, 0x42, 0x05, 0xd2, 0x3f, 0x02,
	0x08, 0x01, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x03, 0x74, 0x75, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x75, 0x6e,
	0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x74, 0x75, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x6e, 0x66, 0x18, 0x08, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x12, 0x23,
	0x0a, 0x05, 0x63, 0x6f, 0x6e, 0x66, 0x34, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x73, 0x79, 0x73, 0x63, 0x74, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x63, 0x6f,
	0x6e, 0x66, 0x34, 0x12, 0x23, 0x0a, 0x05, 0x63, 0x6f, 0x6e, 0x66, 0x36, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x79, 0x73, 0x63, 0x74, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x63, 0x6f, 0x6e, 0x66, 0x36, 0x12, 0x2d, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x76,
	0x6c, 0x61, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x63, 0x76,
	0x6c, 0x61, 0x6e, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x6d, 0x61, 0x63, 0x76, 0x6c, 0x61, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x65, 0x65, 0x72, 0x5f,
	0x69, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70,
	0x65, 0x65, 0x72, 0x49, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65,
	0x65, 0x72, 0x5f, 0x6e, 0x73, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70,
	0x65, 0x65, 0x72, 0x4e, 0x73, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x03, 0x73, 0x69, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x73,
	0x69, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x73, 0x69, 0x74, 0x22, 0x51, 0x0a,
	0x08, 0x6e, 0x65, 0x74, 0x6e, 0x73, 0x5f, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x6e, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e, 0x73, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6e,
	0x65, 0x74, 0x6e, 0x73, 0x69, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x0c, 0x6e, 0x65, 0x74, 0x6e, 0x73, 0x69, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0xd9, 0x02, 0x0a, 0x0b, 0x6e, 0x65, 0x74, 0x6e, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x64, 0x65, 0x66, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x07, 0x64, 0x65, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x19, 0x0a, 0x08, 0x61,
	0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x61,
	0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x2a, 0x0a, 0x09, 0x64, 0x65, 0x66, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x34, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x79, 0x73, 0x63,
	0x74, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x64, 0x65, 0x66, 0x43, 0x6f, 0x6e,
	0x66, 0x34, 0x12, 0x2a, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x34, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x79, 0x73, 0x63, 0x74, 0x6c, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x34, 0x12, 0x2a,
	0x0a, 0x09, 0x64, 0x65, 0x66, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x36, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x79, 0x73, 0x63, 0x74, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x64, 0x65, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x36, 0x12, 0x2a, 0x0a, 0x09, 0x61, 0x6c,
	0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x36, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x73, 0x79, 0x73, 0x63, 0x74, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x61, 0x6c,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x36, 0x12, 0x1f, 0x0a, 0x05, 0x6e, 0x73, 0x69, 0x64, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6e, 0x65, 0x74, 0x6e, 0x73, 0x5f, 0x69, 0x64,
	0x52, 0x05, 0x6e, 0x73, 0x69, 0x64, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x78, 0x74, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x78, 0x74, 0x4b, 0x65, 0x79,
	0x12, 0x2a, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x79, 0x73, 0x63, 0x74, 0x6c, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x2a, 0x64, 0x0a, 0x07,
	0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x4f, 0x4f, 0x50, 0x42,
	0x41, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x45, 0x54, 0x48, 0x10, 0x02, 0x12,
	0x07, 0x0a, 0x03, 0x54, 0x55, 0x4e, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x54, 0x4c,
	0x49, 0x4e, 0x4b, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x45, 0x4e, 0x45, 0x54, 0x10, 0x05,
	0x12, 0x0a, 0x0a, 0x06, 0x42, 0x52, 0x49, 0x44, 0x47, 0x45, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07,
	0x4d, 0x41, 0x43, 0x56, 0x4c, 0x41, 0x4e, 0x10, 0x07, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x49, 0x54,
	0x10, 0x08,
}

var (
	file_netdev_proto_rawDescOnce sync.Once
	file_netdev_proto_rawDescData = file_netdev_proto_rawDesc
)

func file_netdev_proto_rawDescGZIP() []byte {
	file_netdev_proto_rawDescOnce.Do(func() {
		file_netdev_proto_rawDescData = protoimpl.X.CompressGZIP(file_netdev_proto_rawDescData)
	})
	return file_netdev_proto_rawDescData
}

var file_netdev_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_netdev_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_netdev_proto_goTypes = []interface{}{
	(NdType)(0),                      // 0: nd_type
	(*NetDeviceEntry)(nil),           // 1: net_device_entry
	(*NetnsId)(nil),                  // 2: netns_id
	(*NetnsEntry)(nil),               // 3: netns_entry
	(*tun.TunLinkEntry)(nil),         // 4: tun_link_entry
	(*sysctl.SysctlEntry)(nil),       // 5: sysctl_entry
	(*macvlan.MacvlanLinkEntry)(nil), // 6: macvlan_link_entry
	(*sit.SitEntry)(nil),             // 7: sit_entry
}
var file_netdev_proto_depIdxs = []int32{
	0,  // 0: net_device_entry.type:type_name -> nd_type
	4,  // 1: net_device_entry.tun:type_name -> tun_link_entry
	5,  // 2: net_device_entry.conf4:type_name -> sysctl_entry
	5,  // 3: net_device_entry.conf6:type_name -> sysctl_entry
	6,  // 4: net_device_entry.macvlan:type_name -> macvlan_link_entry
	7,  // 5: net_device_entry.sit:type_name -> sit_entry
	5,  // 6: netns_entry.def_conf4:type_name -> sysctl_entry
	5,  // 7: netns_entry.all_conf4:type_name -> sysctl_entry
	5,  // 8: netns_entry.def_conf6:type_name -> sysctl_entry
	5,  // 9: netns_entry.all_conf6:type_name -> sysctl_entry
	2,  // 10: netns_entry.nsids:type_name -> netns_id
	5,  // 11: netns_entry.unix_conf:type_name -> sysctl_entry
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_netdev_proto_init() }
func file_netdev_proto_init() {
	if File_netdev_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_netdev_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetDeviceEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_netdev_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetnsId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_netdev_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetnsEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_netdev_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_netdev_proto_goTypes,
		DependencyIndexes: file_netdev_proto_depIdxs,
		EnumInfos:         file_netdev_proto_enumTypes,
		MessageInfos:      file_netdev_proto_msgTypes,
	}.Build()
	File_netdev_proto = out.File
	file_netdev_proto_rawDesc = nil
	file_netdev_proto_goTypes = nil
	file_netdev_proto_depIdxs = nil
}