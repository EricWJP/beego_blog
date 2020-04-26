package controllers

import (
	"b_blog/models"
	_ "encoding/hex"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

// SessionsController operations for Sessions
type SessionsController struct {
	BaseController
}

// URLMapping ...
func (c *SessionsController) URLMapping() {
	c.Mapping("CreateSession", c.Post)
}

// Post ...
// @Title Create
// @Description create Sessions
// @Param	body		body 	models.Sessions	true		"body for Sessions content"
// @Success 201 {object} models.Sessions
// @Failure 403 body is empty
// @router / [post]
func (c *SessionsController) Post() {
	flash := beego.NewFlash()
	if c.userId > 0 {
		c.redirect(beego.URLFor("MainController.Get"))
	}
	if c.is("post") {
		errorMsg := ""
		session := map[string]string{}
		beego.ReadFromRequest(&c.Controller)
		//formData := c.Ctx.Request.Form
		userName := strings.TrimSpace(c.GetString("Name"))
		userPassword := strings.TrimSpace(c.GetString("Password"))
		c.Data["Name"] = userName
		if userName != "" && userPassword != "" {
			user, err := models.GetUsersByName(userName)
			if err != nil {
				errorMsg = "用户不存在！！！"
			} else if !user.Authenticate(userPassword) {
				errorMsg = "帐号或密码错误！！！"
			} else if !user.Status {
				errorMsg = "帐号已被禁用！！！"
			} else {
				session["userId"] = strconv.Itoa(int(user.Id))
				session["userName"] = user.Name
				session["Status"] = strconv.FormatBool(user.Status)
				session["IsAdmin"] = strconv.FormatBool(user.IsAdmin)
				session["pageSize"] = "20"
				c.session = session
				if userId, err := strconv.Atoi(c.session["userId"]); err == nil {
					c.userId = int64(userId)
				}
				c.userName, _ = c.session["userName"]
				c.Status, _ = strconv.ParseBool(c.session["Status"])
				c.IsAdmin, _ = strconv.ParseBool(c.session["IsAdmin"])
				c.pageSize = 20
				//c.flash.Success("Sign In Successfully!")
				flash.Success("Sign In Successfully!")
				flash.Store(&c.Controller)
				c.Finish()
				c.redirect(beego.URLFor("MainController.Get"))
			}
		} else {
			errorMsg = "帐号或密码不能空！！!"
		}
		//c.flash.Set("danger", errorMsg)
		flash.Set("danger", errorMsg)
		//c.flash.Store(&c.Controller)
		//flash.Store(&c.Controller)
	}
	c.TplName = "create.tpl"
}

func (c *SessionsController) Delete() {
	c.session = nil
	c.userId = 0
	c.Finish()
	c.redirect(beego.URLFor("MainController.Get"))
}
