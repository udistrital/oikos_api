package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:AsignacionEspacioFisicoDependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:AsignacionEspacioFisicoDependenciaV2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:AsignacionEspacioFisicoDependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:AsignacionEspacioFisicoDependenciaV2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:AsignacionEspacioFisicoDependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:AsignacionEspacioFisicoDependenciaV2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:AsignacionEspacioFisicoDependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:AsignacionEspacioFisicoDependenciaV2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:AsignacionEspacioFisicoDependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:AsignacionEspacioFisicoDependenciaV2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:CampoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:CampoV2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:CampoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:CampoV2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:CampoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:CampoV2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:CampoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:CampoV2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:CampoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:CampoV2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreV2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreV2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreV2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreV2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreV2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreV2Controller"],
        beego.ControllerComments{
            Method: "ArbolDependenciasV2",
            Router: "/ArbolDependencias",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreV2Controller"],
        beego.ControllerComments{
            Method: "FacultadesConProyectosV2",
            Router: "/FacultadesConProyectos",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaPadreV2Controller"],
        beego.ControllerComments{
            Method: "TRDependenciaPadreV2",
            Router: "/tr_dependencia_padre",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaTipoDependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaTipoDependenciaV2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaTipoDependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaTipoDependenciaV2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaTipoDependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaTipoDependenciaV2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaTipoDependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaTipoDependenciaV2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaTipoDependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaTipoDependenciaV2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"],
        beego.ControllerComments{
            Method: "GetDependenciasHijasById",
            Router: "/get_dependencias_hijas_by_id/:dependencia",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"],
        beego.ControllerComments{
            Method: "GetDependenciasPadresById",
            Router: "/get_dependencias_padres_by_id/:dependencia",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"],
        beego.ControllerComments{
            Method: "ProyectosPorFacultadNivelAcademico",
            Router: "/proyectosPorFacultad/:id_facultad",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:DependenciaV2Controller"],
        beego.ControllerComments{
            Method: "ProyectosPorFacultad",
            Router: "/proyectosPorFacultad/:id_facultad/:nivel_academico",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoCampoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoCampoV2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoCampoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoCampoV2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoCampoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoCampoV2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoCampoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoCampoV2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoCampoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoCampoV2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoPadreV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoPadreV2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoPadreV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoPadreV2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoPadreV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoPadreV2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoPadreV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoPadreV2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoPadreV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoPadreV2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "EspaciosHuerfanos",
            Router: "/EspaciosHuerfanos/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "GetEspaciosFisicosHijosById",
            Router: "/get_espacios_fisicos_hijos_by_id/:espacio_fisico",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:EspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "GetEspaciosFisicosPadresById",
            Router: "/get_espacios_fisicos_padres_by_id/:espacio_fisico",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:ProyectoCurricularController"],
        beego.ControllerComments{
            Method: "GetAllProyectosByFacultadId",
            Router: "/get_all_proyectos_by_facultad_id/:id_facultad",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:ProyectoCurricularController"],
        beego.ControllerComments{
            Method: "GetAllProyectosByFacultades",
            Router: "/get_all_proyectos_by_facultades",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoDependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoDependenciaV2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoDependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoDependenciaV2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoDependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoDependenciaV2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoDependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoDependenciaV2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoDependenciaV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoDependenciaV2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoEspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoEspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoEspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoEspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoEspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoEspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoEspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoEspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoEspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoEspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoEspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoEspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoEspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoEspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoEspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoEspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoEspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoEspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoEspacioFisicoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoEspacioFisicoV2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoV2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoV2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoV2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoV2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoV2Controller"] = append(beego.GlobalControllerRouter["github.com/udistrital/oikos_api/controllers:TipoUsoV2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
