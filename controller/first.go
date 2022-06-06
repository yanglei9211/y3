package controller

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"gopkg.in/mgo.v2/bson"
	"y3/db_model"
	"y3/http_model"
)

type WelcomeController struct {
	BaseController
}

func (this *WelcomeController) Get() {
	ret := map[string]interface{}{}
	ret["welcome"] = "y3"
	this.writeResponse(ret)
}

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	x, _ := this.GetInt("x")
	y, _ := this.GetInt("y")
	ret := map[string]interface{}{}
	ret["ans"] = x + y
	this.writeResponse(ret)
}

type ItemInfoController struct {
	BaseController
}

func (this *ItemInfoController) Get() {
	strItemId := this.GetString("item_id")
	subjectName := this.GetString("subject")
	logs.Info(fmt.Sprintf("get args: item_id: %s, subject: %s", strItemId, subjectName))
	itemId := bson.ObjectIdHex(strItemId)
	itemInfo := db_model.ItemFindOneById(subjectName, itemId)
	ret := map[string]interface{}{}
	ret["item_info"] = itemInfo
	this.writeResponse(ret)
}

type ItemDetailsController struct {
	BaseController
}

func (this *ItemDetailsController) Get() {
	strItemId := this.GetString("item_id")
	subjectName := this.GetString("subject")
	logs.Info(fmt.Sprintf("get args: item_id: %s, subject: %s", strItemId, subjectName))
	var ret http_model.ItemDetailsResponse
	ret = http_model.GetItemDetails(subjectName, strItemId)
	retData := map[string]interface{}{}
	retData["item_details"] = ret.Data
	this.writeResponse(retData)
}
