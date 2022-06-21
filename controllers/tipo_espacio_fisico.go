package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/udistrital/oikos_api/models"
)

// TipoEspacioFisicoController oprations for TipoEspacioFisico
type TipoEspacioFisicoController struct {
	beego.Controller
}

// URLMapping ...
func (c *TipoEspacioFisicoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create TipoEspacioFisico
// @Param	body		body 	models.TipoEspacioFisico	true		"body for TipoEspacioFisico content"
// @Success 201 {int} models.TipoEspacioFisico
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TipoEspacioFisicoController) Post() {
	var v models.TipoEspacioFisico
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		//-------------- Temporal: Cambio por transición ------- //

		temp := models.TipoEspacioFisicoV2{
			Id:                v.Id,
			Nombre:            v.Nombre,
			Descripcion:       "Descripción",
			CodigoAbreviacion: "TU_" + v.Nombre,
			Activo:            true,
			FechaCreacion:     time.Now(),
			FechaModificacion: time.Now(),
		}
		//-------------- Temporal: Cambio por transición ------- //
		if _, err := models.AddTipoEspacioFisico(&temp); err == nil {
			//if _, err := models.AddTipoEspacioFisico(&v); err == nil {
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
// @Description get TipoEspacioFisico by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.TipoEspacioFisico
// @Failure 404 not found resource
// @router /:id [get]
func (c *TipoEspacioFisicoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTipoEspacioFisicoById(id)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		//-------------- Temporal: Cambio por transición ------- //

		temp := models.TipoEspacioFisico{
			Id:                v.Id,
			Nombre:            v.Nombre,
			Descripcion:       v.Descripcion,
			CodigoAbreviacion: v.CodigoAbreviacion,
			Activo:            v.Activo,
			FechaCreacion:     v.FechaCreacion,
			FechaModificacion: v.FechaModificacion,
		}
		c.Data["json"] = temp
		//-------------- Temporal: Cambio por transición ------- //
		//c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get TipoEspacioFisico
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.TipoEspacioFisico
// @Failure 404 not found resource
// @router / [get]
func (c *TipoEspacioFisicoController) GetAll() {
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

	l, err := models.GetAllTipoEspacioFisico(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
			c.Data["json"] = l
		} else {
			//-------------- Temporal: Cambio por transición ------- //
			var temp []models.TipoEspacioFisico
			for _, i := range l {
				field, _ := i.(models.TipoEspacioFisicoV2)
				x := models.TipoEspacioFisico{
					Id:                field.Id,
					Nombre:            field.Nombre,
					Descripcion:       field.Descripcion,
					CodigoAbreviacion: field.CodigoAbreviacion,
					Activo:            field.Activo,
					FechaCreacion:     field.FechaCreacion,
					FechaModificacion: field.FechaModificacion,
				}

				temp = append(temp, x)
			}

			c.Data["json"] = temp
		}

		//-------------- Temporal: Cambio por transición ------- //

		//c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the TipoEspacioFisico
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TipoEspacioFisico	true		"body for TipoEspacioFisico content"
// @Success 200 {object} models.TipoEspacioFisico
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *TipoEspacioFisicoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	//-------------- Temporal: Cambio por transición ------- //
	infoDep, _ := models.GetTipoEspacioFisicoById(id)
	v := models.TipoEspacioFisico{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v2 := models.TipoEspacioFisicoV2{
			Id:                id,
			Nombre:            v.Nombre,
			Descripcion:       infoDep.Descripcion,
			CodigoAbreviacion: infoDep.CodigoAbreviacion,
			Activo:            infoDep.Activo,
			FechaCreacion:     infoDep.FechaCreacion,
			FechaModificacion: time.Now(),
		}

		if err := models.UpdateTipoEspacioFisicoById(&v2); err == nil {
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
// @Description delete the TipoEspacioFisico
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *TipoEspacioFisicoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteTipoEspacioFisico(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}
