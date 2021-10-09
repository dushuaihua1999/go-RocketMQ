// 负责namesrv的gRPC客户端功能
package server

import (
	"context"
	"errors"
	"go-RocketMQ/cmd/namesrv/proto"
)

type AcquireRegistryServer struct{}

func (a *AcquireRegistryServer) ListRegisteredBroker(ctx context.Context, in *proto.ReqReg) (*proto.RegistryMes, error) {
	if in == nil {
		return nil, errors.New("request is nil")
	}

	// 开始取出数据,看看是生产者还是消费者
	role := in.GetRole()
	// 集群选择 -- > 集群选择器
	// 选择调度器 --> role类型的调度器
	// 选择broker --> 根据调度器选择broker

}
