// Code generated by Kitex v0.9.0. DO NOT EDIT.

package servicea

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	service "grpc_multi_service/kitex_gen/multi/service"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ChatA": kitex.NewMethodInfo(
		chatAHandler,
		newChatAArgs,
		newChatAResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	serviceAServiceInfo                = NewServiceInfo()
	serviceAServiceInfoForClient       = NewServiceInfoForClient()
	serviceAServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return serviceAServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return serviceAServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return serviceAServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "ServiceA"
	handlerType := (*service.ServiceA)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "multiservice",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.0",
		Extra:           extra,
	}
	return svcInfo
}

func chatAHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(service.RequestA)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(service.ServiceA).ChatA(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ChatAArgs:
		success, err := handler.(service.ServiceA).ChatA(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ChatAResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newChatAArgs() interface{} {
	return &ChatAArgs{}
}

func newChatAResult() interface{} {
	return &ChatAResult{}
}

type ChatAArgs struct {
	Req *service.RequestA
}

func (p *ChatAArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(service.RequestA)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ChatAArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ChatAArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ChatAArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ChatAArgs) Unmarshal(in []byte) error {
	msg := new(service.RequestA)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ChatAArgs_Req_DEFAULT *service.RequestA

func (p *ChatAArgs) GetReq() *service.RequestA {
	if !p.IsSetReq() {
		return ChatAArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChatAArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ChatAArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ChatAResult struct {
	Success *service.Reply
}

var ChatAResult_Success_DEFAULT *service.Reply

func (p *ChatAResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(service.Reply)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ChatAResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ChatAResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ChatAResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ChatAResult) Unmarshal(in []byte) error {
	msg := new(service.Reply)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChatAResult) GetSuccess() *service.Reply {
	if !p.IsSetSuccess() {
		return ChatAResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChatAResult) SetSuccess(x interface{}) {
	p.Success = x.(*service.Reply)
}

func (p *ChatAResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChatAResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ChatA(ctx context.Context, Req *service.RequestA) (r *service.Reply, err error) {
	var _args ChatAArgs
	_args.Req = Req
	var _result ChatAResult
	if err = p.c.Call(ctx, "ChatA", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
