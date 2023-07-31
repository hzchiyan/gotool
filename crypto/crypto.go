package crypto

import (
	"errors"
	"fmt"
	"hash"
	"io"
	"os"
)

type Crypto struct {
	hash    hash.Hash
	sumByte []byte
}

func New(hash hash.Hash) *Crypto {
	return &Crypto{hash: hash}
}
func (c *Crypto) WithSumByte(sumByte []byte) *Crypto {
	c.sumByte = sumByte
	return c
}
func (c *Crypto) HashSum(data []byte) ([]byte, error) {
	if _, err := c.hash.Write(data); err != nil {
		return nil, errors.New(`hash.Write failed`)
	}
	return c.hash.Sum(c.sumByte), nil
}

// EncryptBytes 加密字节
func (c *Crypto) EncryptBytes(data []byte) (string, error) {
	sum, err := c.HashSum(data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sum), nil
}

// EncryptString 加密字符串
func (c *Crypto) EncryptString(data string) (string, error) {
	return c.EncryptBytes([]byte(data))
}

// EncryptFile 加密文件
func (c *Crypto) EncryptFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf(`os.Open failed for name "%s"`, path)
	}
	defer f.Close()
	h := c.hash
	_, err = io.Copy(h, f)
	if err != nil {
		return "", errors.New(`io.Copy failed`)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
