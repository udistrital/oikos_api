package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/udistrital/oikos_api/models"

	"github.com/astaxie/beego"

)

// DependenciaController oprations for Dependencia
type DependenciaController struct {
	beego.Controller
}

// URLMapping ...
func (c *DependenciaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("ProyectosPorFacultad", c.ProyectosPorFacultad)
	c.Mapping("GetDependenciasHijasById", c.GetDependenciasHijasById)
}

// Post ...
// @Title Post
// @Description create Dependencia
// @Param	body		body 	models.Dependencia	true		"body for Dependencia content"
// @Success 201 {int} models.Dependencia
// @Failure 403 body is empty
// @router / [post]
func (c *DependenciaController) Post() {
	var v models.Dependencia
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddDependencia(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Dependencia by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Dependencia
// @Failure 403 :id is empty
// @router /:id [get]
func (c *DependenciaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetDependenciaById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Dependencia
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Dependencia
// @Failure 403
// @router / [get]
func (c *DependenciaController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllDependencia(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Dependencia
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Dependencia	true		"body for Dependencia content"
// @Success 200 {object} models.Dependencia
// @Failure 403 :id is not int
// @router /:id [put]
func (c *DependenciaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Dependencia{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateDependenciaById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Dependencia
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *DependenciaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteDependencia(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// ProyectosPorFacultad ...
// @Title ProyectosPorFacultad
// @Description Get curricular projects by faculty
// @Param	id_facultad		path 	int	true		"El id de la facultad a consultar sus proyectos curriculares"
// @Param	nivel_academico		path 	string	true		"El nivel académico a consultar de acuerdo a la facultad"
// @Success 200 {object} models.Dependencia
// @Failure 403 :id_facultad is empty
// @router /proyectosPorFacultad/:id_facultad/:nivel_academico [get]
func (c *DependenciaController) ProyectosPorFacultad() {
	//Se crea variable que contiene el id con tipo de dato string
	idStr := c.Ctx.Input.Param(":id_facultad")
	nivel_academico := c.Ctx.Input.Param(":nivel_academico")
	//Se nombra la variable id, en la cual se hizo la conversión de string a int
	id_facultad, _ := strconv.Atoi(idStr)

	//Construcción Json menus
	l := models.ProyectosPorFacultad(id_facultad, nivel_academico)
	fmt.Println("Este es el resultado de la consulta")
	fmt.Println(l)

	c.Data["json"] = l
	//Generera el Json con los datos obtenidos
	c.ServeJSON()
}

// ProyectosPorFacultadNivelAcademico ...
// @Title ProyectosPorFacultadNivelAcademico
// @Description Get curricular projects by faculty and academic level
// @Param	id_facultad		path 	int	true		"El id de la facultad a consultar sus proyectos curriculares"
// @Success 200 {object} models.Dependencia
// @Failure 403 :id_facultad is empty
// @router /proyectosPorFacultad/:id_facultad [get]
func (c *DependenciaController) ProyectosPorFacultadNivelAcademico() {
	//Se crea variable que contiene el id con tipo de dato string
	idStr := c.Ctx.Input.Param(":id_facultad")
	//Se nombra la variable id, en la cual se hizo la conversión de string a int
	id_facultad, _ := strconv.Atoi(idStr)

	//Construcción Json menus
	l := models.ProyectosPorFacultad(id_facultad, "undefined")
	fmt.Println("Este es el resultado de la consulta")
	fmt.Println(l)

	c.Data["json"] = l
	//Generera el Json con los datos obtenidos
	c.ServeJSON()
}


// GetDependenciasHijasById ...
// @Title GetDependenciasHijasById
// @Description A partir de una dependencia dada, se obtienen las hijas de ella en una estructura de árbol.
// @Param	dependencia	path 	int	true		"Id de la dependencia"
// @Success 200 {object} models.DependenciaPadre
// @Failure 403 :dependencia_padre is empty
// @router /get_dependencias_hijas_by_id/:dependencia [get]
func (c *DependenciaController) GetDependenciasHijasById() {
	//Se crea variable que contiene el id con tipo de dato string
	dependenciaPadre := c.Ctx.Input.Param(":dependencia")
	depPadreint,_:= strconv.Atoi(dependenciaPadre)
	l,err := models.GetDependenciasHijasById(depPadreint)
	if err != nil {
		beego.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
		
	}else{
		
		c.Data["json"] = map[string]interface{}{"Body": l, "Type": "success"}
	}
	
	//Generera el Json con los datos obtenidos
	c.ServeJSON()
}


// GetDependenciasPadresById ...
// @Title GetDependenciasPadresById
// @Description A partir de una dependencia dada, se obtienen todos sus predecesores en una estructura de árbol.
// @Param	dependencia	path 	string	true		"Id de la dependencia"
// @Success 200 {object} models.DependenciaPadre
// @Failure 404 :dependencia is empty
// @router /get_dependencias_padres_by_id/:dependencia [get]
func (c *DependenciaController) GetDependenciasPadresById() {
	//Se crea variable que contiene el id con tipo de dato string
	dependenciaPadre := c.Ctx.Input.Param(":dependencia")
	
	l,err := models.GetDependenciasPadresById(dependenciaPadre)
	if err != nil {
		beego.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
		
	}else{
		
		c.Data["json"] = map[string]interface{}{"Body": l, "Type": "success"}
	}
	
	//Generera el Json con los datos obtenidos
	c.ServeJSON()
}