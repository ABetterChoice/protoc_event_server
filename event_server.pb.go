// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: event_server.proto

package protoc_event_server

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ExposureType int32

const (
	ExposureType_EXPOSURE_TYPE_UNKNOWN   ExposureType = 0
	ExposureType_EXPOSURE_TYPE_AUTOMATIC ExposureType = 1 // sdk 自动记录曝光
	ExposureType_EXPOSURE_TYPE_MANUAL    ExposureType = 2 // 手动记录曝光，由开发者调用记录曝光的 api
)

// Enum value maps for ExposureType.
var (
	ExposureType_name = map[int32]string{
		0: "EXPOSURE_TYPE_UNKNOWN",
		1: "EXPOSURE_TYPE_AUTOMATIC",
		2: "EXPOSURE_TYPE_MANUAL",
	}
	ExposureType_value = map[string]int32{
		"EXPOSURE_TYPE_UNKNOWN":   0,
		"EXPOSURE_TYPE_AUTOMATIC": 1,
		"EXPOSURE_TYPE_MANUAL":    2,
	}
)

func (x ExposureType) Enum() *ExposureType {
	p := new(ExposureType)
	*p = x
	return p
}

func (x ExposureType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExposureType) Descriptor() protoreflect.EnumDescriptor {
	return file_event_server_proto_enumTypes[0].Descriptor()
}

func (ExposureType) Type() protoreflect.EnumType {
	return &file_event_server_proto_enumTypes[0]
}

func (x ExposureType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExposureType.Descriptor instead.
func (ExposureType) EnumDescriptor() ([]byte, []int) {
	return file_event_server_proto_rawDescGZIP(), []int{0}
}

type Code int32

const (
	Code_CODE_SUCCESS            Code = 0    // 正常返回
	Code_CODE_NO_PERMISSION      Code = 1001 // 无权限
	Code_CODE_TRAFFIC_LIMIT      Code = 1002 // 限流返回
	Code_CODE_INVALID_PROJECT_ID Code = 1003 // 入参 projectID 出错
	Code_CODE_SERVER_ERR         Code = 1004 // 服务器处理异常
	Code_CODE_INVALID_PARAM      Code = 1005 // 非法参数
	Code_CODE_SAME_VERSION       Code = 2001 // 版本未更新
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:    "CODE_SUCCESS",
		1001: "CODE_NO_PERMISSION",
		1002: "CODE_TRAFFIC_LIMIT",
		1003: "CODE_INVALID_PROJECT_ID",
		1004: "CODE_SERVER_ERR",
		1005: "CODE_INVALID_PARAM",
		2001: "CODE_SAME_VERSION",
	}
	Code_value = map[string]int32{
		"CODE_SUCCESS":            0,
		"CODE_NO_PERMISSION":      1001,
		"CODE_TRAFFIC_LIMIT":      1002,
		"CODE_INVALID_PROJECT_ID": 1003,
		"CODE_SERVER_ERR":         1004,
		"CODE_INVALID_PARAM":      1005,
		"CODE_SAME_VERSION":       2001,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_event_server_proto_enumTypes[1].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_event_server_proto_enumTypes[1]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_event_server_proto_rawDescGZIP(), []int{1}
}

// 曝光上报
type Exposure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 必填
	UnitId    string `protobuf:"bytes,1,opt,name=unit_id,json=unitId,proto3" json:"unit_id,omitempty"`          // unitID，同一个 UnitID 稳定命中同一个实验版本
	GroupId   int64  `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`      // 命中的实验版本 ID
	ProjectId string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"` // 项目 ID
	Time      int64  `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`                           // 时间戳，精确到秒
	// optional
	LayerKey     string            `protobuf:"bytes,5,opt,name=layer_key,json=layerKey,proto3" json:"layer_key,omitempty"`                                                                                             // 层 key
	ExpKey       string            `protobuf:"bytes,6,opt,name=exp_key,json=expKey,proto3" json:"exp_key,omitempty"`                                                                                                   // 实验 key
	UnitType     string            `protobuf:"bytes,7,opt,name=unit_type,json=unitType,proto3" json:"unit_type,omitempty"`                                                                                             // unitID 类型，预留
	ClusterId    string            `protobuf:"bytes,8,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`                                                                                          // clusterID 网络实验群 ID，高阶实验预留
	SdkType      string            `protobuf:"bytes,9,opt,name=sdk_type,json=sdkType,proto3" json:"sdk_type,omitempty"`                                                                                                // sdk 类型 golang、cpp、java
	SdkVersion   string            `protobuf:"bytes,10,opt,name=sdk_version,json=sdkVersion,proto3" json:"sdk_version,omitempty"`                                                                                      // sdk 版本
	ExposureType ExposureType      `protobuf:"varint,11,opt,name=exposure_type,json=exposureType,proto3,enum=opensource.tab.cache_server.ExposureType" json:"exposure_type,omitempty"`                                 // 曝光方式
	Device       *Device           `protobuf:"bytes,12,opt,name=device,proto3" json:"device,omitempty"`                                                                                                                // 终端设备元数据，可选，也可通过 http header 传入
	ExtraData    map[string]string `protobuf:"bytes,15,rep,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 拓展字段
}

func (x *Exposure) Reset() {
	*x = Exposure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exposure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exposure) ProtoMessage() {}

func (x *Exposure) ProtoReflect() protoreflect.Message {
	mi := &file_event_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exposure.ProtoReflect.Descriptor instead.
func (*Exposure) Descriptor() ([]byte, []int) {
	return file_event_server_proto_rawDescGZIP(), []int{0}
}

func (x *Exposure) GetUnitId() string {
	if x != nil {
		return x.UnitId
	}
	return ""
}

func (x *Exposure) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *Exposure) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Exposure) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Exposure) GetLayerKey() string {
	if x != nil {
		return x.LayerKey
	}
	return ""
}

func (x *Exposure) GetExpKey() string {
	if x != nil {
		return x.ExpKey
	}
	return ""
}

func (x *Exposure) GetUnitType() string {
	if x != nil {
		return x.UnitType
	}
	return ""
}

func (x *Exposure) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *Exposure) GetSdkType() string {
	if x != nil {
		return x.SdkType
	}
	return ""
}

func (x *Exposure) GetSdkVersion() string {
	if x != nil {
		return x.SdkVersion
	}
	return ""
}

func (x *Exposure) GetExposureType() ExposureType {
	if x != nil {
		return x.ExposureType
	}
	return ExposureType_EXPOSURE_TYPE_UNKNOWN
}

func (x *Exposure) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *Exposure) GetExtraData() map[string]string {
	if x != nil {
		return x.ExtraData
	}
	return nil
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectVersion string `protobuf:"bytes,1,opt,name=project_version,json=projectVersion,proto3" json:"project_version,omitempty"` // 业务版本
	Platform       string `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`                                   // 平台类型 Windows/MacOS/iPhone/Android/iPod/iPad/Linux/UNIX/Unknown
	OsModel        string `protobuf:"bytes,3,opt,name=os_model,json=osModel,proto3" json:"os_model,omitempty"`                      // 声明浏览器操作系统或者设备类型，QQBrowser/UCBrowser
	OsVersion      string `protobuf:"bytes,4,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`                // 系统平台版本
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_event_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_event_server_proto_rawDescGZIP(), []int{1}
}

func (x *Device) GetProjectVersion() string {
	if x != nil {
		return x.ProjectVersion
	}
	return ""
}

func (x *Device) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *Device) GetOsModel() string {
	if x != nil {
		return x.OsModel
	}
	return ""
}

func (x *Device) GetOsVersion() string {
	if x != nil {
		return x.OsVersion
	}
	return ""
}

// 通用事件上报
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 必填
	EventName string            `protobuf:"bytes,1,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`                                                                       // 事件名
	ProjectId string            `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`                                                                       // 项目 ID
	Time      int64             `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`                                                                                                 // 时间戳，精确到秒
	Metadata  map[string]string `protobuf:"bytes,10,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 一条日志里多个 kv 组合
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_event_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_event_server_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *Event) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Event) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Event) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// 批量上报曝光
type ExposureGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exposures []*Exposure `protobuf:"bytes,1,rep,name=exposures,proto3" json:"exposures,omitempty"` // 多条曝光记录
}

func (x *ExposureGroup) Reset() {
	*x = ExposureGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExposureGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExposureGroup) ProtoMessage() {}

func (x *ExposureGroup) ProtoReflect() protoreflect.Message {
	mi := &file_event_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExposureGroup.ProtoReflect.Descriptor instead.
func (*ExposureGroup) Descriptor() ([]byte, []int) {
	return file_event_server_proto_rawDescGZIP(), []int{3}
}

func (x *ExposureGroup) GetExposures() []*Exposure {
	if x != nil {
		return x.Exposures
	}
	return nil
}

// 批量上报事件
type EventGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"` // 多条事件记录
}

func (x *EventGroup) Reset() {
	*x = EventGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventGroup) ProtoMessage() {}

func (x *EventGroup) ProtoReflect() protoreflect.Message {
	mi := &file_event_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventGroup.ProtoReflect.Descriptor instead.
func (*EventGroup) Descriptor() ([]byte, []int) {
	return file_event_server_proto_rawDescGZIP(), []int{4}
}

func (x *EventGroup) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type CommonResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code   `protobuf:"varint,1,opt,name=code,proto3,enum=opensource.tab.cache_server.Code" json:"code,omitempty"` // 错误码
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`                                  // 详细信息描述
}

func (x *CommonResp) Reset() {
	*x = CommonResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResp) ProtoMessage() {}

func (x *CommonResp) ProtoReflect() protoreflect.Message {
	mi := &file_event_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResp.ProtoReflect.Descriptor instead.
func (*CommonResp) Descriptor() ([]byte, []int) {
	return file_event_server_proto_rawDescGZIP(), []int{5}
}

func (x *CommonResp) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_CODE_SUCCESS
}

func (x *CommonResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_event_server_proto protoreflect.FileDescriptor

var file_event_server_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x74, 0x61, 0x62, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbf, 0x04, 0x0a, 0x08, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x6e, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x6e, 0x69, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x12, 0x17, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x78, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e,
	0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x64, 0x6b, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x64, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x64, 0x6b, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x64, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x4e, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x74, 0x61, 0x62, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x3b, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x74,
	0x61, 0x62, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53,
	0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0f, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x74, 0x61, 0x62, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x3c, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x87, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a,
	0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xe4, 0x01, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x74, 0x61, 0x62, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x54, 0x0a, 0x0d, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x43, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x74, 0x61, 0x62, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x52, 0x09, 0x65,
	0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x73, 0x22, 0x48, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3a, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x74, 0x61, 0x62, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x5d, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x35, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x74, 0x61, 0x62, 0x2e,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2a, 0x60, 0x0a, 0x0c, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x58, 0x50, 0x4f, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17,
	0x45, 0x58, 0x50, 0x4f, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x55,
	0x54, 0x4f, 0x4d, 0x41, 0x54, 0x49, 0x43, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x58, 0x50,
	0x4f, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x4e, 0x55, 0x41,
	0x4c, 0x10, 0x02, 0x2a, 0xaf, 0x01, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x0c,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x17,
	0x0a, 0x12, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0xe9, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0xea, 0x07,
	0x12, 0x1c, 0x0a, 0x17, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x49, 0x44, 0x10, 0xeb, 0x07, 0x12, 0x14,
	0x0a, 0x0f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52,
	0x52, 0x10, 0xec, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x10, 0xed, 0x07, 0x12, 0x16, 0x0a,
	0x11, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x41, 0x4d, 0x45, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49,
	0x4f, 0x4e, 0x10, 0xd1, 0x0f, 0x32, 0x9e, 0x02, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x8a, 0x01, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x45, 0x78, 0x70,
	0x6f, 0x73, 0x75, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2a, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x74, 0x61, 0x62, 0x2e, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x27, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x74, 0x61, 0x62, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70,
	0x6f, 0x73, 0x75, 0x72, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x6c, 0x6f, 0x67, 0x3a,
	0x01, 0x2a, 0x12, 0x81, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x27, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x74, 0x61, 0x62, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x27, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x74, 0x61, 0x62, 0x2e, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f,
	0x6c, 0x6f, 0x67, 0x3a, 0x01, 0x2a, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x6f, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e,
	0x74, 0x5f, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_server_proto_rawDescOnce sync.Once
	file_event_server_proto_rawDescData = file_event_server_proto_rawDesc
)

func file_event_server_proto_rawDescGZIP() []byte {
	file_event_server_proto_rawDescOnce.Do(func() {
		file_event_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_server_proto_rawDescData)
	})
	return file_event_server_proto_rawDescData
}

var file_event_server_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_event_server_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_event_server_proto_goTypes = []interface{}{
	(ExposureType)(0),     // 0: opensource.tab.cache_server.ExposureType
	(Code)(0),             // 1: opensource.tab.cache_server.Code
	(*Exposure)(nil),      // 2: opensource.tab.cache_server.Exposure
	(*Device)(nil),        // 3: opensource.tab.cache_server.Device
	(*Event)(nil),         // 4: opensource.tab.cache_server.Event
	(*ExposureGroup)(nil), // 5: opensource.tab.cache_server.ExposureGroup
	(*EventGroup)(nil),    // 6: opensource.tab.cache_server.EventGroup
	(*CommonResp)(nil),    // 7: opensource.tab.cache_server.CommonResp
	nil,                   // 8: opensource.tab.cache_server.Exposure.ExtraDataEntry
	nil,                   // 9: opensource.tab.cache_server.Event.MetadataEntry
}
var file_event_server_proto_depIdxs = []int32{
	0, // 0: opensource.tab.cache_server.Exposure.exposure_type:type_name -> opensource.tab.cache_server.ExposureType
	3, // 1: opensource.tab.cache_server.Exposure.device:type_name -> opensource.tab.cache_server.Device
	8, // 2: opensource.tab.cache_server.Exposure.extra_data:type_name -> opensource.tab.cache_server.Exposure.ExtraDataEntry
	9, // 3: opensource.tab.cache_server.Event.metadata:type_name -> opensource.tab.cache_server.Event.MetadataEntry
	2, // 4: opensource.tab.cache_server.ExposureGroup.exposures:type_name -> opensource.tab.cache_server.Exposure
	4, // 5: opensource.tab.cache_server.EventGroup.events:type_name -> opensource.tab.cache_server.Event
	1, // 6: opensource.tab.cache_server.CommonResp.code:type_name -> opensource.tab.cache_server.Code
	5, // 7: opensource.tab.cache_server.EventServer.LogExposureGroup:input_type -> opensource.tab.cache_server.ExposureGroup
	6, // 8: opensource.tab.cache_server.EventServer.LogEventGroup:input_type -> opensource.tab.cache_server.EventGroup
	7, // 9: opensource.tab.cache_server.EventServer.LogExposureGroup:output_type -> opensource.tab.cache_server.CommonResp
	7, // 10: opensource.tab.cache_server.EventServer.LogEventGroup:output_type -> opensource.tab.cache_server.CommonResp
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_event_server_proto_init() }
func file_event_server_proto_init() {
	if File_event_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Exposure); i {
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
		file_event_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_event_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_event_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExposureGroup); i {
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
		file_event_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventGroup); i {
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
		file_event_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResp); i {
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
			RawDescriptor: file_event_server_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_event_server_proto_goTypes,
		DependencyIndexes: file_event_server_proto_depIdxs,
		EnumInfos:         file_event_server_proto_enumTypes,
		MessageInfos:      file_event_server_proto_msgTypes,
	}.Build()
	File_event_server_proto = out.File
	file_event_server_proto_rawDesc = nil
	file_event_server_proto_goTypes = nil
	file_event_server_proto_depIdxs = nil
}
