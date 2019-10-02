package controllers

import (
	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
	"gostudy/newstudy/blog/models"
	"errors"
	"gostudy/newstudy/blog/syserror"
)

const SEESION_USER_KEY = "SESSION_USER_KEY"

type BaseControllers struct {
	beego.Controller
	User models.User
	IsLogin bool
}

type NextPreparer interface {
	NextPrepare()
}

func (this *BaseControllers) Prepare() {
	this.Data["Path"] = this.Ctx.Request.RequestURI
	u, ok := this.GetSession(SEESION_USER_KEY).(models.User)
	this.IsLogin = false
	if ok {
		this.User = u
		this.IsLogin = true
		this.Data["User"] = this.User
	}
	this.Data["isLogin"] = this.IsLogin
	if a, ok := this.AppController.(NextPreparer); ok {
		a.NextPrepare()
	}
}

func (this *BaseControllers) Abort500(err error) {
	this.Data["error"] = err
	this.Abort("500")
}

func (this *BaseControllers) GetMustString(key, msg string) string{
	email := this.GetString(key)
	if len(email) == 0 {
		this.Abort500(errors.New(msg))
	}
	return email
}

func (this *BaseControllers) MustLogin() {
	if !this.IsLogin {
		this.Abort500(syserror.NoUserError{})
	}
}

type H map [string]interface{}

func (this *BaseControllers) JsonOk(msg, action string) {
	this.Data["json"] = H{
		"code": 0,
		"msg": msg,
		"action": action,
	}
	this.ServeJSON()
}

func (this *BaseControllers) JsonOkH(msg string, data H) {
	if data == nil {
		data = H{}
	}
	data["code"] = 0
	data["msg"] = msg
	this.Data["json"] = data
	this.ServeJSON()
}

func (this *BaseControllers) UUID() string {
	u, err := uuid.NewV4()
	if err != nil {
		this.Abort500(syserror.New("系统错误", err))
	}
	return u.String()
}
