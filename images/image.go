package images

import (
	"errors"
	"golang.org/x/image/webp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func ImageConfig(imagePath string) (config image.Config, err error) {
	var fopen *os.File
	fopen, err = os.Open(imagePath)
	if err != nil {
		return
	}
	defer fopen.Close()
	ext := strings.ToLower(filepath.Ext(imagePath))
	switch ext {
	case ".jpg", ".jpeg":
		config, err = jpeg.DecodeConfig(fopen)
	case ".gif":
		config, err = gif.DecodeConfig(fopen)
	case ".png":
		config, err = png.DecodeConfig(fopen)
	case ".webp":
		config, err = webp.DecodeConfig(fopen)
	default:
		err = errors.New("不支持的图片格式")
	}
	return
}
