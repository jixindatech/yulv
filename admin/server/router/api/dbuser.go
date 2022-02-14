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

type dbUserForm struct {
	Name     string `json:"name" validate:"required,max=254"`
	Username string `json:"username" validate:"required,max=254"`
	Password string `json:"password" validate:"required,max=254"`

	Remark string `json:"remark" validate:"max=254"`
}

func AddDBUser(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     dbUserForm
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

	dbUserSrv := service.DBUser{
		Name:     form.Name,
		Username: form.Username,
		Password: form.Password,
		Remark:   form.Remark,
	}

	err = dbUserSrv.Save()
	if err != nil {
		log.Logger.Error("DBUser", zap.String("add", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.DBUserAddFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func UpdateDBUser(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     dbUserForm
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

	dbUserSrv := service.DBUser{
		ID:       formId.ID,
		Name:     form.Name,
		Username: form.Username,
		Password: form.Password,
		Remark:   form.Remark,
	}

	err = dbUserSrv.Save()
	if err != nil {
		log.Logger.Error("DBUser", zap.String("put", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.DBUserUpdateFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func DeleteDBUser(c *gin.Context) {
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

	dbUserSrv := service.DBUser{
		ID: formId.ID,
	}
	err = dbUserSrv.Delete()
	if err != nil {
		log.Logger.Error("DBUser", zap.String("delete", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.DBDeleteFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func GetDBUser(c *gin.Context) {
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

	dbSrv := service.DBUser{
		ID: formId.ID,
	}
	dbItem, err := dbSrv.Get()
	if err != nil {
		log.Logger.Error("DBUser", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.DBGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["item"] = dbItem
	appG.Response(httpCode, errCode, "", data)
}

type queryDBUserForm struct {
	Name     string `form:"name" validate:"max=254"`
	Page     int    `form:"page" validate:"required,min=1,max=50"`
	PageSize int    `form:"size" validate:"required,min=1"`
}

func GetDBUsers(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryDBUserForm
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

	dbSrv := service.DBUser{
		Name:     form.Name,
		Page:     form.Page,
		PageSize: form.PageSize,
	}
	ips, count, err := dbSrv.GetList()
	if err != nil {
		log.Logger.Error("DBUser", zap.String("get", err.Error()))
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

type userDBForm struct {
	DB []uint `form:"db" validate:"required"`
}

func UpdateDBUserDB(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     userDBForm
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

	dbUserSrv := service.DBUser{
		ID:  formId.ID,
		Dbs: form.DB,
	}

	err = dbUserSrv.UpdateUserDB()
	if err != nil {
		log.Logger.Error("DBUser", zap.String("put", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.DBUserUpdateFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func DistributeDBUser(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)
	srv := service.DBUser{}
	err := srv.Distribute()
	if err != nil {
		log.Logger.Error("DBUser", zap.String("post", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.DBUserDistributeFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}
