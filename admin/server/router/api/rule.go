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

type ruleForm struct {
	Name     string `json:"name" validate:"required,max=254"`
	User     string `json:"user" validate:"omitempty,max=254"`
	IP       string `json:"ip" validate:"omitempty,ip,max=254"`
	Database string `json:"database" validate:"omitempty,max=254"`
	Sql      string `json:"sql" validate:"required,max=254"`

	Remark string `json:"remark" validate:"omitempty,max=254"`
}

func AddRule(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     ruleForm
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

	ruleSrv := service.Rule{
		Name:     form.Name,
		User:     form.User,
		IP:       form.IP,
		Database: form.Database,
		Sql:      form.Sql,
		Remark:   form.Remark,
	}

	err = ruleSrv.Save()
	if err != nil {
		log.Logger.Error("Rule", zap.String("add", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleAddFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func UpdateRule(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		formId   app.IDForm
		form     ruleForm
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

	ruleSrv := service.Rule{
		ID:       formId.ID,
		Name:     form.Name,
		User:     form.User,
		IP:       form.IP,
		Database: form.Database,
		Sql:      form.Sql,
		Remark:   form.Remark,
	}

	err = ruleSrv.Save()
	if err != nil {
		log.Logger.Error("Rule", zap.String("put", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleUpdateFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func DeleteRule(c *gin.Context) {
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

	ruleSrv := service.Rule{
		ID: formId.ID,
	}
	err = ruleSrv.Delete()
	if err != nil {
		log.Logger.Error("Rule", zap.String("delete", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleDeleteFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

func GetRule(c *gin.Context) {
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

	ruleSrv := service.Rule{
		ID: formId.ID,
	}
	ruleItem, err := ruleSrv.Get()
	if err != nil {
		log.Logger.Error("Rule", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["item"] = ruleItem
	appG.Response(httpCode, errCode, "", data)
}

type queryRuleForm struct {
	Name     string `form:"name" validate:"max=254"`
	Page     int    `form:"page" validate:"required,min=1,max=50"`
	PageSize int    `form:"size" validate:"required,min=1"`
}

func GetRules(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		form     queryRuleForm
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

	ruleSrv := service.Rule{
		Name:     form.Name,
		Page:     form.Page,
		PageSize: form.PageSize,
	}
	ips, count, err := ruleSrv.GetList()
	if err != nil {
		log.Logger.Error("Rule", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = ips
	data["total"] = count
	appG.Response(httpCode, errCode, "", data)
}
