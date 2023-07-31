package str

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
)

// GbkToUtf8 gbk->utf8
func GbkToUtf8(str string) string {
	reader := transform.NewReader(
		bytes.NewReader([]byte(str)),
		simplifiedchinese.GBK.NewDecoder(),
	)
	d, _ := io.ReadAll(reader)
	return string(d)
}

// GbkToUtf8ByChinese gbk->utf8 存在中国文字
func GbkToUtf8ByChinese(str string) string {
	if IsChineseChar(str) {
		return str
	}
	return GbkToUtf8(str)
}
