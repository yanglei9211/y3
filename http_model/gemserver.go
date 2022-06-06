package http_model

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type gemserverModel struct {
	Host      string
	Appid     string
	Username2 string //小学账号
	Username3 string //初中账号
	Username4 string //高中账号
	Username  string //无限定账号
}

var gemserverConfigModel gemserverModel

type gemserverTask struct {
	url  string
	data QueryInterface
}

func (s gemserverTask) GetUrl() string {
	furl := fmt.Sprintf("%s%s", gemserverConfigModel.Host, s.url)
	return furl
}

func (s gemserverTask) GetData() QueryInterface {
	return s.data
}

type ItemDetailsQuery struct {
	Appid    string
	Username string
	Subject  string
	ItemIds  []string
}

func (s ItemDetailsQuery) BuildParams() string {
	//ret := fmt.Sprintf("subject=%s&item_ids%s&appid=%s&username=%s",
	//	s.Subject, json_item_ids, s.Appid, s.Username)
	mp := map[string]interface{}{}
	mp["appid"] = s.Appid
	mp["username"] = s.Username
	mp["subject"] = s.Subject
	mp["item_ids"] = s.ItemIds
	mpJson, _ := json.Marshal(mp)
	ret := fmt.Sprintf("data=%s", mpJson)
	return ret
}

type ItemDetail struct {
	Id   string `json:"_id"`
	Data struct {
		Desc string `json:"desc"`
		Stem string `json:"stem"`
		Html string `json:"html"`
		Qs   []struct {
			Desc string      `json:"desc"`
			Exp  string      `json:"exp"`
			Opts []string    `json:"opts"`
			Ans  interface{} `json:"ans"`
		} `json:"qs"`
	} `json:"data"`
}

type ItemDetailsResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Items []ItemDetail `json:"items"`
	} `json:"data"'`
}

func GetItemDetails(subject string, item_id string) ItemDetailsResponse {
	query := ItemDetailsQuery{
		ItemIds:  []string{item_id},
		Subject:  subject,
		Appid:    gemserverConfigModel.Appid,
		Username: gemserverConfigModel.Username,
	}
	url := fmt.Sprintf("/%s/pub/item/batch/details", subject)
	httpTask := gemserverTask{url: url, data: query}
	//httpTask := gemserverTask{fmt.Sprintf("/%s/pub/item/batch/details", subject)}
	retBody := requestGet(httpTask)
	var ret ItemDetailsResponse
	_ = json.Unmarshal(retBody, &ret)
	return ret
}

func InitGemserverModel() {
	appConfig := web.AppConfig
	fmt.Println(gemserverConfigModel)
	gemserverConfigModel.Host, _ = appConfig.String("gemserver_host")
	gemserverConfigModel.Appid, _ = appConfig.String("gemserver_appid")
	gemserverConfigModel.Username, _ = appConfig.String("gemserver_username")
	gemserverConfigModel.Username3, _ = appConfig.String("gemserver_3_username")
	gemserverConfigModel.Username4, _ = appConfig.String("gemserver_4_username")
	fmt.Println(gemserverConfigModel)
}
