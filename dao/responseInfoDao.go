package dao

import (
	"github.com/pkg/errors"
	"github.com/xingliuhua/easyserver/db"
	"github.com/xingliuhua/easyserver/model"
)

func GetResponseInfo(key string) (err error, info *model.ResponseInfo) {
	for _, v := range db.ResponseInfoList {
		if v.Key == key {
			return nil, v
		}
	}
	err = errors.New("not exist")
	return
}

func GetResponseInfoById(id string) (err error, info *model.ResponseInfo) {
	for _, v := range db.ResponseInfoList {
		if v.Id == id {
			return nil, v
		}
	}
	err = errors.New("not found")
	return
}
func DeleteResponseInfo(id string) (err error) {
	for i, v := range db.ResponseInfoList {
		if v.Id == id {
			db.ResponseInfoList = append(db.ResponseInfoList[:i], db.ResponseInfoList[i+1:]...)
			delete(db.ResponseInfoKeyMap, v.Key)
			return
		}
	}
	err = errors.New("data not found")
	return
}
