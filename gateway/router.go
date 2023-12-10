// Code generated by hertz generator.

package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
	"msp/gateway/biz/handler"
	"msp/gateway/biz/middleware"
	"msp/gateway/config"
	"os"
	"path"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...
	registerGateway(r)
}

// registerGateway registers the router of gateway
func registerGateway(r *server.Hertz) {
	group := r.Group("/gateway").Use(middleware.GatewayAuth()...)

	if handler.SvcRouteMap == nil {
		handler.SvcRouteMap = make(map[string]genericclient.Client)
	}

	if handler.SvcFingerPrintMap == nil {
		handler.SvcFingerPrintMap = make(map[string]string)
	}

	resource := config.GlobalServerConfig.GatewayResource
	if len(resource) == 0 {
		hlog.Fatalf("new thrift file provider failed: resource len zero")
	}
	// 从注册中心获取服务
	consulResolver, err := consul.NewConsulResolver(fmt.Sprintf("%s:%d",
		config.GlobalLocalConfig.ConsulConfig.Host,
		config.GlobalLocalConfig.ConsulConfig.Port))

	if err != nil {
		hlog.Fatalf("Resolver connect failed: %v", err)
	}

	// 多个路由可能映射到相同的svr、idl，增加cli缓存
	loadingCacheMap := make(map[string]genericclient.Client)

	includeContent := make(map[string]string)

	for _, pathResource := range resource {
		route := pathResource.Route
		svrName := pathResource.SvrName
		fingerPrint := pathResource.FingerPrint
		idlPath := pathResource.IdlPath
		includePath := pathResource.IncludePath
		parentPath := pathResource.ParentPath
		evalIncludeContent(parentPath, idlPath, includePath, includeContent)
		providerConnect(route, svrName, fingerPrint, idlPath, includeContent, consulResolver, loadingCacheMap)
	}
	group.POST("/*path", handler.Gateway)
	group.GET("/mtop/api/list", handler.GetawayList)

}

func evalIncludeContent(parentPath, baseIdl string, includePath []string, includeContent map[string]string) {
	for _, idlPath := range includePath {
		if _, ok := includeContent[idlPath]; !ok {
			evalContent(parentPath, idlPath, includeContent)
		}
	}
	evalContent(parentPath, baseIdl, includeContent)
}

func evalContent(parentPath, idlPath string, includeContent map[string]string) {
	entry, err := os.ReadFile(path.Join(parentPath, idlPath))
	if err != nil {
		hlog.Fatalf("new thrift includeContent failed: %v", err)
	}

	content := string(entry)
	includeContent[idlPath] = content
}

func providerConnect(route, svrName, fingerPrint, idlPath string, includes map[string]string, resolver discovery.Resolver, loadingCacheMap map[string]genericclient.Client) {

	provider, err := generic.NewThriftContentProvider(includes[idlPath], includes)
	if err != nil {
		hlog.Fatalf("new thrift file provider failed: %v", err)
		return
	}
	g, err := generic.HTTPThriftGeneric(provider)
	if err != nil {
		hlog.Fatal(err)
	}
	// 相同svrName下多idl中均定义service时，不确定会不会有问题，所以粒度更细一些（如果provider为svrName纬度，则cacheKey = svrName）
	cacheKey := svrName + idlPath
	if _, ok := loadingCacheMap[cacheKey]; !ok {
		cli, err := genericclient.NewClient(
			svrName,
			g,
			client.WithResolver(resolver),
			client.WithTransportProtocol(transport.TTHeader),
			client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		)
		if err != nil {
			hlog.Fatal(err)
		}
		loadingCacheMap[cacheKey] = cli
	}

	// 绑定 handler 对应的 泛化服务
	handler.SvcRouteMap[route] = loadingCacheMap[cacheKey]
	// 绑定泛化对应的方法
	handler.SvcFingerPrintMap[route] = fingerPrint
}
