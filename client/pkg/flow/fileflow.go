package flow

import (
	"client/pkg/common"
	"crypto/md5"
	"io/ioutil"
)

func NewFile(fname string) *common.File {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil
	}

	temp := md5.Sum(data)
	sum := make([]byte, 16)
	copy(sum, temp[:])

	return &common.File{FileId: sum, FileContent: data}
}
