// Code generated by Kitex v0.5.2. DO NOT EDIT.

package peoplesservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	orbital2 "github.com/simbayippy/orbital/kitex_gen/orbital2"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return peoplesServiceServiceInfo
}

var peoplesServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "PeoplesService"
	handlerType := (*orbital2.PeoplesService)(nil)
	methods := map[string]kitex.MethodInfo{
		"editPerson": kitex.NewMethodInfo(editPersonHandler, newEditPersonArgs, newEditPersonResult, false),
		"echo":       kitex.NewMethodInfo(echoHandler, newEchoArgs, newEchoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "orbital2",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func editPersonHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(orbital2.Person)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(orbital2.PeoplesService).EditPerson(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *EditPersonArgs:
		success, err := handler.(orbital2.PeoplesService).EditPerson(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*EditPersonResult)
		realResult.Success = success
	}
	return nil
}
func newEditPersonArgs() interface{} {
	return &EditPersonArgs{}
}

func newEditPersonResult() interface{} {
	return &EditPersonResult{}
}

type EditPersonArgs struct {
	Req *orbital2.Person
}

func (p *EditPersonArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(orbital2.Person)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *EditPersonArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *EditPersonArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *EditPersonArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in EditPersonArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *EditPersonArgs) Unmarshal(in []byte) error {
	msg := new(orbital2.Person)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var EditPersonArgs_Req_DEFAULT *orbital2.Person

func (p *EditPersonArgs) GetReq() *orbital2.Person {
	if !p.IsSetReq() {
		return EditPersonArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EditPersonArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *EditPersonArgs) GetFirstArgument() interface{} {
	return p.Req
}

type EditPersonResult struct {
	Success *orbital2.Person
}

var EditPersonResult_Success_DEFAULT *orbital2.Person

func (p *EditPersonResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(orbital2.Person)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *EditPersonResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *EditPersonResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *EditPersonResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in EditPersonResult")
	}
	return proto.Marshal(p.Success)
}

func (p *EditPersonResult) Unmarshal(in []byte) error {
	msg := new(orbital2.Person)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditPersonResult) GetSuccess() *orbital2.Person {
	if !p.IsSetSuccess() {
		return EditPersonResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EditPersonResult) SetSuccess(x interface{}) {
	p.Success = x.(*orbital2.Person)
}

func (p *EditPersonResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EditPersonResult) GetResult() interface{} {
	return p.Success
}

func echoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(orbital2.Request)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(orbital2.PeoplesService).Echo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *EchoArgs:
		success, err := handler.(orbital2.PeoplesService).Echo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*EchoResult)
		realResult.Success = success
	}
	return nil
}
func newEchoArgs() interface{} {
	return &EchoArgs{}
}

func newEchoResult() interface{} {
	return &EchoResult{}
}

type EchoArgs struct {
	Req *orbital2.Request
}

func (p *EchoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(orbital2.Request)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *EchoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *EchoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *EchoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in EchoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *EchoArgs) Unmarshal(in []byte) error {
	msg := new(orbital2.Request)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var EchoArgs_Req_DEFAULT *orbital2.Request

func (p *EchoArgs) GetReq() *orbital2.Request {
	if !p.IsSetReq() {
		return EchoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EchoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *EchoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type EchoResult struct {
	Success *orbital2.Response
}

var EchoResult_Success_DEFAULT *orbital2.Response

func (p *EchoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(orbital2.Response)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *EchoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *EchoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *EchoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in EchoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *EchoResult) Unmarshal(in []byte) error {
	msg := new(orbital2.Response)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EchoResult) GetSuccess() *orbital2.Response {
	if !p.IsSetSuccess() {
		return EchoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EchoResult) SetSuccess(x interface{}) {
	p.Success = x.(*orbital2.Response)
}

func (p *EchoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EchoResult) GetResult() interface{} {
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

func (p *kClient) EditPerson(ctx context.Context, Req *orbital2.Person) (r *orbital2.Person, err error) {
	var _args EditPersonArgs
	_args.Req = Req
	var _result EditPersonResult
	if err = p.c.Call(ctx, "editPerson", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Echo(ctx context.Context, Req *orbital2.Request) (r *orbital2.Response, err error) {
	var _args EchoArgs
	_args.Req = Req
	var _result EchoResult
	if err = p.c.Call(ctx, "echo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
