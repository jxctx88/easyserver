package cache

import (
	"errors"
	"github.com/xingliuhua/easyserver/model"
)

var historyList = make([]model.RequestHistory, 0)
var historyKeyMap = make(map[string]interface{})

const HISTORY_MAX_COUNT = 50

func AddHistory2Cache(history model.RequestHistory) {
	if len(historyList) > HISTORY_MAX_COUNT {
		key := historyList[0].UrlPath
		delete(historyKeyMap, key)
		historyList = historyList[1:]
	}
	historyKeyMap[history.UrlPath] = nil
	historyList = append(historyList, history)
}
func GetHistoryById(historyId string) (err error, history model.RequestHistory) {
	for _, v := range historyList {
		if v.Id == historyId {
			return nil, v
		}
	}
	return errors.New("not found"), history
}
func GetAllHistoryList() []model.RequestHistory {
	return historyList
}
func HasConfigedRequest(key string) bool {
	_, b := responseInfoKeyMap[key]
	return b
}

var responseInfoList = make([]*model.ResponseInfo, 0)
var responseInfoKeyMap = make(map[string]interface{})

const responseInfo_MAX_COUNT = 100

func AddResponseInfo2Cache(info model.ResponseInfo) {

	if len(responseInfoList) > responseInfo_MAX_COUNT {
		key := responseInfoList[0].UrlPath
		delete(responseInfoKeyMap, key)
		responseInfoList = responseInfoList[1:]
	}
	responseInfoKeyMap[info.Key] = nil
	responseInfoList = append(responseInfoList, &info)
}
func UpdateResponseInfo2Cache(info model.ResponseInfo) {
	for _, v := range responseInfoList {
		if v.Id == info.Id {
			v.Description = info.Description
			v.ResponseText = info.ResponseText
			return
		}
	}
}
func GetResponseInfo(key string) (err error, info model.ResponseInfo) {
	for _, v := range responseInfoList {
		if v.Key == key {
			return nil, *v
		}
	}
	err = errors.New("not exist")
	return
}

func GetResponseInfoById(id string) (err error, info model.ResponseInfo) {
	for _, v := range responseInfoList {
		if v.Id == id {
			return nil, *v
		}
	}
	err = errors.New("not found")
	return
}
func GetAllResponseInfo() (list []model.ResponseInfo) {

	for _, v := range responseInfoList {
		list = append(list, *v)
	}
	return
}
func DeleteResponseInfo(id string) {
	for i, v := range responseInfoList {
		if v.Id == id {
			delete(responseInfoKeyMap, v.Key)
			responseInfoList = append(responseInfoList[:i], responseInfoList[i+1:]...)
			return
		}
	}

}
func ClearResponseInfo() {
	responseInfoList = responseInfoList[0:0]

}
