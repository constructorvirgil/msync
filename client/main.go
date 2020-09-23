package main

import (
	"client/global"
	"client/pkg/flow"
	"client/pkg/setting"
	"fmt"
	"github.com/aceld/zinx/znet"
	"log"
	"net"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}
func setupSetting() error {
	settingObj, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = settingObj.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = settingObj.ReadSection("Client", &global.ClientSetting)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	fmt.Println("Client Test ... start")

	ip := global.ServerSetting.IP
	port := global.ServerSetting.Port

	fmt.Printf("ServerIP: %s:%s\n", ip, port)

	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	data := flow.NewFile("/Users/kiasma/WKspace/msync/client/bin/go_build_client")

	src, _ := data.Pack()
	for _, v := range src {
		_ = flow.Send(conn, 1, v)
		dp := znet.NewDataPack()
		_ = flow.Catch(conn, *dp)
	}

	//_ = flow.Send(conn, 0, []byte{})

}
