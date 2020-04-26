package controllers

import (
	"b_blog/models"
	_"time"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	query := map[string]interface{}{}
	//query["user_id"] = c.userId
	//now := time.Now()
	//query["created_at__gte"] = time.Date(now.Year(), now.Month(),
	//	now.Day(), 0, 0, 0, 0,
	//	time.Local).AddDate(0, 0,
	//	(int(time.Monday-now.Weekday())+(-7))%(-7))
	fields := []string{"Id", "Title", "Content", "UserId", "CreatedAt"}
	sortby := []string{"Id"}
	order := []string{"desc"}
	offset := int64(0)
	limit := int64(20)
	microposts, _, _ := models.GetAllMicroposts(query, fields, sortby, order, offset, limit)
	c.Data["microposts"] = microposts
	users, _ := models.GetAllUsers(map[string]interface{}{}, []string{"Id", "Name"},
		[]string{"Id"}, []string{"desc"}, int64(0), int64(5))
	c.Data["users"] = users
	c.Data["currentUser"] = c.currentUser
	c.TplName = "index.tpl"
}

func (c *MainController) Search() {
	//var page int64
	//
	//if pg, err := c.GetInt64("page"); err == nil {
	//	page = pg
	//} else {
	//	page = 1
	//}
	query := map[string]interface{}{}
	search := c.GetString("search")
	c.Data["search"] = ""
	query["title__contains"] = ""
	if search == "" {
		formData := c.Ctx.Request.Form
		if len(formData["search"]) > 0 && formData["search"][0] != "" {
			query["title__contains"] = formData["search"][0]
			c.Data["search"] = formData["search"][0]
		}
	} else {
		c.Data["search"] = search
		query["title__contains"] = search
	}

	//now := time.Now()
	//query["created_at__gte"] = time.Date(now.Year(), now.Month(),
	//	now.Day(), 0, 0, 0, 0,
	//	time.Local).AddDate(0, 0,
	//	(int(time.Monday-now.Weekday())+(-7))%(-7))
	fields := []string{"Id", "Title", "Content", "UserId", "CreatedAt"}
	sortby := []string{"Id"}
	order := []string{"desc"}
	//offset := (page - 1) * c.pageSize
	//limit := c.pageSize
	offset := int64(0)
	limit := int64(20)
	microposts, _, _ := models.GetAllMicroposts(query, fields, sortby, order, offset, limit)
	c.Data["microposts"] = microposts
	users, _ := models.GetAllUsers(map[string]interface{}{}, []string{"Id", "Name"},
		[]string{"Id"}, []string{"desc"}, int64(0), int64(5))
	c.Data["users"] = users
	c.Data["currentUser"] = c.currentUser
	c.Data["offset"] = offset
	c.Data["href"] = "?"
	c.TplName = "index.tpl"

	//query := map[string]interface{}{}
	//query["user_id"] = c.userId
	//now := time.Now()
	//query["created_at__gte"] = time.Date(now.Year(), now.Month(),
	//	now.Day(), 0, 0, 0, 0,
	//	time.Local).AddDate(0, 0,
	//	(int(time.Monday - now.Weekday()) + (-7))%(-7))
	//fields := []string{"Id", "Title", "Content", "UserId", "CreatedAt"}
	//sortby := []string{"Id"}
	//order := []string{"desc"}
	//offset := int64(0)
	//limit := int64(20)
	//microposts, _, _ := models.GetAllMicroposts(query, fields, sortby, order, offset, limit)
	//c.Data["microposts"] = microposts
	//users, _ := models.GetAllUsers(map[string]interface{}{}, []string{"Id", "Name"},
	//	[]string{"Id"}, []string{"desc"}, int64(0), int64(5))
	//c.Data["users"] = users
	//c.Data["currentUser"] = c.currentUser
	//c.TplName = "index.tpl"
}
