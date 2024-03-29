package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:AsignacionEspacioFisicoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:AsignacionEspacioFisicoDependenciaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:AsignacionEspacioFisicoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:AsignacionEspacioFisicoDependenciaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:AsignacionEspacioFisicoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:AsignacionEspacioFisicoDependenciaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:AsignacionEspacioFisicoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:AsignacionEspacioFisicoDependenciaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:AsignacionEspacioFisicoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:AsignacionEspacioFisicoDependenciaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:CampoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:CampoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:CampoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:CampoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:CampoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"],
        beego.ControllerComments{
            Method: "GetDependenciasHijasById",
            Router: "/get_dependencias_hijas_by_id/:dependencia",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"],
        beego.ControllerComments{
            Method: "GetDependenciasPadresById",
            Router: "/get_dependencias_padres_by_id/:dependencia",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"],
        beego.ControllerComments{
            Method: "ProyectosPorFacultadNivelAcademico",
            Router: "/proyectosPorFacultad/:id_facultad",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaController"],
        beego.ControllerComments{
            Method: "ProyectosPorFacultad",
            Router: "/proyectosPorFacultad/:id_facultad/:nivel_academico",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaPadreController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaPadreController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaPadreController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaPadreController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaPadreController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaPadreController"],
        beego.ControllerComments{
            Method: "ArbolDependencias",
            Router: "/ArbolDependencias",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaPadreController"],
        beego.ControllerComments{
            Method: "FacultadesConProyectos",
            Router: "/FacultadesConProyectos",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaPadreController"],
        beego.ControllerComments{
            Method: "TRDependenciaPadre",
            Router: "/tr_dependencia_padre",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaTipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaTipoDependenciaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaTipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaTipoDependenciaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaTipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaTipoDependenciaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaTipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaTipoDependenciaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaTipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:DependenciaTipoDependenciaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoCampoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoCampoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoCampoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoCampoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoCampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoCampoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoController"],
        beego.ControllerComments{
            Method: "EspaciosHuerfanos",
            Router: "/EspaciosHuerfanos/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoController"],
        beego.ControllerComments{
            Method: "GetEspaciosFisicosHijosById",
            Router: "/get_espacios_fisicos_hijos_by_id/:espacio_fisico",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoController"],
        beego.ControllerComments{
            Method: "GetEspaciosFisicosPadresById",
            Router: "/get_espacios_fisicos_padres_by_id/:espacio_fisico",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoPadreController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoPadreController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoPadreController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoPadreController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:EspacioFisicoPadreController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoDependenciaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoDependenciaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoDependenciaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoDependenciaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoDependenciaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoEspacioFisicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoEspacioFisicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoEspacioFisicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoEspacioFisicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoEspacioFisicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoEspacioFisicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoEspacioFisicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoEspacioFisicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoEspacioFisicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoEspacioFisicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers/v1:TipoUsoEspacioFisicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
