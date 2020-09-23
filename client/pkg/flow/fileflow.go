package flow

import (
	"client/pkg/common"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/aceld/zinx/znet"
	"io/ioutil"
	"net"
)

//返回一个32位md5加密后的字符串
func GetMD5Encode(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func NewFile(fName string) *common.File {
	data, err := ioutil.ReadFile(fName)
	if err != nil {
		return nil
	}

	temp := GetMD5Encode(data)
	return &common.File{FileId: []byte(temp), FileContent: data}
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
}
