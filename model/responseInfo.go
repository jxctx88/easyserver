package model

type ResponseInfo struct {
	Id           string
	Time         int64
	Method       string `form:"method,required"`
	Key          string `form:"key,required"`
	Description  string `form:"description,required"`
	UrlPath      string `form:"urlPath,required"`
	Params       map[string][]string
	ResponseText string `form:"responseText,required"`
}
