package db

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/xingliuhua/easyserver/cache"
	"github.com/xingliuhua/easyserver/model"
	"io/ioutil"
	"os"
)

const DATA_DIR_PATH = "./easyserverdata"

func LoadData() (err error) {
	cache.ClearResponseInfo()
	exists := DataDirExists()
	if !exists {
		err = errors.New("file not exists")
		return
	}
	dir, err := ioutil.ReadDir(DATA_DIR_PATH)
	if err != nil {
		return err
	}

	for _, header := range dir {
		if header.IsDir() {
			continue
		}
		file, err := os.Open(DATA_DIR_PATH + string(os.PathSeparator) + header.Name())
		if err != nil {
			continue
		}
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			continue
		}
		responseInfo := &model.ResponseInfo{}
		err = json.Unmarshal(bytes, responseInfo)
		if err != nil {
			continue
		}
		cache.AddResponseInfo2Cache(*responseInfo)
	}
	// 按时间排序
	cache.SortResponseInfoList()
	return
}
func DataDirExists() bool {

	exists := Exists(DATA_DIR_PATH)
	return exists
}
func CreateDataDir() bool {

	err := os.Mkdir(DATA_DIR_PATH, 0711)
	if err != nil {
		return false
	}
	return true
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
func hasFileNumber() int {
	if !DataDirExists() {
		return 0
	}
	dir, err := ioutil.ReadDir(DATA_DIR_PATH)
	if err != nil {
		return 0
	}
	return len(dir)
}
func deleteLastFile() {
	if !DataDirExists() {
		return
	}
	dir, err := ioutil.ReadDir(DATA_DIR_PATH)
	if err != nil {
		return
	}
	var lastHeader os.FileInfo = nil
	for _, header := range dir {
		if lastHeader == nil || header.ModTime().Unix() < lastHeader.ModTime().Unix() {
			lastHeader = header
		}
	}
	if lastHeader != nil {
		err := os.Remove(DATA_DIR_PATH + string(os.PathSeparator) + lastHeader.Name())
		fmt.Println("删除文件结果：", err)
	}
}
