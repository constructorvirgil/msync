package filemanage

import (
	"io/ioutil"
)

type FilePart struct {
	Index    int
	MaxIndex int
	Id       string
	Part     []byte
}

type File struct {
	FileName string
	Content  []byte
}

var fileCh = make(chan File)
var partCh = make(chan FilePart)

func Init() {
	go addWorker(partCh, fileCh) //添加部分文件内容到map，只能有一个worker
	for i := 0; i < 16; i++ {    //写文件，可以有多个worker
		go writeWorker(fileCh)
	}
}

//负责添加部分文件内容到map
func addWorker(partCh chan FilePart, fileCh chan File) {
	m := map[string]*File{}
	for v := range (partCh) {
		if m[v.Id] == nil { //在map中无记录
			m[v.Id] = &File {
				FileName: "file-" + v.Id,
				Content:  v.Part,
			}
		} else {              //在map有记录
			if v.Index == 0 { //Index等于0说明文件是重新传输的，而map有记录说明上一次传输没有正常结束，丢弃上一次的结果，然后重新开始
				m[v.Id] = &File{
					FileName: "file-" + v.Id,
					Content:  v.Part,
				}
			}else {
				m[v.Id].Content = append(m[v.Id].Content, v.Part...)
			}
		}
		if v.Index == v.MaxIndex { //文件的最后一部分已到达
			fileCh <- *m[v.Id]
			m[v.Id] = nil
		}
	}

}

//负责写入文件到磁盘
func writeWorker(fileCh chan File) {
	for v := range (fileCh) {
		_ = ioutil.WriteFile(v.FileName, v.Content, 0644)
	}
}

//给外部调用的接口，添加部分文件内容
func Add(part FilePart) {
	partCh <- part
}
