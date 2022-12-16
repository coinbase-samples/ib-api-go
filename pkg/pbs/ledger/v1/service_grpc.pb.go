// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pkg/pbs/ledger/v1/service.proto

package v1

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

// LedgerClient is the client API for Ledger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LedgerClient interface {
	InitializeAccount(ctx context.Context, in *InitializeAccountRequest, opts ...grpc.CallOption) (*InitializeAccountResponse, error)
	GetAccounts(ctx context.Context, in *GetAccountsRequest, opts ...grpc.CallOption) (*GetAccountsResponse, error)
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
	PartialReleaseHold(ctx context.Context, in *PartialReleaseHoldRequest, opts ...grpc.CallOption) (*PartialReleaseHoldResponse, error)
	FinalizeTransaction(ctx context.Context, in *FinalizeTransactionRequest, opts ...grpc.CallOption) (*FinalizeTransactionResponse, error)
}

type ledgerClient struct {
	cc grpc.ClientConnInterface
}

func NewLedgerClient(cc grpc.ClientConnInterface) LedgerClient {
	return &ledgerClient{cc}
}

func (c *ledgerClient) InitializeAccount(ctx context.Context, in *InitializeAccountRequest, opts ...grpc.CallOption) (*InitializeAccountResponse, error) {
	out := new(InitializeAccountResponse)
	err := c.cc.Invoke(ctx, "/pkg.pbs.ledger.v1.Ledger/InitializeAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerClient) GetAccounts(ctx context.Context, in *GetAccountsRequest, opts ...grpc.CallOption) (*GetAccountsResponse, error) {
	out := new(GetAccountsResponse)
	err := c.cc.Invoke(ctx, "/pkg.pbs.ledger.v1.Ledger/GetAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, "/pkg.pbs.ledger.v1.Ledger/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerClient) PartialReleaseHold(ctx context.Context, in *PartialReleaseHoldRequest, opts ...grpc.CallOption) (*PartialReleaseHoldResponse, error) {
	out := new(PartialReleaseHoldResponse)
	err := c.cc.Invoke(ctx, "/pkg.pbs.ledger.v1.Ledger/PartialReleaseHold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerClient) FinalizeTransaction(ctx context.Context, in *FinalizeTransactionRequest, opts ...grpc.CallOption) (*FinalizeTransactionResponse, error) {
	out := new(FinalizeTransactionResponse)
	err := c.cc.Invoke(ctx, "/pkg.pbs.ledger.v1.Ledger/FinalizeTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LedgerServer is the server API for Ledger service.
// All implementations must embed UnimplementedLedgerServer
// for forward compatibility
type LedgerServer interface {
	InitializeAccount(context.Context, *InitializeAccountRequest) (*InitializeAccountResponse, error)
	GetAccounts(context.Context, *GetAccountsRequest) (*GetAccountsResponse, error)
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
	PartialReleaseHold(context.Context, *PartialReleaseHoldRequest) (*PartialReleaseHoldResponse, error)
	FinalizeTransaction(context.Context, *FinalizeTransactionRequest) (*FinalizeTransactionResponse, error)
	mustEmbedUnimplementedLedgerServer()
}

// UnimplementedLedgerServer must be embedded to have forward compatible implementations.
type UnimplementedLedgerServer struct {
}

func (UnimplementedLedgerServer) InitializeAccount(context.Context, *InitializeAccountRequest) (*InitializeAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitializeAccount not implemented")
}
func (UnimplementedLedgerServer) GetAccounts(context.Context, *GetAccountsRequest) (*GetAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccounts not implemented")
}
func (UnimplementedLedgerServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedLedgerServer) PartialReleaseHold(context.Context, *PartialReleaseHoldRequest) (*PartialReleaseHoldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PartialReleaseHold not implemented")
}
func (UnimplementedLedgerServer) FinalizeTransaction(context.Context, *FinalizeTransactionRequest) (*FinalizeTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinalizeTransaction not implemented")
}
func (UnimplementedLedgerServer) mustEmbedUnimplementedLedgerServer() {}

// UnsafeLedgerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LedgerServer will
// result in compilation errors.
type UnsafeLedgerServer interface {
	mustEmbedUnimplementedLedgerServer()
}

func RegisterLedgerServer(s grpc.ServiceRegistrar, srv LedgerServer) {
	s.RegisterService(&Ledger_ServiceDesc, srv)
}

func _Ledger_InitializeAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitializeAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServer).InitializeAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.pbs.ledger.v1.Ledger/InitializeAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServer).InitializeAccount(ctx, req.(*InitializeAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ledger_GetAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServer).GetAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.pbs.ledger.v1.Ledger/GetAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServer).GetAccounts(ctx, req.(*GetAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ledger_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.pbs.ledger.v1.Ledger/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ledger_PartialReleaseHold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartialReleaseHoldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServer).PartialReleaseHold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.pbs.ledger.v1.Ledger/PartialReleaseHold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServer).PartialReleaseHold(ctx, req.(*PartialReleaseHoldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ledger_FinalizeTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinalizeTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServer).FinalizeTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.pbs.ledger.v1.Ledger/FinalizeTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServer).FinalizeTransaction(ctx, req.(*FinalizeTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ledger_ServiceDesc is the grpc.ServiceDesc for Ledger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ledger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pkg.pbs.ledger.v1.Ledger",
	HandlerType: (*LedgerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitializeAccount",
			Handler:    _Ledger_InitializeAccount_Handler,
		},
		{
			MethodName: "GetAccounts",
			Handler:    _Ledger_GetAccounts_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _Ledger_CreateTransaction_Handler,
		},
		{
			MethodName: "PartialReleaseHold",
			Handler:    _Ledger_PartialReleaseHold_Handler,
		},
		{
			MethodName: "FinalizeTransaction",
			Handler:    _Ledger_FinalizeTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pbs/ledger/v1/service.proto",
}
