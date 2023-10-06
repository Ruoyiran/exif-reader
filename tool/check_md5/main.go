package main

import (
	"flag"
	"fmt"
	_ "github.com/Ruoyiran/exif-reader/logger"
	"github.com/Ruoyiran/exif-reader/utils/file"
	"github.com/Ruoyiran/exif-reader/utils/hash"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	dir               string
	outputPath        string
	duplicateFilePath string
	maxWorkers        int
)

type Md5Path struct {
	Md5  string
	Path string
}

func init() {
	flag.StringVar(&dir, "dir", "", "input dir path")
	flag.StringVar(&outputPath, "out", "./files_md5.txt", "output file path")
	flag.StringVar(&duplicateFilePath, "dup", "./duplicate_files.txt", "output duplicate file path")
	flag.IntVar(&maxWorkers, "max_workers", 10, "max workers")
}

func main() {
	flag.Parse()
	if dir == "" {
		flag.Usage()
		return
	}
	dir, _ = filepath.Abs(dir)
	files := file.ListFilesRecursively(dir)
	if maxWorkers < 1 {
		maxWorkers = 1
	}
	totalFiles := len(files)
	if totalFiles < maxWorkers {
		maxWorkers = totalFiles
	}
	if maxWorkers < 1 {
		return
	}
	start := time.Now()
	logrus.Infof("start process, input dir: %s, output path: %s, max workers: %d", dir, outputPath, maxWorkers)
	fileBlocks := make([][]file.FileInfo, maxWorkers)
	md5Paths := make([][]Md5Path, maxWorkers)
	for i, f := range files {
		index := i % maxWorkers
		fileBlocks[index] = append(fileBlocks[index], f)
	}

	wg := sync.WaitGroup{}
	wg.Add(maxWorkers)
	totalProcessed := atomic.Int32{}
	for i := 0; i < maxWorkers; i++ {
		go func(id int) {
			count := int32(0)
			total := len(fileBlocks[id])
			for idx, f := range fileBlocks[id] {
				if f.IsDir() {
					continue
				}
				if f.FileInfo.Name() == ".DS_Store" {
					continue
				}
				md5, err := hash.MD5SumFromFile(f.Path)
				if err != nil {
					logrus.Errorf("calc file md5 error: %s, path: %s", err.Error(), f.Path)
				} else {
					count += 1
					relativePath := strings.ReplaceAll(f.Path, dir, "")
					if strings.HasPrefix(relativePath, "/") {
						relativePath = relativePath[1:]
					}
					md5Paths[id] = append(md5Paths[id], Md5Path{
						Md5:  md5,
						Path: relativePath,
					})
					logrus.Debugf("worker %d, processed: %d/%d", id, idx+1, total)
				}
			}
			logrus.Infof("worker %d finished, total files: %d, processed: %d", id, total, count)
			totalProcessed.Add(count)
			wg.Done()
		}(i)
	}
	wg.Wait()
	logrus.Infof("write file to %s", outputPath)
	_ = file.CreateDirectoryFromPath(outputPath)
	fp, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		logrus.Error("unable to create file %s, error: %s", outputPath, err.Error())
		return
	}
	defer fp.Close()

	fileMd5Map := make(map[string][]string)
	for _, items := range md5Paths {
		for _, item := range items {
			fileMd5Map[item.Md5] = append(fileMd5Map[item.Md5], item.Path)
			line := fmt.Sprintf("%s\t%s\n", item.Md5, item.Path)
			_, _ = fp.WriteString(line)
		}
	}

	var dupFile *os.File
	for md5, files := range fileMd5Map {
		if len(files) < 2 {
			continue
		}
		logrus.Infof("found duplicate file, md5: %s, file count: %d", md5, len(files))
		if dupFile == nil {
			_ = file.CreateDirectoryFromPath(duplicateFilePath)
			dupFile, err = os.OpenFile(duplicateFilePath, os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				logrus.Error("unable to create file %s, error: %s", duplicateFilePath, err.Error())
				return
			}
		}

		if dupFile != nil {
			line := fmt.Sprintf("%s\t%d\t%s\n", md5, len(files), strings.Join(files, "\t"))
			_, _ = dupFile.WriteString(line)
		}
	}

	if dupFile != nil {
		_ = dupFile.Close()
	}

	logrus.Infof("all done, total files: %d, total processed: %d, time cost: %.2f min", len(files), totalProcessed.Load(), time.Now().Sub(start).Seconds()/60.0)
}
