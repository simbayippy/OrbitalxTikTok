// Code generated by Kitex v0.5.2. DO NOT EDIT.

package peopleservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	orbital "github.com/simbayippy/orbital/kitex_gen/orbital"
)

func serviceInfo() *kitex.ServiceInfo {
	return peopleServiceServiceInfo
}

var peopleServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "PeopleService"
	handlerType := (*orbital.PeopleService)(nil)
	methods := map[string]kitex.MethodInfo{
		"editPerson": kitex.NewMethodInfo(editPersonHandler, newPeopleServiceEditPersonArgs, newPeopleServiceEditPersonResult, false),
		"echo":       kitex.NewMethodInfo(echoHandler, newPeopleServiceEchoArgs, newPeopleServiceEchoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "orbital",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func editPersonHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*orbital.PeopleServiceEditPersonArgs)
	realResult := result.(*orbital.PeopleServiceEditPersonResult)
	success, err := handler.(orbital.PeopleService).EditPerson(ctx, realArg.Person)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPeopleServiceEditPersonArgs() interface{} {
	return orbital.NewPeopleServiceEditPersonArgs()
}

func newPeopleServiceEditPersonResult() interface{} {
	return orbital.NewPeopleServiceEditPersonResult()
}

func echoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*orbital.PeopleServiceEchoArgs)
	realResult := result.(*orbital.PeopleServiceEchoResult)
	success, err := handler.(orbital.PeopleService).Echo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPeopleServiceEchoArgs() interface{} {
	return orbital.NewPeopleServiceEchoArgs()
}

func newPeopleServiceEchoResult() interface{} {
	return orbital.NewPeopleServiceEchoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) EditPerson(ctx context.Context, person *orbital.Person) (r *orbital.Person, err error) {
	var _args orbital.PeopleServiceEditPersonArgs
	_args.Person = person
	var _result orbital.PeopleServiceEditPersonResult
	if err = p.c.Call(ctx, "editPerson", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Echo(ctx context.Context, req *orbital.Request) (r *orbital.Response, err error) {
	var _args orbital.PeopleServiceEchoArgs
	_args.Req = req
	var _result orbital.PeopleServiceEchoResult
	if err = p.c.Call(ctx, "echo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
