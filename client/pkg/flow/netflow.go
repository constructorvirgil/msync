package flow

import (
	"fmt"
	"github.com/aceld/zinx/znet"
	"io"
	"net"
)

func Send(conn net.Conn, id uint32, data []byte) error {
	dp := znet.NewDataPack()
	msg, _ := dp.Pack(znet.NewMsgPackage(id, data))
	_, err := conn.Write(msg)
	if err != nil {
		fmt.Println("Write error err", err)
		return err
	}
	return nil
}

func Catch(conn net.Conn, dp znet.DataPack) error {
	header := make([]byte, dp.GetHeadLen())
	_, err := io.ReadFull(conn, header)
	if err != nil {
		fmt.Println("Read head error")
		return err
	}

	msgHead, err := dp.Unpack(header)
	if err != nil {
		fmt.Println("Server unpack err:", err)
		return err
	}

	if msgHead.GetDataLen() > 0 {
		msg := msgHead.(*znet.Message)
		msg.Data = make([]byte, msg.GetDataLen())

		_, err := io.ReadFull(conn, msg.Data)
		if err != nil {
			fmt.Println("Server unpack data err:", err)
			return err
		}

		//fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
	}

	return nil
}
