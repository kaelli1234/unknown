package v1

import (
	"github.com/astaxie/beego"
	"github.com/json-iterator/go"

	"ccgwf/base"
	Error "ccgwf/error"
	"ccgwf/logs"
	"ccgwf/models"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type ExampleController struct {
	beego.Controller
	base.ClientController
}

func (e *ExampleController) Home() {
	e.Initialize(e.Ctx)
	e.Output()
	return
}

type TestRequest struct {
	AccessToken string `json:"access_token"`
	Cid         string `json:"cid"`
	UID         string `json:"uid"`
	DevType     string `json:"dev_type"`
	ClientVer   string `json:"client_ver"`
}

func (e *ExampleController) Test() {
	e.Initialize(e.Ctx)
	var request TestRequest
	if err := json.Unmarshal(e.Ctx.Input.RequestBody, &request); err != nil {
		logs.Error("Test err %v", err)
		e.Output(Error.UNKNOWN_ERROR)
		return
	}

	datas := []string{}
	e.Output(Error.SUCCESS, datas)
	return
}

type TokenResponse struct {
	ErrCode     string `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
}

func (e *ExampleController) Token() {
	e.Initialize(e.Ctx)
	e.Output(&TokenResponse{})
	return
}

func (e *ExampleController) Mysql() {
	e.Initialize(e.Ctx)
	response := models.GetDemo()
	e.Output(Error.SUCCESS, response)
	return
}
