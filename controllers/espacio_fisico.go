package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/oikos_api/models"

)

// EspacioFisicoController oprations for EspacioFisico
type EspacioFisicoController struct {
	beego.Controller
}

// URLMapping ...
func (c *EspacioFisicoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("EspaciosHuerfanos", c.EspaciosHuerfanos)
}

// Post ...
// @Title Post
// @Description create EspacioFisico
// @Param	body		body 	models.EspacioFisico	true		"body for EspacioFisico content"
// @Success 201 {int} models.EspacioFisico
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *EspacioFisicoController) Post() {
	var v models.EspacioFisico
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

		te := &models.TipoEspacioFisicoV2 {
			Id: v.TipoEspacio.Id,
		}

		temp := models.EspacioFisicoV2 {
			Id  :   v.Id,          
			Nombre:  v.Nombre,    
			Descripcion:  "Descripción - " +v.Nombre,    
			CodigoAbreviacion : v.Codigo,
			Activo   : act,
			FechaCreacion :  time.Now(),   
			FechaModificacion :  time.Now(),		
			TipoEspacioFisicoId : te,	      
		}

		if _, err := models.AddEspacioFisico(&temp); err == nil {

		//if _, err := models.AddEspacioFisico(&v); err == nil {
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
// @Success 200 {object} models.EspacioFisico
// @Failure 404 not found resource
// @router /:id [get]
func (c *EspacioFisicoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetEspacioFisicoById(id)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		//-------------- Temporal: Cambio por transición ------- //
		var act string;
		if (v.Activo == true){
			act = "Activo"
		}else {
			act = "Inactivo"
		}

		te := &models.TipoEspacioFisico {
				Id: v.TipoEspacioFisicoId.Id,
				Nombre: v.TipoEspacioFisicoId.Nombre, 
				Descripcion: v.TipoEspacioFisicoId.Descripcion,
				CodigoAbreviacion: v.TipoEspacioFisicoId.CodigoAbreviacion,
				Activo: v.TipoEspacioFisicoId.Activo,
				FechaCreacion: v.TipoEspacioFisicoId.FechaCreacion,
				FechaModificacion: v.TipoEspacioFisicoId.FechaModificacion,	     		  
		}

		temp := models.EspacioFisico {
			Id: v.Id,
			Nombre: v.Nombre,   
			Codigo: v.CodigoAbreviacion,
			Estado: act,
			Descripcion:  v.Descripcion,    
			FechaCreacion :  v.FechaCreacion,   
			FechaModificacion :  v.FechaModificacion,		
			TipoEspacio : te,	
			//espacio_fisicoTipoespacio_fisico: field.espacio_fisicoTipoespacio_fisico,       
		}

	c.Data["json"] = temp
	//-------------- Temporal: Cambio por transición ------- //
	//c.Data["json"] = v
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
// @Success 200 {object} models.EspacioFisico
// @Failure 404 not found resource
// @router / [get]
func (c *EspacioFisicoController) GetAll() {
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
			if (k == "TipoEspacio.Id"){
				query["TipoEspacioFisicoId.Id"] = v
			}else{
				query[k] = v
			}
			
			
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
			c.Data["json"] = l
		}else{
			var temp []models.EspacioFisico
			for _, i := range l {
				field, _ := i.(models.EspacioFisicoV2)
				var act string;
				if (field.Activo == true){
					act = "Activo"
				}else {
					act = "Inactivo"
				}
			

				te := &models.TipoEspacioFisico {
					Id: field.TipoEspacioFisicoId.Id,
					Nombre: field.TipoEspacioFisicoId.Nombre, 
					Descripcion: field.TipoEspacioFisicoId.Descripcion,
					CodigoAbreviacion: field.TipoEspacioFisicoId.CodigoAbreviacion,
					Activo: field.TipoEspacioFisicoId.Activo,
					FechaCreacion: field.TipoEspacioFisicoId.FechaCreacion,
					FechaModificacion: field.TipoEspacioFisicoId.FechaModificacion,	     		  
				}
				
				x := models.EspacioFisico {
					Id: field.Id,
					Nombre: field.Nombre,   
					Codigo: field.CodigoAbreviacion,
					Estado: act, 
					Descripcion:  field.Descripcion,    
					FechaCreacion :  field.FechaCreacion,   
					FechaModificacion :  field.FechaModificacion,		
					TipoEspacio : te,	
					//espacio_fisicoTipoespacio_fisico: field.espacio_fisicoTipoespacio_fisico,       
				}

				temp = append(temp,x)
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
// @Description update the EspacioFisico
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.EspacioFisico	true		"body for EspacioFisico content"
// @Success 200 {object} models.EspacioFisico
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *EspacioFisicoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	infoDep, _ := models.GetEspacioFisicoById(id)
	v := models.EspacioFisico{Id: id}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		tef := &models.TipoEspacioFisicoV2{
			Id: v.TipoEspacio.Id,
		}

		v2 := models.EspacioFisicoV2{
			Id: id,
			Nombre: v.Nombre,
			Descripcion: v.Descripcion,
			CodigoAbreviacion: v.Codigo,
			Activo : infoDep.Activo,
			FechaCreacion : infoDep.FechaCreacion,
			FechaModificacion  : time.Now(),
			TipoEspacioFisicoId: tef,
		}
	
		if err := models.UpdateEspacioFisicoById(&v2); err == nil {
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
func (c *EspacioFisicoController) Delete() {
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
func (c *EspacioFisicoController) EspaciosHuerfanos() {
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
func (c *EspacioFisicoController) GetEspaciosFisicosHijosById() {
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
func (c *EspacioFisicoController) GetEspaciosFisicosPadresById() {
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