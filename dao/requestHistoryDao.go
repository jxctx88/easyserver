package dao

import (
	"github.com/pkg/errors"
	"github.com/xingliuhua/easyserver/db"
	"github.com/xingliuhua/easyserver/model"
)

func GetHistoryById(id string) (err error, history model.RequestHistory) {
	for _, v := range db.HistoryList {
		if v.Id == id {
			return nil, v
		}
	}
	err = errors.New("not found")
	return
}
