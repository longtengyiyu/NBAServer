package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"nbaserver/internal/config"
	"nbaserver/internal/handler"
	"nbaserver/internal/spider"
	"nbaserver/internal/svc"
)

var configFile = flag.String("f", "etc/nbaserver-api.yaml", "the config file")

func main() {

	spider.Start()

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
