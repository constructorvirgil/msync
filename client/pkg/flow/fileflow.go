package flow

import (
	"client/pkg/common"
	"client/pkg/encode"
	"fmt"
	"github.com/aceld/zinx/znet"
	"io/ioutil"
	"net"
)

func NewFile(fName string) *common.File {
	data, err := ioutil.ReadFile(fName)
	if err != nil {
		return nil
	}

	temp := encode.GetMD5Encode(data)
	key := []byte("573392132@qq.com")
	enFile, err := encode.EnFile(data, key)
	if err != nil {
		return nil
	}
	return &common.File{FileId: []byte(temp), FileContent: enFile}
}

func File2net(fName, ip, port string) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}
	data := NewFile(fName)

	src, _ := data.Pack()
	for _, v := range src {
		_ = Send(conn, 1, v)
		dp := znet.NewDataPack()
		_ = Catch(conn, *dp)
	}

	fmt.Printf("Finished %v", fName)
}

func GetFiles(path string, files []string) ([]string, error) {
	rd, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, fod := range rd {
		if fod.IsDir() {
			files, err = GetFiles(path+"/"+fod.Name(), files)
		} else {
			files = append(files, fod.Name())
		}
	}

	return files, nil
}
