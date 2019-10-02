package routers

import (
	"github.com/astaxie/beego"
	"gostudy/newstudy/blog/controllers"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(
		&controllers.IndexController{},
		&controllers.UserController{},
	)
	beego.AddNamespace(
		beego.NewNamespace(
			"note",
			beego.NSInclude(&controllers.NoteController{}),
		),
		beego.NewNamespace(
			"message",
			beego.NSInclude(&controllers.MessageController{}),
		),
		beego.NewNamespace(
			"praise",
			beego.NSInclude(&controllers.PraiseController{}),
		),
	)
}
