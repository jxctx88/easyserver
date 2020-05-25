package db

import "github.com/xingliuhua/easyserver/model"

var ResponseInfoList = make([]*model.ResponseInfo, 0)
var ResponseInfoKeyMap = make(map[string]interface{})

const ResponseInfo_MAX_COUNT = 100

func AddResponseInfo(ResponseInfo *model.ResponseInfo) {
	if len(ResponseInfoList) > ResponseInfo_MAX_COUNT {
		key := ResponseInfoList[0].UrlPath
		delete(ResponseInfoKeyMap, key)
		ResponseInfoList = ResponseInfoList[1:]
	}
	ResponseInfoKeyMap[ResponseInfo.UrlPath] = nil
	ResponseInfoList = append(ResponseInfoList, ResponseInfo)
}
