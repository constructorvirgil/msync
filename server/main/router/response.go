package router

import (
	"github.com/aceld/zinx/ziface"
	"github.com/constructorvirgil/msync/common"
	"github.com/constructorvirgil/msync/server/main/global"
)

func Response(request ziface.IRequest, msgId int, code int, msg string) error {
	status := common.RespStatus{
		Code: code,
		Msg:  msg,
	}

	bytes, err := status.Pack()
	if err != nil {
		return err
	}

	bytes, err = common.AESEncode(bytes, global.AESKey)
	if err != nil {
		return err
	}

	err = request.GetConnection().SendBuffMsg(uint32(msgId), bytes)
	if err != nil {
		return err
	}
	return nil
}
