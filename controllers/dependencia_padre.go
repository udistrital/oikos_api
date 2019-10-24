package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"fmt"
	"github.com/udistrital/oikos_api/models"
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// DependenciaPadreController oprations for DependenciaPadre
type DependenciaPadreController struct {
	beego.Controller
}

// URLMapping ...
func (c *DependenciaPadreController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("FacultadesConProyectos", c.FacultadesConProyectos)
	c.Mapping("ArbolDependencias", c.ArbolDependencias)
}

// Post ...
// @Title Post
// @Description create DependenciaPadre
// @Param	body		body 	models.DependenciaPadre	true		"body for DependenciaPadre content"
// @Success 201 {int} models.DependenciaPadre
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *DependenciaPadreController) Post() {
	var v models.DependenciaPadre
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		//-------------- Temporal: Cambio por transición ------- //
		
		dp := &models.DependenciaV2 {
			Id: v.Padre.Id,
		}

		dh := &models.DependenciaV2 {
			Id: v.Hija.Id,
		}

		temp := models.DependenciaPadreV2 {
			Id: v.Id,
			PadreId: dp,
			HijaId: dh,
			Activo : true,
			FechaCreacion  : time.Now(),
			FechaModificacion  : time.Now(),
			
		}
		//-------------- Temporal: Cambio por transición ------- //
		if _, err := models.AddDependenciaPadre(&temp); err == nil {
		//if _, err := models.AddDependenciaPadre(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get DependenciaPadre by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.DependenciaPadre
// @Failure 404 not found resource
// @router /:id [get]
func (c *DependenciaPadreController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetDependenciaPadreById(id)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		//-------------- Temporal: Cambio por transición ------- //
		dp := &models.Dependencia {
			Id: v.PadreId.Id,
			Nombre: v.PadreId.Nombre,
			TelefonoDependencia:v.PadreId.TelefonoDependencia,
			CorreoElectronico: v.PadreId.CorreoElectronico,
		}

		dh := &models.Dependencia {
			Id: v.HijaId.Id,
			Nombre: v.HijaId.Nombre,
			TelefonoDependencia:v.HijaId.TelefonoDependencia,
			CorreoElectronico: v.HijaId.CorreoElectronico,
		}
		
		temp := models.DependenciaPadre {
					Id: v.Id,
					Padre: dp,
					Hija: dh,
					Activo: v.Activo,
					FechaCreacion: v.FechaCreacion,
					FechaModificacion: v.FechaModificacion,		  
			
				}

		c.Data["json"] = temp
//		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get DependenciaPadre
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.DependenciaPadre
// @Failure 404 not found resource
// @router / [get]
func (c *DependenciaPadreController) GetAll() {
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

	l, err := models.GetAllDependenciaPadre(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
			c.Data["json"] = l
		}else{
				//-------------- Temporal: Cambio por transición ------- //
			var temp []models.DependenciaPadre
			for _, i := range l {
				field, _ := i.(models.DependenciaPadreV2)
				
				dp := &models.Dependencia {
					Id: field.PadreId.Id,
					Nombre: field.PadreId.Nombre,
					TelefonoDependencia:field.PadreId.TelefonoDependencia,
					CorreoElectronico: field.PadreId.CorreoElectronico,
				}

				dh := &models.Dependencia {
					Id: field.HijaId.Id,
					Nombre: field.HijaId.Nombre,
					TelefonoDependencia:field.HijaId.TelefonoDependencia,
					CorreoElectronico: field.HijaId.CorreoElectronico,
				}

				x := models.DependenciaPadre {
					Id: field.Id,
					Padre: dp,
					Hija: dh,
					Activo: field.Activo,
					FechaCreacion: field.FechaCreacion,
					FechaModificacion: field.FechaModificacion,		  
			
				}

				temp = append(temp,x)
			}

			c.Data["json"] = temp

		}

		
		//c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the DependenciaPadre
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.DependenciaPadre	true		"body for DependenciaPadre content"
// @Success 200 {object} models.DependenciaPadre
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *DependenciaPadreController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.DependenciaPadre{Id: id}
	//-------------- Temporal: Cambio por transición ------- //
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		dp := &models.DependenciaV2{
			Id: v.Padre.Id,
		}
		dh := &models.DependenciaV2{
			Id: v.Hija.Id,
		}	
		v2 := models.DependenciaPadreV2{
			Id: id,
			PadreId: dp,
			HijaId: dh,
			Activo : v.Activo,
			FechaCreacion : v.FechaCreacion,
			FechaModificacion  : time.Now(),
		}

		if err := models.UpdateDependenciaPadreById(&v2); err == nil {
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the DependenciaPadre
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *DependenciaPadreController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteDependenciaPadre(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}

// FacultadesConProyectos ...
// @Title FacultadesConProyectos
// @Description Lista las facultades con sus respectivos proyectos curriculares
// @Success 200 {object} models.DependenciaPadre
// @Failure 403
// @router /FacultadesConProyectos [get]
func (c *DependenciaPadreController) FacultadesConProyectos() {
	//Construcción Json menus
	l := models.Facultades()
	c.Data["json"] = l
	//Generera el Json con los datos obtenidos
	c.ServeJSON()
}

// ArbolDependencias ...
// @Title ArbolDependencias
// @Description ArbolDependencias
// @Success 200 {object} models.Tree
// @Failure 403
// @router /ArbolDependencias [get]
func (c *DependenciaPadreController) ArbolDependencias() {
	//Construcción Json menus
	l := models.ConstruirDependenciasPadre()
	fmt.Println("Este es el resultado de la consulta")
	fmt.Println(l)

	c.Data["json"] = l
	//Generera el Json con los datos obtenidos
	c.ServeJSON()
}
