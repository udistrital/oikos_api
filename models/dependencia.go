package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"container/list"
	"github.com/astaxie/beego/orm"
	"time"
)

var elementMap = make(map[int]DependenciaPadreHijo)
var l = list.New()

type Dependencia struct {
	Id                         int                           `orm:"column(id);pk;auto"`
	Nombre                     string                        `orm:"column(nombre)"`
	TelefonoDependencia        string                        `orm:"column(telefono_dependencia)"`
	CorreoElectronico          string                        `orm:"column(correo_electronico)"`
	DependenciaTipoDependencia []*DependenciaTipoDependencia `orm:"reverse(many)"`

}

type DependenciaV2 struct {
	Id     			 		   int    	`orm:"column(id);pk;auto"`
	Nombre            		   string 	`orm:"column(nombre)"`
	TelefonoDependencia        string    `orm:"column(telefono_dependencia)"`
	CorreoElectronico          string    `orm:"column(correo_electronico)"`
	Activo           		   bool      `orm:"column(activo)"`
	FechaCreacion     		   time.Time `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion          time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
	DependenciaTipoDependencia []*DependenciaTipoDependenciaV2 `orm:"reverse(many)"`

	
}

//Estructura para traer el ID y el nombre de cada proyecto curriculares
type ProyectosCurriculares struct {
	Id     int
	Nombre string
}

func (t *DependenciaV2) TableName() string {
	return "dependencia"
}

func init() {
	//orm.RegisterModel(new(Dependencia))
	orm.RegisterModel(new(DependenciaV2))
}



// AddDependencia insert a new Dependencia into database and returns
// last inserted Id on success.
func AddDependencia(m *DependenciaV2) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDependenciaById retrieves Dependencia by Id. Returns error if
// Id doesn't exist
func GetDependenciaById(id int) (v *DependenciaV2, err error) {
	o := orm.NewOrm()
	v = &DependenciaV2{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDependencia retrieves all Dependencia matches certain condition. Returns empty list if
// no records exist
func GetAllDependencia(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DependenciaV2)).RelatedSel(5)
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

	var l []DependenciaV2
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				o.LoadRelated(&v, "DependenciaTipoDependencia", 5)
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

// UpdateDependencia updates Dependencia by Id and returns error if
// the record to be updated doesn't exist
func UpdateDependenciaById(m *DependenciaV2) (err error) {
	o := orm.NewOrm()
	v := DependenciaV2{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDependencia deletes Dependencia by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDependencia(id int) (err error) {
	o := orm.NewOrm()
	v := DependenciaV2{Id: id}
	
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DependenciaV2{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//Se realiza sobrecarga de la función ProyectosPorFacultad que recibe como parámetros el id de la facultad y el nivel académico
func ProyectosPorFacultad(facultad int, nivel_academico string) (dependencia []ProyectosCurriculares) {
	
		//Conversión de entero a string
		id_facultad := strconv.Itoa(facultad)
		fmt.Println(id_facultad)
		fmt.Println(nivel_academico)
	
		o := orm.NewOrm()
		//Arreglo
		var proyectosCurriculares []ProyectosCurriculares
	
		if nivel_academico == "PREGRADO" {
	
			num, err := o.Raw(`SELECT DISTINCT ON (dh.id) dh.id AS id, dh.nombre AS nombre
												 FROM oikos.dependencia d INNER JOIN oikos.dependencia_padre dp ON d.id = dp.padre_id
												 INNER JOIN oikos.dependencia dh ON dh.id = dp.hija_id
												 INNER JOIN oikos.dependencia_tipo_dependencia dtd ON dh.id = dtd.dependencia_id
												 WHERE d.id = ` + id_facultad + ` AND dtd.tipo_dependencia_id = 14`).QueryRows(&proyectosCurriculares)
	
			if err == nil {
				fmt.Println("Proyectos curriculares encontrados: ", num)
			} else {
				fmt.Println("Este es el error ", err)
			}
	
		} else if nivel_academico == "POSGRADO" {
			num, err := o.Raw(`SELECT DISTINCT ON (dh.id) dh.id AS id, dh.nombre AS nombre
												 FROM oikos.dependencia d INNER JOIN oikos.dependencia_padre dp ON d.id = dp.padre_id
												 INNER JOIN oikos.dependencia dh ON dh.id = dp.hija_id
												 INNER JOIN oikos.dependencia_tipo_dependencia dtd ON dh.id = dtd.dependencia_id
												 WHERE d.id = ` + id_facultad + ` AND dtd.tipo_dependencia_id = 15`).QueryRows(&proyectosCurriculares)
	
			if err == nil {
				fmt.Println("Proyectos curriculares encontrados: ", num)
			} else {
				fmt.Println("Este es el error ", err)
			}
		} else if nivel_academico == "undefined" {
			num, err := o.Raw(`SELECT DISTINCT ON (dh.id) dh.id AS id, dh.nombre AS nombre
												 FROM oikos.dependencia d INNER JOIN oikos.dependencia_padre dp ON d.id = dp.padre_id
												 INNER JOIN oikos.dependencia dh ON dh.id = dp.hija_id
												 INNER JOIN oikos.dependencia_tipo_dependencia dtd ON dh.id = dtd.dependencia_id
												 WHERE d.id = ` + id_facultad + ` AND dtd.tipo_dependencia_id IN (1,14,15)`).QueryRows(&proyectosCurriculares)
	
			if err == nil {
				fmt.Println("Proyectos curriculares encontrados: ", num)
			} else {
				fmt.Println("Este es el error ", err)
			}
		}
	
		return proyectosCurriculares
	}
	
	//Funcion recursiva que busca las dependencias hijas a partir de un id de la dependencia padre
	func getDependenciasHijas(Padre *DependenciaPadreHijo,padre int)(dep []DependenciaPadreHijo){ 
		
		for _,element := range elementMap{
					
			if(padre == element.Padre){
				var x DependenciaPadreHijo
				x.Id = element.Id
				x.Nombre = element.Nombre
				x.Padre = element.Padre
				j := &x
				x.Opciones = getDependenciasHijas(j,element.Id)
				Padre.Opciones = append(Padre.Opciones,x)
		
			}
		   
		}
	
		return Padre.Opciones
	
	}
	
	
	//Funcion recursiva que busca las dependencias padre a partir de un id de la dependencia hija (hoja)
	func getDependenciasPadres(Hija DependenciaPadreHijo)(dep DependenciaPadreHijo){ 
	
	
		var x DependenciaPadreHijo
		for _,element := range elementMap{
	
			if(Hija.Padre == element.Hija){
				x.Id = element.Id
				x.Nombre = element.Nombre
				x.Padre = element.Padre
				x.Hija = element.Hija
	
				l.PushFront(x)
				getDependenciasPadres(x)
			}
		}
		
		return x
		
	}
	
	func buscarDep(dep int)(padre DependenciaPadreHijo){
	
		var x DependenciaPadreHijo
		for _,element := range elementMap{
	
			if(dep == element.Id){
				x.Id = element.Id
				x.Nombre = element.Nombre
				x.Padre = element.Padre
				x.Hija = element.Hija
			
			}
		}
	
		return x
	}
	
	func GetDependenciasPadresById(dependenciaHija int)(dependencias []DependenciaPadreHijo, e error){
		
		var dependenciaPadres []DependenciaPadreHijo
		var listaDependencias []DependenciaPadreHijo
		l.Init()
	
		qb, _ := orm.NewQueryBuilder("mysql")
		//buscar todos
		qb.Select("de.id",
			"de.nombre",
			"dep.padre_id",
			"dep.hija_id").
			From("oikos.dependencia as de").
			LeftJoin("oikos.dependencia_padre as dep").On("de.id = dep.hija_id").
			OrderBy("de.id")
	
		sql := qb.String()
	
		o := orm.NewOrm()
		_,err:=o.Raw(sql).QueryRows(&dependenciaPadres)
	
	
		//TO MAP
		for _, s := range dependenciaPadres {  
			elementMap[s.Id] = s 
		}
	
	   
		 //Obtener informacion sobre dependencia que se busca
		 var Cola DependenciaPadreHijo
		 Cola.Id = elementMap[dependenciaHija].Id;
		 Cola.Nombre = elementMap[dependenciaHija].Nombre
		 Cola.Padre = elementMap[dependenciaHija].Padre;
		 Cola.Hija = elementMap[dependenciaHija].Hija;
		
		 if (Cola.Hija != 0){
			getDependenciasPadres(Cola)
			l.PushBack(Cola)
	   
			//Buscar cabeza de la lista
			p := buscarDep(l.Front().Value.(DependenciaPadreHijo).Padre)
			l.PushFront(p)
	   
			for temp := l.Front(); temp != nil; temp = temp.Next() {
			   listaDependencias = append(listaDependencias,temp.Value.(DependenciaPadreHijo))
		   }
		 }
		
	
		
	
		return listaDependencias, err
	}
	
	func GetDependenciasHijasById(dependenciaPadre int)(dependencias *DependenciaPadreHijo, e error){
	
		
		var dependenciaHijas []DependenciaPadreHijo
	
		qb, _ := orm.NewQueryBuilder("mysql")
		//buscar todos
		qb.Select("de.id",
			"de.nombre",
			"dep.padre_id",
			"dep.hija_id").
			From("oikos.dependencia as de").
			LeftJoin("oikos.dependencia_padre as dep").On("de.id = dep.hija_id").
			OrderBy("de.id")
	
		sql := qb.String()
	
		o := orm.NewOrm()
		_,err:=o.Raw(sql).QueryRows(&dependenciaHijas)
	
		//TO MAP
		for _, s := range dependenciaHijas {  
			elementMap[s.Id] = s 
		}
		
		c:= buscarDep(dependenciaPadre)
		Cabeza := &c
		getDependenciasHijas(Cabeza,dependenciaPadre)
		 
	
		return Cabeza, err
	}
