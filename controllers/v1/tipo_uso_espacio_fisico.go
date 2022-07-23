package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"github.com/udistrital/oikos_api/models"
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
// @Success 201 {object} models.TipoUsoEspacioFisico
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TipoUsoEspacioFisicoController) Post() {
	var v models.TipoUsoEspacioFisico
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		//-------------- Temporal: Cambio por transición ------- //
		var temp models.TipoUsoEspacioFisicoV2
		temp.FromV1(v)
		temp.Activo = true
		t := time.Now()
		temp.FechaCreacion = t
		temp.FechaModificacion = t
		if _, err := models.AddTipoUsoEspacioFisico(&temp); err == nil {
			c.Ctx.Output.SetStatus(201)
			v.Id = temp.Id
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
// @Description get TipoUsoEspacioFisico by id
// @Param	id		path 	int	true		"The key for staticblock"
// @Success 200 {object} models.TipoUsoEspacioFisico
// @Failure 404 not found resource
// @router /:id [get]
func (c *TipoUsoEspacioFisicoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTipoUsoEspacioFisicoById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		//-------------- Temporal: Cambio por transición ------- //
		var temp models.TipoUsoEspacioFisico
		v.ToV1(&temp)
		c.Data["json"] = temp
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
// @Param	limit	query	int	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	int	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []models.TipoUsoEspacioFisico
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
		c.Data["json"] = err.Error()
	} else {
		//-------------- Temporal: Cambio por transición ------- //
		var temp []interface{}
		for _, i := range l {
			switch v := i.(type) {
			case map[string]interface{}:
				temp = append(temp, v)
			case models.TipoUsoEspacioFisicoV2:
				var x models.TipoUsoEspacioFisico
				v.ToV1(&x)
				temp = append(temp, x)
				// default:
				// 	// SIN MANEJAR!
			}
		}
		c.Data["json"] = temp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the TipoUsoEspacioFisico
// @Param	id		path 	int	true		"The id you want to update"
// @Param	body		body 	models.TipoUsoEspacioFisico	true		"body for TipoUsoEspacioFisico content"
// @Success 200 {string} update success!
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *TipoUsoEspacioFisicoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	//-------------- Temporal: Cambio por transición ------- //
	v2, _ := models.GetTipoUsoEspacioFisicoById(id)
	v := models.TipoUsoEspacioFisico{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v2.TipoUsoId = &models.TipoUsoV2{Id: v.TipoUsoId.Id}
		v2.EspacioFisicoId = &models.EspacioFisicoV2{Id: v.EspacioFisicoId.Id}
		v2.FechaModificacion = time.Now()
		if err := models.UpdateTipoUsoEspacioFisicoById(v2); err == nil {
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
// @Description delete the TipoUsoEspacioFisico
// @Param	id		path 	int	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *TipoUsoEspacioFisicoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteTipoUsoEspacioFisico(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
