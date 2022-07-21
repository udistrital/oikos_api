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

// CampoController oprations for Campo
type CampoController struct {
	beego.Controller
}

// URLMapping ...
func (c *CampoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Campo
// @Param	body		body 	models.Campo	true		"body for Campo content"
// @Success 201 {object} models.Campo
// @Failure 403 body is empty
// @router / [post]
func (c *CampoController) Post() {
	var v models.Campo
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		//-------------- Temporal: Cambio por transición ------- //

		temp := models.CampoV2{
			Id:                v.Id,
			Nombre:            v.Nombre,
			Descripcion:       v.Descripcion,
			CodigoAbreviacion: "C_" + v.Nombre,
			Activo:            true,
			FechaCreacion:     time.Now(),
			FechaModificacion: time.Now(),
		}
		if _, err := models.AddCampo(&temp); err == nil {
			//if _, err := models.AddCampo(&v); err == nil {
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
// @Description get Campo by id
// @Param	id		path 	int	true		"The key for staticblock"
// @Success 200 {object} models.Campo
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CampoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetCampoById(id)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	} else {

		temp := models.Campo{
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
// @Description get Campo
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	int	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	int	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []models.Campo
// @Failure 403
// @router / [get]
func (c *CampoController) GetAll() {
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

	l, err := models.GetAllCampo(query, fields, sortby, order, offset, limit)
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
			var temp []models.Campo
			for _, i := range l {
				field, _ := i.(models.CampoV2)

				x := models.Campo{
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

		//c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Campo
// @Param	id		path 	int	true		"The id you want to update"
// @Param	body		body 	models.Campo	true		"body for Campo content"
// @Success 200 {object} models.Campo
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CampoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	infoDep, _ := models.GetCampoById(id)
	v := models.Campo{Id: id}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v2 := models.CampoV2{
			Id:                id,
			Nombre:            v.Nombre,
			Descripcion:       v.Descripcion,
			CodigoAbreviacion: infoDep.CodigoAbreviacion,
			Activo:            infoDep.Activo,
			FechaCreacion:     infoDep.FechaCreacion,
			FechaModificacion: time.Now(),
		}

		if err := models.UpdateCampoById(&v2); err == nil {
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
// @Description delete the Campo
// @Param	id		path 	int	true		"The id you want to delete"
// @Success 200 {object} models.Deleted
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CampoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteCampo(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}
