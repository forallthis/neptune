package cryptoutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// 填充key
func paddingkey(key []byte) []byte {
	var buffer bytes.Buffer
	buffer.Write(key)
	for i := len(key); i < 16; i++ {
		buffer.WriteByte('0')
	}

	return buffer.Bytes()
}

// PKCS7Padding 填充
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding 去除填充
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesEncryptByte  aes byte加密
func AesEncryptByte(origData, key []byte) ([]byte, error) {
	key = paddingkey(key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// AesDecryptByte aes byte 解密
func AesDecryptByte(crypted, key []byte) ([]byte, error) {
	key = paddingkey(key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

// AesEncrypt 加密
func AesEncrypt(origData, key string) (string, error) {
	result, err := AesEncryptByte([]byte(origData), []byte(key))
	crypted := base64.URLEncoding.EncodeToString(result)
	return crypted, err
}

// AesDecrypt 解密
func AesDecrypt(encoded, key string) (string, error) {
	crypted, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	origData, err := AesDecryptByte(crypted, []byte(key))
	return string(origData), err
}
