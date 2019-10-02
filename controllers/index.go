package controllers

import (
	"gostudy/newstudy/blog/models"
	"gostudy/newstudy/blog/syserror"
)

type IndexController struct {
	BaseControllers
}

// @router / [get]
func (this *IndexController) Get() {
	limit := 10
	page, err := this.GetInt("page", 1)
	if err != nil || page <= 0{
		page = 1
	}

	title := this.GetString("title")
	notes, err := models.QueryNotesByPage(page, limit, title)
	if err != nil {
		this.Abort500(err)
	}
	this.Data["notes"] = notes

	count, err := models.QueryNotesCount(title)
	if err != nil {
		this.Abort500(err)
	}
	totpage := count / limit
	if count % limit != 0 {
		totpage++
	}
	this.Data["totpage"] = totpage
	this.Data["page"] = page
	this.Data["title"] = title

	this.Data["Path"] = this.Ctx.Request.RequestURI
	this.TplName = "index.html"
}


// @router /message [get]
func (this *IndexController) GetMessage() {
	this.TplName = "message.html"
}

// @router /about [get]
func (this *IndexController) GetAbout() {
	this.TplName = "about.html"
}

// @router /user [get]
func (this *IndexController) GetUser() {
	this.TplName = "user.html"
}

// @router /reg [get]
func (this *IndexController) GetReg() {
	this.TplName = "reg.html"
}

// @router /comment/:key
func (this *IndexController) GetComment() {
	key := this.Ctx.Input.Param(":key")
	note, err := models.QueryNoteByKey(key)
	if err != nil {
		this.Abort500(syserror.New("文章不存在", err))
	}
	this.Data["note"] = note
	this.TplName = "comment.html"
}

// @router /details/:key [get]
func (this *IndexController) GetDetails() {
	key := this.Ctx.Input.Param(":key")
	note, err := models.QueryNoteByKey(key)
	if err != nil {
		this.Abort500(syserror.New("文章不存在", err))
	}
	this.Data["note"] = note

	ms, err := models.QueryMessagesByNotKey(key)
	if err != nil {
		this.Abort500(syserror.New("文章不存在", err))
	}
	this.Data["message"] = ms

	this.TplName = "details.html"
}

