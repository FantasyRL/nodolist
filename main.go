// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/logger/accesslog"
	"nodolist/biz/dal"
	"nodolist/pkg/conf"
	"nodolist/pkg/utils/jwt"
)

// @title           NodoList
// @version         1.0
// @description     video website

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /nodolist
func main() {

	dal.Init(conf.Init())
	jwt.Init()

	h := server.New(
		server.WithHostPorts(conf.ServerAddr),
		server.WithStreamBody(true),
		//server.WithMaxRequestBodySize(16*1024*1024), //最大字节数
	)
	h.Use(accesslog.New())

	register(h)
	h.Spin()
}
