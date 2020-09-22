package common

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
)

func TestGzipEncode(t *testing.T) {
	src := "this is a test"
	zip,_ := GzipEncode([]byte(src))
	unzip,_ := GzipDecode(zip)
	//fmt.Println("src: ", []byte(src))
	//fmt.Println("zip: ", zip)
	//fmt.Println("unzip: ", unzip)
	if !bytes.Equal([]byte(src), unzip) {
		t.Fail()
	}
}

func TestFile_Pack(t *testing.T) {
	file := File{}
	file.FileId = 0
	file.FileContent = genRandomBytes(1027)

	bb, err := file.Pack()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	nFile := File{}
	err = nFile.UnPack(bb)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	//fmt.Println(file.FileContent)
	//fmt.Println(nFile.FileContent)

	if !bytes.Equal(file.FileContent, nFile.FileContent) {
		fmt.Println(err)
		t.Fail()
		return
	}
}

func genRandomBytes(n int)[]byte{
	b := []byte{}
	for i:=0;i<n;i++{
		b = append(b, byte(rand.Int()))
	}
	return b
}