package app

import (
	"admin/core/rbac"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net"
	"regexp"
	"strings"
)

type IDForm struct {
	ID uint `uri:"id" validate:"required,gte=1"`
}

var validate *validator.Validate

var phoneReg *regexp.Regexp

func ValidatePhone(fl validator.FieldLevel) bool {
	var err error
	if phoneReg == nil {
		phoneReg, err = regexp.Compile("^1[3456789]\\d{9}$")
		if err != nil {
			return false
		}
	}

	phone := fl.Field().String()
	return phoneReg.Match([]byte(phone))
}

func ValidateRole(fl validator.FieldLevel) bool {
	role := fl.Field().String()
	if rbac.ROLES != nil && rbac.ROLES[role] != false {
		return true
	}

	return false
}

func ValidateIPs(fl validator.FieldLevel) bool {
	ips := fl.Field().Interface().([]string)
	for _, ip := range ips {
		cidr := strings.Contains(ip, "/")
		if cidr {
			_, _, err := net.ParseCIDR(ip)
			if err != nil {
				return false
			}
		} else {
			address := net.ParseIP(ip)
			if address == nil {
				return false
			}
		}
	}

	return true
}

func ValidateIP(fl validator.FieldLevel) bool {
	ip := fl.Field().String()
	address := net.ParseIP(ip)
	if address == nil {
		return false
	}

	return true
}

func SetupValidate() error {
	var err error

	validate = validator.New()
	err = validate.RegisterValidation("phone", ValidatePhone)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("role", ValidateRole)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("ips", ValidateIPs)
	if err != nil {
		return err
	}

	return nil
}

func BindAndValid(c *gin.Context, form interface{}) error {
	err := c.Bind(form)
	if err != nil {
		return err
	}

	if validate != nil {
		return validate.Struct(form)
	}

	return fmt.Errorf("%s", "invalid validate")
}

func BindUriAndValid(c *gin.Context, form interface{}) error {
	err := c.BindUri(form)
	if err != nil {
		return err
	}

	if validate != nil {
		return validate.Struct(form)
	}

	return fmt.Errorf("%s", "invalid validate")
}
