package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func testAes() {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("sfe023f_9fd&fwfl")
	result, err := AesEncrypt([]byte("polaris@studygolang"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := AesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	fmt.Printf("%x, %v\n", origData, len(origData))
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func testccc() {
	encBytes := []byte("el3%YC3*nrQQZ7aB")
	//ivBytes := []byte("1234567890abcdef")
	crypted := "f95259000216f9e2dc0c44a22838dbd2fdfcf26ac58a641fec028d005e957a1154f44fd2091cf4ad4eff00306ed18fcce9fcddc72ec7cc25e79d80d6715f0910b42b11c9aac6365a23886b5528e7471e"
	cryptedBytes, err := hex.DecodeString(crypted)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%x, %v\n", cryptedBytes, len(cryptedBytes))

	c, _ := aes.NewCipher(encBytes)
	decrypter := cipher.NewCBCDecrypter(c, cryptedBytes[:c.BlockSize()])

	data := make([]byte, len(cryptedBytes))
	copy(data, cryptedBytes)

	decrypter.CryptBlocks(data, data)
	fmt.Printf("%x, %v\n", data, len(data))

	//testAes()
}

func ExampleNewCBCDecrypter() {
	key := []byte("el3%YC3*nrQQZ7aB")
	ciphertext, _ := hex.DecodeString("f95259000216f9e2dc0c44a22838dbd2fdfcf26ac58a641fec028d005e957a1154f44fd2091cf4ad4eff00306ed18fcce9fcddc72ec7cc25e79d80d6715f0910b42b11c9aac6365a23886b5528e7471e")

	fmt.Printf("%x, %v\n", ciphertext, len(ciphertext))
	block, _ := aes.NewCipher(key)

	index := len(ciphertext)-aes.BlockSize
	iv := ciphertext[index:]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(ciphertext, ciphertext)

	fmt.Printf("%x, %v\n", ciphertext, len(ciphertext))
}

func main() {
	ExampleNewCBCDecrypter()
}
