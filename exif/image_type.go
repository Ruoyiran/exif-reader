package exif

import (
	"path"
	"strings"
)

type ImageType int

const (
	ImageTypeUnknown ImageType = -1
	ImageTypeJpeg    ImageType = 1
	ImageTypePng     ImageType = 2
	ImageTypeHeic    ImageType = 3
)

func GetImageTypeByFileName(fileName string) ImageType {
	ext := strings.ToLower(path.Ext(fileName))
	switch ext {
	case ".jpg", ".jpeg":
		return ImageTypeJpeg
	case ".png":
		return ImageTypePng
	case ".heic":
		return ImageTypeHeic
	}
	return ImageTypeUnknown
}

func (t ImageType) IsJpegType() bool {
	return t == ImageTypeJpeg
}

func (t ImageType) IsPngType() bool {
	return t == ImageTypePng
}

func (t ImageType) IsHeicType() bool {
	return t == ImageTypeHeic
}

func (t ImageType) IsUnknownType() bool {
	return t == ImageTypeUnknown
}
