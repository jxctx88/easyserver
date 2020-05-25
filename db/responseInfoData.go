package db

import "github.com/xingliuhua/easyserver/model"

var HistoryList = make([]model.RequestHistory, 0)
var HistoryKeyMap = make(map[string]interface{})

const HISTORY_MAX_COUNT = 50

func AddHistory(history model.RequestHistory) {
	if len(HistoryList) > HISTORY_MAX_COUNT {
		key := HistoryList[0].UrlPath
		delete(HistoryKeyMap, key)
		HistoryList = HistoryList[1:]
	}
	HistoryKeyMap[history.UrlPath] = nil
	HistoryList = append(HistoryList, history)
}
