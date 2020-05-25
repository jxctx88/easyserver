package util

import "github.com/xingliuhua/easyserver/model"

func Reverse(s []model.RequestHistory) []model.RequestHistory {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
