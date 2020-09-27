package common

import "encoding/json"

type RespStatus struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
}

//序列化
func (c *RespStatus) Pack() ([]byte, error) {
	bytes, err := json.Marshal(*c)
	if err != nil{
		return nil, err 
	}
	return bytes, err
}

//反序列化
func (c *RespStatus) Unpack(bytes []byte) error {
	err := json.Unmarshal(bytes, c)
	if err != nil {
		return err
	}
	return nil
}