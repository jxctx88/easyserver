package model

import (
	"bytes"
	"crypto/md5"
	"fmt"
)

type RequestHistory struct {
	Id           string
	Method       string
	UrlPath      string
	Key          string
	Time         int64
	Params       map[string][]string
	ResponseText string
}

func (history *RequestHistory) GetKey() string {
	content := history.Method + history.UrlPath + getParamsStr(history.Params)
	has := md5.Sum([]byte(content))
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	history.Key = md5str
	return md5str
}
func getParamsStr(m map[string][]string) string {
	if m == nil {
		return ""
	}
	bufferString := bytes.NewBufferString("")
	for k, v := range m {
		value := ""
		for _, str := range v {
			value += str
		}
		bufferString.WriteString(k + value)
	}
	return bufferString.String()
}
