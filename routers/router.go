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
	v1 "github.com/udistrital/oikos_api/controllers/v1"

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
				&v1.TipoUsoController{},
			),
		),

		beego.NSNamespace("/dependencia",
			beego.NSInclude(
				&v1.DependenciaController{},
			),
		),

		beego.NSNamespace("/espacio_fisico",
			beego.NSInclude(
				&v1.EspacioFisicoController{},
			),
		),

		beego.NSNamespace("/espacio_fisico_padre",
			beego.NSInclude(
				&v1.EspacioFisicoPadreController{},
			),
		),

		beego.NSNamespace("/tipo_uso_espacio_fisico",
			beego.NSInclude(
				&v1.TipoUsoEspacioFisicoController{},
			),
		),

		beego.NSNamespace("/dependencia_padre",
			beego.NSInclude(
				&v1.DependenciaPadreController{},
			),
		),

		beego.NSNamespace("/tipo_espacio_fisico",
			beego.NSInclude(
				&v1.TipoEspacioFisicoController{},
			),
		),

		beego.NSNamespace("/asignacion_espacio_fisico_dependencia",
			beego.NSInclude(
				&v1.AsignacionEspacioFisicoDependenciaController{},
			),
		),

		beego.NSNamespace("/dependencia_tipo_dependencia",
			beego.NSInclude(
				&v1.DependenciaTipoDependenciaController{},
			),
		),

		beego.NSNamespace("/tipo_dependencia",
			beego.NSInclude(
				&v1.TipoDependenciaController{},
			),
		),

		beego.NSNamespace("/campo",
			beego.NSInclude(
				&v1.CampoController{},
			),
		),

		beego.NSNamespace("/espacio_fisico_campo",
			beego.NSInclude(
				&v1.EspacioFisicoCampoController{},
			),
		),
	)

	ns2 := beego.NewNamespace("/v2",

		beego.NSNamespace("/proyecto_curricular",
			beego.NSInclude(
				&controllers.ProyectoCurricularController{},
			),
		),

		beego.NSNamespace("/tipo_uso_espacio_fisico",
			beego.NSInclude(
				&controllers.TipoUsoEspacioFisicoV2Controller{},
			),
		),

		beego.NSNamespace("/tipo_uso",
			beego.NSInclude(
				&controllers.TipoUsoV2Controller{},
			),
		),

		beego.NSNamespace("/tipo_espacio_fisico",
			beego.NSInclude(
				&controllers.TipoEspacioFisicoV2Controller{},
			),
		),

		beego.NSNamespace("/espacio_fisico_padre",
			beego.NSInclude(
				&controllers.EspacioFisicoPadreV2Controller{},
			),
		),

		beego.NSNamespace("/espacio_fisico_campo",
			beego.NSInclude(
				&controllers.EspacioFisicoCampoV2Controller{},
			),
		),

		beego.NSNamespace("/tipo_dependencia",
			beego.NSInclude(
				&controllers.TipoDependenciaV2Controller{},
			),
		),

		beego.NSNamespace("/espacio_fisico",
			beego.NSInclude(
				&controllers.EspacioFisicoV2Controller{},
			),
		),

		beego.NSNamespace("/dependencia_tipo_dependencia",
			beego.NSInclude(
				&controllers.DependenciaTipoDependenciaV2Controller{},
			),
		),

		beego.NSNamespace("/dependencia_padre",
			beego.NSInclude(
				&controllers.DependenciaPadreV2Controller{},
			),
		),

		beego.NSNamespace("/dependencia",
			beego.NSInclude(
				&controllers.DependenciaV2Controller{},
			),
		),

		beego.NSNamespace("/campo",
			beego.NSInclude(
				&controllers.CampoV2Controller{},
			),
		),

		beego.NSNamespace("/asignacion_espacio_fisico_dependencia",
			beego.NSInclude(
				&controllers.AsignacionEspacioFisicoDependenciaV2Controller{},
			),
		),
	)

	beego.AddNamespace(ns)
	beego.AddNamespace(ns2)

}
