package exif

import (
	"github.com/evanoberholster/imagemeta"
	"github.com/evanoberholster/imagemeta/exif2"
	"github.com/evanoberholster/imagemeta/meta"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type ExifExtractorImpl struct {
	exif     *exif2.Exif
	filePath string
}

func (e *ExifExtractorImpl) Decode() error {
	if e.exif != nil {
		return nil
	}
	f, err := os.Open(e.filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	imageType := GetImageTypeByFileName(e.filePath)
	if imageType.IsJpegType() {
		exif, err := imagemeta.DecodeJPEG(f)
		if err != nil {
			return err
		}
		e.exif = &exif
	} else if imageType.IsHeicType() {
		exif, err := imagemeta.DecodeHeif(f)
		if err != nil {
			return err
		}
		e.exif = &exif
	} else if imageType.IsPngType() {
		exif, err := imagemeta.DecodePng(f)
		if err != nil {
			return err
		}
		e.exif = &exif
	} else {
		exif, err := imagemeta.Decode(f)
		if err != nil {
			return err
		}
		e.exif = &exif
	}

	if isDebug && e.exif != nil {
		logrus.Debug(e.exif.String())
	}
	return nil
}

func (e *ExifExtractorImpl) GetModel() string {
	if e.exif == nil {
		return ""
	}
	return e.exif.Model
}

func (e *ExifExtractorImpl) GetMake() string {
	if e.exif == nil {
		return ""
	}
	return e.exif.Make
}

func (e *ExifExtractorImpl) GetSoftware() string {
	if e.exif == nil {
		return ""
	}
	return e.exif.Software
}

func (e *ExifExtractorImpl) GetImageWidth() int {
	if e.exif == nil {
		return 0
	}
	return int(e.exif.ImageWidth)
}

func (e *ExifExtractorImpl) GetImageHeight() int {
	if e.exif == nil {
		return 0
	}
	return int(e.exif.ImageHeight)
}

func (e *ExifExtractorImpl) GetOrientation() meta.Orientation {
	if e.exif == nil {
		return 0
	}
	return e.exif.Orientation
}

func (e *ExifExtractorImpl) GetLatLong() (lat, long float64, err error) {
	if e.exif == nil {
		return 0, 0, nil
	}
	return e.exif.GPS.Latitude(), e.exif.GPS.Longitude(), nil
}

func (e *ExifExtractorImpl) GetTime() *time.Time {
	if e.exif == nil {
		return nil
	}
	t := e.exif.DateTimeOriginal()
	return &t
}
