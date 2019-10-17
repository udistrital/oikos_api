package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"github.com/astaxie/beego/orm"
)

type EspacioFisico struct {
	Id         		    int                `orm:"column(id);pk;auto"`
	Nombre      		string             `orm:"column(nombre)"`
	Descripcion      	string             `orm:"column(descripcion);null"`
	Area              	float64            `orm:"column(area)"`
	Capacidad         	int                `orm:"column(capacidad)"`
	Codigo      		string             `orm:"column(codigo)"`
	Estado      		string             `orm:"column(estado)"`
	FechaCreacion       time.Time          `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion   time.Time           `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
	TipoEspacio        *TipoEspacioFisico   `orm:"column(tipo_espacio);rel(fk)"`

}

type EspacioFisicoV2 struct {

	Id                int    			 		`orm:"column(id);pk;auto"`
	Nombre            string             		`orm:"column(nombre)"`
	Descripcion       string             		`orm:"column(descripcion);null"`
	Area              float64           		`orm:"column(area)"`
	Capacidad         int                		`orm:"column(capacidad)"`
	CodigoAbreviacion string             		`orm:"column(codigo_abreviacion);null"`
	Activo            bool               		`orm:"column(activo)"`
	FechaCreacion     time.Time          		`orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion time.Time          		`orm:"column(fecha_modificacion);type(timestamp without time zone)"`
	TipoEspacio 	  *TipoEspacioFisicoV2      `orm:"column(tipo_espacio_fisico_id);rel(fk)"`
}

func (t *EspacioFisicoV2) TableName() string {
	return "espacio_fisico"
}

func init() {
	orm.RegisterModel(new(EspacioFisicoV2))
}

// AddEspacioFisico insert a new EspacioFisico into database and returns
// last inserted Id on success.
func AddEspacioFisico(m *EspacioFisico) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEspacioFisicoById retrieves EspacioFisico by Id. Returns error if
// Id doesn't exist
func GetEspacioFisicoById(id int) (v *EspacioFisico, err error) {
	o := orm.NewOrm()
	v = &EspacioFisico{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEspacioFisico retrieves all EspacioFisico matches certain condition. Returns empty list if
// no records exist
func GetAllEspacioFisico(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EspacioFisico)).RelatedSel(5)
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []EspacioFisico
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateEspacioFisico updates EspacioFisico by Id and returns error if
// the record to be updated doesn't exist
func UpdateEspacioFisicoById(m *EspacioFisico) (err error) {
	o := orm.NewOrm()
	v := EspacioFisico{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEspacioFisico deletes EspacioFisico by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEspacioFisico(id int) (err error) {
	o := orm.NewOrm()
	v := EspacioFisico{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EspacioFisico{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//Función que construye los menús
func EspacioFisicosHuerfanos(tipo_espacio int) (espacios []EspacioFisico) {
	o := orm.NewOrm()
	//Conversión de entero a string
	//tipo_espacio_fisico := strconv.Itoa(tipo_espacio)
	fmt.Print(tipo_espacio)
	//Arreglo vacio que se llenará con los espacios físicos huerfanos
	var espaciosHuerfanos []EspacioFisico
	//Consulta SQL que busca los espacios físicos huerfanos
	num, err := o.Raw(`SELECT es.id, es.nombre AS nombre, es.codigo AS codigo, es.tipo_espacio AS tipo, es.estado AS estado
										 FROM oikos.espacio_fisico es WHERE es.tipo_espacio = ? AND es.id NOT IN
										 (SELECT DISTINCT hijo FROM oikos.espacio_fisico_padre)`, tipo_espacio).QueryRows(&espaciosHuerfanos)

	if err == nil {
		fmt.Println("Espacio físicos huerfanos encontrados: ", num)
	}
	return espaciosHuerfanos
}
