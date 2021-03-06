// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/userService.proto

package userService

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	//user
	AddUser(ctx context.Context, in *UserModel, opts ...client.CallOption) (*Response, error)
	UpdateUser(ctx context.Context, in *UserModel, opts ...client.CallOption) (*Response, error)
	DeleteUser(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetUser(ctx context.Context, in *Request, opts ...client.CallOption) (*UserModel, error)
	GetAllUser(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	//end
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (UserService_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (UserService_PingPongService, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) AddUser(ctx context.Context, in *UserModel, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.AddUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateUser(ctx context.Context, in *UserModel, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.UpdateUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) DeleteUser(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.DeleteUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUser(ctx context.Context, in *Request, opts ...client.CallOption) (*UserModel, error) {
	req := c.c.NewRequest(c.name, "UserService.GetUser", in)
	out := new(UserModel)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetAllUser(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.GetAllUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (UserService_StreamService, error) {
	req := c.c.NewRequest(c.name, "UserService.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &userServiceStream{stream}, nil
}

type UserService_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type userServiceStream struct {
	stream client.Stream
}

func (x *userServiceStream) Close() error {
	return x.stream.Close()
}

func (x *userServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *userServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userService) PingPong(ctx context.Context, opts ...client.CallOption) (UserService_PingPongService, error) {
	req := c.c.NewRequest(c.name, "UserService.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &userServicePingPong{stream}, nil
}

type UserService_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type userServicePingPong struct {
	stream client.Stream
}

func (x *userServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *userServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *userServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *userServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Call(context.Context, *Request, *Response) error
	//user
	AddUser(context.Context, *UserModel, *Response) error
	UpdateUser(context.Context, *UserModel, *Response) error
	DeleteUser(context.Context, *Request, *Response) error
	GetUser(context.Context, *Request, *UserModel) error
	GetAllUser(context.Context, *Request, *Response) error
	//end
	Stream(context.Context, *StreamingRequest, UserService_StreamStream) error
	PingPong(context.Context, UserService_PingPongStream) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Call(ctx context.Context, in *Request, out *Response) error
		AddUser(ctx context.Context, in *UserModel, out *Response) error
		UpdateUser(ctx context.Context, in *UserModel, out *Response) error
		DeleteUser(ctx context.Context, in *Request, out *Response) error
		GetUser(ctx context.Context, in *Request, out *UserModel) error
		GetAllUser(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.Call(ctx, in, out)
}

func (h *userServiceHandler) AddUser(ctx context.Context, in *UserModel, out *Response) error {
	return h.UserServiceHandler.AddUser(ctx, in, out)
}

func (h *userServiceHandler) UpdateUser(ctx context.Context, in *UserModel, out *Response) error {
	return h.UserServiceHandler.UpdateUser(ctx, in, out)
}

func (h *userServiceHandler) DeleteUser(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.DeleteUser(ctx, in, out)
}

func (h *userServiceHandler) GetUser(ctx context.Context, in *Request, out *UserModel) error {
	return h.UserServiceHandler.GetUser(ctx, in, out)
}

func (h *userServiceHandler) GetAllUser(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.GetAllUser(ctx, in, out)
}

func (h *userServiceHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.UserServiceHandler.Stream(ctx, m, &userServiceStreamStream{stream})
}

type UserService_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type userServiceStreamStream struct {
	stream server.Stream
}

func (x *userServiceStreamStream) Close() error {
	return x.stream.Close()
}

func (x *userServiceStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *userServiceStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServiceStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServiceStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *userServiceHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.UserServiceHandler.PingPong(ctx, &userServicePingPongStream{stream})
}

type UserService_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type userServicePingPongStream struct {
	stream server.Stream
}

func (x *userServicePingPongStream) Close() error {
	return x.stream.Close()
}

func (x *userServicePingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *userServicePingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServicePingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServicePingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *userServicePingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
