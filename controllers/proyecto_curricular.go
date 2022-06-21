package controllers

import (
	"strconv"

	"github.com/astaxie/beego"

	"github.com/udistrital/oikos_api/models"
)

// ProyectoCurricularController oprations for Facultad
type ProyectoCurricularController struct {
	beego.Controller
}

// URLMapping ...
func (c *ProyectoCurricularController) URLMapping() {
	c.Mapping("GetAllProyectosByFacultades", c.GetAllProyectosByFacultades)
	c.Mapping("GetAllProyectosByFacultadId", c.GetAllProyectosByFacultadId)
}

// GetAllProyectosByFacultades ...
// @Title GetAllProyectosByFacultades
// @Description Obtener una lista de todas las facultades y sus respectivos proyectos curriculares
// @Success 200 {object} []models.DependenciaPadreHijo
// @Failure 403
// @router /get_all_proyectos_by_facultades [get]
func (c *ProyectoCurricularController) GetAllProyectosByFacultades() {
	l, err := models.GetAllProyectosByFacultades()
	if err != nil {
		beego.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")

	} else {

		c.Data["json"] = map[string]interface{}{"Body": l, "Type": "success"}
	}

	//Generera el Json con los datos obtenidos
	c.ServeJSON()
}

// GetAllProyectosByFacultadId ...
// @Title  GetAllProyectosByFacultadId
// @Description Se obtienen los proyectos curriculares, dado un id de una facultad
// @Param	id_facultad		path 	int	true		"El id de la facultad a consultar sus proyectos curriculares"
// @Success 200 {object} []models.DependenciaPadreHijo
// @Failure 403
// @router /get_all_proyectos_by_facultad_id/:id_facultad [get]
func (c *ProyectoCurricularController) GetAllProyectosByFacultadId() {
	idStr := c.Ctx.Input.Param(":id_facultad")
	id_facultad, _ := strconv.Atoi(idStr)
	l, err := models.GetAllProyectosByFacultadId(id_facultad)
	if err != nil {
		beego.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")

	} else {

		c.Data["json"] = map[string]interface{}{"Body": l, "Type": "success"}
	}

	//Generera el Json con los datos obtenidos
	c.ServeJSON()
}
