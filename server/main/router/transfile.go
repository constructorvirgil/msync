package router

import (
	"encoding/json"
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/constructorvirgil/msync/common"
	"github.com/constructorvirgil/msync/server/main/filemanage"
	"github.com/constructorvirgil/msync/server/main/global"
)


//传输文件路由
type TransFileRouter struct {
	znet.BaseRouter
}

//传输文件回调
func (this *TransFileRouter) Handle(request ziface.IRequest) {
	var err error
	data := request.GetData()

	//AES  解密得到明文
	//fmt.Printf("data=%v|aeskey=%v\n", string(data), string(global.AESKey))  //打印密文
	data, err = common.AESDecode(data, global.AESKey)
	if err != nil {
		fmt.Println("aes decode failure: ", err)
		return
	}
	//fmt.Println("data after decode: ", string(data))  //解密之后得到的明文

	filePart := common.FilePart{}
	err = json.Unmarshal(data, &filePart)
	if err != nil {
		fmt.Println("json unmarshal failure: ", err)
		return
	}

	//fmt.Printf("id=%s|index=%d|maxIndex=%d|length of cont=%d, pack recv...\n", string(filePart.FileId), filePart.FileIndex, filePart.FileMaxIndex, len(filePart.PartContent))
	filemanage.Add(filemanage.FilePart{
		Id:       string(filePart.FileId),
		Index:    filePart.FileIndex,
		MaxIndex: filePart.FileMaxIndex,
		Path:     filePart.Path,
		Part:     filePart.PartContent,
	})

	err = request.GetConnection().SendBuffMsg(MsgIdTransFile, []byte("recv file..."))
	if err != nil {
		fmt.Println(err)
	}
}