package global

import (
	"github.com/constructorvirgil/msync/common"
)

var AESKey []byte

func Init() error{
	var err error
	AESKey, err = common.ReadKeyFromFile("aeskey.txt")
	if err != nil {
		return err
	}
	return nil
}
