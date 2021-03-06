package common

import "encoding/json"

type File struct {
	FileId      []byte
	Path        string
	FileContent []byte
}

type FilePart struct {
	FileId       []byte `json:"fileId"`
	FileIndex    int    `json:"fileIndex"`
	FileMaxIndex int    `json:"fileMaxIndex"`
	Path         string `json:"path,ommited"`
	PartContent  []byte `json:"partContent"`
}

const (
	MaxFilePartSize = 2048
)

//把一个文件拆分成多个包
func (c *File) Pack() ([][]byte, error) {
	bb := [][]byte{}

	if len(c.FileContent) == 0 {
		return bb, nil
	}

	partSize := len(c.FileContent) / MaxFilePartSize
	if len(c.FileContent)%MaxFilePartSize != 0 {
		partSize++
	}

	parts := make([]FilePart, partSize) //分成了partSize个包

	for i := 0; i < len(parts); i++ {
		parts[i].FileId = c.FileId
		parts[i].FileIndex = i
		parts[i].FileMaxIndex = len(parts) - 1

		b := []byte{}
		for j := i * MaxFilePartSize; j < len(c.FileContent) && j < (i+1)*MaxFilePartSize; j++ {
			b = append(b, c.FileContent[j])
		}
		parts[i].PartContent = b
	}

	//只有第一个pack存储着文件的路径
	parts[0].Path = c.Path

	for _, p := range parts {
		jbyte, err := json.Marshal(p)
		if err != nil {
			return bb, err
		}
		bb = append(bb, jbyte)
	}

	return bb, nil
}

//从多个包还原回一个文件
func (c *File) UnPack(bb [][]byte) error {
	if len(bb) == 0 {
		return nil
	}
	parts := make([]FilePart, len(bb))
	for i := 0; i < len(bb); i++ {
		err := json.Unmarshal(bb[i], &parts[i])
		if err != nil {
			return err
		}
	}

	c.FileId = parts[0].FileId

	//还原文件路径信息
	c.Path = parts[0].Path

	for _, p := range parts {
		c.FileContent = append(c.FileContent, p.PartContent...)
	}
	return nil
}
