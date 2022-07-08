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

// TipoDependenciaController oprations for TipoDependencia
type TipoDependenciaController struct {
	beego.Controller
}

// URLMapping ...
func (c *TipoDependenciaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create TipoDependencia
// @Param	body		body 	models.TipoDependencia	true		"body for TipoDependencia content"
// @Success 201 {object} models.TipoDependencia
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TipoDependenciaController) Post() {
	var v models.TipoDependencia
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		//-------------- Temporal: Cambio por transición ------- //

		// TODO: Revisar lo siguiente ...:
		temp := models.TipoDependenciaV2{
			Id:                v.Id,
			Nombre:            v.Nombre,
			Descripcion:       "Descripción",
			CodigoAbreviacion: "TU_" + v.Nombre,
			Activo:            true,
			FechaCreacion:     time.Now(),
			FechaModificacion: time.Now(),
		}
		// ... debería bastar con:
		// var temp models.TipoDependenciaV2
		// temp.FromV1(v)
		//-------------- Temporal: Cambio por transición ------- //
		if _, err := models.AddTipoDependencia(&temp); err == nil {
			//if _, err := models.AddTipoDependencia(&v); err == nil {
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
// @Description get TipoDependencia by id
// @Param	id		path 	int	true		"The key for staticblock"
// @Success 200 {object} models.TipoDependencia
// @Failure 404 not found resource
// @router /:id [get]
func (c *TipoDependenciaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTipoDependenciaById(id)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		//-------------- Temporal: Cambio por transición ------- //
		var temp models.TipoDependencia
		v.ToV1(&temp)
		c.Data["json"] = temp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get TipoDependencia
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	int	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	int	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []models.TipoDependencia
// @Failure 404 not found resource
// @router / [get]
func (c *TipoDependenciaController) GetAll() {
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

	l, err := models.GetAllTipoDependencia(query, fields, sortby, order, offset, limit)
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
			var temp []interface{}
			for _, i := range l {
				switch v := i.(type) {
				case map[string]interface{}:
					temp = append(temp, v)
				case models.TipoDependenciaV2:
					var x models.TipoDependencia
					v.ToV1(&x)
					temp = append(temp, v)
					// default:
					// 	// SIN MANEJAR!
				}
			}
			c.Data["json"] = temp
		}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the TipoDependencia
// @Param	id		path 	int	true		"The id you want to update"
// @Param	body		body 	models.TipoDependencia	true		"body for TipoDependencia content"
// @Success 200 {object} models.TipoDependencia
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *TipoDependenciaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	//-------------- Temporal: Cambio por transición ------- //
	v2, _ := models.GetTipoDependenciaById(id)
	v := models.TipoDependencia{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		// TODO: Revisar lo siguiente...:
		v2.Id = id
		v2.Nombre = v.Nombre
		v2.FechaModificacion = time.Now()
		// ... debería bastar con:
		// v2.FromV1(v)
		if err := models.UpdateTipoDependenciaById(v2); err == nil {
			v2.ToV1(&v)
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
// @Description delete the TipoDependencia
// @Param	id		path 	int	true		"The id you want to delete"
// @Success 200 {object} models.Deleted
// @Failure 404 not found resource
// @router /:id [delete]
func (c *TipoDependenciaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteTipoDependencia(id); err == nil {
		c.Data["json"] = models.Deleted{Id: id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}
