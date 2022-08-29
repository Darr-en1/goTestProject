package base

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"sync"
)

var (
	client *api.Client
	once   sync.Once
)

func init() {
	GetClient()
}

func GetClient() *api.Client {
	once.Do(func() {
		cfg := api.DefaultConfig()
		cfg.Address = "127.0.0.1:8500"
		var err error
		client, err = api.NewClient(cfg)
		if err != nil {
			panic(err)
		}
	})
	return client

}

func Register(address string, port int, name string, tags []string, id string) {

	registration := api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Tags:    tags,
		Port:    port,
		Address: address,
		Check: &api.AgentServiceCheck{ // 健康检查
			HTTP:                           fmt.Sprintf("https://%s:%d/health", address, port), // 服务的健康检查的接口
			Timeout:                        "5s",
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: "10s",
		},
	}
	err := client.Agent().ServiceRegister(&registration)
	if err != nil {
		panic(err)
	}
}

// RegisterWithGRPCHealthCheck 需要grpc 先注册健康监测的方法 grpc_health_v1.RegisterHealthServer(s, health.NewServer())
func RegisterWithGRPCHealthCheck(address string, port int, name string, tags []string, id string) {
	registration := api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Tags:    tags,
		Port:    port,
		Address: address,
		Check: &api.AgentServiceCheck{ // 健康检查
			GRPC:                           fmt.Sprintf("%s:%d", address, port),
			Timeout:                        "5s",
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: "10s",
		},
	}
	err := client.Agent().ServiceRegister(&registration)
	if err != nil {
		panic(err)
	}
}

func AllServices() {
	services, err := client.Agent().Services()

	if err != nil {
		panic(err)
	}

	for key, service := range services {
		fmt.Println(key, service)
	}
}

func ServicesWithFilter() {
	// 服务名称等于darr_en1
	services, err := client.Agent().ServicesWithFilter(`Service=="darr_en1"`)

	if err != nil {
		panic(err)
	}

	for key, service := range services {
		fmt.Println(key, service)
	}
}

// GetServicesWithFilter 从注册中心获取到别的服务的信息
func GetServicesWithFilter(filter string) (string, int) {
	// 服务名称等于darr_en1
	services, err := client.Agent().ServicesWithFilter(filter)

	if err != nil {
		panic(err)
	}

	for _, service := range services {
		return service.Address, service.Port
	}
	return "", 0
}
