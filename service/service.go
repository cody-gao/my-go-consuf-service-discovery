package service

import (
	"ch6-discovery/config"
	"ch6-discovery/discover"
	"context"
	"errors"
)

var ErrNotServiceInstances = errors.New("instance are not existed")

type Service interface {
	//健康检查接口
	HealthCheck() bool

	//打招呼接口
	SayHello() string

	//服务发现接口
	DiscoveryService(ctx context.Context, serviceName string) ([]interface{}, error)
}

type DiscoveryServiceImpl struct {
	discoveryClient discover.DiscoveryClient
}

func NewDiscoveryServiceImpl(discoveryClient discover.DiscoveryClient) Service {
	return &DiscoveryServiceImpl{
		discoveryClient: discoveryClient,
	}
}

func (*DiscoveryServiceImpl) SayHello() string {
	return "Hello World!"
}

func (service *DiscoveryServiceImpl) DiscoveryService(ctx context.Context, serviceName string) ([]interface{}, error) {
	//从consful中根据服务名获取服务实例列表
	instances := service.discoveryClient.DiscoverServices(serviceName, config.Logger)
	if instances == nil || len(instances) == 0 {
		return nil, ErrNotServiceInstances
	}

	return instances, nil
}

// HealthCheck implement Service method
// 用于检查服务的健康状态，这里仅仅返回true
func (*DiscoveryServiceImpl) HealthCheck() bool {
	return true
}
