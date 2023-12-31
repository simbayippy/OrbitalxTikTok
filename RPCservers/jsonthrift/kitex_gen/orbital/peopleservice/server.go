// Code generated by Kitex v0.5.2. DO NOT EDIT.
package peopleservice

import (
	server "github.com/cloudwego/kitex/server"
	orbital "github.com/simbayippy/OrbitalxTiktok/RPCservers/jsonthrift/kitex_gen/orbital"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler orbital.PeopleService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
