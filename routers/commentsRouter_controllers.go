package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetAbout",
            Router: `/about`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetComment",
            Router: `/comment/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetDetails",
            Router: `/details/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetMessage",
            Router: `/message`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetReg",
            Router: `/reg`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: `/user`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:MessageController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:MessageController"],
        beego.ControllerComments{
            Method: "Count",
            Router: `/count`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:MessageController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:MessageController"],
        beego.ControllerComments{
            Method: "NewMessage",
            Router: `/new/?:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:MessageController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:MessageController"],
        beego.ControllerComments{
            Method: "Query",
            Router: `/query`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:NoteController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:NoteController"],
        beego.ControllerComments{
            Method: "Del",
            Router: `/del/:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:NoteController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:NoteController"],
        beego.ControllerComments{
            Method: "EditPage",
            Router: `/edit/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:NoteController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:NoteController"],
        beego.ControllerComments{
            Method: "NewPage",
            Router: `/new`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:NoteController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:NoteController"],
        beego.ControllerComments{
            Method: "Save",
            Router: `/save/:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:PraiseController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:PraiseController"],
        beego.ControllerComments{
            Method: "Praise",
            Router: `/:type/:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:UserController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:UserController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:UserController"] = append(beego.GlobalControllerRouter["gostudy/newstudy/blog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Reg",
            Router: `/reg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
