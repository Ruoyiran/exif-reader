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
	extractor := exif.GetExifExtractor(filePath)
	err := extractor.Decode()
	if err != nil {
		logrus.Fatalf("%s", err.Error())
	}

	info := extractor.GetExifInfo()
	logrus.Debugf("Make: %s", info.Make)
	logrus.Debugf("Model: %s", info.Model)
	logrus.Debugf("Software: %s", info.Software)
	logrus.Debugf("Width: %d", info.ImageWidth)
	logrus.Debugf("Height: %d", info.ImageHeight)
	logrus.Debugf("Orientation: %s(%d)", info.Orientation.String(), info.Orientation)
	logrus.Debugf("GPS: %f,%f", info.GPSInfo.Latitude, info.GPSInfo.Longitude)
	logrus.Debugf("Time: %s", info.Time.String())
}
