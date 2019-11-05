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

// EspacioFisicoPadreController oprations for EspacioFisicoPadre
type EspacioFisicoPadreController struct {
	beego.Controller
}

// URLMapping ...
func (c *EspacioFisicoPadreController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create EspacioFisicoPadre
// @Param	body		body 	models.EspacioFisicoPadre	true		"body for EspacioFisicoPadre content"
// @Success 201 {int} models.EspacioFisicoPadre
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *EspacioFisicoPadreController) Post() {
	var v models.EspacioFisicoPadre
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		//-------------- Temporal: Cambio por transición ------- //
		
		efp := &models.EspacioFisicoV2 {
			Id: v.Padre.Id,
		}

		efh := &models.EspacioFisicoV2 {
			Id: v.Hijo.Id,
		}

		temp := models.EspacioFisicoPadreV2 {
			Id: v.Id,
			PadreId: efp,
			HijoId: efh,
			FechaCreacion  : time.Now(),
			FechaModificacion  : time.Now(),
			
		}
		//-------------- Temporal: Cambio por transición ------- //
		if _, err := models.AddEspacioFisicoPadre(&temp); err == nil {
		//if _, err := models.AddEspacioFisicoPadre(&v); err == nil {
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
// @Description get EspacioFisicoPadre by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.EspacioFisicoPadre
// @Failure 404 not found resource
// @router /:id [get]
func (c *EspacioFisicoPadreController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetEspacioFisicoPadreById(id)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
	//-------------- Temporal: Cambio por transición ------- //

	efp := &models.EspacioFisico {
		Id : v.PadreId.Id,            			 		
	}

	efh := &models.EspacioFisico {
		Id : v.HijoId.Id,            			 		
	}
	
	temp := models.EspacioFisicoPadre {
				Id: v.Id,
				Padre: efp,
				Hijo: efh,
				FechaCreacion: v.FechaCreacion,
				FechaModificacion: v.FechaModificacion,		  
		}

	c.Data["json"] = temp	
		//c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get EspacioFisicoPadre
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.EspacioFisicoPadre
// @Failure 404 not found resource
// @router / [get]
func (c *EspacioFisicoPadreController) GetAll() {
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

	l, err := models.GetAllEspacioFisicoPadre(query, fields, sortby, order, offset, limit)
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
		var temp []models.EspacioFisicoPadre
		var act string;
		for _, i := range l {
			field, _ := i.(models.EspacioFisicoPadreV2)
			
			tefp := &models.TipoEspacioFisico {
				Id: field.PadreId.TipoEspacioFisicoId.Id,
				Nombre: field.PadreId.TipoEspacioFisicoId.Nombre,      		  
				Descripcion: field.PadreId.TipoEspacioFisicoId.Descripcion,
				CodigoAbreviacion: field.PadreId.TipoEspacioFisicoId.CodigoAbreviacion,
				Activo : field.PadreId.TipoEspacioFisicoId.Activo,
				FechaCreacion  : field.PadreId.TipoEspacioFisicoId.FechaCreacion,
				FechaModificacion  : field.PadreId.TipoEspacioFisicoId.FechaModificacion,
				
			}
		
			tefh := &models.TipoEspacioFisico {
				Id: field.HijoId.TipoEspacioFisicoId.Id,
				Nombre: field.HijoId.TipoEspacioFisicoId.Nombre,      		  
				Descripcion: field.HijoId.TipoEspacioFisicoId.Descripcion,
				CodigoAbreviacion: field.HijoId.TipoEspacioFisicoId.CodigoAbreviacion,
				Activo : field.HijoId.TipoEspacioFisicoId.Activo,
				FechaCreacion  : field.HijoId.TipoEspacioFisicoId.FechaCreacion,
				FechaModificacion  : field.HijoId.TipoEspacioFisicoId.FechaModificacion,
				
			}
			
			if (field.PadreId.Activo == true){
				act = "Activo"
			}else {
				act = "Inactivo"
			}

			efp := &models.EspacioFisico {
				Id : field.PadreId.Id,            			 		
				Nombre  : field.PadreId.Nombre,
				Estado: act,
				Codigo: field.PadreId.CodigoAbreviacion ,        	
				Descripcion :  field.PadreId.Descripcion    ,      		
				FechaCreacion :  field.PadreId.FechaCreacion    ,
				FechaModificacion:  field.PadreId.FechaModificacion ,
				TipoEspacio: tefp ,
			}
			
			if (field.HijoId.Activo == true){
				act = "Activo"
			}else {
				act = "Inactivo"
			}

			efh := &models.EspacioFisico {
				Id : field.HijoId.Id,            			 		
				Nombre  : field.HijoId.Nombre,
				Estado: act,
				Codigo: field.HijoId.CodigoAbreviacion,         	
				Descripcion :  field.HijoId.Descripcion    ,      		
				FechaCreacion :  field.HijoId.FechaCreacion    ,
				FechaModificacion: field.HijoId.FechaModificacion ,
				TipoEspacio: tefh ,
			}
			
			x := models.EspacioFisicoPadre {
						Id: field.Id,
						Padre: efp,
						Hijo: efh,
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
// @Description update the EspacioFisicoPadre
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.EspacioFisicoPadre	true		"body for EspacioFisicoPadre content"
// @Success 200 {object} models.EspacioFisicoPadre
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *EspacioFisicoPadreController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.EspacioFisicoPadre{Id: id}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
	//-------------- Temporal: Cambio por transición ------- //
		efp := &models.EspacioFisicoV2{
			Id: v.Padre.Id,
		}
		efh := &models.EspacioFisicoV2{
			Id: v.Hijo.Id,
		}	

		v2 := models.EspacioFisicoPadreV2{
			Id: id,
			PadreId: efp,
			HijoId: efh,
			FechaCreacion : v.FechaCreacion,
			FechaModificacion  : time.Now(),
		}

		if err := models.UpdateEspacioFisicoPadreById(&v2); err == nil {
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
// @Description delete the EspacioFisicoPadre
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *EspacioFisicoPadreController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteEspacioFisicoPadre(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}
