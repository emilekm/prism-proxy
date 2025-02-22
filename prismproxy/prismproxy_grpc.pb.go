// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: prismproxy/prismproxy.proto

package prismproxy

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

const (
	Proxy_GetServerDetails_FullMethodName       = "/prismproxy.Proxy/GetServerDetails"
	Proxy_ServerDetailsUpdates_FullMethodName   = "/prismproxy.Proxy/ServerDetailsUpdates"
	Proxy_GetGameplayDetails_FullMethodName     = "/prismproxy.Proxy/GetGameplayDetails"
	Proxy_GameplayDetailsUpdates_FullMethodName = "/prismproxy.Proxy/GameplayDetailsUpdates"
	Proxy_APIAdmin_FullMethodName               = "/prismproxy.Proxy/APIAdmin"
	Proxy_RACommand_FullMethodName              = "/prismproxy.Proxy/RACommand"
	Proxy_ChatMessages_FullMethodName           = "/prismproxy.Proxy/ChatMessages"
	Proxy_KillMessages_FullMethodName           = "/prismproxy.Proxy/KillMessages"
	Proxy_GetPlayers_FullMethodName             = "/prismproxy.Proxy/GetPlayers"
	Proxy_PlayersUpdates_FullMethodName         = "/prismproxy.Proxy/PlayersUpdates"
	Proxy_PlayerLeaveUpdates_FullMethodName     = "/prismproxy.Proxy/PlayerLeaveUpdates"
	Proxy_UserList_FullMethodName               = "/prismproxy.Proxy/UserList"
	Proxy_AddUser_FullMethodName                = "/prismproxy.Proxy/AddUser"
	Proxy_ChangeUser_FullMethodName             = "/prismproxy.Proxy/ChangeUser"
	Proxy_DeleteUser_FullMethodName             = "/prismproxy.Proxy/DeleteUser"
)

// ProxyClient is the client API for Proxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProxyClient interface {
	// Get the server details
	GetServerDetails(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServerDetails, error)
	// Continous stream of ServerDetails
	ServerDetailsUpdates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Proxy_ServerDetailsUpdatesClient, error)
	// Get gameplay details
	GetGameplayDetails(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GameplayDetails, error)
	// Continous stream of GameplayDetails
	GameplayDetailsUpdates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Proxy_GameplayDetailsUpdatesClient, error)
	// APIAdmin command
	APIAdmin(ctx context.Context, in *APIAdminReq, opts ...grpc.CallOption) (*APIAdminResp, error)
	// RACommand
	RACommand(ctx context.Context, in *RACommandReq, opts ...grpc.CallOption) (*RACommandResp, error)
	// Continous stream of ChatMessages
	ChatMessages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Proxy_ChatMessagesClient, error)
	// Continous stream of KillMessages
	KillMessages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Proxy_KillMessagesClient, error)
	// Get all players
	GetPlayers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Players, error)
	// Continous stream of Players
	PlayersUpdates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Proxy_PlayersUpdatesClient, error)
	// Continous stream of leaving players
	PlayerLeaveUpdates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Proxy_PlayerLeaveUpdatesClient, error)
	// Get all users
	UserList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserListResp, error)
	// Add a user
	AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*UserListResp, error)
	// Change a user
	ChangeUser(ctx context.Context, in *ChangeUserReq, opts ...grpc.CallOption) (*UserListResp, error)
	// Delete a user
	DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*UserListResp, error)
}

type proxyClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyClient(cc grpc.ClientConnInterface) ProxyClient {
	return &proxyClient{cc}
}

func (c *proxyClient) GetServerDetails(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServerDetails, error) {
	out := new(ServerDetails)
	err := c.cc.Invoke(ctx, Proxy_GetServerDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) ServerDetailsUpdates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Proxy_ServerDetailsUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Proxy_ServiceDesc.Streams[0], Proxy_ServerDetailsUpdates_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyServerDetailsUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Proxy_ServerDetailsUpdatesClient interface {
	Recv() (*ServerDetails, error)
	grpc.ClientStream
}

type proxyServerDetailsUpdatesClient struct {
	grpc.ClientStream
}

func (x *proxyServerDetailsUpdatesClient) Recv() (*ServerDetails, error) {
	m := new(ServerDetails)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *proxyClient) GetGameplayDetails(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GameplayDetails, error) {
	out := new(GameplayDetails)
	err := c.cc.Invoke(ctx, Proxy_GetGameplayDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) GameplayDetailsUpdates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Proxy_GameplayDetailsUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Proxy_ServiceDesc.Streams[1], Proxy_GameplayDetailsUpdates_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyGameplayDetailsUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Proxy_GameplayDetailsUpdatesClient interface {
	Recv() (*GameplayDetails, error)
	grpc.ClientStream
}

type proxyGameplayDetailsUpdatesClient struct {
	grpc.ClientStream
}

func (x *proxyGameplayDetailsUpdatesClient) Recv() (*GameplayDetails, error) {
	m := new(GameplayDetails)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *proxyClient) APIAdmin(ctx context.Context, in *APIAdminReq, opts ...grpc.CallOption) (*APIAdminResp, error) {
	out := new(APIAdminResp)
	err := c.cc.Invoke(ctx, Proxy_APIAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) RACommand(ctx context.Context, in *RACommandReq, opts ...grpc.CallOption) (*RACommandResp, error) {
	out := new(RACommandResp)
	err := c.cc.Invoke(ctx, Proxy_RACommand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) ChatMessages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Proxy_ChatMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Proxy_ServiceDesc.Streams[2], Proxy_ChatMessages_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyChatMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Proxy_ChatMessagesClient interface {
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type proxyChatMessagesClient struct {
	grpc.ClientStream
}

func (x *proxyChatMessagesClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *proxyClient) KillMessages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Proxy_KillMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Proxy_ServiceDesc.Streams[3], Proxy_KillMessages_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyKillMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Proxy_KillMessagesClient interface {
	Recv() (*KillMessage, error)
	grpc.ClientStream
}

type proxyKillMessagesClient struct {
	grpc.ClientStream
}

func (x *proxyKillMessagesClient) Recv() (*KillMessage, error) {
	m := new(KillMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *proxyClient) GetPlayers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Players, error) {
	out := new(Players)
	err := c.cc.Invoke(ctx, Proxy_GetPlayers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) PlayersUpdates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Proxy_PlayersUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Proxy_ServiceDesc.Streams[4], Proxy_PlayersUpdates_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyPlayersUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Proxy_PlayersUpdatesClient interface {
	Recv() (*PlayersUpdate, error)
	grpc.ClientStream
}

type proxyPlayersUpdatesClient struct {
	grpc.ClientStream
}

func (x *proxyPlayersUpdatesClient) Recv() (*PlayersUpdate, error) {
	m := new(PlayersUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *proxyClient) PlayerLeaveUpdates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Proxy_PlayerLeaveUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Proxy_ServiceDesc.Streams[5], Proxy_PlayerLeaveUpdates_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyPlayerLeaveUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Proxy_PlayerLeaveUpdatesClient interface {
	Recv() (*PlayerLeave, error)
	grpc.ClientStream
}

type proxyPlayerLeaveUpdatesClient struct {
	grpc.ClientStream
}

func (x *proxyPlayerLeaveUpdatesClient) Recv() (*PlayerLeave, error) {
	m := new(PlayerLeave)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *proxyClient) UserList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserListResp, error) {
	out := new(UserListResp)
	err := c.cc.Invoke(ctx, Proxy_UserList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*UserListResp, error) {
	out := new(UserListResp)
	err := c.cc.Invoke(ctx, Proxy_AddUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) ChangeUser(ctx context.Context, in *ChangeUserReq, opts ...grpc.CallOption) (*UserListResp, error) {
	out := new(UserListResp)
	err := c.cc.Invoke(ctx, Proxy_ChangeUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*UserListResp, error) {
	out := new(UserListResp)
	err := c.cc.Invoke(ctx, Proxy_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyServer is the server API for Proxy service.
// All implementations must embed UnimplementedProxyServer
// for forward compatibility
type ProxyServer interface {
	// Get the server details
	GetServerDetails(context.Context, *Empty) (*ServerDetails, error)
	// Continous stream of ServerDetails
	ServerDetailsUpdates(*Empty, Proxy_ServerDetailsUpdatesServer) error
	// Get gameplay details
	GetGameplayDetails(context.Context, *Empty) (*GameplayDetails, error)
	// Continous stream of GameplayDetails
	GameplayDetailsUpdates(*Empty, Proxy_GameplayDetailsUpdatesServer) error
	// APIAdmin command
	APIAdmin(context.Context, *APIAdminReq) (*APIAdminResp, error)
	// RACommand
	RACommand(context.Context, *RACommandReq) (*RACommandResp, error)
	// Continous stream of ChatMessages
	ChatMessages(*Empty, Proxy_ChatMessagesServer) error
	// Continous stream of KillMessages
	KillMessages(*Empty, Proxy_KillMessagesServer) error
	// Get all players
	GetPlayers(context.Context, *Empty) (*Players, error)
	// Continous stream of Players
	PlayersUpdates(*Empty, Proxy_PlayersUpdatesServer) error
	// Continous stream of leaving players
	PlayerLeaveUpdates(*Empty, Proxy_PlayerLeaveUpdatesServer) error
	// Get all users
	UserList(context.Context, *Empty) (*UserListResp, error)
	// Add a user
	AddUser(context.Context, *AddUserReq) (*UserListResp, error)
	// Change a user
	ChangeUser(context.Context, *ChangeUserReq) (*UserListResp, error)
	// Delete a user
	DeleteUser(context.Context, *DeleteUserReq) (*UserListResp, error)
	mustEmbedUnimplementedProxyServer()
}

// UnimplementedProxyServer must be embedded to have forward compatible implementations.
type UnimplementedProxyServer struct {
}

func (UnimplementedProxyServer) GetServerDetails(context.Context, *Empty) (*ServerDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerDetails not implemented")
}
func (UnimplementedProxyServer) ServerDetailsUpdates(*Empty, Proxy_ServerDetailsUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerDetailsUpdates not implemented")
}
func (UnimplementedProxyServer) GetGameplayDetails(context.Context, *Empty) (*GameplayDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameplayDetails not implemented")
}
func (UnimplementedProxyServer) GameplayDetailsUpdates(*Empty, Proxy_GameplayDetailsUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method GameplayDetailsUpdates not implemented")
}
func (UnimplementedProxyServer) APIAdmin(context.Context, *APIAdminReq) (*APIAdminResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method APIAdmin not implemented")
}
func (UnimplementedProxyServer) RACommand(context.Context, *RACommandReq) (*RACommandResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RACommand not implemented")
}
func (UnimplementedProxyServer) ChatMessages(*Empty, Proxy_ChatMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method ChatMessages not implemented")
}
func (UnimplementedProxyServer) KillMessages(*Empty, Proxy_KillMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method KillMessages not implemented")
}
func (UnimplementedProxyServer) GetPlayers(context.Context, *Empty) (*Players, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayers not implemented")
}
func (UnimplementedProxyServer) PlayersUpdates(*Empty, Proxy_PlayersUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method PlayersUpdates not implemented")
}
func (UnimplementedProxyServer) PlayerLeaveUpdates(*Empty, Proxy_PlayerLeaveUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method PlayerLeaveUpdates not implemented")
}
func (UnimplementedProxyServer) UserList(context.Context, *Empty) (*UserListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (UnimplementedProxyServer) AddUser(context.Context, *AddUserReq) (*UserListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedProxyServer) ChangeUser(context.Context, *ChangeUserReq) (*UserListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUser not implemented")
}
func (UnimplementedProxyServer) DeleteUser(context.Context, *DeleteUserReq) (*UserListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedProxyServer) mustEmbedUnimplementedProxyServer() {}

// UnsafeProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProxyServer will
// result in compilation errors.
type UnsafeProxyServer interface {
	mustEmbedUnimplementedProxyServer()
}

func RegisterProxyServer(s grpc.ServiceRegistrar, srv ProxyServer) {
	s.RegisterService(&Proxy_ServiceDesc, srv)
}

func _Proxy_GetServerDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).GetServerDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Proxy_GetServerDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).GetServerDetails(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_ServerDetailsUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProxyServer).ServerDetailsUpdates(m, &proxyServerDetailsUpdatesServer{stream})
}

type Proxy_ServerDetailsUpdatesServer interface {
	Send(*ServerDetails) error
	grpc.ServerStream
}

type proxyServerDetailsUpdatesServer struct {
	grpc.ServerStream
}

func (x *proxyServerDetailsUpdatesServer) Send(m *ServerDetails) error {
	return x.ServerStream.SendMsg(m)
}

func _Proxy_GetGameplayDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).GetGameplayDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Proxy_GetGameplayDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).GetGameplayDetails(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_GameplayDetailsUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProxyServer).GameplayDetailsUpdates(m, &proxyGameplayDetailsUpdatesServer{stream})
}

type Proxy_GameplayDetailsUpdatesServer interface {
	Send(*GameplayDetails) error
	grpc.ServerStream
}

type proxyGameplayDetailsUpdatesServer struct {
	grpc.ServerStream
}

func (x *proxyGameplayDetailsUpdatesServer) Send(m *GameplayDetails) error {
	return x.ServerStream.SendMsg(m)
}

func _Proxy_APIAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIAdminReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).APIAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Proxy_APIAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).APIAdmin(ctx, req.(*APIAdminReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_RACommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RACommandReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).RACommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Proxy_RACommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).RACommand(ctx, req.(*RACommandReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_ChatMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProxyServer).ChatMessages(m, &proxyChatMessagesServer{stream})
}

type Proxy_ChatMessagesServer interface {
	Send(*ChatMessage) error
	grpc.ServerStream
}

type proxyChatMessagesServer struct {
	grpc.ServerStream
}

func (x *proxyChatMessagesServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Proxy_KillMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProxyServer).KillMessages(m, &proxyKillMessagesServer{stream})
}

type Proxy_KillMessagesServer interface {
	Send(*KillMessage) error
	grpc.ServerStream
}

type proxyKillMessagesServer struct {
	grpc.ServerStream
}

func (x *proxyKillMessagesServer) Send(m *KillMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Proxy_GetPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).GetPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Proxy_GetPlayers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).GetPlayers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_PlayersUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProxyServer).PlayersUpdates(m, &proxyPlayersUpdatesServer{stream})
}

type Proxy_PlayersUpdatesServer interface {
	Send(*PlayersUpdate) error
	grpc.ServerStream
}

type proxyPlayersUpdatesServer struct {
	grpc.ServerStream
}

func (x *proxyPlayersUpdatesServer) Send(m *PlayersUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func _Proxy_PlayerLeaveUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProxyServer).PlayerLeaveUpdates(m, &proxyPlayerLeaveUpdatesServer{stream})
}

type Proxy_PlayerLeaveUpdatesServer interface {
	Send(*PlayerLeave) error
	grpc.ServerStream
}

type proxyPlayerLeaveUpdatesServer struct {
	grpc.ServerStream
}

func (x *proxyPlayerLeaveUpdatesServer) Send(m *PlayerLeave) error {
	return x.ServerStream.SendMsg(m)
}

func _Proxy_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Proxy_UserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).UserList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Proxy_AddUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).AddUser(ctx, req.(*AddUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_ChangeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).ChangeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Proxy_ChangeUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).ChangeUser(ctx, req.(*ChangeUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Proxy_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).DeleteUser(ctx, req.(*DeleteUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Proxy_ServiceDesc is the grpc.ServiceDesc for Proxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Proxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prismproxy.Proxy",
	HandlerType: (*ProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServerDetails",
			Handler:    _Proxy_GetServerDetails_Handler,
		},
		{
			MethodName: "GetGameplayDetails",
			Handler:    _Proxy_GetGameplayDetails_Handler,
		},
		{
			MethodName: "APIAdmin",
			Handler:    _Proxy_APIAdmin_Handler,
		},
		{
			MethodName: "RACommand",
			Handler:    _Proxy_RACommand_Handler,
		},
		{
			MethodName: "GetPlayers",
			Handler:    _Proxy_GetPlayers_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _Proxy_UserList_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _Proxy_AddUser_Handler,
		},
		{
			MethodName: "ChangeUser",
			Handler:    _Proxy_ChangeUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Proxy_DeleteUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerDetailsUpdates",
			Handler:       _Proxy_ServerDetailsUpdates_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GameplayDetailsUpdates",
			Handler:       _Proxy_GameplayDetailsUpdates_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ChatMessages",
			Handler:       _Proxy_ChatMessages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "KillMessages",
			Handler:       _Proxy_KillMessages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PlayersUpdates",
			Handler:       _Proxy_PlayersUpdates_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PlayerLeaveUpdates",
			Handler:       _Proxy_PlayerLeaveUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "prismproxy/prismproxy.proto",
}
