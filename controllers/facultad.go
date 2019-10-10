package controllers

import (
	//"encoding/json"
	//"errors"
	//"fmt"
	//"strconv"
	//"strings"
	"github.com/udistrital/oikos_api/models"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/logs"
)

// FacultadController oprations for Facultad
type FacultadController struct {
	beego.Controller
}

// URLMapping ...
func (c *FacultadController) URLMapping() {
	c.Mapping("GetFacultadesYProyectos", c.GetFacultadesYProyectos)
}

// GetFacultadesYProyectos ...
// @Title PGetFacultadesYProyectos
// @Description Obtener una lista de todas las facultades y sus respectivos proyectos curriculares
// @Success 200 {object} models.Dependencia
// @Failure 403 
// @router /get_facultades_y_proyectos [get]
func (c *FacultadController) GetFacultadesYProyectos() {
	l, err := models.GetFacultadesYProyectos()
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

