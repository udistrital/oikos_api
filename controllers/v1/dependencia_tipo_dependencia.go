package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/udistrital/oikos_api/models"
	"github.com/udistrital/utils_oas/formatdata"
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
			c.Ctx.Output.SetStatus(201)
			temp.ToV1(&v)
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
		c.Data["json"] = err.Error()
	} else {
		//-------------- Temporal: Cambio por transición ------- //
		var temp models.DependenciaTipoDependencia
		v.ToV1(&temp)
		c.Data["json"] = temp
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

	aux := models.DependenciaTipoDependenciaV2{}
	l, err := models.GetAllDependenciaTipoDependencia(
		aux.QueryFromV1(query),
		aux.SelectorsFromV1(fields),
		aux.SelectorsFromV1(sortby), order,
		offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		//-------------- Temporal: Cambio por transición ------- //
		var temp []interface{}
		for _, i := range l {
			switch v := i.(type) {
			case map[string]interface{}:
				// len(fields) > 0
				var (
					v2    models.DependenciaTipoDependenciaV2
					v1aux models.DependenciaTipoDependencia
					v1    map[string]interface{}
					err   error
				)
				formatdata.FillStruct(v, &v2)                     // convertir a estructura v2 ...
				v2.ToV1(&v1aux)                                   // ... para poder convertir a v1
				formatdata.FillStruct(v1aux, &v1)                 // Luego a un mapeo auxiliar a ser...
				if v1, err = FilterKeys(v1, fields); err != nil { // ...filtrado
					logs.Error(err)
					c.Abort(fmt.Sprint(http.StatusInternalServerError))
				}
				temp = append(temp, v1)
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
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the DependenciaTipoDependencia
// @Param	id		path 	int	true		"The id you want to update"
// @Param	body		body 	models.DependenciaTipoDependencia	true		"body for DependenciaTipoDependencia content"
// @Success 200 {string} update success!
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *DependenciaTipoDependenciaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v2, _ := models.GetDependenciaTipoDependenciaById(id)
	v := models.DependenciaTipoDependencia{Id: id}
	//-------------- Temporal: Cambio por transición ------- //
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v2.TipoDependenciaId = &models.TipoDependenciaV2{Id: v.TipoDependenciaId.Id}
		v2.DependenciaId = &models.DependenciaV2{Id: v.DependenciaId.Id}
		v2.FechaModificacion = time.Now()
		if err := models.UpdateDependenciaTipoDependenciaById(v2); err == nil {
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
// @Description delete the DependenciaTipoDependencia
// @Param	id		path 	int	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *DependenciaTipoDependenciaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteDependenciaTipoDependencia(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
