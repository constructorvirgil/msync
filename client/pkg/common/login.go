package common

import (
	"encoding/json"
	"errors"
)

const (
	ByteLenUser = 32
	ByteLenPwd  = 32
)

var (
	ErrLengthIllegal = errors.New("length is illegal")
)

type MsgLogin struct {
	User string `json:"user"`
	Pwd  string `json:"pwd"`
}

func (c *MsgLogin) Pack() ([]byte, error) {
	if len([]byte(c.User)) > ByteLenUser || len([]byte(c.Pwd)) > ByteLenPwd {
		return nil, ErrLengthIllegal
	}

	jbyte, err := json.Marshal(*c)
	if err != nil {
		return nil, err
	}

	return jbyte, nil
}

func (c *MsgLogin) UnPack(b []byte) error {
	return json.Unmarshal(b, c)
}
