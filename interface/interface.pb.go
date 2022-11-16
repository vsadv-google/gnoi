//
// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.9
// source: interface/interface.proto

package intf

import (
	types "github.com/openconfig/gnoi/types"
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

// SetLoopbackModeRequest requests the provide interface be have its loopback
// mode set to the specified mode. The mode may be vendor specific. For example,
// on a transport device, available modes are "none", "mac", "phy",
// "phy_remote", "framer_facility", and "framer_terminal".
type SetLoopbackModeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interface *types.Path `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	Mode      string      `protobuf:"bytes,2,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *SetLoopbackModeRequest) Reset() {
	*x = SetLoopbackModeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_interface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLoopbackModeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLoopbackModeRequest) ProtoMessage() {}

func (x *SetLoopbackModeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interface_interface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLoopbackModeRequest.ProtoReflect.Descriptor instead.
func (*SetLoopbackModeRequest) Descriptor() ([]byte, []int) {
	return file_interface_interface_proto_rawDescGZIP(), []int{0}
}

func (x *SetLoopbackModeRequest) GetInterface() *types.Path {
	if x != nil {
		return x.Interface
	}
	return nil
}

func (x *SetLoopbackModeRequest) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

type SetLoopbackModeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetLoopbackModeResponse) Reset() {
	*x = SetLoopbackModeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_interface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLoopbackModeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLoopbackModeResponse) ProtoMessage() {}

func (x *SetLoopbackModeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interface_interface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLoopbackModeResponse.ProtoReflect.Descriptor instead.
func (*SetLoopbackModeResponse) Descriptor() ([]byte, []int) {
	return file_interface_interface_proto_rawDescGZIP(), []int{1}
}

type GetLoopbackModeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interface *types.Path `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
}

func (x *GetLoopbackModeRequest) Reset() {
	*x = GetLoopbackModeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_interface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoopbackModeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoopbackModeRequest) ProtoMessage() {}

func (x *GetLoopbackModeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interface_interface_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoopbackModeRequest.ProtoReflect.Descriptor instead.
func (*GetLoopbackModeRequest) Descriptor() ([]byte, []int) {
	return file_interface_interface_proto_rawDescGZIP(), []int{2}
}

func (x *GetLoopbackModeRequest) GetInterface() *types.Path {
	if x != nil {
		return x.Interface
	}
	return nil
}

type GetLoopbackModeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode string `protobuf:"bytes,1,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *GetLoopbackModeResponse) Reset() {
	*x = GetLoopbackModeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_interface_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoopbackModeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoopbackModeResponse) ProtoMessage() {}

func (x *GetLoopbackModeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interface_interface_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoopbackModeResponse.ProtoReflect.Descriptor instead.
func (*GetLoopbackModeResponse) Descriptor() ([]byte, []int) {
	return file_interface_interface_proto_rawDescGZIP(), []int{3}
}

func (x *GetLoopbackModeResponse) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

type ClearInterfaceCountersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interface []*types.Path `protobuf:"bytes,1,rep,name=interface,proto3" json:"interface,omitempty"`
}

func (x *ClearInterfaceCountersRequest) Reset() {
	*x = ClearInterfaceCountersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_interface_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearInterfaceCountersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearInterfaceCountersRequest) ProtoMessage() {}

func (x *ClearInterfaceCountersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interface_interface_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearInterfaceCountersRequest.ProtoReflect.Descriptor instead.
func (*ClearInterfaceCountersRequest) Descriptor() ([]byte, []int) {
	return file_interface_interface_proto_rawDescGZIP(), []int{4}
}

func (x *ClearInterfaceCountersRequest) GetInterface() []*types.Path {
	if x != nil {
		return x.Interface
	}
	return nil
}

type ClearInterfaceCountersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearInterfaceCountersResponse) Reset() {
	*x = ClearInterfaceCountersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_interface_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearInterfaceCountersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearInterfaceCountersResponse) ProtoMessage() {}

func (x *ClearInterfaceCountersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interface_interface_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearInterfaceCountersResponse.ProtoReflect.Descriptor instead.
func (*ClearInterfaceCountersResponse) Descriptor() ([]byte, []int) {
	return file_interface_interface_proto_rawDescGZIP(), []int{5}
}

var File_interface_interface_proto protoreflect.FileDescriptor

var file_interface_interface_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x6e, 0x6f,
	0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x1a, 0x11, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c,
	0x0a, 0x16, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6e,
	0x6f, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x09, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x19, 0x0a, 0x17,
	0x53, 0x65, 0x74, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x48, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x22, 0x2d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b,
	0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x22, 0x4f, 0x0a, 0x1d, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x22, 0x20, 0x0a, 0x1e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xd2, 0x02, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x12, 0x64, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63,
	0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67,
	0x6e, 0x6f, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x65,
	0x74, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x2e, 0x67, 0x6e, 0x6f,
	0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x79, 0x0a,
	0x16, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x67, 0x6e, 0x6f, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x3b, 0x69, 0x6e, 0x74, 0x66, 0xd2, 0x3e, 0x05, 0x30, 0x2e, 0x31, 0x2e, 0x30, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interface_interface_proto_rawDescOnce sync.Once
	file_interface_interface_proto_rawDescData = file_interface_interface_proto_rawDesc
)

func file_interface_interface_proto_rawDescGZIP() []byte {
	file_interface_interface_proto_rawDescOnce.Do(func() {
		file_interface_interface_proto_rawDescData = protoimpl.X.CompressGZIP(file_interface_interface_proto_rawDescData)
	})
	return file_interface_interface_proto_rawDescData
}

var file_interface_interface_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_interface_interface_proto_goTypes = []interface{}{
	(*SetLoopbackModeRequest)(nil),         // 0: gnoi.interface.SetLoopbackModeRequest
	(*SetLoopbackModeResponse)(nil),        // 1: gnoi.interface.SetLoopbackModeResponse
	(*GetLoopbackModeRequest)(nil),         // 2: gnoi.interface.GetLoopbackModeRequest
	(*GetLoopbackModeResponse)(nil),        // 3: gnoi.interface.GetLoopbackModeResponse
	(*ClearInterfaceCountersRequest)(nil),  // 4: gnoi.interface.ClearInterfaceCountersRequest
	(*ClearInterfaceCountersResponse)(nil), // 5: gnoi.interface.ClearInterfaceCountersResponse
	(*types.Path)(nil),                     // 6: gnoi.types.Path
}
var file_interface_interface_proto_depIdxs = []int32{
	6, // 0: gnoi.interface.SetLoopbackModeRequest.interface:type_name -> gnoi.types.Path
	6, // 1: gnoi.interface.GetLoopbackModeRequest.interface:type_name -> gnoi.types.Path
	6, // 2: gnoi.interface.ClearInterfaceCountersRequest.interface:type_name -> gnoi.types.Path
	0, // 3: gnoi.interface.Interface.SetLoopbackMode:input_type -> gnoi.interface.SetLoopbackModeRequest
	2, // 4: gnoi.interface.Interface.GetLoopbackMode:input_type -> gnoi.interface.GetLoopbackModeRequest
	4, // 5: gnoi.interface.Interface.ClearInterfaceCounters:input_type -> gnoi.interface.ClearInterfaceCountersRequest
	1, // 6: gnoi.interface.Interface.SetLoopbackMode:output_type -> gnoi.interface.SetLoopbackModeResponse
	3, // 7: gnoi.interface.Interface.GetLoopbackMode:output_type -> gnoi.interface.GetLoopbackModeResponse
	5, // 8: gnoi.interface.Interface.ClearInterfaceCounters:output_type -> gnoi.interface.ClearInterfaceCountersResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_interface_interface_proto_init() }
func file_interface_interface_proto_init() {
	if File_interface_interface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interface_interface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLoopbackModeRequest); i {
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
		file_interface_interface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLoopbackModeResponse); i {
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
		file_interface_interface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoopbackModeRequest); i {
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
		file_interface_interface_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoopbackModeResponse); i {
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
		file_interface_interface_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearInterfaceCountersRequest); i {
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
		file_interface_interface_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearInterfaceCountersResponse); i {
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
			RawDescriptor: file_interface_interface_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interface_interface_proto_goTypes,
		DependencyIndexes: file_interface_interface_proto_depIdxs,
		MessageInfos:      file_interface_interface_proto_msgTypes,
	}.Build()
	File_interface_interface_proto = out.File
	file_interface_interface_proto_rawDesc = nil
	file_interface_interface_proto_goTypes = nil
	file_interface_interface_proto_depIdxs = nil
}
