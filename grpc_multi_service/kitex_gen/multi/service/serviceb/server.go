// Code generated by Kitex v0.8.0. DO NOT EDIT.
package serviceb

import (
	server "github.com/cloudwego/kitex/server"
	service "grpc_multi_service/kitex_gen/multi/service"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler service.ServiceB, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
