package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"fmt"
	"github.com/udistrital/oikos_api/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// EspacioFisicoV2Controller operations for EspacioFisico
type EspacioFisicoV2Controller struct {
	beego.Controller
}

// URLMapping ...
func (c *EspacioFisicoV2Controller) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create EspacioFisico
// @Param	body		body 	models.EspacioFisicoV2	true		"body for EspacioFisico content"
// @Success 201 {int} models.EspacioFisicoV2
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *EspacioFisicoV2Controller) Post() {
	var v models.EspacioFisicoV2
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddEspacioFisico(&v); err == nil {
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
// @Description get EspacioFisico by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.EspacioFisicoV2
// @Failure 404 not found resource
// @router /:id [get]
func (c *EspacioFisicoV2Controller) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetEspacioFisicoById(id)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get EspacioFisico
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.EspacioFisicoV2
// @Failure 404 not found resource
// @router / [get]
func (c *EspacioFisicoV2Controller) GetAll() {
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

	l, err := models.GetAllEspacioFisico(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the EspacioFisico
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.EspacioFisicoV2	true		"body for EspacioFisico content"
// @Success 200 {object} models.EspacioFisicoV2
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *EspacioFisicoV2Controller) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.EspacioFisicoV2{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateEspacioFisicoById(&v); err == nil {
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
// @Description delete the EspacioFisico
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *EspacioFisicoV2Controller) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteEspacioFisico(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}

// EspaciosHuerfanos ...
// @Title EspaciosHuerfanos
// @Description Función para cargar los espacios físicos huerfanos
// @Param	id		path 	string	true		"Id del espacio físico"
// @Success 200 {object} models.EspacioFisico
// @Failure 403 id is empty
// @router /EspaciosHuerfanos/:id [get]
//Función para cargar los espacios físicos huerfanos
func (c *EspacioFisicoV2Controller) EspaciosHuerfanos() {
	fmt.Println("tipo ", c.Ctx.Input.Param(":id"))
	tipo := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(tipo)
	//perfiles := ("Admin_Arka")
	//perfilesR := strings.NewReplacer(",", "','")

	//Construcción Json Menús Huerfanos
	l := models.EspacioFisicosHuerfanos(id)
	fmt.Println("Este es el resultado de la consulta")
	fmt.Println(l)

	c.Data["json"] = l
	//Generera el Json con los datos obtenidos
	c.ServeJSON()
}

// GetEspaciosFisicosHijosById ...
// @Title GetEspaciosFisicosHijosById
// @Description A partir de un espacio físico dado, se obtienen los hijas de él en una estructura de árbol.
// @Param	espacio_fisico	path 	int	true		"Id del espacio físico"
// @Success 200 {object} models.EspacioFisicoPadreHijo
// @Failure 403 :espacio_fisico is empty
// @router /get_espacios_fisicos_hijos_by_id/:espacio_fisico [get]
func (c *EspacioFisicoV2Controller) GetEspaciosFisicosHijosById() {
	//Se crea variable que contiene el id con tipo de dato string
	espacioFisicoPadre := c.Ctx.Input.Param(":espacio_fisico")
	EFPadreint, _ := strconv.Atoi(espacioFisicoPadre)
	l, err := models.GetEspaciosFisicosHijosById(EFPadreint)
	if err != nil {
		beego.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")

	} else {

		c.Data["json"] = map[string]interface{}{"Body": l, "Type": "success"}
	}

	//Generera el Json con los datos obtenidos
	c.ServeJSON()
}

// GetEspaciosFisicosPadresById ...
// @Title GetEspaciosFisicosPadresById
// @Description A partir de una espacio_fisico dado, se obtienen todos sus predecesores en una estructura de árbol.
// @Param	espacio_fisico	path 	string	true		"Id de la espacio_fisico"
// @Success 200 {object} models.EspafioFisicoPadreHijo
// @Failure 404 :espacio_fisico is empty
// @router /get_espacios_fisicos_padres_by_id/:espacio_fisico [get]
func (c *EspacioFisicoV2Controller) GetEspaciosFisicosPadresById() {
	//Se crea variable que contiene el id con tipo de dato string
	espacioFisicoHijo := c.Ctx.Input.Param(":espacio_fisico")
	EFHijoint, _ := strconv.Atoi(espacioFisicoHijo)
	l, err := models.GetEspaciosFisicosPadresById(EFHijoint)
	if err != nil {
		beego.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")

	} else {

		c.Data["json"] = map[string]interface{}{"Body": l, "Type": "success"}
	}

	//Generera el Json con los datos obtenidos
	c.ServeJSON()
}