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

type dbForm struct {
	Name     string `json:"name" validate:"required,max=254"`
	User     string `json:"user" validate:"required,max=254"`
	Password string `json:"password" validate:"required,max=254"`
	Host     string `json:"host" validate:"required,max=254"`
	Port     uint32 `json:"port" validate:"required,min=1,max=65535"`

	Remark string `json:"remark" validate:"max=254"`
}

func AddDB(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     dbForm
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

	dbSrv := service.DB{
		Name:     form.Name,
		User:     form.User,
		Password: form.Password,
		Host:     form.Host,
		Port:     form.Port,
		Remark:   form.Remark,
	}

	err = dbSrv.Save()
	if err != nil {
		log.Logger.Error("DB", zap.String("add", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.DBAddFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func UpdateDB(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     dbForm
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)

	err := app.BindUriAndValid(c, &formId)
	if err != nil {
		httpCode = e.InvalidParams
		errCode = e.ERROR
		appG.Response(httpCode, errCode, err.Error(), nil)
		return
	}

	err = app.BindAndValid(c, &form)
	if err != nil {
		httpCode = e.InvalidParams
		errCode = e.ERROR
		appG.Response(httpCode, errCode, err.Error(), nil)
		return
	}

	dbSrv := service.DB{
		ID:       formId.ID,
		Name:     form.Name,
		User:     form.User,
		Password: form.Password,
		Host:     form.Host,
		Port:     form.Port,
		Remark:   form.Remark,
	}

	err = dbSrv.Save()
	if err != nil {
		log.Logger.Error("DB", zap.String("put", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.DBUpdateFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func DeleteDB(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		httpCode = http.StatusOK
		formId   app.IDForm
		errCode  = e.SUCCESS
	)

	err := app.BindUriAndValid(c, &formId)
	if err != nil {
		httpCode = e.InvalidParams
		errCode = e.ERROR
		appG.Response(httpCode, errCode, err.Error(), nil)
		return
	}

	dbSrv := service.DB{
		ID: formId.ID,
	}
	err = dbSrv.Delete()
	if err != nil {
		log.Logger.Error("DB", zap.String("delete", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.DBDeleteFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func GetDB(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		httpCode = http.StatusOK
		formId   app.IDForm
		errCode  = e.SUCCESS
	)

	err := app.BindUriAndValid(c, &formId)
	if err != nil {
		httpCode = e.InvalidParams
		errCode = e.ERROR
		appG.Response(httpCode, errCode, err.Error(), nil)
		return
	}

	dbSrv := service.DB{
		ID: formId.ID,
	}
	dbItem, err := dbSrv.Get()
	if err != nil {
		log.Logger.Error("DB", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.DBGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["item"] = dbItem
	appG.Response(httpCode, errCode, "", data)
}

type queryDBForm struct {
	Name     string `form:"name" validate:"max=254"`
	Page     int    `form:"page" validate:"required,min=1,max=50"`
	PageSize int    `form:"size" validate:"required,min=1"`
}

func GetDBs(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryDBForm
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

	dbSrv := service.DB{
		Name:     form.Name,
		Page:     form.Page,
		PageSize: form.PageSize,
	}
	ips, count, err := dbSrv.GetList()
	if err != nil {
		log.Logger.Error("DB", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.DBGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = ips
	data["total"] = count
	appG.Response(httpCode, errCode, "", data)
}
