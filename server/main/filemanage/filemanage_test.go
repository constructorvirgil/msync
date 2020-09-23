package filemanage

import (
	"fmt"
	"github.com/constructorvirgil/msync/common"
	"io/ioutil"
	"testing"
	"time"
)

func TestAddAndFlush(t *testing.T) {
	var err error
	file := common.File{}
	file.FileId = []byte("test file id")
	file.FileContent,err = ioutil.ReadFile("C:\\Users\\Administrator\\Documents\\GitHub\\msync\\server\\msync_linux")
	if err != nil {
		fmt.Println("read file failure: ", err)
		return
	}

	bb, err := file.Pack()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	nFile := common.File{}
	err = nFile.UnPack(bb)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	AddAndFlush("id1", nFile.FileContent)
	time.Sleep(time.Second*3)  //等待写入文件
}
