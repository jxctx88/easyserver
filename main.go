package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/gin-gonic/gin"
	"github.com/xingliuhua/easyserver/db"
	"github.com/xingliuhua/easyserver/handler"
	"html/template"
	"net/http"
)

var runed bool

func main() {
	r := gin.Default()
	r.POST("/mock/responses", handler.AddResponseInfo)
	r.DELETE("/mock/responses/:id", handler.DelResponseInfo)
	r.PUT("/mock/responses", handler.UpdateResponseInfo)
	r.Static("/mock/static", "./static")
	r.SetFuncMap(template.FuncMap{
		"subLongText": handler.SubLongText,
		"formatTime":  handler.FormatTime,
	})
	r.LoadHTMLGlob("static/html/*.html")
	//r.LoadHTMLFiles("static/html/index.html", "static/html/add.html", "static/html/config.html")
	r.GET("/mock/index", handler.IndexHtml)
	r.GET("/mock/add", handler.AddOrUpdateResponseInfoHtml)
	r.GET("/mock/edit", handler.UpdateResponseInfoHtml)
	r.GET("/mock/config", handler.ConfigHtml)

	// not found router
	r.NoRoute(handler.HandleAll)

	var server *http.Server

	app := app.New()
	tvStatus := widget.NewLabel("------ wait... ------")
	w := app.NewWindow("Easy server")
	etPort := widget.NewEntry()
	etPort.Text = "8080"
	var btnSwitch *widget.Button
	btnSwitch = widget.NewButton("Run", func() {
		if runed {
			err := server.Close()
			if err != nil {
				return
			}
			runed = false
			btnSwitch.Text = "Run"
			tvStatus.SetText("------ wait... ------")

		} else {
			server = &http.Server{
				Handler: r,
			}
			server.Addr = ":" + etPort.Text
			go server.ListenAndServe()
			runed = true

			btnSwitch.Text = "Stop"
			tvStatus.SetText("------ servering... ------")
			// 读取文件加载已经配置的数据
			db.LoadData()
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
