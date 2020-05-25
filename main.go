package main

import (
	"context"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/gin-gonic/gin"
	"github.com/xingliuhua/easyserver/db"
	"github.com/xingliuhua/easyserver/handler"
	"github.com/xingliuhua/easyserver/util"
	"html/template"
	"net/http"
)

var runed bool

func main() {
	r := gin.Default()
	r.POST("/mock/responses", handler.AddResponseInfo)
	r.DELETE("/mock/responses/:id", handler.DelResponseInfo)
	r.POST("/mock/editResponse", handler.UpdateResponseInfo)
	//r.PUT("/mock/responses", handler.UpdateResponseInfo)

	r.Static("/mock/static", "./static")
	r.SetFuncMap(template.FuncMap{
		"jiequ": handler.Jiequ,
		"formatTime": handler.FormatTime,
	})
	r.LoadHTMLGlob("static/html/*.html")
	//r.LoadHTMLFiles("static/html/index.html", "static/html/add.html", "static/html/config.html")
	r.GET("/mock/index", func(c *gin.Context) {
		tmp:=util.Reverse(db.HistoryList)
		c.HTML(http.StatusOK, "index.html", tmp)
	})
	r.GET("/mock/add", handler.AddOrUpdateResponseInfoHtml)
	r.GET("/mock/edit", handler.UpdateResponseInfoHtml)
	r.GET("/mock/config", handler.ConfigHtml)

	// 未知路由处理
	r.NoRoute(handler.HandleAll)

	server := &http.Server{
		Handler: r,
	}

	app := app.New()
	tvStatus := widget.NewLabel("wait...")
	w := app.NewWindow("Hello")
	etPort := widget.NewEntry()
	etPort.Text = "8080"
	var btnSwitch *widget.Button
	btnSwitch = widget.NewButton("Run", func() {
		if runed {
			server.Shutdown(context.Background())
			runed = false
			btnSwitch.Text = "Run"
			tvStatus.Text = "wait..."

		} else {
			server.Addr = ":" + etPort.Text
			go server.ListenAndServe()
			runed = true
			btnSwitch.Text = "Stop"
			tvStatus.Text = "servering..."

		}
	})
	box := widget.NewVBox(
		tvStatus,
		layout.NewSpacer(),
		etPort,
		layout.NewSpacer(),
		btnSwitch,
	)
	w.SetContent(box)

	w.ShowAndRun()
}
