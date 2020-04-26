package routers

import (
	"b_blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.Router("/search", &controllers.MainController{}, "get,post:Search")
	beego.Router("/users", &controllers.UsersController{}, "get:GetAll")
	beego.Router("/users/:id", &controllers.UsersController{}, "get:GetOne")
	beego.Router("/users/delete", &controllers.UsersController{}, "delete:Delete")
	beego.Router("/users/:id/follow", &controllers.UsersController{}, "post:Follow")
	beego.Router("/users/:id/unfollow", &controllers.UsersController{}, "post:Unfollow")
	beego.Router("/microposts/:id/follow", &controllers.UsersController{}, "post:Follow")
	beego.Router("/microposts/:id/unfollow", &controllers.UsersController{}, "post:Unfollow")

	beego.Router("/microposts", &controllers.MicropostsController{}, "get:GetAll")
	beego.Router("/microposts/:id", &controllers.MicropostsController{}, "get:GetOne")
	beego.Router("/microposts/:id/edit", &controllers.MicropostsController{}, "get,post,put:Put")
	beego.Router("/microposts/create", &controllers.MicropostsController{}, "get,post:Post")
	beego.Router("/microposts/:id/delete", &controllers.MicropostsController{}, "delete:Delete")

	beego.Router("/my_microposts", &controllers.MicropostsController{}, "get:MyMicroposts")

	beego.Router("/setup", &controllers.UsersController{}, "get,post:Put")
	beego.Router("/reset_password", &controllers.UsersController{}, "get,post:ResetPassword")
	//beego.Router("/sign_up", &controllers.UsersController{}, "get,post:CreateSession")
	beego.Router("/sign_up", &controllers.UsersController{}, "get,post:Post")
	beego.Router("/sign_in", &controllers.SessionsController{}, "get,post:Post")
	beego.Router("/sign_out", &controllers.SessionsController{}, "*:Delete")
	//beego.AutoRouter()
}
