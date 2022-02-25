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

type queryEventForm struct {
	Database string `form:"database" validate:"omitempty,max=254"`
	Start    int64  `form:"start" validate:"required,min=1"`
	End      int64  `form:"end" validate:"required,min=1"`
	Page     int    `form:"page" validate:"required,min=1"`
	PageSize int    `form:"size" validate:"required,min=1,max=50"`
}

func GetAccessEvents(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryEventForm
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

	srv := service.Event{
		Start:    form.Start,
		End:      form.End,
		Page:     form.Page,
		PageSize: form.PageSize,
	}
	events, err := srv.GetAccessList()
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

type queryEventInfoForm struct {
	Database string `form:"database" validate:"omitempty,max=254"`
	Start    int64  `form:"start" validate:"required,min=1"`
	End      int64  `form:"end" validate:"required,min=1"`
}

func GetAccessEventsInfo(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryEventInfoForm
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

	srv := service.Event{
		Database: form.Database,
		Start:    form.Start,
		End:      form.End,
	}

	info, err := srv.GetAccessInfo()
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

func GetRuleEvents(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryEventForm
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

	srv := service.Event{
		Start:    form.Start,
		End:      form.End,
		Page:     form.Page,
		PageSize: form.PageSize,
	}
	events, err := srv.GetRuleList()
	if err != nil {
		log.Logger.Error("RuleEvent", zap.String("get", err.Error()))
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

func GetRuleEventsInfo(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryEventInfoForm
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

	srv := service.Event{
		Database: form.Database,
		Start:    form.Start,
		End:      form.End,
	}

	info, err := srv.GetRuleInfo()
	if err != nil {
		log.Logger.Error("RuleEvent", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.AccessEventGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["data"] = info
	appG.Response(httpCode, errCode, "", data)
}
