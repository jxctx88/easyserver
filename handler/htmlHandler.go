package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xingliuhua/easyserver/cache"
	"net/http"
	"time"
)

func AddOrUpdateResponseInfoHtml(c *gin.Context) {
	historyId, b := c.GetQuery("id")
	if !b {
		WriterError(errors.New("not has id"), c)
		return
	}
	err, history := cache.GetHistoryById(historyId)
	if err != nil {
		WriterError(err, c)
		return
	}

	err, resInfo := cache.GetResponseInfo(history.Key)
	tmp := struct {
		HistoryId      string
		ResponseInfoId string
		Method         string
		Description    string
		UrlPath        string
		Params         map[string][]string
		ResponseText   string
	}{}
	tmp.HistoryId = history.Id
	tmp.Method = history.Method
	tmp.UrlPath = history.UrlPath
	tmp.Params = history.Params
	tmp.ResponseText = history.ResponseText
	if err == nil {
		// 已经配置过
		tmp.ResponseInfoId = resInfo.Id
		tmp.Description = resInfo.Description
	}

	c.HTML(http.StatusOK, "add.html", tmp)
}

func IndexHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", cache.GetAllHistoryList())
}
func UpdateResponseInfoHtml(c *gin.Context) {
	responseId, _ := c.GetQuery("id")
	_, info := cache.GetResponseInfoById(responseId)
	tmp := struct {
		ResponseInfoId string
		Method         string
		Description    string
		UrlPath        string
		Params         map[string][]string
		ResponseText   string
	}{}
	tmp.ResponseInfoId = info.Id
	tmp.Description = info.Description
	tmp.Method = info.Method
	tmp.UrlPath = info.UrlPath
	tmp.Params = info.Params
	tmp.ResponseText = info.ResponseText

	c.HTML(http.StatusOK, "edit.html", tmp)
}
func ConfigHtml(c *gin.Context) {

	c.HTML(http.StatusOK, "config.html", cache.GetAllResponseInfo())
}
func SubLongText(str string) string {
	if len(str) > 40 {
		return str[:40] + "..."
	}
	return str
}
func FormatTime(t int64) string {
	tm := time.Unix(t, 0)
	format := tm.Format("2006-01-02 15:04:05")
	return format
}
func FormatParams(p map[string][]string) (str string) {
	sp := ""
	for k, v := range p {
		bytes, _ := json.Marshal(v)
		str = str + sp + k + ":" + string(bytes)
		sp = "&"
	}
	return
}
