package filemanage

import (
	"fmt"
	"io/ioutil"
	"sync"
)

type data struct {
	content []byte
}

var globalMap = map[string]*data{}
var globalMutex = sync.Mutex{}

func Add(id string, cont []byte){
	globalMutex.Lock()
	defer globalMutex.Unlock()

	if globalMap[id] == nil {
		globalMap[id] = &data{}
	}
	globalMap[id].content = append(globalMap[id].content, cont...)
}

func AddAndFlush(id string, cont []byte) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	if globalMap[id] == nil {
		globalMap[id] = &data{}
	}
	globalMap[id].content = append(globalMap[id].content, cont...)
	ch := make(chan int)
	go func(ch chan int) {  //把该数据写入文件
		_ = <-ch  //等待切片添加完毕
		_ = ioutil.WriteFile(fmt.Sprintf("filemanage-id%s", id), globalMap[id].content, 0644)
		Clear(id)
	}(ch)
	ch <- 0
}

func Clear(id string) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	globalMap[id] = nil
}
