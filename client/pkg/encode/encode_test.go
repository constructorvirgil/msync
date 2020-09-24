package encode

import (
	"io/ioutil"
	"testing"
)

func TestEnFile(t *testing.T) {
	src, _ := ioutil.ReadFile("/Users/kiasma/WKspace/msync/client/bin/test3.txt")
	ans := GetMD5Encode(src)

	encodeFile, _ := EnFile(src, []byte("573392132@qq.com"))
	decodeFile, _ := DeFile(encodeFile, []byte("573392132@qq.com"))
	res := GetMD5Encode(decodeFile)

	t.Log("res:\t" + res)
	t.Log("ans:\t" + ans)

	if res != ans {
		t.Fail()
	}

}
