package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
	"github.com/udistrital/oikos_api/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// TipoUsoEspacioFisicoController oprations for TipoUsoEspacioFisico
type TipoUsoEspacioFisicoController struct {
	beego.Controller
}

// URLMapping ...
func (c *TipoUsoEspacioFisicoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create TipoUsoEspacioFisico
// @Param	body		body 	models.TipoUsoEspacioFisico	true		"body for TipoUsoEspacioFisico content"
// @Success 201 {int} models.TipoUsoEspacioFisico
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TipoUsoEspacioFisicoController) Post() {
	var v models.TipoUsoEspacioFisico
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		//-------------- Temporal: Cambio por transición ------- //
		ef := &models.EspacioFisicoV2 {
			Id: v.EspacioFisicoId.Id,
		}

		tu := &models.TipoUsoV2 {
			Id: v.TipoUsoId.Id,
		}

		temp := models.TipoUsoEspacioFisicoV2 {
			Id: v.Id,
			TipoUsoId: tu,
			EspacioFisicoId: ef,
			Activo : true,
			FechaCreacion  : time.Now(),
			FechaModificacion  : time.Now(),
			
		}
		//-------------- Temporal: Cambio por transición ------- //
		if _, err := models.AddTipoUsoEspacioFisico(&temp); err == nil {
	//	if _, err := models.AddTipoUsoEspacioFisico(&v); err == nil {
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
// @Description get TipoUsoEspacioFisico by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.TipoUsoEspacioFisico
// @Failure 404 not found resource
// @router /:id [get]
func (c *TipoUsoEspacioFisicoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTipoUsoEspacioFisicoById(id)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		//-------------- Temporal: Cambio por transición ------- //
		te := &models.TipoEspacioFisico {
			Id: v.EspacioFisicoId.TipoEspacio.Id,
			Nombre: v.EspacioFisicoId.TipoEspacio.Nombre, 
			Descripcion: v.EspacioFisicoId.TipoEspacio.Descripcion,
			CodigoAbreviacion: v.EspacioFisicoId.TipoEspacio.CodigoAbreviacion,
			Activo: v.EspacioFisicoId.TipoEspacio.Activo,
			FechaCreacion: v.EspacioFisicoId.TipoEspacio.FechaCreacion,
			FechaModificacion: v.EspacioFisicoId.TipoEspacio.FechaModificacion,	     		  
		}

		ef := &models.EspacioFisico {
			Id: v.EspacioFisicoId.Id,
			Nombre: v.EspacioFisicoId.Nombre,   
			Codigo: v.EspacioFisicoId.CodigoAbreviacion,
			Estado: "ACTIVO",  //v.Activo
			Descripcion:  v.EspacioFisicoId.Descripcion,    
			FechaCreacion :  v.EspacioFisicoId.FechaCreacion,   
			FechaModificacion :  v.EspacioFisicoId.FechaModificacion,		
			TipoEspacio : te,	
			//DependenciaTipoDependencia: field.DependenciaTipoDependencia,       
		}

		tu := &models.TipoUso {
			Id: v.TipoUsoId .Id,
			Nombre: v.TipoUsoId.Nombre,      
			Descripcion: v.TipoUsoId.Descripcion,
			CodigoAbreviacion: v.TipoUsoId.CodigoAbreviacion,
			Activo: v.TipoUsoId.Activo,
			FechaCreacion: v.TipoUsoId.FechaCreacion,
			FechaModificacion: v.TipoUsoId.FechaModificacion,		
		}

		temp := models.TipoUsoEspacioFisico {
			Id: v.Id,
			TipoUsoId : tu,
			EspacioFisicoId : ef,
			Activo : true,
			FechaCreacion  : v.FechaCreacion,
			FechaModificacion  : v.FechaModificacion,
			
		}

		c.Data["json"] = temp

		//c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get TipoUsoEspacioFisico
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.TipoUsoEspacioFisico
// @Failure 404 not found resource
// @router / [get]
func (c *TipoUsoEspacioFisicoController) GetAll() {
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

	l, err := models.GetAllTipoUsoEspacioFisico(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		//-------------- Temporal: Cambio por transición ------- //
		var temp []models.TipoUsoEspacioFisico
		for _, i := range l {
			field, _ := i.(models.TipoUsoEspacioFisicoV2)
			
			te := &models.TipoEspacioFisico {
				Id: field.EspacioFisicoId.TipoEspacio.Id,
				Nombre: field.EspacioFisicoId.TipoEspacio.Nombre, 
				Descripcion: field.EspacioFisicoId.TipoEspacio.Descripcion,
				CodigoAbreviacion: field.EspacioFisicoId.TipoEspacio.CodigoAbreviacion,
				Activo: field.EspacioFisicoId.TipoEspacio.Activo,
				FechaCreacion: field.EspacioFisicoId.TipoEspacio.FechaCreacion,
				FechaModificacion: field.EspacioFisicoId.TipoEspacio.FechaModificacion,	     		  
			}
	
			ef := &models.EspacioFisico {
				Id: field.EspacioFisicoId.Id,
				Nombre: field.EspacioFisicoId.Nombre,   
				Codigo: field.EspacioFisicoId.CodigoAbreviacion,
				Estado: "ACTIVO",  //field.Activo
				Descripcion:  field.EspacioFisicoId.Descripcion,    
				FechaCreacion :  field.EspacioFisicoId.FechaCreacion,   
				FechaModificacion :  field.EspacioFisicoId.FechaModificacion,		
				TipoEspacio : te,	
				//DependenciaTipoDependencia: field.DependenciaTipoDependencia,       
			}
	
			tu := &models.TipoUso {
				Id: field.TipoUsoId .Id,
				Nombre: field.TipoUsoId.Nombre,      
				Descripcion: field.TipoUsoId.Descripcion,
				CodigoAbreviacion: field.TipoUsoId.CodigoAbreviacion,
				Activo: field.TipoUsoId.Activo,
				FechaCreacion: field.TipoUsoId.FechaCreacion,
				FechaModificacion: field.TipoUsoId.FechaModificacion,		
			}
	
			x := models.TipoUsoEspacioFisico {
				Id: field.Id,
				TipoUsoId : tu,
				EspacioFisicoId : ef,
				Activo : true,
				FechaCreacion  : field.FechaCreacion,
				FechaModificacion  : field.FechaModificacion,
				
			}

			temp = append(temp,x)
		}

		c.Data["json"] = temp

		//c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the TipoUsoEspacioFisico
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TipoUsoEspacioFisico	true		"body for TipoUsoEspacioFisico content"
// @Success 200 {object} models.TipoUsoEspacioFisico
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *TipoUsoEspacioFisicoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	//-------------- Temporal: Cambio por transición ------- //
	infoDep, _ := models.GetTipoUsoEspacioFisicoById(id)
	v := models.TipoUsoEspacioFisicoV2{
		Id: id,
		Activo : true,
		FechaCreacion : infoDep.FechaCreacion,
		FechaModificacion  : time.Now(),
	}

	//v := models.TipoUsoEspacioFisico{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateTipoUsoEspacioFisicoById(&v); err == nil {
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
// @Description delete the TipoUsoEspacioFisico
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *TipoUsoEspacioFisicoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteTipoUsoEspacioFisico(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}
