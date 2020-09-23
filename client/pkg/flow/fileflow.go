package flow

import (
	"client/pkg/common"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
)

//返回一个32位md5加密后的字符串
func GetMD5Encode(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func NewFile(fname string) *common.File {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil
	}

	temp := GetMD5Encode(data)
	return &common.File{FileId: []byte(temp), FileContent: data}
}
