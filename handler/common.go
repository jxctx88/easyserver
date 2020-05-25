package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xingliuhua/easyserver/dao"
	"github.com/xingliuhua/easyserver/db"
	"github.com/xingliuhua/easyserver/model"
	"github.com/xingliuhua/easyserver/util"
	"net/http"
	"time"
)

func HandleAll(c *gin.Context) {
	url := c.Request.URL
	m := make(map[string][]string)
	//s, i := c.GetPostForm("name")
	//fmt.Println(s,i)
	c.Request.ParseForm()
	c.Request.ParseMultipartForm(10000)
	switch c.Request.Method {
	case http.MethodGet:
		for k, v := range c.Request.Form {
			m[k] = v
		}
	case http.MethodPost:
		for k, v := range c.Request.PostForm {
			m[k] = v
		}
	default:
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "not support method"})
		return
	}

	history := model.RequestHistory{
		Id:      util.GenUUID(),
		Method:  c.Request.Method,
		UrlPath: url.Path,
		Params:  m,
		Time:    time.Now().Unix(),
	}
	history.GetKey() // 生成key

	// 判断是否配置的有该请求，如果有则设置响应，没有则返回未配置
	_, b := db.ResponseInfoKeyMap[history.Key]
	if !b {
		hs := gin.H{"code": 0, "msg": "not set data"}
		bytes, _ := json.Marshal(hs)
		history.ResponseText = string(bytes)
		fmt.Println(history)
		db.AddHistory(history)
		c.JSON(http.StatusOK, hs)
		return
	}
	_, info := dao.GetResponseInfo(history.Key)
	history.ResponseText = info.ResponseText
	db.AddHistory(history)

	c.String(http.StatusOK, info.ResponseText)
}

func WriterError(err error, ctx *gin.Context) {
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})

	}
}
