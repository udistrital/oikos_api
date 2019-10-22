package controllers

import (
	"encoding/json"
	"errors"
	"github.com/udistrital/oikos_api/models"
	"strconv"
	"strings"
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// EspacioFisicoCampoController oprations for EspacioFisicoCampo
type EspacioFisicoCampoController struct {
	beego.Controller
}

// URLMapping ...
func (c *EspacioFisicoCampoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create EspacioFisicoCampo
// @Param	body		body 	models.EspacioFisicoCampo	true		"body for EspacioFisicoCampo content"
// @Success 201 {int} models.EspacioFisicoCampo
// @Failure 403 body is empty
// @router / [post]
func (c *EspacioFisicoCampoController) Post() {
	var v models.EspacioFisicoCampo
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		ca := &models.CampoV2 {
			Id: v.Campo.Id,
		}

		ef := &models.EspacioFisicoV2 {
			Id: v.EspacioFisico.Id,
		}

		temp := models.EspacioFisicoCampoV2 {
			Id: v.Id,
			Valor: v.Valor,
			CampoId: ca,
			EspacioFisicoId: ef,
			FechaInicio: time.Now(),
			Activo : true,
			FechaCreacion  : time.Now(),
			FechaModificacion  : time.Now(),
			
		}
		//-------------- Temporal: Cambio por transición ------- //
		if _, err := models.AddEspacioFisicoCampo(&temp); err == nil {
		//if _, err := models.AddEspacioFisicoCampo(&v); err == nil {
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
// @Description get EspacioFisicoCampo by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.EspacioFisicoCampo
// @Failure 403 :id is empty
// @router /:id [get]
func (c *EspacioFisicoCampoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetEspacioFisicoCampoById(id)
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

		ca := &models.Campo {
			Id: v.CampoId.Id,
			Nombre: v.CampoId.Nombre,      
			Descripcion: v.CampoId.Descripcion,
			CodigoAbreviacion: v.CampoId.CodigoAbreviacion,
			Activo: v.CampoId.Activo,
			FechaCreacion: v.CampoId.FechaCreacion,
			FechaModificacion: v.CampoId.FechaModificacion,		
		}

		ef := &models.EspacioFisico {
			Id: v.EspacioFisicoId.Id,
			Nombre: v.EspacioFisicoId.Nombre,      
			Descripcion: v.EspacioFisicoId.Descripcion,
			Codigo: v.EspacioFisicoId.CodigoAbreviacion,
			Estado: "v.EspacioFisicoId.Activo",
			FechaCreacion: v.EspacioFisicoId.FechaCreacion,
			FechaModificacion: v.EspacioFisicoId.FechaModificacion,		
			TipoEspacio: te,
		}

		temp := models.EspacioFisicoCampo {
			Id: v.Id,
			Valor: v.Valor,
			Campo: ca,
			EspacioFisico: ef,
			FechaInicio: v.FechaInicio,
			FechaFin: v.FechaFin,
			Activo : v.Activo,
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
// @Description get EspacioFisicoCampo
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.EspacioFisicoCampo
// @Failure 403
// @router / [get]
func (c *EspacioFisicoCampoController) GetAll() {
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

	l, err := models.GetAllEspacioFisicoCampo(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		//-------------- Temporal: Cambio por transición ------- //
		var temp []models.EspacioFisicoCampo
		for _, i := range l {
			field, _ := i.(models.EspacioFisicoCampoV2)
			
			te := &models.TipoEspacioFisico {
				Id: field.EspacioFisicoId.TipoEspacio.Id,
				Nombre: field.EspacioFisicoId.TipoEspacio.Nombre, 
				Descripcion: field.EspacioFisicoId.TipoEspacio.Descripcion,
				CodigoAbreviacion: field.EspacioFisicoId.TipoEspacio.CodigoAbreviacion,
				Activo: field.EspacioFisicoId.TipoEspacio.Activo,
				FechaCreacion: field.EspacioFisicoId.TipoEspacio.FechaCreacion,
				FechaModificacion: field.EspacioFisicoId.TipoEspacio.FechaModificacion,	     		  
			}
	
			c := &models.Campo {
				Id: field.CampoId.Id,
				Nombre: field.CampoId.Nombre,      
				Descripcion: field.CampoId.Descripcion,
				CodigoAbreviacion: field.CampoId.CodigoAbreviacion,
				Activo: field.CampoId.Activo,
				FechaCreacion: field.CampoId.FechaCreacion,
				FechaModificacion: field.CampoId.FechaModificacion,		
			}
	
			ef := &models.EspacioFisico {
				Id: field.EspacioFisicoId.Id,
				Nombre: field.EspacioFisicoId.Nombre,      
				Descripcion: field.EspacioFisicoId.Descripcion,
				Codigo: field.EspacioFisicoId.CodigoAbreviacion,
				Estado: "field.EspacioFisicoId.Activo",
				FechaCreacion: field.EspacioFisicoId.FechaCreacion,
				FechaModificacion: field.EspacioFisicoId.FechaModificacion,		
				TipoEspacio: te,
			}
	
			x := models.EspacioFisicoCampo {
				Id: field.Id,
				Valor: field.Valor,
				Campo: c,
				EspacioFisico: ef,
				FechaInicio: field.FechaInicio,
				FechaFin: field.FechaFin,
				Activo : field.Activo,
				FechaCreacion  : field.FechaCreacion,
				FechaModificacion  : field.FechaModificacion,
				
			}

			temp = append(temp,x)
		}

		c.Data["json"] = map[string]interface{}{"Status": "200", "Body": temp, "Type": "success"}

		//c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the EspacioFisicoCampo
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.EspacioFisicoCampo	true		"body for EspacioFisicoCampo content"
// @Success 200 {object} models.EspacioFisicoCampo
// @Failure 403 :id is not int
// @router /:id [put]
func (c *EspacioFisicoCampoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	//-------------- Temporal: Cambio por transición ------- //
	infoDep, _ := models.GetEspacioFisicoCampoById(id)
	v := models.EspacioFisicoCampoV2{
		Id: id,
		Activo : true,
		FechaCreacion : infoDep.FechaCreacion,
		FechaModificacion  : time.Now(),
	}
	//v := models.EspacioFisicoCampo{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateEspacioFisicoCampoById(&v); err == nil {
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
// @Description delete the EspacioFisicoCampo
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *EspacioFisicoCampoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteEspacioFisicoCampo(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}