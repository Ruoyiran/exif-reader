package exif

import (
	"github.com/evanoberholster/imagemeta/meta"
	"time"
)

const (
	isDebug = false
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

	GetTime() time.Time
}

func GetExifExtractor(filePath string) ExifExtractor {
	return &ExifExtractorImpl{filePath: filePath}
}
