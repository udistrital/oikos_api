package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoController"],
		beego.ControllerComments{
			Method:           "EspaciosHuerfanos",
			Router:           `EspaciosHuerfanos/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

}
