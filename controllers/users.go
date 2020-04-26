package controllers

import (
	"b_blog/models"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//  UsersController operations for Users
type UsersController struct {
	BaseController
}

// URLMapping ...
func (c *UsersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Users
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 201 {int} models.Users
// @Failure 403 body is empty
// @router / [post]
func (c *UsersController) Post() {
	//flash := beego.NewFlash()
	flash := beego.ReadFromRequest(&c.Controller)
	if c.is("post") {
		var v models.Users
		formData := c.Ctx.Request.Form
		if formData["Password"][0] != formData["PasswordConfirmation"][0] {
			for k, v := range formData {
				c.Data[k] = v[0]
			}
			//c.flash.Set("danger", "两次密码不一致！！！")
			flash.Set("danger", "两次密码不一致！！！")
			c.Data["flash"] = flash.Data
			//flash.Store(&c.Controller)
			//c.Data["flash"] = c.flash.Data
			c.TplName = "create.tpl"
			c.Render()
			c.StopRun()
		}
		v.Name = formData["Name"][0]
		if len(formData["Phone"]) != 0 {
			v.Phone = formData["Phone"][0]
		}
		if len(formData["Gender"]) != 0 {
			v.Gender = formData["Gender"][0]
		}
		if len(formData["Phone"]) != 0 {
			v.Phone = formData["Phone"][0]
		}
		v.Email = formData["Email"][0]
		v.Password = models.GetEncryptPassword(formData["Password"][0])
		v.Status = true
		v.IsAdmin = false
		v.FollowedCount = 0
		v.Comment = formData["Comment"][0]
		currentTime := time.Now()
		v.UpdatedAt, v.CreatedAt = currentTime, currentTime
		if err := c.ParseForm(&v); err != nil {
			c.Data["ParseFormErr"] = "数据解析到结构体错误"
		} else {
			user := v
			valid := validation.Validation{} //创建验证数据对象
			//验证用户名不能为空且最小长度为6
			// Message 是自定义消息
			valid.Required(user.Name, "Name").Message("用户名不能为空")
			//valid.MinSize(user.Name, 6,"Name").Message("用户名最短为6位" )
			valid.Required(user.Password, "Password").Message("密码不能为空")
			valid.MinSize(user.Password, 6, "Pwd").Message("密码最短为6位")
			//valid.Numeric(user.Age, "Age").Message("年龄只能为数字" )
			//valid.Length(user.IdCard, 18,"IdCard").Message("身份证格式非法" )
			valid.Required(user.Email, "Email").Message("邮箱不能为空")
			valid.Email(user.Email, "Email").Message("邮箱格式非法")
			valid.Required(user.Phone, "Phone").Message("手机不能为空")
			valid.Mobile(user.Phone, "Phone").Message("手机格式非法")
			if valid.HasErrors() {
				// 如果有错误信息，证明验证没通过
				// 打印错误信息
				for _, err := range valid.Errors {
					flash.Set("danger", flash.Data["danger"], err.Key+err.Message+"！")
				}
				re := regexp.MustCompile(`(%!!\(string=)|(\)!\(string=)|(\)\(EXTRA string=)|\)`)
				flash.Set("danger", re.ReplaceAllLiteralString(flash.Data["danger"], ""))
				c.Data["flash"] = flash.Data
				for k, v := range formData {
					c.Data[k] = v[0]
				}
				c.TplName = "create.tpl"
			} else {
				if _, err := models.AddUsers(&v); err == nil {
					c.Ctx.Output.SetStatus(200)
					flash.Success("Sign Up Successfully!")
					flash.Store(&c.Controller)
					c.redirect(beego.URLFor("SessionsController.Post"))
				} else {
					for k, v := range formData {
						c.Data[k] = v[0]
					}
					flash.Set("danger", flash.Data["danger"], "\n", err.Error())
					c.TplName = "create.tpl"
				}
			}
		}
	} else {
		c.TplName = "create.tpl"
	}
}

// GetOne ...
// @Title Get One
// @Description get Users by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Users
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UsersController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetUsersById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Users
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Users
// @Failure 403
// @router / [get]
func (c *UsersController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]interface{})
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllUsers(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Users
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 200 {object} models.Users
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UsersController) Put() {
	//flash := beego.NewFlash()
	flash := beego.ReadFromRequest(&c.Controller)
	if c.userId == 0 {
		//c.flash.Set("warning", "您还未登录，请登录后再操作！！！")
		flash.Set("warning", "您还未登录，请登录后再操作！！！")
		flash.Store(&c.Controller)
		c.redirect(beego.URLFor("SessionsController.Post"))
	}
	if c.is("post") {
		v := c.currentUser
		formData := c.Ctx.Request.Form
		v.Gender = formData["Gender"][0]
		v.Phone = formData["Phone"][0]
		v.Email = formData["Email"][0]
		v.Status = true
		v.IsAdmin = false
		v.FollowedCount = 0
		v.Comment = formData["Comment"][0]
		v.UpdatedAt = time.Now()
		if err := models.UpdateUsersById(v); err == nil {
			flash.Success("Updated Successfully!")
			c.Data["flash"] = flash.Data
			//flash.Store(&c.Controller)
			c.Data["json"] = v
		} else {
			c.Data["json"] = v
			flash.Set("danger", err.Error())
			c.Data["flash"] = flash.Data
			//flash.Store(&c.Controller)
		}
	} else {
		c.Data["json"] = c.currentUser
	}
	c.TplName = "edit.tpl"
}

// Delete ...
// @Title Delete
// @Description delete the Users
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UsersController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteUsers(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *UsersController) ResetPassword() {
	//flash := beego.NewFlash()
	flash := beego.ReadFromRequest(&c.Controller)
	if c.userId == 0 {
		//c.flash.Set("warning", "您还未登录，请登录后再操作！！！")
		flash.Set("warning", "您还未登录，请登录后再操作！！！")
		flash.Store(&c.Controller)
		c.redirect(beego.URLFor("SessionsController.Post"))
	}
	c.Data["Name"] = c.userName
	c.TplName = "reset_password.tpl"
	if c.is("post") {
		var user *models.Users
		//beego.ReadFromRequest(&c.Controller)
		//userId := strings.TrimSpace(c.GetString("id", 0))
		userPassword := strings.TrimSpace(c.GetString("Password"))
		userPasswordConfirmation := strings.TrimSpace(c.GetString("PasswordConfirmation"))
		if userPassword != "" && userPassword == userPasswordConfirmation {
			user, _ = models.GetUsersById(c.userId)
			user.Password = models.GetEncryptPassword(userPassword)
			if err := models.UpdateUsersById(user); err == nil {
				c.Data["Name"] = c.userName
				//c.flash.Success("修改成功！")
				flash.Success("修改成功！")
				c.Data["flash"] = flash.Data
				//flash.Store(&c.Controller)
			}
		} else {
			//c.flash.Set("danger", "密码不一致！！！")
			flash.Set("danger", "密码不一致！！！")
			c.Data["flash"] = flash.Data
			//flash.Store(&c.Controller)
		}
	}
}

func (c *UsersController) Follow() {

}

func (c *UsersController) Unfollow() {

}
