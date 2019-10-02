package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/goquery-master"
	"github.com/jinzhu/gorm"
	"gostudy/newstudy/blog/models"
	"gostudy/newstudy/blog/syserror"
	"time"
)

type NoteController struct {
	BaseControllers
}

func (this *NoteController) NextPrepare() {
	this.MustLogin()
	if this.User.Role != 0 {
		this.Abort500(errors.New("权限不足"))
	}
}


// @router /new [get]
func (this *NoteController) NewPage() {
	this.Data["key"] = this.UUID()
	this.TplName = "note_new.html"
}

// @router /edit/:key [get]
func (this *NoteController) EditPage() {
	key := this.Ctx.Input.Param(":key")
	note, err := models.QueryNoteByKeyAndUserId(key, int(this.User.ID))
	if err != nil {
		this.Abort500(syserror.New("文章不存在", err))
	}
	this.Data["note"] = note
	this.Data["key"] = key
	this.TplName = "note_new.html"
}

// @router /save/:key [post]
func (ctx *NoteController) Save() {
	key := ctx.Ctx.Input.Param(":key")
	title := ctx.GetMustString("title", "标题不能为空！")
	content := ctx.GetMustString("content", "内容不能为空！")
	summary, _ := getSummary(content)
	note, err := models.QueryNoteByKeyAndUserId(key, int(ctx.User.ID))
	var n models.Note
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			ctx.Abort500(syserror.New("保存失败！", err))
		}
		n = models.Note{
			Key:     key,
			Summary: summary,
			Title:   title,
			Content: content,
			UserID:  int(ctx.User.ID),
		}
	} else {
		n = note
		n.Title = title
		n.Content = content
		n.Summary = summary
		n.UpdatedAt = time.Now()
	}

	if err := models.SaveNote(&n); err != nil {
		ctx.Abort500(syserror.New("保存失败！", err))
	}
	ctx.JsonOk("成功", fmt.Sprintf("/details/%s", key))
}

func getSummary(html string) (string, error) {
	var bufbytes bytes.Buffer
	_, err := bufbytes.Write([]byte(html))
	if err != nil {
		return "", err
	}
	doc, err := goquery.NewDocumentFromReader(&bufbytes)
	if err != nil {
		return "", err
	}
	htmlstr := doc.Find("body").Text()
	if len(htmlstr) > 600 {
		htmlstr = htmlstr[:600]
	}
	return htmlstr, err
}

// @router /del/:key [post]
func (this *NoteController) Del() {
	key := this.Ctx.Input.Param(":key")
	if err := models.DeleteNoteByUserIdAndKey(key, int(this.User.ID)); err != nil {
		this.Abort500(syserror.New("新增失败", err))
	}
	this.JsonOk("删除成功", "/")
}
