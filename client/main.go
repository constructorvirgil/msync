package main

import (
	"client/global"
	"client/pkg/flow"
	"client/pkg/setting"
	"fmt"
	"log"
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

	ip := global.ServerSetting.IP
	port := global.ServerSetting.Port

	fmt.Printf("ServerIP: %s:%s\n", ip, port)

	go func() {
		flow.File2net("/Users/kiasma/WKspace/msync/client/bin/test3.txt", ip, port)
	}()

	go func() {
		flow.File2net("/Users/kiasma/WKspace/msync/client/bin/go_build_client", ip, port)
	}()

	for true {
		time.Sleep(time.Second)
	}
}
