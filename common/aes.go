package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

//返回一个32位md5加密后的字符串
func GetMD5Encode(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func EnPack(data [][]byte, key []byte) ([][]byte, error) {
	var enPack [][]byte

	for _, i := range data {
		encodei, err := aesEncrypt(i, key)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		enPack = append(enPack, encodei)
	}

	return enPack, nil
}

func DePack(data [][]byte, key []byte) ([][]byte, error) {
	var dePack [][]byte

	for _, i := range data {
		decodei, err := aesDecrypt(i, key)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		dePack = append(dePack, decodei)
	}

	return dePack, nil
}

func aesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = pkcs5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func aesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pkcs5UnPadding(origData)
	return origData, nil
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
