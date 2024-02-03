// Code generated by Kitex v0.7.3. DO NOT EDIT.

package chatservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	gpt "msp/biz_server/kitex_gen/gpt"
)

func serviceInfo() *kitex.ServiceInfo {
	return chatServiceServiceInfo
}

var chatServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ChatService"
	handlerType := (*gpt.ChatService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Chat":                kitex.NewMethodInfo(chatHandler, newChatServiceChatArgs, newChatServiceChatResult, false),
		"ChatApplicationList": kitex.NewMethodInfo(chatApplicationListHandler, newChatServiceChatApplicationListArgs, newChatServiceChatApplicationListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "gpt",
		"ServiceFilePath": `../idl/rpc/gpt.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.3",
		Extra:           extra,
	}
	return svcInfo
}

func chatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*gpt.ChatServiceChatArgs)
	realResult := result.(*gpt.ChatServiceChatResult)
	success, err := handler.(gpt.ChatService).Chat(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newChatServiceChatArgs() interface{} {
	return gpt.NewChatServiceChatArgs()
}

func newChatServiceChatResult() interface{} {
	return gpt.NewChatServiceChatResult()
}

func chatApplicationListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*gpt.ChatServiceChatApplicationListArgs)
	realResult := result.(*gpt.ChatServiceChatApplicationListResult)
	success, err := handler.(gpt.ChatService).ChatApplicationList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newChatServiceChatApplicationListArgs() interface{} {
	return gpt.NewChatServiceChatApplicationListArgs()
}

func newChatServiceChatApplicationListResult() interface{} {
	return gpt.NewChatServiceChatApplicationListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Chat(ctx context.Context, req *gpt.ApplicationReq) (r *gpt.ApplicationResp, err error) {
	var _args gpt.ChatServiceChatArgs
	_args.Req = req
	var _result gpt.ChatServiceChatResult
	if err = p.c.Call(ctx, "Chat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChatApplicationList(ctx context.Context, req *gpt.ApplicationListReq) (r *gpt.ApplicationListResp, err error) {
	var _args gpt.ChatServiceChatApplicationListArgs
	_args.Req = req
	var _result gpt.ChatServiceChatApplicationListResult
	if err = p.c.Call(ctx, "ChatApplicationList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
