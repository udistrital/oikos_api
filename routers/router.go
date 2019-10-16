// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/oikos_api/controllers"

	"github.com/astaxie/beego"
	//Libreria de middleware
	//"github.com/udistrital/auditoria"
)

func init() {

	//Iniciar middleware
	//auditoria.InitMiddleware()

	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_uso",
			beego.NSInclude(
				&controllers.TipoUsoController{},
			),
		),

		beego.NSNamespace("/dependencia",
			beego.NSInclude(
				&controllers.DependenciaController{},
			),
		),

		beego.NSNamespace("/espacio_fisico",
			beego.NSInclude(
				&controllers.EspacioFisicoController{},
			),
		),

		beego.NSNamespace("/espacio_fisico_padre",
			beego.NSInclude(
				&controllers.EspacioFisicoPadreController{},
			),
		),

		beego.NSNamespace("/tipo_uso_espacio_fisico",
			beego.NSInclude(
				&controllers.TipoUsoEspacioFisicoController{},
			),
		),

		beego.NSNamespace("/dependencia_padre",
			beego.NSInclude(
				&controllers.DependenciaPadreController{},
			),
		),

		beego.NSNamespace("/tipo_espacio_fisico",
			beego.NSInclude(
				&controllers.TipoEspacioFisicoController{},
			),
		),

		beego.NSNamespace("/asignacion_espacio_fisico_dependencia",
			beego.NSInclude(
				&controllers.AsignacionEspacioFisicoDependenciaController{},
			),
		),

		beego.NSNamespace("/dependencia_tipo_dependencia",
			beego.NSInclude(
				&controllers.DependenciaTipoDependenciaController{},
			),
		),

		beego.NSNamespace("/tipo_dependencia",
			beego.NSInclude(
				&controllers.TipoDependenciaController{},
			),
		),

	)

	ns2 := beego.NewNamespace("/v2",

		beego.NSNamespace("/proyecto_curricular",
			beego.NSInclude(
				&controllers.ProyectoCurricularController{},
			),
		),
	)

	beego.AddNamespace(ns)
	beego.AddNamespace(ns2)

}
