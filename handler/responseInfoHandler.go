package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xingliuhua/easyserver/dao"
	"github.com/xingliuhua/easyserver/db"
	"github.com/xingliuhua/easyserver/model"
	"github.com/xingliuhua/easyserver/util"
	"net/http"
	"time"
)

func AddResponseInfo(c *gin.Context) {
	responseId, hasResponseId := c.GetPostForm("responseInfoId")
	historyId, b := c.GetPostForm("historyId")
	description, b := c.GetPostForm("description")
	if !b {
		return
	}
	responseText, b := c.GetPostForm("responseText")
	if !b {
		return
	}

	err, history := dao.GetHistoryById(historyId)
	if err != nil {
		return
	}
	if hasResponseId && responseId != ""{
		// 修改
		_, responseInfo := dao.GetResponseInfoById(responseId)
		responseInfo.ResponseText = responseText
		responseInfo.Description = description
	} else {
		// 添加
		responseInfoNew := &model.ResponseInfo{
			Id:           util.GenUUID(),
			Method:       history.Method,
			Time:         time.Now().Unix(),
			UrlPath:      history.UrlPath,
			Params:       history.Params,
			ResponseText: responseText,
			Key:          history.Key,
			Description:  description,
		}
		db.AddResponseInfo(responseInfoNew)
		db.ResponseInfoKeyMap[history.Key] = nil
	}
	c.JSON(http.StatusOK,gin.H{"code":1,"msg":"success"})

}
func DelResponseInfo(c *gin.Context) {
	id:= c.Param("id")
	fmt.Println("要删除的id",id)
	err := dao.DeleteResponseInfo(id)
	if err!=nil{
		WriterError(err,c)
	}
	c.JSON(http.StatusOK,gin.H{"code":1,"msg":"success"})
}

func UpdateResponseInfo(c *gin.Context) {
	responseId, hasResponseId := c.GetPostForm("responseInfoId")
	if !hasResponseId {
		WriterError(errors.New("id not exist"), c)
		return
	}
	_, responseInfo := dao.GetResponseInfoById(responseId)
	description, b := c.GetPostForm("description")
	if !b {
		return
	}
	responseText, b := c.GetPostForm("responseText")
	if !b {
		return
	}

	responseInfo.ResponseText = responseText
	responseInfo.Description = description
	c.JSON(http.StatusOK,gin.H{"code":1,"msg":"success"})
}
