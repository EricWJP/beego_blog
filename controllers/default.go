package controllers

import (
	"b_blog/models"
	"fmt"
	"time"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	query := map[string]interface{}{}
	//query["user_id"] = c.userId
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
	users, _ := models.GetAllUsers(map[string]interface{}{}, []string{"Id", "Name"},
		[]string{"Id"}, []string{"desc"}, int64(0), int64(5))
	c.Data["users"] = users
	c.Data["currentUser"] = c.currentUser
	c.TplName = "index.tpl"
}

func (c *MainController) Search() {
	formData := c.Ctx.Request.Form
	fmt.Println(formData)
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
	//microposts, _ := models.GetAllMicroposts(query, fields, sortby, order, offset, limit)
	//c.Data["microposts"] = microposts
	//users, _ := models.GetAllUsers(map[string]interface{}{}, []string{"Id", "Name"},
	//	[]string{"Id"}, []string{"desc"}, int64(0), int64(5))
	//c.Data["users"] = users
	//c.Data["currentUser"] = c.currentUser
	//c.TplName = "index.tpl"
}
