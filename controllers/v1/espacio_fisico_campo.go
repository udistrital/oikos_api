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
// @Success 201 {object} models.EspacioFisicoCampo
// @Failure 403 body is empty
// @router / [post]
func (c *EspacioFisicoCampoController) Post() {
	var v models.EspacioFisicoCampo
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		var temp models.EspacioFisicoCampoV2
		temp.FromV1(v)
		temp.Activo = true
		t := time.Now()
		temp.FechaInicio = t
		temp.FechaCreacion = t
		temp.FechaModificacion = t
		if _, err := models.AddEspacioFisicoCampo(&temp); err == nil {
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
// @Description get EspacioFisicoCampo by id
// @Param	id		path 	int	true		"The key for staticblock"
// @Success 200 {object} models.EspacioFisicoCampo
// @Failure 403 :id is empty
// @router /:id [get]
func (c *EspacioFisicoCampoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetEspacioFisicoCampoById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		var temp models.EspacioFisicoCampo
		v.ToV1(&temp)
		c.Data["json"] = temp
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
// @Param	limit	query	int	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	int	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []models.EspacioFisicoCampo
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

	aux := models.EspacioFisicoCampoV2{}
	l, err := models.GetAllEspacioFisicoCampo(
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
					v2    models.EspacioFisicoCampoV2
					v1aux models.EspacioFisicoCampo
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
			case models.EspacioFisicoCampoV2:
				var x models.EspacioFisicoCampo
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
// @Description update the EspacioFisicoCampo
// @Param	id		path 	int	true		"The id you want to update"
// @Param	body		body 	models.EspacioFisicoCampo	true		"body for EspacioFisicoCampo content"
// @Success 200 {string} update success!
// @Failure 403 :id is not int
// @router /:id [put]
func (c *EspacioFisicoCampoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	//-------------- Temporal: Cambio por transición ------- //
	v2, _ := models.GetEspacioFisicoCampoById(id)
	v := models.EspacioFisicoCampo{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v2.Valor = v.Valor
		v2.EspacioFisicoId = &models.EspacioFisicoV2{Id: v.EspacioFisico.Id}
		v2.CampoId = &models.CampoV2{Id: v.Campo.Id}
		v2.FechaModificacion = time.Now()
		if err := models.UpdateEspacioFisicoCampoById(v2); err == nil {
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
// @Description delete the EspacioFisicoCampo
// @Param	id		path 	int	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *EspacioFisicoCampoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteEspacioFisicoCampo(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
