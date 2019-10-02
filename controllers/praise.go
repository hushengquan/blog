package controllers

import (
	"errors"
	"gostudy/newstudy/blog/models"
	"gostudy/newstudy/blog/syserror"
)

type PraiseController struct {
	BaseControllers
}

func (this *PraiseController) NextPrepare() {
	this.MustLogin()
}

// @router /:type/:key [post]
func (this *PraiseController) Praise() {
	ttype := this.Ctx.Input.Param(":type")
	key := this.Ctx.Input.Param(":key")

	table := "notes"
	switch ttype {
	case "message":
		table = "messages"
	case "note":
		table = "notes"
	default:
		this.Abort500(errors.New("未知类型"))
	}

	pcnt, err := models.UpdatePraise(table, key, int(this.User.ID))
	if err != nil {
		if e2, ok := err.(syserror.HasPraiseError); ok {
			this.Abort500(e2)
		} else {
			this.Abort500(syserror.New("点赞失败", err))
		}
	}

	this.JsonOkH("点赞成功", H{"praise":pcnt})
}