package exif

import (
	"github.com/evanoberholster/imagemeta"
	"github.com/evanoberholster/imagemeta/exif2"
	"github.com/sirupsen/logrus"
	"os"
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

func (e *ExifExtractorImpl) GetExifInfo() *ExifInfo {
	if e.exif == nil {
		return nil
	}
	return &ExifInfo{
		Make:        e.exif.Make,
		Model:       e.exif.Model,
		Software:    e.exif.Software,
		ImageWidth:  int(e.exif.ImageWidth),
		ImageHeight: int(e.exif.ImageHeight),
		Time:        e.exif.DateTimeOriginal(),
		Orientation: e.exif.Orientation,
		GPSInfo: GPSInfo{
			Latitude:  e.exif.GPS.Latitude(),
			Longitude: e.exif.GPS.Longitude(),
		},
	}
}
