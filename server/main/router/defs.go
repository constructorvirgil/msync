package router

//msgId 定义
const (
	MsgIdLogin = iota
	MsgIdTransFile
)

//响应码
const(
	RespCodeOK = iota
	RespCodeDecodeError
)