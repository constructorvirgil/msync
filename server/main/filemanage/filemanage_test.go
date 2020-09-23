package filemanage

import (
	"fmt"
	"testing"
)

func TestAddAndFlush(t *testing.T) {
	Add("id1", []byte("this is id1 content"))
	Add("id2", []byte("this is id2 content"))
	Add("id3", []byte("this is id3 content"))
	fmt.Println(globalMap)
	AddAndFlush("id1",[]byte("follow id1 content"))
	AddAndFlush("id2",[]byte("follow id2 content"))
}
