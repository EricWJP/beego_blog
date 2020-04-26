package controllers

import (
	"b_blog/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

//  MicropostsController operations for Microposts
type MicropostsController struct {
	BaseController
}

// URLMapping ...
func (c *MicropostsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Microposts
// @Param	body		body 	models.Microposts	true		"body for Microposts content"
// @Success 201 {int} models.Microposts
// @Failure 403 body is empty
// @router / [post]
func (c *MicropostsController) Post() {
	//flash := beego.NewFlash()
	flash := beego.ReadFromRequest(&c.Controller)
	if c.is("post") {
		var v models.Microposts
		formData := c.Ctx.Request.Form
		v.Title = formData["Title"][0]
		v.Content = formData["Content"][0]
		v.Status = true
		v.UserId = c.userId
		currentTime := time.Now()
		v.UpdatedAt, v.CreatedAt = currentTime, currentTime
		if _, err := models.AddMicroposts(&v); err == nil {
			c.Ctx.Output.SetStatus(200)
			//c.flash.Success("Created Successfully!")
			flash.Success("Created Successfully!")
			flash.Store(&c.Controller)
			//c.Data["flash"] = c.flash.Data
			c.redirect(beego.URLFor("MicropostsController.GetAll"))
		} else {
			for k, v := range formData {
				c.Data[k] = v[0]
			}
			//c.flash.Set("danger", err.Error())
			flash.Set("danger", err.Error())
			c.Data["flash"] = flash.Data
			//flash.Store(&c.Controller)
			//c.Data["flash"] = c.flash.Data
			c.TplName = "create.tpl"
		}
	} else {
		c.TplName = "create.tpl"
	}
}

// GetOne ...
// @Title Get One
// @Description get Microposts by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Microposts
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MicropostsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetMicropostsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
		c.Data["userName"] = c.userName
	}
	c.TplName = "show.tpl"
}

// GetAll ...
// @Title Get All
// @Description get Microposts
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Microposts
// @Failure 403
// @router / [get]
func (c *MicropostsController) GetAll() {
	//c.ServeJSON()
	query := map[string]interface{}{}
	fields := []string{"Id", "Title", "Content", "UserId", "CreatedAt"}
	sortby := []string{"id"}
	order := []string{"desc"}
	offset := int64(0)
	limit := int64(20)
	microposts, _ := models.GetAllMicroposts(query, fields, sortby, order, offset, limit)
	c.Data["microposts"] = microposts
	c.TplName = "index.tpl"
}

// Put ...
// @Title Put
// @Description update the Microposts
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Microposts	true		"body for Microposts content"
// @Success 200 {object} models.Microposts
// @Failure 403 :id is not int
// @router /:id [put]
func (c *MicropostsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Microposts{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateMicropostsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Microposts
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *MicropostsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteMicroposts(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *MicropostsController) MyMicroposts() {
	//flash := beego.NewFlash()
	flash := beego.ReadFromRequest(&c.Controller)
	if c.userId == 0 {
		//c.flash.Set("warning", "您还未登录，请登录后再操作！！！")
		flash.Set("warning", "您还未登录，请登录后再操作！！！")
		flash.Store(&c.Controller)
		c.redirect(beego.URLFor("SessionsController.Post"))
	}
	query := map[string]interface{}{}
	query["user_id"] = c.userId
	now := time.Now()
	query["created_at__gte"] = time.Date(now.Year(), now.Month(),
		now.Day(), 0, 0, 0, 0,
		time.Local).AddDate(0, 0,
		(int(time.Monday-now.Weekday())+(-7))%(-7))
	fields := []string{"Id", "Title", "Content", "UserId", "CreatedAt"}
	sortby := []string{"Id"}
	order := []string{"desc"}
	offset := int64(0)
	limit := int64(20)
	microposts, _ := models.GetAllMicroposts(query, fields, sortby, order, offset, limit)
	c.Data["microposts"] = microposts
	c.Data["currentUser"] = c.currentUser
	c.TplName = "my_microposts.tpl"
}
