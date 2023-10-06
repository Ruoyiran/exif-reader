package exif

import (
	"fmt"
	"github.com/Ruoyiran/exif-reader/utils/file"
	"github.com/evanoberholster/imagemeta/meta"
	"time"
)

const (
	isDebug = true
)

type ExifExtractor interface {
	Decode() error
	GetMake() string
	GetModel() string
	GetSoftware() string
	GetImageWidth() int
	GetImageHeight() int

	// GetOrientation Valid EXIF values are 1 to 8. Values outside this range (for example, 0) are ignored.
	GetOrientation() meta.Orientation

	GetLatLong() (lat, long float64, err error)

	GetTime() *time.Time
}

func GetExifExtractor(filePath string) (ExifExtractor, error) {
	if !file.IsFileExist(filePath) {
		return nil, fmt.Errorf("file %s not exist", filePath)
	}
	return &ExifExtractorImpl{filePath: filePath}, nil
}
