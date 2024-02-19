package svc

import (
	"DP/DP/internal/config"
	"DP/rpc/Bms/bmsclient"
	"DP/rpc/Sms/smsclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	SmsRpcClient smsclient.Sms
	BmsRpcClient bmsclient.Bms
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		SmsRpcClient: smsclient.NewSms(zrpc.MustNewClient(c.SmsRPCConf)),
		BmsRpcClient: bmsclient.NewBms(zrpc.MustNewClient(c.BmsRPCConf)),
	}
}
