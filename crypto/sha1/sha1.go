package sha1

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

// EncryptBytes encrypts `data` using MD5 algorithms.
func EncryptBytes(data []byte) (encrypt string, err error) {
	h := sha1.New()
	if _, err = h.Write(data); err != nil {
		return "", errors.New(`hash.Write failed`)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
func MustEncryptBytes(data []byte) string {
	result, err := EncryptBytes(data)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptString encrypts string `data` using MD5 algorithms.
func EncryptString(data string) (encrypt string, err error) {
	return EncryptBytes([]byte(data))
}

// MustEncryptString encrypts string `data` using MD5 algorithms.
// It panics if any error occurs.
func MustEncryptString(data string) string {
	result, err := EncryptString(data)
	if err != nil {
		panic(err)
	}
	return result
}

// EncryptFile encrypts file content of `path` using SHA1 algorithms.
func EncryptFile(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", errors.New(fmt.Sprintf(`os.Open failed for name "%s"`, path))
	}
	defer f.Close()
	h := sha1.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", errors.New(`io.Copy failed`)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// MustEncryptFile encrypts file content of `path` using SHA1 algorithms.
// It panics if any error occurs.
func MustEncryptFile(path string) string {
	result, err := EncryptFile(path)
	if err != nil {
		panic(err)
	}
	return result
}
