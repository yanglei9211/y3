package http_model

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/valyala/fasthttp"
)

type QueryInterface interface {
	BuildParams() string
}
type HttpTaskInterface interface {
	GetData() QueryInterface
	GetUrl() string
}

func buildQueryUrl(url string, query QueryInterface) string {
	ret := fmt.Sprintf("%s?%s", url, query.BuildParams())
	return ret
}

func requestGet(s HttpTaskInterface) []byte {
	query := s.GetData()
	url := s.GetUrl()
	furl := buildQueryUrl(url, query)
	fmt.Println(furl)
	statusCode, body, err := fasthttp.Get(nil, furl)
	fmt.Println(statusCode, err)
	if err != nil || statusCode != 200 {
		logs.Error("request get: %s error, code: %d, error: %s",
			furl, statusCode, err.Error())
		return []byte{}
	} else {
		return body
	}

}

func requestPost(s HttpTaskInterface) map[string]interface{} {
	return nil
}

func InitHttpModel() {
	InitGemserverModel()
}
