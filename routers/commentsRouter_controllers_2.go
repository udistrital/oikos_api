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

	beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method:           "ArbolDependencias",
			Router:           `ArbolDependencias/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaController"],
		beego.ControllerComments{
			Method:           "ProyectosPorFacultad",
			Router:           `ProyectosPorFacultad/:id_facultad`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

}
