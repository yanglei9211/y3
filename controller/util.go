package controller

import (
	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
}

func (this *BaseController) writeResponse(data map[string]interface{}) {
	retData := map[string]interface{}{}
	retData["status"] = 1
	retData["data"] = data
	this.Data["json"] = retData
	this.ServeJSON()
	//response, err := json.Marshal(
	//	map[string]interface{}{
	//		"status": 1,
	//		"data":   data,
	//	})
	//if err != nil {
	//	panic(err)
	//}
	//this.Ctx.WriteString(string(response))
	//this.Ctx.Request.Header.Set("content-type", "application/json")
}

func (this *BaseController) writeException() {

}
