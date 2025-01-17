// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.15.6
// source: terraform1.proto

package terraform1

import (
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

type Diagnostic_Severity int32

const (
	Diagnostic_INVALID Diagnostic_Severity = 0
	Diagnostic_ERROR   Diagnostic_Severity = 1
	Diagnostic_WARNING Diagnostic_Severity = 2
)

// Enum value maps for Diagnostic_Severity.
var (
	Diagnostic_Severity_name = map[int32]string{
		0: "INVALID",
		1: "ERROR",
		2: "WARNING",
	}
	Diagnostic_Severity_value = map[string]int32{
		"INVALID": 0,
		"ERROR":   1,
		"WARNING": 2,
	}
)

func (x Diagnostic_Severity) Enum() *Diagnostic_Severity {
	p := new(Diagnostic_Severity)
	*p = x
	return p
}

func (x Diagnostic_Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Diagnostic_Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_terraform1_proto_enumTypes[0].Descriptor()
}

func (Diagnostic_Severity) Type() protoreflect.EnumType {
	return &file_terraform1_proto_enumTypes[0]
}

func (x Diagnostic_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Diagnostic_Severity.Descriptor instead.
func (Diagnostic_Severity) EnumDescriptor() ([]byte, []int) {
	return file_terraform1_proto_rawDescGZIP(), []int{2, 0}
}

// Represents a selected or available version of a provider, from either a
// dependency lock object (selected) or a provider cache object (available).
//
// This message type corresponds in meaning with a single "provider" block in a
// dependency lock file, but not all messages of this type directly represent
// such a physical block.
type ProviderPackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of the provider using the canonical form of the provider
	// source address syntax.
	SourceAddr string `protobuf:"bytes,1,opt,name=source_addr,json=sourceAddr,proto3" json:"source_addr,omitempty"`
	// The version number of this provider package. Unset for (and only for)
	// built-in providers; callers may use the set-ness of this field to
	// distinguish installable vs. built-in providers without having to
	// parse the source address syntax.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// The hash strings that Terraform knows about for this provider package,
	// using the same "scheme:hash" syntax used in Terraform's dependency
	// lock file format.
	//
	// For a message representing a "selected" provider package this enumerates
	// all of the checksums that were previously loaded from a dependency
	// lock file or otherwise inserted into a dependency locks object, which
	// are usually (but not necessarily) originally obtained from the
	// provider's origin registry and then cached in the lock file.
	//
	// For a message representing an "available" provider package this
	// describes only the actual package on disk, and so will typically
	// include only the subset of the checksums from the corresponding
	// "selected" package that are relevant to the current platform where
	// Terraform Core is running.
	Hashes []string `protobuf:"bytes,3,rep,name=hashes,proto3" json:"hashes,omitempty"`
}

func (x *ProviderPackage) Reset() {
	*x = ProviderPackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderPackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderPackage) ProtoMessage() {}

func (x *ProviderPackage) ProtoReflect() protoreflect.Message {
	mi := &file_terraform1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderPackage.ProtoReflect.Descriptor instead.
func (*ProviderPackage) Descriptor() ([]byte, []int) {
	return file_terraform1_proto_rawDescGZIP(), []int{0}
}

func (x *ProviderPackage) GetSourceAddr() string {
	if x != nil {
		return x.SourceAddr
	}
	return ""
}

func (x *ProviderPackage) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ProviderPackage) GetHashes() []string {
	if x != nil {
		return x.Hashes
	}
	return nil
}

// A source address in the same form as it would appear in a Terraform
// configuration: a source string combined with an optional version constraint
// string, where the latter is valid only for registry module addresses.
//
// This is not used for "final" source addresses that have already been reduced
// to an exact version selection. For those we just directly encode the string
// representation of the final address, including a version number if necessary.
type SourceAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source   string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Versions string `protobuf:"bytes,2,opt,name=versions,proto3" json:"versions,omitempty"`
}

func (x *SourceAddress) Reset() {
	*x = SourceAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceAddress) ProtoMessage() {}

func (x *SourceAddress) ProtoReflect() protoreflect.Message {
	mi := &file_terraform1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceAddress.ProtoReflect.Descriptor instead.
func (*SourceAddress) Descriptor() ([]byte, []int) {
	return file_terraform1_proto_rawDescGZIP(), []int{1}
}

func (x *SourceAddress) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *SourceAddress) GetVersions() string {
	if x != nil {
		return x.Versions
	}
	return ""
}

type Diagnostic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Severity Diagnostic_Severity `protobuf:"varint,1,opt,name=severity,proto3,enum=terraform1.Diagnostic_Severity" json:"severity,omitempty"`
	Summary  string              `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	Detail   string              `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	Subject  *SourceRange        `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Context  *SourceRange        `protobuf:"bytes,5,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *Diagnostic) Reset() {
	*x = Diagnostic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Diagnostic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Diagnostic) ProtoMessage() {}

func (x *Diagnostic) ProtoReflect() protoreflect.Message {
	mi := &file_terraform1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Diagnostic.ProtoReflect.Descriptor instead.
func (*Diagnostic) Descriptor() ([]byte, []int) {
	return file_terraform1_proto_rawDescGZIP(), []int{2}
}

func (x *Diagnostic) GetSeverity() Diagnostic_Severity {
	if x != nil {
		return x.Severity
	}
	return Diagnostic_INVALID
}

func (x *Diagnostic) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Diagnostic) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *Diagnostic) GetSubject() *SourceRange {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *Diagnostic) GetContext() *SourceRange {
	if x != nil {
		return x.Context
	}
	return nil
}

type SourceRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceAddr string     `protobuf:"bytes,1,opt,name=source_addr,json=sourceAddr,proto3" json:"source_addr,omitempty"`
	Start      *SourcePos `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End        *SourcePos `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *SourceRange) Reset() {
	*x = SourceRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceRange) ProtoMessage() {}

func (x *SourceRange) ProtoReflect() protoreflect.Message {
	mi := &file_terraform1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceRange.ProtoReflect.Descriptor instead.
func (*SourceRange) Descriptor() ([]byte, []int) {
	return file_terraform1_proto_rawDescGZIP(), []int{3}
}

func (x *SourceRange) GetSourceAddr() string {
	if x != nil {
		return x.SourceAddr
	}
	return ""
}

func (x *SourceRange) GetStart() *SourcePos {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *SourceRange) GetEnd() *SourcePos {
	if x != nil {
		return x.End
	}
	return nil
}

type SourcePos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Byte   int64 `protobuf:"varint,1,opt,name=byte,proto3" json:"byte,omitempty"`
	Line   int64 `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	Column int64 `protobuf:"varint,3,opt,name=column,proto3" json:"column,omitempty"`
}

func (x *SourcePos) Reset() {
	*x = SourcePos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourcePos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourcePos) ProtoMessage() {}

func (x *SourcePos) ProtoReflect() protoreflect.Message {
	mi := &file_terraform1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourcePos.ProtoReflect.Descriptor instead.
func (*SourcePos) Descriptor() ([]byte, []int) {
	return file_terraform1_proto_rawDescGZIP(), []int{4}
}

func (x *SourcePos) GetByte() int64 {
	if x != nil {
		return x.Byte
	}
	return 0
}

func (x *SourcePos) GetLine() int64 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *SourcePos) GetColumn() int64 {
	if x != nil {
		return x.Column
	}
	return 0
}

var File_terraform1_proto protoreflect.FileDescriptor

var file_terraform1_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x22, 0x64,
	0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64,
	0x64, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x68, 0x61,
	0x73, 0x68, 0x65, 0x73, 0x22, 0x43, 0x0a, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x92, 0x02, 0x0a, 0x0a, 0x44, 0x69,
	0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x12, 0x3b, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74,
	0x69, 0x63, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61,
	0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65,
	0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x2f, 0x0a,
	0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x22, 0x84,
	0x01, 0x0a, 0x0b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12,
	0x2b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x27, 0x0a, 0x03,
	0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x72, 0x72,
	0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x73,
	0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x4b, 0x0a, 0x09, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50,
	0x6f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x79, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x62, 0x79, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_terraform1_proto_rawDescOnce sync.Once
	file_terraform1_proto_rawDescData = file_terraform1_proto_rawDesc
)

func file_terraform1_proto_rawDescGZIP() []byte {
	file_terraform1_proto_rawDescOnce.Do(func() {
		file_terraform1_proto_rawDescData = protoimpl.X.CompressGZIP(file_terraform1_proto_rawDescData)
	})
	return file_terraform1_proto_rawDescData
}

var file_terraform1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_terraform1_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_terraform1_proto_goTypes = []any{
	(Diagnostic_Severity)(0), // 0: terraform1.Diagnostic.Severity
	(*ProviderPackage)(nil),  // 1: terraform1.ProviderPackage
	(*SourceAddress)(nil),    // 2: terraform1.SourceAddress
	(*Diagnostic)(nil),       // 3: terraform1.Diagnostic
	(*SourceRange)(nil),      // 4: terraform1.SourceRange
	(*SourcePos)(nil),        // 5: terraform1.SourcePos
}
var file_terraform1_proto_depIdxs = []int32{
	0, // 0: terraform1.Diagnostic.severity:type_name -> terraform1.Diagnostic.Severity
	4, // 1: terraform1.Diagnostic.subject:type_name -> terraform1.SourceRange
	4, // 2: terraform1.Diagnostic.context:type_name -> terraform1.SourceRange
	5, // 3: terraform1.SourceRange.start:type_name -> terraform1.SourcePos
	5, // 4: terraform1.SourceRange.end:type_name -> terraform1.SourcePos
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_terraform1_proto_init() }
func file_terraform1_proto_init() {
	if File_terraform1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_terraform1_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ProviderPackage); i {
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
		file_terraform1_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SourceAddress); i {
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
		file_terraform1_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Diagnostic); i {
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
		file_terraform1_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SourceRange); i {
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
		file_terraform1_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SourcePos); i {
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
			RawDescriptor: file_terraform1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_terraform1_proto_goTypes,
		DependencyIndexes: file_terraform1_proto_depIdxs,
		EnumInfos:         file_terraform1_proto_enumTypes,
		MessageInfos:      file_terraform1_proto_msgTypes,
	}.Build()
	File_terraform1_proto = out.File
	file_terraform1_proto_rawDesc = nil
	file_terraform1_proto_goTypes = nil
	file_terraform1_proto_depIdxs = nil
}
