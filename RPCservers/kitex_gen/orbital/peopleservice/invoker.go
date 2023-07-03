// Code generated by Kitex v0.5.2. DO NOT EDIT.

package peopleservice

import (
	server "github.com/cloudwego/kitex/server"
	orbital "github.com/simbayippy/OrbitalxTiktok/RPCservers/kitex_gen/orbital"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler orbital.PeopleService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
