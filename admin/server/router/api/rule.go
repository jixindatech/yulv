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
	Action   int    `json:"action" validate:"required,min=1,max=3"`
	User     string `json:"user" validate:"omitempty,max=254"`
	IP       string `json:"ip" validate:"omitempty,ip,max=254"`
	Database string `json:"database" validate:"omitempty,max=254"`
	Type     string `json:"type" validate:"omitempty,max=254"`
	Sql      string `json:"sql" validate:"omitempty,max=254"`
	Match    string `json:"match" validate:"omitempty,max=254"`
	Pattern  string `json:"pattern" validate:"omitempty,max=254"`
	Rows     int    `json:"rows" validate:"omitempty,min=1"`

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

	if len(form.Match) > 0 {
		if len(form.Pattern) == 0 {
			httpCode = e.InvalidParams
			errCode = e.ERROR
			appG.Response(httpCode, errCode, "pattern needs content", nil)
			return
		}
	}

	ruleSrv := service.Rule{
		Name:     form.Name,
		Action:   form.Action,
		User:     form.User,
		IP:       form.IP,
		Database: form.Database,
		Type:     form.Type,
		Sql:      form.Sql,
		Match:    form.Match,
		Pattern:  form.Pattern,
		Rows:     form.Rows,

		Remark: form.Remark,
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

	if len(form.Match) > 0 {
		if len(form.Pattern) == 0 {
			httpCode = e.InvalidParams
			errCode = e.ERROR
			appG.Response(httpCode, errCode, "pattern needs content", nil)
			return
		}
	}

	ruleSrv := service.Rule{
		ID:       formId.ID,
		Name:     form.Name,
		Action:   form.Action,
		User:     form.User,
		IP:       form.IP,
		Database: form.Database,
		Type:     form.Type,
		Sql:      form.Sql,
		Match:    form.Match,
		Pattern:  form.Pattern,
		Rows:     form.Rows,

		Remark: form.Remark,
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
	Action   int    `form:"action" validate:"omitempty,min=1,max=3"`
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
		Action:   form.Action,
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

func DistributeRule(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		httpCode = http.StatusOK
		errCode  = e.SUCCESS
	)
	srv := service.Rule{}
	err := srv.Distribute()
	if err != nil {
		log.Logger.Error("ip", zap.String("post", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.IPDistributeFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	appG.Response(httpCode, errCode, "", nil)
}

type queryRuleTestForm struct {
	Sql string `form:"sql" validate:"required,max=254"`
}

func GetRuleTest(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		httpCode = http.StatusOK
		form     queryRuleTestForm
		errCode  = e.SUCCESS
	)

	err := app.BindAndValid(c, &form)
	if err != nil {
		httpCode = e.InvalidParams
		errCode = e.ERROR
		appG.Response(httpCode, errCode, err.Error(), nil)
		return
	}

	srv := service.Rule{}
	res, err := srv.Test(form.Sql)
	if err != nil {
		log.Logger.Error("Rule", zap.String("get", err.Error()))
		httpCode = http.StatusInternalServerError
		errCode = e.RuleGetFailed
		appG.Response(httpCode, errCode, "", nil)
		return
	}

	data := make(map[string]interface{})
	data["item"] = res
	appG.Response(httpCode, errCode, "", data)
}
