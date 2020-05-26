package db

import (
	"encoding/json"
	"github.com/xingliuhua/easyserver/model"
	"os"
)

const ResponseInfo_MAX_COUNT = 3

func AddResponseInfo(responseInfo model.ResponseInfo) bool {
	if !DataDirExists() {
		CreateDataDir()
	}
	// 如果文件夹下文件超过限制就删除
	if hasFileNumber() >= ResponseInfo_MAX_COUNT {
		deleteLastFile()
	}
	filePath := DATA_DIR_PATH + "/" + responseInfo.Id + ".json"
	file, err := os.Create(filePath)
	if err != nil {
		return false
	}
	defer file.Close()
	bytes, _ := json.Marshal(responseInfo)
	file.Write(bytes)
	return true
}
func UpdateResponseInfo(responseInfo model.ResponseInfo) bool {

	if !DataDirExists() {
		CreateDataDir()
	}

	filePath := DATA_DIR_PATH + "/" + responseInfo.Id + ".json"
	file, err := os.Create(filePath)
	if err != nil {
		return false
	}
	defer file.Close()
	bytes, _ := json.Marshal(responseInfo)
	file.Write(bytes)
	return true
}
func DeleteResponseInfo(id string) (err error) {
	if !DataDirExists() {
		return
	}
	filePath := DATA_DIR_PATH + "/" + id + ".json"
	err = os.Remove(filePath)
	return
}
