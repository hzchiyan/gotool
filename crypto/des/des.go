package des

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
	"fmt"
)

const (
	NOPADDING = iota
	PKCS5PADDING
)

// EncryptECB encrypts `plainText` using ECB mode.
func EncryptECB(plainText []byte, key []byte, padding int) ([]byte, error) {
	text, err := textPadding(plainText, padding)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, len(text))
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf(`des.NewCipher failed for key "%s"`, key)
	}

	blockSize := block.BlockSize()
	for i, count := 0, len(text)/blockSize; i < count; i++ {
		begin, end := i*blockSize, i*blockSize+blockSize
		block.Encrypt(cipherText[begin:end], text[begin:end])
	}
	return cipherText, nil
}

// DecryptECB decrypts `cipherText` using ECB mode.
func DecryptECB(cipherText []byte, key []byte, padding int) ([]byte, error) {
	text := make([]byte, len(cipherText))
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(`des.NewCipher failed for key "%s"`, key))
	}

	blockSize := block.BlockSize()
	for i, count := 0, len(text)/blockSize; i < count; i++ {
		begin, end := i*blockSize, i*blockSize+blockSize
		block.Decrypt(text[begin:end], cipherText[begin:end])
	}

	plainText, err := textUnPadding(text, padding)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

// EncryptECBTriple encrypts `plainText` using TripleDES and ECB mode.
// The length of the `key` should be either 16 or 24 bytes.
func EncryptECBTriple(plainText []byte, key []byte, padding int) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 {
		return nil, errors.New("key length error")
	}

	text, err := textPadding(plainText, padding)
	if err != nil {
		return nil, err
	}

	var newKey []byte
	if len(key) == 16 {
		newKey = append([]byte{}, key...)
		newKey = append(newKey, key[:8]...)
	} else {
		newKey = append([]byte{}, key...)
	}

	block, err := des.NewTripleDESCipher(newKey)
	if err != nil {
		return nil, fmt.Errorf(`des.NewTripleDESCipher failed for key "%s"`, newKey)
	}

	blockSize := block.BlockSize()
	cipherText := make([]byte, len(text))
	for i, count := 0, len(text)/blockSize; i < count; i++ {
		begin, end := i*blockSize, i*blockSize+blockSize
		block.Encrypt(cipherText[begin:end], text[begin:end])
	}
	return cipherText, nil
}

// DecryptECBTriple decrypts `cipherText` using TripleDES and ECB mode.
// The length of the `key` should be either 16 or 24 bytes.
func DecryptECBTriple(cipherText []byte, key []byte, padding int) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 {
		return nil, errors.New("key length error")
	}

	var newKey []byte
	if len(key) == 16 {
		newKey = append([]byte{}, key...)
		newKey = append(newKey, key[:8]...)
	} else {
		newKey = append([]byte{}, key...)
	}

	block, err := des.NewTripleDESCipher(newKey)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(`des.NewTripleDESCipher failed for key "%s"`, newKey))
	}

	blockSize := block.BlockSize()
	text := make([]byte, len(cipherText))
	for i, count := 0, len(text)/blockSize; i < count; i++ {
		begin, end := i*blockSize, i*blockSize+blockSize
		block.Decrypt(text[begin:end], cipherText[begin:end])
	}

	plainText, err := textUnPadding(text, padding)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

// EncryptCBC encrypts `plainText` using CBC mode.
func EncryptCBC(plainText []byte, key []byte, iv []byte, padding int) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		err = errors.New(fmt.Sprintf(`des.NewCipher failed for key "%s"`, key))
		return nil, err
	}

	if len(iv) != block.BlockSize() {
		return nil, errors.New("invalid iv length")
	}

	text, err := textPadding(plainText, padding)
	if err != nil {
		return nil, err
	}
	cipherText := make([]byte, len(text))

	encryptor := cipher.NewCBCEncrypter(block, iv)
	encryptor.CryptBlocks(cipherText, text)

	return cipherText, nil
}

// DecryptCBC decrypts `cipherText` using CBC mode.
func DecryptCBC(cipherText []byte, key []byte, iv []byte, padding int) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		err = errors.New(fmt.Sprintf(`des.NewCipher failed for key "%s"`, key))
		return nil, err
	}

	if len(iv) != block.BlockSize() {
		return nil, errors.New("iv length invalid")
	}

	text := make([]byte, len(cipherText))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(text, cipherText)

	plainText, err := textUnPadding(text, padding)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

// EncryptCBCTriple encrypts `plainText` using TripleDES and CBC mode.
func EncryptCBCTriple(plainText []byte, key []byte, iv []byte, padding int) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 {
		return nil, errors.New("key length invalid")
	}

	var newKey []byte
	if len(key) == 16 {
		newKey = append([]byte{}, key...)
		newKey = append(newKey, key[:8]...)
	} else {
		newKey = append([]byte{}, key...)
	}

	block, err := des.NewTripleDESCipher(newKey)
	if err != nil {
		err = errors.New(fmt.Sprintf(`des.NewTripleDESCipher failed for key "%s"`, newKey))
		return nil, err
	}

	if len(iv) != block.BlockSize() {
		return nil, errors.New("invalid iv length")
	}

	text, err := textPadding(plainText, padding)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, len(text))
	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(cipherText, text)

	return cipherText, nil
}

// DecryptCBCTriple decrypts `cipherText` using TripleDES and CBC mode.
func DecryptCBCTriple(cipherText []byte, key []byte, iv []byte, padding int) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 {
		return nil, errors.New("key length invalid")
	}

	var newKey []byte
	if len(key) == 16 {
		newKey = append([]byte{}, key...)
		newKey = append(newKey, key[:8]...)
	} else {
		newKey = append([]byte{}, key...)
	}

	block, err := des.NewTripleDESCipher(newKey)
	if err != nil {
		err = errors.New(fmt.Sprintf(`des.NewTripleDESCipher failed for key "%s"`, newKey))
		return nil, err
	}

	if len(iv) != block.BlockSize() {
		return nil, errors.New("invalid iv length")
	}

	text := make([]byte, len(cipherText))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(text, cipherText)

	plainText, err := textUnPadding(text, padding)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

func paddingPKCS5(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, padText...)
}

func unPaddingPKCS5(text []byte) []byte {
	length := len(text)
	padText := int(text[length-1])
	return text[:(length - padText)]
}

func textPadding(text []byte, padding int) ([]byte, error) {
	switch padding {
	case NOPADDING:
		if len(text)%8 != 0 {
			return nil, errors.New("invalid text length")
		}
	case PKCS5PADDING:
		return paddingPKCS5(text, 8), nil
	default:
		return nil, errors.New(fmt.Sprintf(`unsupported padding type "%d"`, padding))
	}
	return text, nil
}

func textUnPadding(text []byte, padding int) ([]byte, error) {
	switch padding {
	case NOPADDING:
		if len(text)%8 != 0 {
			return nil, errors.New("invalid text length")
		}
	case PKCS5PADDING:
		return unPaddingPKCS5(text), nil
	default:
		return nil, errors.New(fmt.Sprintf(`unsupported padding type "%d"`, padding))
	}
	return text, nil
}
