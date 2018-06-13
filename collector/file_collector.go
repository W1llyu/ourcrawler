package collector

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type FileCollector struct {
	DirPath string
}

func NewFileCollector(path string) *FileCollector {
	return &FileCollector{
		DirPath: path,
	}
}

func (f *FileCollector) Process(resultItems ResultItems) error {
	if f.DirPath == "" {
		return errors.New("no director path")
	}
	now := time.Now()
	dirPath := fmt.Sprintf("%s/%d/%d/%d/%d", f.DirPath, now.Year(), now.Month(), now.Day(), now.Hour())
	_, err := os.Stat(dirPath)
	// 目录不存在
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0777)
		if err != nil {
			return err
		}
	}
	b, _ := json.Marshal(resultItems.Results)
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	var d = []byte(out.String())
	return ioutil.WriteFile(fmt.Sprintf("%s/%s", dirPath, generateFileName(resultItems)), d, 0666)
}

func generateFileName(resultItems ResultItems) string {
	now := time.Now()
	timeStamp := now.Unix()
	seed := fmt.Sprintf("ourcrawler%s%d", resultItems.Request.URL.String(), timeStamp)
	name := fmt.Sprintf("%x.json", md5.Sum([]byte(seed)))

	return name
}
