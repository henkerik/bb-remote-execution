// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: pkg/proto/configuration/filesystem/virtual/virtual.proto

package virtual

import (
	eviction "github.com/buildbarn/bb-storage/pkg/proto/configuration/eviction"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MountConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MountPath string `protobuf:"bytes,1,opt,name=mount_path,json=mountPath,proto3" json:"mount_path,omitempty"`
	// Types that are assignable to Backend:
	//
	//	*MountConfiguration_Fuse
	//	*MountConfiguration_Nfsv4
	Backend isMountConfiguration_Backend `protobuf_oneof:"backend"`
}

func (x *MountConfiguration) Reset() {
	*x = MountConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountConfiguration) ProtoMessage() {}

func (x *MountConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountConfiguration.ProtoReflect.Descriptor instead.
func (*MountConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_filesystem_virtual_virtual_proto_rawDescGZIP(), []int{0}
}

func (x *MountConfiguration) GetMountPath() string {
	if x != nil {
		return x.MountPath
	}
	return ""
}

func (m *MountConfiguration) GetBackend() isMountConfiguration_Backend {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (x *MountConfiguration) GetFuse() *FUSEMountConfiguration {
	if x, ok := x.GetBackend().(*MountConfiguration_Fuse); ok {
		return x.Fuse
	}
	return nil
}

func (x *MountConfiguration) GetNfsv4() *NFSv4MountConfiguration {
	if x, ok := x.GetBackend().(*MountConfiguration_Nfsv4); ok {
		return x.Nfsv4
	}
	return nil
}

type isMountConfiguration_Backend interface {
	isMountConfiguration_Backend()
}

type MountConfiguration_Fuse struct {
	Fuse *FUSEMountConfiguration `protobuf:"bytes,2,opt,name=fuse,proto3,oneof"`
}

type MountConfiguration_Nfsv4 struct {
	Nfsv4 *NFSv4MountConfiguration `protobuf:"bytes,3,opt,name=nfsv4,proto3,oneof"`
}

func (*MountConfiguration_Fuse) isMountConfiguration_Backend() {}

func (*MountConfiguration_Nfsv4) isMountConfiguration_Backend() {}

type FUSEMountConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DirectoryEntryValidity                           *durationpb.Duration `protobuf:"bytes,2,opt,name=directory_entry_validity,json=directoryEntryValidity,proto3" json:"directory_entry_validity,omitempty"`
	InodeAttributeValidity                           *durationpb.Duration `protobuf:"bytes,3,opt,name=inode_attribute_validity,json=inodeAttributeValidity,proto3" json:"inode_attribute_validity,omitempty"`
	MaximumDirtyPagesPercentage                      int32                `protobuf:"varint,4,opt,name=maximum_dirty_pages_percentage,json=maximumDirtyPagesPercentage,proto3" json:"maximum_dirty_pages_percentage,omitempty"`
	AllowOther                                       bool                 `protobuf:"varint,6,opt,name=allow_other,json=allowOther,proto3" json:"allow_other,omitempty"`
	DirectMount                                      bool                 `protobuf:"varint,7,opt,name=direct_mount,json=directMount,proto3" json:"direct_mount,omitempty"`
	InHeaderAuthenticationMetadataJmespathExpression string               `protobuf:"bytes,8,opt,name=in_header_authentication_metadata_jmespath_expression,json=inHeaderAuthenticationMetadataJmespathExpression,proto3" json:"in_header_authentication_metadata_jmespath_expression,omitempty"`
}

func (x *FUSEMountConfiguration) Reset() {
	*x = FUSEMountConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FUSEMountConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FUSEMountConfiguration) ProtoMessage() {}

func (x *FUSEMountConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FUSEMountConfiguration.ProtoReflect.Descriptor instead.
func (*FUSEMountConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_filesystem_virtual_virtual_proto_rawDescGZIP(), []int{1}
}

func (x *FUSEMountConfiguration) GetDirectoryEntryValidity() *durationpb.Duration {
	if x != nil {
		return x.DirectoryEntryValidity
	}
	return nil
}

func (x *FUSEMountConfiguration) GetInodeAttributeValidity() *durationpb.Duration {
	if x != nil {
		return x.InodeAttributeValidity
	}
	return nil
}

func (x *FUSEMountConfiguration) GetMaximumDirtyPagesPercentage() int32 {
	if x != nil {
		return x.MaximumDirtyPagesPercentage
	}
	return 0
}

func (x *FUSEMountConfiguration) GetAllowOther() bool {
	if x != nil {
		return x.AllowOther
	}
	return false
}

func (x *FUSEMountConfiguration) GetDirectMount() bool {
	if x != nil {
		return x.DirectMount
	}
	return false
}

func (x *FUSEMountConfiguration) GetInHeaderAuthenticationMetadataJmespathExpression() string {
	if x != nil {
		return x.InHeaderAuthenticationMetadataJmespathExpression
	}
	return ""
}

type NFSv4MountConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to OperatingSystem:
	//
	//	*NFSv4MountConfiguration_Darwin
	OperatingSystem      isNFSv4MountConfiguration_OperatingSystem `protobuf_oneof:"operating_system"`
	EnforcedLeaseTime    *durationpb.Duration                      `protobuf:"bytes,2,opt,name=enforced_lease_time,json=enforcedLeaseTime,proto3" json:"enforced_lease_time,omitempty"`
	AnnouncedLeaseTime   *durationpb.Duration                      `protobuf:"bytes,3,opt,name=announced_lease_time,json=announcedLeaseTime,proto3" json:"announced_lease_time,omitempty"`
	SystemAuthentication *RPCv2SystemAuthenticationConfiguration   `protobuf:"bytes,4,opt,name=system_authentication,json=systemAuthentication,proto3" json:"system_authentication,omitempty"`
}

func (x *NFSv4MountConfiguration) Reset() {
	*x = NFSv4MountConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NFSv4MountConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NFSv4MountConfiguration) ProtoMessage() {}

func (x *NFSv4MountConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NFSv4MountConfiguration.ProtoReflect.Descriptor instead.
func (*NFSv4MountConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_filesystem_virtual_virtual_proto_rawDescGZIP(), []int{2}
}

func (m *NFSv4MountConfiguration) GetOperatingSystem() isNFSv4MountConfiguration_OperatingSystem {
	if m != nil {
		return m.OperatingSystem
	}
	return nil
}

func (x *NFSv4MountConfiguration) GetDarwin() *NFSv4DarwinMountConfiguration {
	if x, ok := x.GetOperatingSystem().(*NFSv4MountConfiguration_Darwin); ok {
		return x.Darwin
	}
	return nil
}

func (x *NFSv4MountConfiguration) GetEnforcedLeaseTime() *durationpb.Duration {
	if x != nil {
		return x.EnforcedLeaseTime
	}
	return nil
}

func (x *NFSv4MountConfiguration) GetAnnouncedLeaseTime() *durationpb.Duration {
	if x != nil {
		return x.AnnouncedLeaseTime
	}
	return nil
}

func (x *NFSv4MountConfiguration) GetSystemAuthentication() *RPCv2SystemAuthenticationConfiguration {
	if x != nil {
		return x.SystemAuthentication
	}
	return nil
}

type isNFSv4MountConfiguration_OperatingSystem interface {
	isNFSv4MountConfiguration_OperatingSystem()
}

type NFSv4MountConfiguration_Darwin struct {
	Darwin *NFSv4DarwinMountConfiguration `protobuf:"bytes,1,opt,name=darwin,proto3,oneof"`
}

func (*NFSv4MountConfiguration_Darwin) isNFSv4MountConfiguration_OperatingSystem() {}

type NFSv4DarwinMountConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SocketPath                              string               `protobuf:"bytes,1,opt,name=socket_path,json=socketPath,proto3" json:"socket_path,omitempty"`
	MinimumDirectoriesAttributeCacheTimeout *durationpb.Duration `protobuf:"bytes,2,opt,name=minimum_directories_attribute_cache_timeout,json=minimumDirectoriesAttributeCacheTimeout,proto3" json:"minimum_directories_attribute_cache_timeout,omitempty"`
	MaximumDirectoriesAttributeCacheTimeout *durationpb.Duration `protobuf:"bytes,3,opt,name=maximum_directories_attribute_cache_timeout,json=maximumDirectoriesAttributeCacheTimeout,proto3" json:"maximum_directories_attribute_cache_timeout,omitempty"`
}

func (x *NFSv4DarwinMountConfiguration) Reset() {
	*x = NFSv4DarwinMountConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NFSv4DarwinMountConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NFSv4DarwinMountConfiguration) ProtoMessage() {}

func (x *NFSv4DarwinMountConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NFSv4DarwinMountConfiguration.ProtoReflect.Descriptor instead.
func (*NFSv4DarwinMountConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_filesystem_virtual_virtual_proto_rawDescGZIP(), []int{3}
}

func (x *NFSv4DarwinMountConfiguration) GetSocketPath() string {
	if x != nil {
		return x.SocketPath
	}
	return ""
}

func (x *NFSv4DarwinMountConfiguration) GetMinimumDirectoriesAttributeCacheTimeout() *durationpb.Duration {
	if x != nil {
		return x.MinimumDirectoriesAttributeCacheTimeout
	}
	return nil
}

func (x *NFSv4DarwinMountConfiguration) GetMaximumDirectoriesAttributeCacheTimeout() *durationpb.Duration {
	if x != nil {
		return x.MaximumDirectoriesAttributeCacheTimeout
	}
	return nil
}

type RPCv2SystemAuthenticationConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetadataJmespathExpression string                          `protobuf:"bytes,1,opt,name=metadata_jmespath_expression,json=metadataJmespathExpression,proto3" json:"metadata_jmespath_expression,omitempty"`
	MaximumCacheSize           int32                           `protobuf:"varint,2,opt,name=maximum_cache_size,json=maximumCacheSize,proto3" json:"maximum_cache_size,omitempty"`
	CacheReplacementPolicy     eviction.CacheReplacementPolicy `protobuf:"varint,3,opt,name=cache_replacement_policy,json=cacheReplacementPolicy,proto3,enum=buildbarn.configuration.eviction.CacheReplacementPolicy" json:"cache_replacement_policy,omitempty"`
}

func (x *RPCv2SystemAuthenticationConfiguration) Reset() {
	*x = RPCv2SystemAuthenticationConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCv2SystemAuthenticationConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCv2SystemAuthenticationConfiguration) ProtoMessage() {}

func (x *RPCv2SystemAuthenticationConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCv2SystemAuthenticationConfiguration.ProtoReflect.Descriptor instead.
func (*RPCv2SystemAuthenticationConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_filesystem_virtual_virtual_proto_rawDescGZIP(), []int{4}
}

func (x *RPCv2SystemAuthenticationConfiguration) GetMetadataJmespathExpression() string {
	if x != nil {
		return x.MetadataJmespathExpression
	}
	return ""
}

func (x *RPCv2SystemAuthenticationConfiguration) GetMaximumCacheSize() int32 {
	if x != nil {
		return x.MaximumCacheSize
	}
	return 0
}

func (x *RPCv2SystemAuthenticationConfiguration) GetCacheReplacementPolicy() eviction.CacheReplacementPolicy {
	if x != nil {
		return x.CacheReplacementPolicy
	}
	return eviction.CacheReplacementPolicy(0)
}

var File_pkg_proto_configuration_filesystem_virtual_virtual_proto protoreflect.FileDescriptor

var file_pkg_proto_configuration_filesystem_virtual_virtual_proto_rawDesc = []byte{
	0x0a, 0x38, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x2f, 0x76, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x65, 0x76, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x12, 0x4d, 0x6f, 0x75, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x58, 0x0a,
	0x04, 0x66, 0x75, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x2e, 0x46, 0x55, 0x53, 0x45, 0x4d, 0x6f, 0x75,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x04, 0x66, 0x75, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x05, 0x6e, 0x66, 0x73, 0x76, 0x34,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61,
	0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x69, 0x72, 0x74,
	0x75, 0x61, 0x6c, 0x2e, 0x4e, 0x46, 0x53, 0x76, 0x34, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x05, 0x6e,
	0x66, 0x73, 0x76, 0x34, 0x42, 0x09, 0x0a, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x22,
	0xc2, 0x03, 0x0a, 0x16, 0x46, 0x55, 0x53, 0x45, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x18, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12,
	0x53, 0x0a, 0x18, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16, 0x69, 0x6e,
	0x6f, 0x64, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x1e, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f,
	0x64, 0x69, 0x72, 0x74, 0x79, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1b, 0x6d, 0x61,
	0x78, 0x69, 0x6d, 0x75, 0x6d, 0x44, 0x69, 0x72, 0x74, 0x79, 0x50, 0x61, 0x67, 0x65, 0x73, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x5f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x6f, 0x0a,
	0x35, 0x69, 0x6e, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x6a, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x30, 0x69, 0x6e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4a, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x74, 0x68, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4a, 0x04,
	0x08, 0x05, 0x10, 0x06, 0x22, 0xb4, 0x03, 0x0a, 0x17, 0x4e, 0x46, 0x53, 0x76, 0x34, 0x4d, 0x6f,
	0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x63, 0x0a, 0x06, 0x64, 0x61, 0x72, 0x77, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x49, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x2e, 0x4e, 0x46,
	0x53, 0x76, 0x34, 0x44, 0x61, 0x72, 0x77, 0x69, 0x6e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x64,
	0x61, 0x72, 0x77, 0x69, 0x6e, 0x12, 0x49, 0x0a, 0x13, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x64, 0x5f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x65,
	0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x4b, 0x0a, 0x14, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x61, 0x6e, 0x6e, 0x6f, 0x75,
	0x6e, 0x63, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x87, 0x01,
	0x0a, 0x15, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x52, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x2e, 0x52, 0x50, 0x43, 0x76, 0x32,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x14, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x12, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0xb2, 0x02, 0x0a, 0x1d,
	0x4e, 0x46, 0x53, 0x76, 0x34, 0x44, 0x61, 0x72, 0x77, 0x69, 0x6e, 0x4d, 0x6f, 0x75, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x77,
	0x0a, 0x2b, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x27,
	0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x77, 0x0a, 0x2b, 0x6d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6d, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x27, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x22, 0x8c, 0x02, 0x0a, 0x26, 0x52, 0x50, 0x43, 0x76, 0x32, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x1c, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6a, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x74, 0x68,
	0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x1a, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4a, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x74, 0x68, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a,
	0x12, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6d, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x72, 0x0a, 0x18, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x16, 0x63, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42,
	0x55, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2f, 0x62, 0x62, 0x2d, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x2d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_configuration_filesystem_virtual_virtual_proto_rawDescOnce sync.Once
	file_pkg_proto_configuration_filesystem_virtual_virtual_proto_rawDescData = file_pkg_proto_configuration_filesystem_virtual_virtual_proto_rawDesc
)

func file_pkg_proto_configuration_filesystem_virtual_virtual_proto_rawDescGZIP() []byte {
	file_pkg_proto_configuration_filesystem_virtual_virtual_proto_rawDescOnce.Do(func() {
		file_pkg_proto_configuration_filesystem_virtual_virtual_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_configuration_filesystem_virtual_virtual_proto_rawDescData)
	})
	return file_pkg_proto_configuration_filesystem_virtual_virtual_proto_rawDescData
}

var file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_proto_configuration_filesystem_virtual_virtual_proto_goTypes = []interface{}{
	(*MountConfiguration)(nil),                     // 0: buildbarn.configuration.filesystem.virtual.MountConfiguration
	(*FUSEMountConfiguration)(nil),                 // 1: buildbarn.configuration.filesystem.virtual.FUSEMountConfiguration
	(*NFSv4MountConfiguration)(nil),                // 2: buildbarn.configuration.filesystem.virtual.NFSv4MountConfiguration
	(*NFSv4DarwinMountConfiguration)(nil),          // 3: buildbarn.configuration.filesystem.virtual.NFSv4DarwinMountConfiguration
	(*RPCv2SystemAuthenticationConfiguration)(nil), // 4: buildbarn.configuration.filesystem.virtual.RPCv2SystemAuthenticationConfiguration
	(*durationpb.Duration)(nil),                    // 5: google.protobuf.Duration
	(eviction.CacheReplacementPolicy)(0),           // 6: buildbarn.configuration.eviction.CacheReplacementPolicy
}
var file_pkg_proto_configuration_filesystem_virtual_virtual_proto_depIdxs = []int32{
	1,  // 0: buildbarn.configuration.filesystem.virtual.MountConfiguration.fuse:type_name -> buildbarn.configuration.filesystem.virtual.FUSEMountConfiguration
	2,  // 1: buildbarn.configuration.filesystem.virtual.MountConfiguration.nfsv4:type_name -> buildbarn.configuration.filesystem.virtual.NFSv4MountConfiguration
	5,  // 2: buildbarn.configuration.filesystem.virtual.FUSEMountConfiguration.directory_entry_validity:type_name -> google.protobuf.Duration
	5,  // 3: buildbarn.configuration.filesystem.virtual.FUSEMountConfiguration.inode_attribute_validity:type_name -> google.protobuf.Duration
	3,  // 4: buildbarn.configuration.filesystem.virtual.NFSv4MountConfiguration.darwin:type_name -> buildbarn.configuration.filesystem.virtual.NFSv4DarwinMountConfiguration
	5,  // 5: buildbarn.configuration.filesystem.virtual.NFSv4MountConfiguration.enforced_lease_time:type_name -> google.protobuf.Duration
	5,  // 6: buildbarn.configuration.filesystem.virtual.NFSv4MountConfiguration.announced_lease_time:type_name -> google.protobuf.Duration
	4,  // 7: buildbarn.configuration.filesystem.virtual.NFSv4MountConfiguration.system_authentication:type_name -> buildbarn.configuration.filesystem.virtual.RPCv2SystemAuthenticationConfiguration
	5,  // 8: buildbarn.configuration.filesystem.virtual.NFSv4DarwinMountConfiguration.minimum_directories_attribute_cache_timeout:type_name -> google.protobuf.Duration
	5,  // 9: buildbarn.configuration.filesystem.virtual.NFSv4DarwinMountConfiguration.maximum_directories_attribute_cache_timeout:type_name -> google.protobuf.Duration
	6,  // 10: buildbarn.configuration.filesystem.virtual.RPCv2SystemAuthenticationConfiguration.cache_replacement_policy:type_name -> buildbarn.configuration.eviction.CacheReplacementPolicy
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_pkg_proto_configuration_filesystem_virtual_virtual_proto_init() }
func file_pkg_proto_configuration_filesystem_virtual_virtual_proto_init() {
	if File_pkg_proto_configuration_filesystem_virtual_virtual_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountConfiguration); i {
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
		file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FUSEMountConfiguration); i {
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
		file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NFSv4MountConfiguration); i {
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
		file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NFSv4DarwinMountConfiguration); i {
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
		file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCv2SystemAuthenticationConfiguration); i {
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
	file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MountConfiguration_Fuse)(nil),
		(*MountConfiguration_Nfsv4)(nil),
	}
	file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*NFSv4MountConfiguration_Darwin)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_configuration_filesystem_virtual_virtual_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_configuration_filesystem_virtual_virtual_proto_goTypes,
		DependencyIndexes: file_pkg_proto_configuration_filesystem_virtual_virtual_proto_depIdxs,
		MessageInfos:      file_pkg_proto_configuration_filesystem_virtual_virtual_proto_msgTypes,
	}.Build()
	File_pkg_proto_configuration_filesystem_virtual_virtual_proto = out.File
	file_pkg_proto_configuration_filesystem_virtual_virtual_proto_rawDesc = nil
	file_pkg_proto_configuration_filesystem_virtual_virtual_proto_goTypes = nil
	file_pkg_proto_configuration_filesystem_virtual_virtual_proto_depIdxs = nil
}
