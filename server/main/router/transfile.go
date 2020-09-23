package router

import (
	"encoding/json"
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/constructorvirgil/msync/common"
	"github.com/constructorvirgil/msync/main/filemanage"
)


//传输文件路由
type TransFileRouter struct {
	znet.BaseRouter
}

//传输文件回调
func (this *TransFileRouter) Handle(request ziface.IRequest) {
	data := request.GetData()

	filePart := common.FilePart{}
	err := json.Unmarshal(data, &filePart)
	if err != nil {
		fmt.Println("json unmarshal failure: ", err)
		return
	}

	fmt.Printf("id=%s|index=%d|maxIndex=%d|length of cont=%d, pack recv...\n", string(filePart.FileId), filePart.FileIndex, filePart.FileMaxIndex, len(filePart.PartContent))
	if filePart.FileMaxIndex != filePart.FileMaxIndex {
		filemanage.Add(string(filePart.FileId), filePart.PartContent)
	}else{
		filemanage.AddAndFlush(string(filePart.FileId), filePart.PartContent)
	}

	err = request.GetConnection().SendBuffMsg(MsgIdTransFile, []byte("recv file..."))
	if err != nil {
		fmt.Println(err)
	}
}