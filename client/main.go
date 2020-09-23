package main

import (
	"client/global"
	"client/pkg/flow"
	"client/pkg/setting"
	"fmt"
	"log"
	"net"
	"time"
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
	//3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)

	ip := global.ServerSetting.IP
	port := global.ServerSetting.Port

	fmt.Printf("ServerIP: %s:%s\n", ip, port)

	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	//data := flow.NewFile("/Users/kiasma/WKspace/msync/testFile")
	//
	//src, _ := data.Pack()
	//for _, v := range src {
	//	_ = flow.Send(conn, 0, )
	//}

	flow.Send(conn, 0, []byte{})
}
