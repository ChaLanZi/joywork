// Code generated by goctl. DO NOT EDIT!
// Source: account.proto

package main

import (
	"flag"
	"fmt"

	"account/rpc/internal/config"
	"account/rpc/internal/server"
	"account/rpc/internal/svc"
	account "account/rpc/pb"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/rpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	accountServiceSrv := server.NewAccountServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		account.RegisterAccountServiceServer(grpcServer, accountServiceSrv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
