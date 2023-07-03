// Code generated by Kitex v0.5.2. DO NOT EDIT.
package peoplesservice

import (
	server "github.com/cloudwego/kitex/server"
	orbital2 "github.com/simbayippy/OrbitalxTiktok/APIGateway/kitex_gen/orbital2"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler orbital2.PeoplesService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}