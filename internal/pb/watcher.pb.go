// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: internal/pb/watcher.proto

package pb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type StartWatchingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileAlias string `protobuf:"bytes,1,opt,name=file_alias,json=fileAlias,proto3" json:"file_alias,omitempty"`
	FilePath  string `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
}

func (x *StartWatchingRequest) Reset() {
	*x = StartWatchingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_watcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartWatchingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartWatchingRequest) ProtoMessage() {}

func (x *StartWatchingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_watcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartWatchingRequest.ProtoReflect.Descriptor instead.
func (*StartWatchingRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_watcher_proto_rawDescGZIP(), []int{0}
}

func (x *StartWatchingRequest) GetFileAlias() string {
	if x != nil {
		return x.FileAlias
	}
	return ""
}

func (x *StartWatchingRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type GetChangesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileAlias string `protobuf:"bytes,1,opt,name=file_alias,json=fileAlias,proto3" json:"file_alias,omitempty"`
}

func (x *GetChangesRequest) Reset() {
	*x = GetChangesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_watcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChangesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChangesRequest) ProtoMessage() {}

func (x *GetChangesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_watcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChangesRequest.ProtoReflect.Descriptor instead.
func (*GetChangesRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_watcher_proto_rawDescGZIP(), []int{1}
}

func (x *GetChangesRequest) GetFileAlias() string {
	if x != nil {
		return x.FileAlias
	}
	return ""
}

type Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileAlias string               `protobuf:"bytes,1,opt,name=file_alias,json=fileAlias,proto3" json:"file_alias,omitempty"`
	EventName string               `protobuf:"bytes,2,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
	Timestamp *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Change) Reset() {
	*x = Change{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_watcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Change) ProtoMessage() {}

func (x *Change) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_watcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Change.ProtoReflect.Descriptor instead.
func (*Change) Descriptor() ([]byte, []int) {
	return file_internal_pb_watcher_proto_rawDescGZIP(), []int{2}
}

func (x *Change) GetFileAlias() string {
	if x != nil {
		return x.FileAlias
	}
	return ""
}

func (x *Change) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *Change) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type GetChangesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Changes []*Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *GetChangesResponse) Reset() {
	*x = GetChangesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_watcher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChangesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChangesResponse) ProtoMessage() {}

func (x *GetChangesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_watcher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChangesResponse.ProtoReflect.Descriptor instead.
func (*GetChangesResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_watcher_proto_rawDescGZIP(), []int{3}
}

func (x *GetChangesResponse) GetChanges() []*Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

var File_internal_pb_watcher_proto protoreflect.FileDescriptor

var file_internal_pb_watcher_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a,
	0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x41,
	0x6c, 0x69, 0x61, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x61,
	0x6c, 0x69, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x41, 0x6c, 0x69, 0x61, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x3a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x32, 0x90, 0x01, 0x0a, 0x0e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pb_watcher_proto_rawDescOnce sync.Once
	file_internal_pb_watcher_proto_rawDescData = file_internal_pb_watcher_proto_rawDesc
)

func file_internal_pb_watcher_proto_rawDescGZIP() []byte {
	file_internal_pb_watcher_proto_rawDescOnce.Do(func() {
		file_internal_pb_watcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pb_watcher_proto_rawDescData)
	})
	return file_internal_pb_watcher_proto_rawDescData
}

var file_internal_pb_watcher_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_pb_watcher_proto_goTypes = []interface{}{
	(*StartWatchingRequest)(nil), // 0: pb.StartWatchingRequest
	(*GetChangesRequest)(nil),    // 1: pb.GetChangesRequest
	(*Change)(nil),               // 2: pb.Change
	(*GetChangesResponse)(nil),   // 3: pb.GetChangesResponse
	(*timestamp.Timestamp)(nil),  // 4: google.protobuf.Timestamp
	(*empty.Empty)(nil),          // 5: google.protobuf.Empty
}
var file_internal_pb_watcher_proto_depIdxs = []int32{
	4, // 0: pb.Change.timestamp:type_name -> google.protobuf.Timestamp
	2, // 1: pb.GetChangesResponse.changes:type_name -> pb.Change
	0, // 2: pb.WatcherService.StartWatching:input_type -> pb.StartWatchingRequest
	1, // 3: pb.WatcherService.GetChanges:input_type -> pb.GetChangesRequest
	5, // 4: pb.WatcherService.StartWatching:output_type -> google.protobuf.Empty
	3, // 5: pb.WatcherService.GetChanges:output_type -> pb.GetChangesResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_pb_watcher_proto_init() }
func file_internal_pb_watcher_proto_init() {
	if File_internal_pb_watcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pb_watcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartWatchingRequest); i {
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
		file_internal_pb_watcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChangesRequest); i {
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
		file_internal_pb_watcher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Change); i {
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
		file_internal_pb_watcher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChangesResponse); i {
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
			RawDescriptor: file_internal_pb_watcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_pb_watcher_proto_goTypes,
		DependencyIndexes: file_internal_pb_watcher_proto_depIdxs,
		MessageInfos:      file_internal_pb_watcher_proto_msgTypes,
	}.Build()
	File_internal_pb_watcher_proto = out.File
	file_internal_pb_watcher_proto_rawDesc = nil
	file_internal_pb_watcher_proto_goTypes = nil
	file_internal_pb_watcher_proto_depIdxs = nil
}
