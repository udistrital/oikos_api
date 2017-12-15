// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego/plugins/cors"
	//"github.com/ojardila/beego_middleware"

	"github.com/udistrital/oikos_api/controllers"

	"github.com/astaxie/beego"
	//Libreria de middleware
)

func init() {

	//Iniciar middleware
	//beego_middleware.InitMiddleware()

	//Incluyendo el CORS
	beego.Debug("Filters init...")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

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

		beego.NSNamespace("/campo",
			beego.NSInclude(
				&controllers.CampoController{},
			),
		),

		beego.NSNamespace("/espacio_fisico_campo",
			beego.NSInclude(
				&controllers.EspacioFisicoCampoController{},
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
	beego.AddNamespace(ns)
}
