package controllers

import (
	"b_blog/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

//  BaseController operations for Base
type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	//user           *models.Admin
	session     map[string]string
	userId      int64
	userName    string
	Status      bool
	IsAdmin     bool
	pageSize    int
	allowUrl    string
	noLayout    bool
	hexKey      string
	currentUser *models.Users
	//flash 			*beego.FlashData
}

//前期准备
func (c *BaseController) Prepare() {
	c.Data["signed"] = false
	//c.flash = beego.NewFlash()
	flash := beego.NewFlash()
	c.hexKey = beego.AppConfig.String("hexkey")
	controllerName, actionName := c.GetControllerAndAction()
	c.controllerName = controllerName
	c.actionName = actionName
	c.Layout = "layouts/layout.tpl"
	c.TplPrefix = strings.ToLower(strings.ReplaceAll(c.controllerName, "Controller", "/"))

	session, ok := c.Ctx.GetSecureCookie(c.hexKey, "session")
	if !ok || session == "null" {
		if !c.is("get") && c.controllerName != "SessionsController" &&
			!(c.controllerName == "UsersController" && c.actionName == "Post") {
			//c.flash.Set("info", "请登录")
			flash.Set("info", "请登录")
			flash.Store(&c.Controller)
			c.redirect(beego.URLFor("SessionsController.Post"))
		}
	} else {
		if err := json.Unmarshal([]byte(session), &(c.session)); err == nil {
			if userId, err := strconv.Atoi(c.session["userId"]); err == nil {
				c.userId = int64(userId)
			}
			c.userName, _ = c.session["userName"]
			c.Status, _ = strconv.ParseBool(c.session["Status"])
			c.IsAdmin, _ = strconv.ParseBool(c.session["IsAdmin"])
			if c.pageSize, err = strconv.Atoi(c.session["pageSize"]); err != nil {
				c.pageSize = 20
			}
			c.Data["signed"] = true
			c.currentUser, _ = models.GetUsersById(c.userId)
		} else {
			//c.flash.Set("danger", "cookie解析出错，请清空本地缓存重新登录")
			flash.Set("danger", "cookie解析出错，请清空本地缓存重新登录")
			flash.Store(&c.Controller)
			c.redirect(beego.URLFor("SessionsController.Post"))
		}
	}

	//fmt.Println(beego.BConfig.WebConfig.FlashName)
}

//后期准备
func (c *BaseController) Finish() {
	session, _ := json.Marshal(c.session)
	c.Ctx.SetSecureCookie(c.hexKey, "session", string(session))
}

// 关键字判断
func (c *BaseController) is(_method string) bool {
	return c.Ctx.Request.Method == strings.ToUpper(_method)
}

//获取用户IP地址
func (c *BaseController) getClientIp() string {
	s := strings.Split(c.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

// 重定向
func (c *BaseController) redirect302(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

func (c *BaseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}
