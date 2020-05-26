package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xingliuhua/easyserver/cache"
	"github.com/xingliuhua/easyserver/dao"
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
		WriterError(errors.New("description missed"), c)
		return
	}
	responseText, b := c.GetPostForm("responseText")
	if !b {
		WriterError(errors.New("responseText missed"), c)
		return
	}

	err, history := cache.GetHistoryById(historyId)
	if err != nil {
		return
	}
	if hasResponseId && responseId != "" {
		// 修改
		_, responseInfo := cache.GetResponseInfoById(responseId)
		responseInfo.ResponseText = responseText
		responseInfo.Description = description
		// chache层修改
		cache.UpdateResponseInfo2Cache(responseInfo)
		// dao层修改
		r := dao.UpdateResponseInfo(responseInfo)
		if !r {
			WriterError(errors.New("edit failed"), c)
			return
		}
	} else {
		// 添加
		responseInfoNew := model.ResponseInfo{
			Id:           util.GenUUID(),
			Method:       history.Method,
			Time:         time.Now().Unix(),
			UrlPath:      history.UrlPath,
			Params:       history.Params,
			ResponseText: responseText,
			Key:          history.Key,
			Description:  description,
		}
		cache.AddResponseInfo2Cache(responseInfoNew)
		//  dao层添加
		r := dao.AddResponseInfo(responseInfoNew)
		if !r {
			WriterError(errors.New("add failed"), c)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "success"})

}
func DelResponseInfo(c *gin.Context) {
	id := c.Param("id")
	cache.DeleteResponseInfo(id)
	err := dao.DeleteResponseInfo(id)
	if err != nil {
		WriterError(err, c)
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "success"})
}

func UpdateResponseInfo(c *gin.Context) {
	responseId, hasResponseId := c.GetPostForm("responseInfoId")
	if !hasResponseId {
		WriterError(errors.New("id not exist"), c)
		return
	}
	_, responseInfo := cache.GetResponseInfoById(responseId)
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
	// cache层修改
	cache.UpdateResponseInfo2Cache(responseInfo)
	// dao层修改
	r := dao.UpdateResponseInfo(responseInfo)
	if !r {
		WriterError(errors.New("edit failed"), c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "success"})
}
