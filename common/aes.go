package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

//error code
var (
	ErrCodeAESEncodeFail = errors.New("aes encode error")
	ErrCodeAESDecodeFail = errors.New("aes decode error")
)

//返回一个32位md5加密后的字符串
func GetMD5Encode(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func AESEncode(data []byte, key []byte) (encodePack []byte, err error) {
	//避免panic程序崩溃
	defer func(){
		if recErr:=recover();recErr!=nil{
			encodePack = nil
			err = ErrCodeAESEncodeFail
		}
	}()

	encodePack, err = aesEncrypt(data, key)
	if err != nil {
		return nil, err
	}

	return encodePack, nil
}

func AESDecode(data []byte, key []byte) (decodePack []byte, err error) {
	//避免panic程序崩溃
	defer func(){
		if recErr:=recover();recErr!=nil{
			decodePack = nil
			err = ErrCodeAESDecodeFail
		}
	}()

	decodePack, err = aesDecrypt(data, key)
	if err != nil {
		return nil, err
	}

	return decodePack, nil
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
