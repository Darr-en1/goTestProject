package testEnvironment

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"testing"
)

const (
	Image        = "mongo:latest"
	ExposedPorts = "27017/tcp"
)

// RunWithMongoInDocker runs the tests with a mongodb instance in a docker container
func RunWithMongoInDocker(m *testing.M, mongoURI *string) int {
	c, err := client.NewClientWithOpts()
	if err != nil {
		panic(c)
	}

	ctx := context.Background()

	resp, err := c.ContainerCreate(ctx, &container.Config{
		Image: Image,
		ExposedPorts: nat.PortSet{
			ExposedPorts: {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			ExposedPorts: []nat.PortBinding{
				{
					HostIP:   "127.0.0.1",
					HostPort: "0", // 不设置或者设置默认值，会自动分配一个可以用的 port,这样可以防止端口被暂用
				},
			},
		},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	containerID := resp.ID

	defer func() {
		fmt.Println("killing container")
		// docker rm -F
		err = c.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{Force: true})
		if err != nil {
			panic(err)
		}
		fmt.Println("killing container finish")

	}()

	// 创建一个docker container  docker run
	err = c.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("container started")

	// 通过容器检查可以获取动态映射的端口
	inspect, err := c.ContainerInspect(ctx, containerID)
	if err != nil {
		panic(err)
	}
	host := inspect.NetworkSettings.Ports[ExposedPorts][0]

	*mongoURI = fmt.Sprintf("mongodb://%s:%s", host.HostIP, host.HostPort)

	fmt.Printf("listening at %+v\n", host)

	return m.Run()
}
