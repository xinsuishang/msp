// Code generated by Kitex v0.7.3. DO NOT EDIT.
package chatservice

import (
	server "github.com/cloudwego/kitex/server"
	gpt "msp/biz_server/kitex_gen/gpt"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler gpt.ChatService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
