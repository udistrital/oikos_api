package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"github.com/astaxie/beego/orm"
)

type DependenciaPadre struct {
	Id    int          `orm:"column(id);pk;auto"`
	Padre *Dependencia `orm:"column(padre);rel(fk)"`
	Hija  *Dependencia `orm:"column(hija);rel(fk)"`
}



func (t *DependenciaPadre) TableName() string {
	return "dependencia_padre"
}

func init() {
	orm.RegisterModel(new(DependenciaPadre))
}

// AddDependenciaPadre insert a new DependenciaPadre into database and returns
// last inserted Id on success.
func AddDependenciaPadre(m *DependenciaPadre) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDependenciaPadreById retrieves DependenciaPadre by Id. Returns error if
// Id doesn't exist
func GetDependenciaPadreById(id int) (v *DependenciaPadre, err error) {
	o := orm.NewOrm()
	v = &DependenciaPadre{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDependenciaPadre retrieves all DependenciaPadre matches certain condition. Returns empty list if
// no records exist
func GetAllDependenciaPadre(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DependenciaPadre)).RelatedSel(5)
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

	var l []DependenciaPadre
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

// UpdateDependenciaPadre updates DependenciaPadre by Id and returns error if
// the record to be updated doesn't exist
func UpdateDependenciaPadreById(m *DependenciaPadre) (err error) {
	o := orm.NewOrm()
	v := DependenciaPadre{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDependenciaPadre deletes DependenciaPadre by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDependenciaPadre(id int) (err error) {
	o := orm.NewOrm()
	v := DependenciaPadre{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DependenciaPadre{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//Función que busca las dependencias de tipo facultad
func Facultades() (facultad []Tree) {
	
		//Declaración objeto ORM
		o := orm.NewOrm()
	
		//Arreglo que tendra las facultades encontradas
		var facultades []Tree
	
		_, err := o.Raw(`SELECT dh.id AS id, dh.nombre AS nombre
											 FROM oikos.dependencia d INNER JOIN oikos.dependencia_padre dp ON d.id = dp.padre
						 INNER JOIN oikos.dependencia dh ON dh.id = dp.hija
											 INNER JOIN oikos.dependencia_tipo_dependencia dtd ON dh.id = dtd.dependencia_id
											 WHERE dtd.tipo_dependencia_id = 2`).QueryRows(&facultades)
	
		if err == nil {
		
			//For para que recorra los Ids en busca de hijos
			for i := 0; i < len(facultades); i++ {
				//Me verifica que los Id tengan hijos
				ProyectosCurricularesPorFacultad(&facultades[i])
			}
		}
		return facultades
	}
	
	//Función que busca las dependencias de tipo facultad
	func ProyectosCurricularesPorFacultad(Facultad *Tree) (proyectos []Tree) {
	
		//Declaración objeto ORM
		o := orm.NewOrm()
	
		//Conversión de entero a string
		padre := strconv.Itoa(Facultad.Id)
	
		//Arreglo que tendra las facultades encontradas
		var proyectos_curriculares []Tree
	
		_, err := o.Raw(`SELECT DISTINCT de.id, de.nombre, dep.padre, dep.hija
											 FROM oikos.dependencia AS de
											 LEFT JOIN oikos.dependencia_padre AS dep ON de.id = dep.hija
											 INNER JOIN oikos.dependencia_tipo_dependencia dtd ON dep.hija = dtd.dependencia_id
											 WHERE dep.padre =` + padre + ` AND dtd.tipo_dependencia_id IN (1,14,15) ORDER BY de.id`).QueryRows(&proyectos_curriculares)
	
		if err == nil {
				
			//Llena el elemento Opciones en la estructura del menú padre
			Facultad.Opciones = &proyectos_curriculares
		}
		return proyectos_curriculares
	}
