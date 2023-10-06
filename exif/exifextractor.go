package exif

import (
	"github.com/evanoberholster/imagemeta/meta"
	"time"
)

const (
	isDebug = false
)

type GPSInfo struct {
	Latitude  float64
	Longitude float64
}

type ExifInfo struct {
	Make        string
	Model       string
	Software    string
	ImageWidth  int
	ImageHeight int
	Time        time.Time
	Orientation meta.Orientation
	GPSInfo     GPSInfo
}

type ExifExtractor interface {
	Decode() error
	GetExifInfo() *ExifInfo
}

func GetExifExtractor(filePath string) ExifExtractor {
	return &ExifExtractorImpl{filePath: filePath}
}
