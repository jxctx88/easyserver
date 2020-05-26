package dao

import (
	"github.com/xingliuhua/easyserver/db"
	"github.com/xingliuhua/easyserver/model"
)

func DeleteResponseInfo(id string) (err error) {
	return db.DeleteResponseInfo(id)
}
func AddResponseInfo(info model.ResponseInfo) bool {
	return db.AddResponseInfo(info)
}
func UpdateResponseInfo(info model.ResponseInfo) bool {
	return db.UpdateResponseInfo(info)
}
