package common

import (
	"testing"
)

func TestGzipEncode(t *testing.T) {
	src := "this is a test"
	zip,_ := GzipEncode([]byte(src))
	unzip,_ := GzipDecode(zip)
	//fmt.Println("src: ", []byte(src))
	//fmt.Println("zip: ", zip)
	//fmt.Println("unzip: ", unzip)
	if !bytesEqual([]byte(src), unzip) {
		t.Fail()
	}
}

func bytesEqual(a []byte, b []byte)bool{
	if len(a) != len(b) {
		return false
	}
	for i:=0;i<len(a);i++{
		if a[i] != b[i] {
			return false
		}
	}
	return true
}