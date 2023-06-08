package main

import (
	"context"
	"log"
	"fmt"

	rpc "github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc/imservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
    rdb = &RedisClient{} // make the RedisClient with global visibility in the 'main' scope
)

func main() {
	ctx := context.Background() // https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go
	
	// init client on "redis-master:6379" while building this image to deploy in k8s
	// err := rdb.InitClient(ctx, "redis-master:6379", "")
	err := rdb.InitClient(ctx, "redis:6379", "")
    if err != nil {
       errMsg := fmt.Sprintf("failed to init Redis client, err: %v", err)
       log.Fatal(errMsg)
    }

	r, err := etcd.NewEtcdRegistry([]string{"etcd:2379"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	svr := rpc.NewServer(new(IMServiceImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "demo.rpc.server",
	}))

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
