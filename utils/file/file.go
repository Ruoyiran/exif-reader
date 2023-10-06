package file

import (
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
)

type FileInfo struct {
	os.FileInfo
	Path string
}

func IsFileExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func CreateDirectory(dir string) error {
	if IsFileExist(dir) != true {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			logrus.Errorf("failed to create directory, dir: %s", dir)
			return err
		}
	}
	return nil
}

func CreateDirectoryFromPath(filepath string) error {
	dirPath, _ := path.Split(filepath)
	return CreateDirectory(dirPath)
}
func ListFilesRecursively(path string) []FileInfo {
	fileInfos := make([]FileInfo, 0)
	_ = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logrus.Error(err)
			return nil
		}
		if !info.IsDir() {
			fileInfos = append(fileInfos, FileInfo{
				FileInfo: info,
				Path:     path,
			})
		}
		return nil
	})

	return fileInfos
}
