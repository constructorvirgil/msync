package flow

import (
	"fmt"
	"os"
)

func file2Byte(fname string) ([]byte, error) {
	//Open File
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	data := make([]byte, stats.Size())
	fLen, err := file.Read(data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Read file %s len: %d\n", fname, fLen)
	return data, nil
}
