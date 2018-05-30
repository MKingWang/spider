package controller

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

type SpiderController struct {
	beego.Controller
}

type spiderInfo struct {
	ID     int         `json:"id"`
	Status int         `json:"status"`
	Url    string      `json:"url"`
	Worker string      `json:"worker"`
	Result interface{} `json:"result"`
}

func (s *SpiderController) SpiderApi() {
	var spiderinfo spiderInfo
	json.Unmarshal(s.Ctx.Input.RequestBody, &spiderinfo)

	s.Data["json"] = spiderinfo
	s.ServeJSON()

}
