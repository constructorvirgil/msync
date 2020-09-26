package main

import (
	"client/global"
	"client/pkg/flow"
	"client/pkg/setting"
	"fmt"
	"log"
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

	var files []string
	files, _ = flow.GetFiles(".", files)
	fmt.Println(files)

	//ch := make(chan int)
	for _, file := range files {
		//go func(file string, ch chan int) {
		fmt.Println("in go routine file name: ", file)
		flow.File2net(file, ip, port)
		//ch <- 0
		//}(file, ch)
	}

	//for i := 0; i < len(files); i++ {
	//	_ = <-ch
	//}

}
