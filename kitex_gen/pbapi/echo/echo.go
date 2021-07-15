// Code generated by Kitex v1.4.2. DO NOT EDIT.

package echo

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex-examples/kitex_gen/pbapi"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return echoServiceInfo
}

var echoServiceInfo = newServiceInfo()

func newServiceInfo() *kitex.ServiceInfo {
	serviceName := "Echo"
	handlerType := (*pbapi.Echo)(nil)
	methods := map[string]kitex.MethodInfo{
		"Echo": kitex.NewMethodInfo(echoHandler, newEchoArgs, newEchoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "pbapi",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v1.4.2",
		Extra:           extra,
	}
	return svcInfo
}

func echoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*EchoArgs)
	realResult := result.(*EchoResult)
	success, err := handler.(pbapi.Echo).Echo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEchoArgs() interface{} {
	return &EchoArgs{}
}

func newEchoResult() interface{} {
	return &EchoResult{}
}

type EchoArgs struct {
	Req *pbapi.Request
}

func (p *EchoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in EchoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *EchoArgs) Unmarshal(in []byte) error {
	msg := new(pbapi.Request)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var EchoArgs_Req_DEFAULT *pbapi.Request

func (p *EchoArgs) GetReq() *pbapi.Request {
	if !p.IsSetReq() {
		return EchoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EchoArgs) IsSetReq() bool {
	return p.Req != nil
}

type EchoResult struct {
	Success *pbapi.Response
}

var EchoResult_Success_DEFAULT *pbapi.Response

func (p *EchoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in EchoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *EchoResult) Unmarshal(in []byte) error {
	msg := new(pbapi.Response)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EchoResult) GetSuccess() *pbapi.Response {
	if !p.IsSetSuccess() {
		return EchoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EchoResult) SetSuccess(x interface{}) {
	p.Success = x.(*pbapi.Response)
}

func (p *EchoResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Echo(ctx context.Context, Req *pbapi.Request) (r *pbapi.Response, err error) {
	var _args EchoArgs
	_args.Req = Req
	var _result EchoResult
	if err = p.c.Call(ctx, "Echo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
