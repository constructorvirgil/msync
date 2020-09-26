package common

import (
	"io/ioutil"
)

//用于读取aes加密中的key
func ReadKeyFromFile(path string) ([]byte, error){
	key, err := ioutil.ReadFile(path)
	if err != nil {
		return key, err
	}
	return key, nil
}
