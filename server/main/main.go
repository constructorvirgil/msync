package main

import (
	"fmt"
	"github.com/aceld/zinx/znet"
	"github.com/constructorvirgil/msync/server/main/filemanage"
	"github.com/constructorvirgil/msync/server/main/global"
	"github.com/constructorvirgil/msync/server/main/router"
)

func main() {
	var err error

	//创建服务器句柄
	s := znet.NewServer()

	//配置路由
	s.AddRouter(router.MsgIdLogin, &router.LoginRouter{})
	s.AddRouter(router.MsgIdTransFile, &router.TransFileRouter{})

	//初始化全局资源
	err = global.Init()
	if err != nil {
		fmt.Println("init global resource failure: ", err)
		return
	}

	//启动文件管理routine
	filemanage.Init()

	//开启服务
	s.Serve()
}