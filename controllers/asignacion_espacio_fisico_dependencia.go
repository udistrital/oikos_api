package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"github.com/udistrital/oikos_api/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

// AsignacionEspacioFisicoDependenciaController oprations for AsignacionEspacioFisicoDependencia
type AsignacionEspacioFisicoDependenciaController struct {
	beego.Controller
}

// URLMapping ...
func (c *AsignacionEspacioFisicoDependenciaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create AsignacionEspacioFisicoDependencia
// @Param	body		body 	models.AsignacionEspacioFisicoDependencia	true		"body for AsignacionEspacioFisicoDependencia content"
// @Success 201 {int} models.AsignacionEspacioFisicoDependencia
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *AsignacionEspacioFisicoDependenciaController) Post() {
	var v models.AsignacionEspacioFisicoDependencia
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		//-------------- Temporal: Cambio por transición ------- //
		var act bool;
		if (v.Estado == "Activo"){
			act = true
		}else if (v.Estado == "Inactivo"){
			act = false
		}else{
			act = true
		}

		dc,_ := strconv.Atoi(v.DocumentoSoporte)
		ef := &models.EspacioFisicoV2 {
			Id: v.EspacioFisicoId.Id,
		}
		d := &models.DependenciaV2 {
			Id: v.DependenciaId.Id,
		}

		temp := models.AsignacionEspacioFisicoDependenciaV2 {
					
					Id: v.Id,
					EspacioFisicoId: ef,
					DependenciaId: d,
					FechaInicio:  v.FechaInicio,
					FechaFin: v.FechaFin,
					DocumentoSoporte: dc, 
	  				Activo : act ,
					FechaCreacion  : time.Now(),
					FechaModificacion  : time.Now(),
					
		}
			
		if _, err := models.AddAsignacionEspacioFisicoDependencia(&temp); err == nil {
		//-------------- Temporal: Cambio por transición ------- //	
		//	if _, err := models.AddAsignacionEspacioFisicoDependencia(&v); err == nil {
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
// @Description get AsignacionEspacioFisicoDependencia by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.AsignacionEspacioFisicoDependencia
// @Failure 404 not found resource
// @router /:id [get]
func (c *AsignacionEspacioFisicoDependenciaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetAsignacionEspacioFisicoDependenciaById(id)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		//-------------- Temporal: Cambio por transición ------- //
		dc := strconv.Itoa(v.DocumentoSoporte)
		ef := &models.EspacioFisico {
			Id: v.EspacioFisicoId.Id,
		}
		d := &models.Dependencia {
			Id: v.DependenciaId.Id,
		}

		var act string;
				if (v.Activo == true){
					act = "Activo"
				}else {
					act = "Inactivo"
				}
			

		temp := models.AsignacionEspacioFisicoDependencia {
					
			  Id: v.Id,
			  Estado: act,
			  FechaInicio:  v.FechaInicio,
			  FechaFin: v.FechaFin,
			  DocumentoSoporte: dc, 
			  EspacioFisicoId: ef,
			  DependenciaId: d,

		}
		c.Data["json"] = temp
		//-------------- Temporal: Cambio por transición ------- //
		//c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get AsignacionEspacioFisicoDependencia
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.AsignacionEspacioFisicoDependencia
// @Failure 404 not found resource
// @router / [get]
func (c *AsignacionEspacioFisicoDependenciaController) GetAll() {
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

	l, err := models.GetAllAsignacionEspacioFisicoDependencia(query, fields, sortby, order, offset, limit)
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
			var temp []models.AsignacionEspacioFisicoDependencia
			for _, i := range l {
				field, _ := i.(models.AsignacionEspacioFisicoDependenciaV2)
				dc := strconv.Itoa(field.DocumentoSoporte)

				tef := &models.TipoEspacioFisico {
					Id: field.EspacioFisicoId.TipoEspacioFisicoId.Id,
					Nombre: field.EspacioFisicoId.TipoEspacioFisicoId.Nombre, 
					Descripcion: field.EspacioFisicoId.TipoEspacioFisicoId.Descripcion,
					CodigoAbreviacion: field.EspacioFisicoId.TipoEspacioFisicoId.CodigoAbreviacion,
					Activo: field.EspacioFisicoId.TipoEspacioFisicoId.Activo,
					FechaCreacion: field.EspacioFisicoId.TipoEspacioFisicoId.FechaCreacion,
					FechaModificacion: field.EspacioFisicoId.TipoEspacioFisicoId.FechaModificacion,	
				}

				var act string;
				if (field.EspacioFisicoId.Activo == true){
					act = "Activo"
				}else {
					act = "Inactivo"
				}
			
				
				ef := &models.EspacioFisico {
					Id: field.EspacioFisicoId.Id,
					Nombre: field.EspacioFisicoId.Nombre,
					Descripcion: field.EspacioFisicoId.Descripcion,
					Codigo: field.EspacioFisicoId.CodigoAbreviacion,
					Estado:act,
					TipoEspacio : tef,
				}

				d := &models.Dependencia {
					Id: field.DependenciaId.Id,
					Nombre: field.DependenciaId.Nombre,      		  
					TelefonoDependencia: field.DependenciaId.TelefonoDependencia,
					CorreoElectronico: field.DependenciaId.CorreoElectronico, 
					       
				}

				x := models.AsignacionEspacioFisicoDependencia {
					Id: field.Id,
					Estado: "TRUE",
					FechaInicio:  field.FechaInicio,
					FechaFin: field.FechaFin,
					EspacioFisicoId: ef,
					DependenciaId: d,
					DocumentoSoporte: dc, 
				
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
// @Description update the AsignacionEspacioFisicoDependencia
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.AsignacionEspacioFisicoDependencia	true		"body for AsignacionEspacioFisicoDependencia content"
// @Success 200 {object} models.AsignacionEspacioFisicoDependencia
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *AsignacionEspacioFisicoDependenciaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	//-------------- Temporal: Cambio por transición ------- //
	infoDep, _ := models.GetAsignacionEspacioFisicoDependenciaById(id)
	v := models.AsignacionEspacioFisicoDependenciaV2{
		Id: id,
		Activo : true ,
		FechaCreacion : infoDep.FechaCreacion,
		FechaModificacion  : time.Now(),
	}
	//-------------- Temporal: Cambio por transición ------- //
	//v := models.AsignacionEspacioFisicoDependencia{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateAsignacionEspacioFisicoDependenciaById(&v); err == nil {
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
// @Description delete the AsignacionEspacioFisicoDependencia
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *AsignacionEspacioFisicoDependenciaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteAsignacionEspacioFisicoDependencia(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}
