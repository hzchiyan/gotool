package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

var IVDefaultValue = "I Love Go Crypto!"

// Encrypt is alias of EncryptCBC.
func Encrypt(plainText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	return EncryptCBC(plainText, key, iv...)
}

// Decrypt is alias of DecryptCBC.
func Decrypt(cipherText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	return DecryptCBC(cipherText, key, iv...)
}

// EncryptCBC encrypts `plainText` using CBC mode.
// Note that the key must be 16/24/32 bit length.
// The parameter `iv` initialization vector is unnecessary.
func EncryptCBC(plainText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plainText = pKCS5Padding(plainText, blockSize)
	ivValue := ([]byte)(nil)
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = []byte(IVDefaultValue)
	}
	blockMode := cipher.NewCBCEncrypter(block, ivValue)
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)

	return cipherText, nil
}

// DecryptCBC decrypts `cipherText` using CBC mode.
// Note that the key must be 16/24/32 bit length.
// The parameter `iv` initialization vector is unnecessary.
func DecryptCBC(cipherText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {

		return nil, err
	}
	blockSize := block.BlockSize()
	if len(cipherText) < blockSize {
		return nil, errors.New("cipherText too short")
	}
	ivValue := ([]byte)(nil)
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = []byte(IVDefaultValue)
	}
	if len(cipherText)%blockSize != 0 {
		return nil, errors.New("cipherText is not a multiple of the block size")
	}
	blockModel := cipher.NewCBCDecrypter(block, ivValue)
	plainText := make([]byte, len(cipherText))
	blockModel.CryptBlocks(plainText, cipherText)
	plainText, e := pKCS5UnPadding(plainText, blockSize)
	if e != nil {
		return nil, e
	}
	return plainText, nil
}

func pKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func pKCS5UnPadding(src []byte, blockSize int) ([]byte, error) {
	length := len(src)
	if blockSize <= 0 {

		return nil, errors.New("invalid blocklen")
	}

	if length%blockSize != 0 || length == 0 {
		return nil, errors.New("invalid data len")
	}

	unpadding := int(src[length-1])
	if unpadding > blockSize || unpadding == 0 {
		return nil, errors.New("invalid padding")
	}

	padding := src[length-unpadding:]
	for i := 0; i < unpadding; i++ {
		if padding[i] != byte(unpadding) {
			return nil, errors.New("invalid padding")
		}
	}

	return src[:(length - unpadding)], nil
}

// EncryptCFB encrypts `plainText` using CFB mode.
// Note that the key must be 16/24/32 bit length.
// The parameter `iv` initialization vector is unnecessary.
func EncryptCFB(plainText []byte, key []byte, padding *int, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf(`aes.NewCipher failed for key "%s"`, key)
	}
	blockSize := block.BlockSize()
	plainText, *padding = zeroPadding(plainText, blockSize)
	ivValue := ([]byte)(nil)
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = []byte(IVDefaultValue)
	}
	stream := cipher.NewCFBEncrypter(block, ivValue)
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)
	return cipherText, nil
}

// DecryptCFB decrypts `plainText` using CFB mode.
// Note that the key must be 16/24/32 bit length.
// The parameter `iv` initialization vector is unnecessary.
func DecryptCFB(cipherText []byte, key []byte, unPadding int, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf(`aes.NewCipher failed for key "%s"`, key)
	}
	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("cipherText too short")
	}
	ivValue := ([]byte)(nil)
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = []byte(IVDefaultValue)
	}
	stream := cipher.NewCFBDecrypter(block, ivValue)
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)
	plainText = zeroUnPadding(plainText, unPadding)
	return plainText, nil
}

func zeroPadding(cipherText []byte, blockSize int) ([]byte, int) {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(0)}, padding)
	return append(cipherText, padText...), padding
}

func zeroUnPadding(plaintext []byte, unPadding int) []byte {
	length := len(plaintext)
	return plaintext[:(length - unPadding)]
}
