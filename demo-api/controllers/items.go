package controllers

import (
	"github.com/hackranger/demo-api/models"
	"encoding/json"

	"github.com/astaxie/beego"
	"strconv"
	"log"
)

// Operations about Users
type ItemsController struct {
	beego.Controller
}

// @Title CreateItem
// @Description create item
// @Param	body		body 	models.Item	true		"body for item content"
// @Success 200 {int} models.Item.Id
// @Failure 403 body is empty
// @router / [post]
func (i *ItemsController) Post() {
	var item models.Item
	json.Unmarshal(i.Ctx.Input.RequestBody, &item)
	models.ItemList[item.Id] = &item
	i.Data["json"] = item
	i.ServeJSON()
}

// @Title GetAll
// @Description get all items
// @Success 200 {object} models.User
// @router / [get]
func (i *ItemsController) GetAll() {	
	i.Data["json"] = models.ItemList
	i.ServeJSON()
}

// @Title Get
// @Description get item by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (i *ItemsController) Get() {
	id := i.GetString(":id")
	log.Println("ID: " , id)

	if id != "" {
		itemid,_ := strconv.Atoi(id)
		item := models.ItemList[itemid]
		i.Data["json"] = item
	} else {
		i.Data["json"] = nil
	}

	i.ServeJSON()
}