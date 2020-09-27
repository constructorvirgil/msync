package flow

import (
	"client/pkg/key"
	"fmt"
	"github.com/aceld/zinx/znet"
	"github.com/constructorvirgil/msync/common"
	"io/ioutil"
	"net"
)

func NewFile(fName string) *common.File {
	data, err := ioutil.ReadFile(fName)
	if err != nil {
		fmt.Printf("Can't find file name: %v err: %v", fName, err)
		return nil
	}

	fmt.Println("read file: ", len(data))
	temp := common.GetMD5Encode(data)
	fmt.Println("file md5: ", temp)

	fmt.Println("encode file content: ", len(data))
	return &common.File{FileId: []byte(temp), Path: fName, FileContent: data}
}

func File2net(fName, ip, port string) {
	conn, err := net.Dial("tcp", ip+":"+port)

	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}
	defer conn.Close()

	data := NewFile(fName)
	if data == nil {
		return
	}

	aesKey := key.GetKey()

	src, _ := data.Pack()
	var encodePack [][]byte

	for _, pack := range src {
		i, err := common.AESEncode(pack, aesKey)
		if err != nil {
			fmt.Println("AESEncode ERROR:", err)
			return
		}
		encodePack = append(encodePack, i)
	}

	for _, v := range encodePack {
		_ = Send(conn, 1, v)
		dp := znet.NewDataPack()
		_ = Catch(conn, *dp)
	}

	fmt.Printf("Finished %v\n\n", fName)
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
			files = append(files, path+"/"+fod.Name())
		}
	}

	return files, nil
}
