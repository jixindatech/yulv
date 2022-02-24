package api

import (
	"admin/core/log"
	"admin/server/pkg/app"
	"admin/server/pkg/e"
	"admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type queryAccessEventForm struct {
	Database string `form:"database" validate:"omitempty,max=254"`
	Start    int64  `form:"start" validate:"required"`
	End      int64  `form:"end" validate:"required"`
	Page     int    `form:"page" validate:"required,min=1"`
	PageSize int    `form:"size" validate:"required,min=1,max=50"`
}

func GetAccessEvents(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryAccessEventForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindAndValid(c, &form)
	if err != nil {
		httpCode = e.InvalidParams
		errCode = e.ERROR
		appG.Response(httpCode, errCode, err.Error(), nil)
		return
	}

	srv := service.AccessEvent{
		Start:    form.Start,
		End:      form.End,
		Page:     form.Page,
		PageSize: form.PageSize,
	}
	events, err := srv.GetList()
	if err != nil {
		log.Logger.Error("AccessEvent", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.AccessEventGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = events["data"]
	data["total"] = events["count"]
	appG.Response(httpCode, errCode, "", data)
}

type queryAccessEventInfoForm struct {
	Database string `form:"database" validate:"omitempty,max=254"`
	Start    int64  `form:"start" validate:"required"`
	End      int64  `form:"end" validate:"required"`
}

func GetAccessEventsInfo(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryAccessEventInfoForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindAndValid(c, &form)
	if err != nil {
		httpCode = e.InvalidParams
		errCode = e.ERROR
		appG.Response(httpCode, errCode, err.Error(), nil)
		return
	}

	srv := service.AccessEvent{
		Database: form.Database,
		Start:    form.Start,
		End:      form.End,
	}

	info, err := srv.GetInfo()
	if err != nil {
		log.Logger.Error("AccessEvent", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.AccessEventGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["data"] = info
	appG.Response(httpCode, errCode, "", data)
}
