package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	SmsRPCConf zrpc.RpcClientConf
	BmsRPCConf zrpc.RpcClientConf
	AmsRPCConf zrpc.RpcClientConf
	//JWT的认证
	Auth struct {
		AccessSecret string
		//AccessExpire int64
	}
}
