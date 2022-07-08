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

// DependenciaTipoDependenciaController oprations for DependenciaTipoDependencia
type DependenciaTipoDependenciaController struct {
	beego.Controller
}

// URLMapping ...
func (c *DependenciaTipoDependenciaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create DependenciaTipoDependencia
// @Param	body		body 	models.DependenciaTipoDependencia	true		"body for DependenciaTipoDependencia content"
// @Success 201 {object} models.DependenciaTipoDependencia
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *DependenciaTipoDependenciaController) Post() {
	var v models.DependenciaTipoDependencia
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		//-------------- Temporal: Cambio por transición ------- //
		var temp models.DependenciaTipoDependenciaV2
		temp.FromV1(v)
		temp.Activo = true
		t := time.Now()
		temp.FechaCreacion = t
		temp.FechaModificacion = t
		//-------------- Temporal: Cambio por transición ------- //
		if _, err := models.AddDependenciaTipoDependencia(&temp); err == nil {
			//if _, err := models.AddDependenciaTipoDependencia(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			temp.ToV1(&v)
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
// @Description get DependenciaTipoDependencia by id
// @Param	id		path 	int	true		"The key for staticblock"
// @Success 200 {object} models.DependenciaTipoDependencia
// @Failure 404 not found resource
// @router /:id [get]
func (c *DependenciaTipoDependenciaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetDependenciaTipoDependenciaById(id)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		//-------------- Temporal: Cambio por transición ------- //
		var temp models.DependenciaTipoDependencia
		v.ToV1(&temp)
		c.Data["json"] = temp

		//c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get DependenciaTipoDependencia
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	int	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	int	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []models.DependenciaTipoDependencia
// @Failure 404 not found resource
// @router / [get]
func (c *DependenciaTipoDependenciaController) GetAll() {
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

	l, err := models.GetAllDependenciaTipoDependencia(query, fields, sortby, order, offset, limit)
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
				case models.DependenciaTipoDependenciaV2:
					var x models.DependenciaTipoDependencia
					v.ToV1(&x)
					temp = append(temp, x)
					// default:
					// 	// SIN MANEJAR!
				}
			}
			c.Data["json"] = temp
		}

		//c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the DependenciaTipoDependencia
// @Param	id		path 	int	true		"The id you want to update"
// @Param	body		body 	models.DependenciaTipoDependenciaV2	true		"body for DependenciaTipoDependencia content"
// @Success 200 {object} models.DependenciaTipoDependenciaV2
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *DependenciaTipoDependenciaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v2, _ := models.GetDependenciaTipoDependenciaById(id)
	v := models.DependenciaTipoDependencia{Id: id}
	//-------------- Temporal: Cambio por transición ------- //
	// TODO: Revisar lo siguiente ...:
	v2.Id = id
	v2.FechaModificacion = time.Now()
	// ... debería bastar con:
	// v2.FromV1(v)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateDependenciaTipoDependenciaById(v2); err == nil {
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
// @Description delete the DependenciaTipoDependencia
// @Param	id		path 	int	true		"The id you want to delete"
// @Success 200 {object} models.Deleted
// @Failure 404 not found resource
// @router /:id [delete]
func (c *DependenciaTipoDependenciaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteDependenciaTipoDependencia(id); err == nil {
		c.Data["json"] = models.Deleted{Id: id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}
