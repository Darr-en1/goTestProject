package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"time"
)

func main() {
	c, err := client.NewClientWithOpts()
	if err != nil {
		panic(c)
	}

	ctx := context.Background()
	resp, err := c.ContainerCreate(ctx, &container.Config{
		Image: "mongo:latest",
		ExposedPorts: nat.PortSet{
			"27017/tcp": {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"27017/tcp": []nat.PortBinding{
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
	// 创建一个docker container  docker run
	err = c.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("container started")
	time.Sleep(5 * time.Second)

	// 通过容器检查可以获取动态映射的端口
	inspect, err := c.ContainerInspect(ctx, resp.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("listening at %+v\n",
		inspect.NetworkSettings.Ports["27017/tcp"][0],
	)

	fmt.Println("killing container")
	// c.ContainerStop    c.ContainerKill
	// docker rm
	err = c.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{
		Force: true, // -f 参数   相当与docker rm -f
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("killing container finish")

}
