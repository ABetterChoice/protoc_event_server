// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: event_server.proto

package protoc_event_server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EventServerClientV2 is the client API for EventServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventServerClientV2 interface {
	LogExposureGroup(ctx context.Context, in *ExposureGroup, opts ...grpc.CallOption) (*CommonResp, error)
	LogEventGroup(ctx context.Context, in *EventGroup, opts ...grpc.CallOption) (*CommonResp, error)
	LogMonitorEventGroup(ctx context.Context, in *MonitorEventGroup, opts ...grpc.CallOption) (*CommonResp, error)
}

type eventServerClientV2 struct {
	cc grpc.ClientConnInterface
}

func NewEventServerClientV2(cc grpc.ClientConnInterface) EventServerClientV2 {
	return &eventServerClientV2{cc}
}

func (c *eventServerClientV2) LogExposureGroup(ctx context.Context, in *ExposureGroup, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, "/opensource.tab.cache_server.EventServer/LogExposureGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServerClientV2) LogEventGroup(ctx context.Context, in *EventGroup, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, "/opensource.tab.cache_server.EventServer/LogEventGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServerClientV2) LogMonitorEventGroup(ctx context.Context, in *MonitorEventGroup, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, "/opensource.tab.cache_server.EventServer/LogMonitorEventGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServerServerV2 is the server API for EventServer service.
// All implementations should embed UnimplementedEventServerServerV2
// for forward compatibility
type EventServerServerV2 interface {
	LogExposureGroup(context.Context, *ExposureGroup, *CommonResp) error
	LogEventGroup(context.Context, *EventGroup, *CommonResp) error
	LogMonitorEventGroup(context.Context, *MonitorEventGroup, *CommonResp) error
}

// UnimplementedEventServerServerV2 should be embedded to have forward compatible implementations.
type UnimplementedEventServerServerV2 struct {
}

func (UnimplementedEventServerServerV2) LogExposureGroup(context.Context, *ExposureGroup, *CommonResp) error {
	return status.Errorf(codes.Unimplemented, "method LogExposureGroup not implemented")
}
func (UnimplementedEventServerServerV2) LogEventGroup(context.Context, *EventGroup, *CommonResp) error {
	return status.Errorf(codes.Unimplemented, "method LogEventGroup not implemented")
}
func (UnimplementedEventServerServerV2) LogMonitorEventGroup(context.Context, *MonitorEventGroup, *CommonResp) error {
	return status.Errorf(codes.Unimplemented, "method LogMonitorEventGroup not implemented")
}

// UnsafeEventServerServerV2 may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServerServerV2 will
// result in compilation errors.
type UnsafeEventServerServerV2 interface {
	mustEmbedUnimplementedEventServerServerV2()
}

func RegisterEventServerServerV2(s grpc.ServiceRegistrar, srv EventServerServerV2) {
	s.RegisterService(&EventServer_ServiceDescV2, srv)
}

func _EventServer_LogExposureGroup_HandlerV2(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExposureGroup)
	out := new(CommonResp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		err := srv.(EventServerServerV2).LogExposureGroup(ctx, in, out)
		return out, err
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensource.tab.cache_server.EventServer/LogExposureGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		rsp := new(CommonResp)
		err := srv.(EventServerServerV2).LogExposureGroup(ctx, req.(*ExposureGroup), rsp)
		return rsp, err
	}
	return interceptor(ctx, in, info, handler)
}

func _EventServer_LogEventGroup_HandlerV2(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventGroup)
	out := new(CommonResp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		err := srv.(EventServerServerV2).LogEventGroup(ctx, in, out)
		return out, err
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensource.tab.cache_server.EventServer/LogEventGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		rsp := new(CommonResp)
		err := srv.(EventServerServerV2).LogEventGroup(ctx, req.(*EventGroup), rsp)
		return rsp, err
	}
	return interceptor(ctx, in, info, handler)
}

func _EventServer_LogMonitorEventGroup_HandlerV2(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonitorEventGroup)
	out := new(CommonResp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		err := srv.(EventServerServerV2).LogMonitorEventGroup(ctx, in, out)
		return out, err
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensource.tab.cache_server.EventServer/LogMonitorEventGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		rsp := new(CommonResp)
		err := srv.(EventServerServerV2).LogMonitorEventGroup(ctx, req.(*MonitorEventGroup), rsp)
		return rsp, err
	}
	return interceptor(ctx, in, info, handler)
}

// EventServer_ServiceDescV2 is the grpc.ServiceDesc for EventServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventServer_ServiceDescV2 = grpc.ServiceDesc{
	ServiceName: "opensource.tab.cache_server.EventServer",
	HandlerType: (*EventServerServerV2)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogExposureGroup",
			Handler:    _EventServer_LogExposureGroup_HandlerV2,
		},
		{
			MethodName: "LogEventGroup",
			Handler:    _EventServer_LogEventGroup_HandlerV2,
		},
		{
			MethodName: "LogMonitorEventGroup",
			Handler:    _EventServer_LogMonitorEventGroup_HandlerV2,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event_server.proto",
}
