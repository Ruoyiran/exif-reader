package main

import (
	"github.com/Ruoyiran/exif-reader/exif"
	_ "github.com/Ruoyiran/exif-reader/logger"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please give filename as argument")
	}
	filePath := os.Args[1]
	extractor, err := exif.GetExifExtractor(filePath)
	if err != nil {
		logrus.Fatalf("%s", err.Error())
	}
	err = extractor.Decode()
	if err != nil {
		logrus.Fatalf("%s", err.Error())
	}

	logrus.Debugf("Make: %s", extractor.GetMake())
	logrus.Debugf("Model: %s", extractor.GetModel())
	logrus.Debugf("Software: %s", extractor.GetSoftware())
	logrus.Debugf("Width: %d", extractor.GetImageWidth())
	logrus.Debugf("Height: %d", extractor.GetImageHeight())
	logrus.Debugf("Orientation: %s(%d)", extractor.GetOrientation().String(), extractor.GetOrientation())
	lat, long, err := extractor.GetLatLong()
	if err == nil {
		logrus.Debugf("GPS: %f,%f", lat, long)
	}
	logrus.Debugf("Time: %s", extractor.GetTime().String())
}
