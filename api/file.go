//  Simple Admin File
//
//  This is simple admin file manager api doc
//
//  Schemes: http, https
//  Host: localhost
//  BasePath: /api
//  Version: 0.0.1
//  Contact: yuansu.china.work@gmail.com
//
//  Consumes:
//    - application/json
//
//  Produces:
//    - application/json
//
// swagger:meta
package main

import (
	"flag"
	"fmt"

	"github.com/suyuan32/simple-admin-file/api/internal/config"
	"github.com/suyuan32/simple-admin-file/api/internal/handler"
	"github.com/suyuan32/simple-admin-file/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/file.yaml", "the config file")

func main() {
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
