package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
)

type DependenciaPadre struct {
	Id    int          `orm:"column(id);pk;auto"`
	Padre *Dependencia `orm:"column(padre);rel(fk)"`
	Hija  *Dependencia `orm:"column(hija);rel(fk)"`
}

func (d *DependenciaPadreV2) FromV1(in DependenciaPadre) error {
	if err := formatdata.FillStruct(in, &d); err != nil {
		return err
	}
	if in.Padre != nil {
		var dep DependenciaV2
		dep.FromV1(*in.Padre)
		d.PadreId = &dep
	}
	if in.Hija != nil {
		var dep DependenciaV2
		dep.FromV1(*in.Hija)
		d.HijaId = &dep
	}
	return nil
}
func (d *DependenciaPadreV2) ToV1(out *DependenciaPadre) error {
	if err := formatdata.FillStruct(d, out); err != nil {
		return err
	}
	if d.PadreId != nil {
		var dep Dependencia
		d.PadreId.ToV1(&dep)
		out.Padre = &dep
	}
	if d.HijaId != nil {
		var dep Dependencia
		d.HijaId.ToV1(&dep)
		out.Hija = &dep
	}
	return nil
}

type DependenciaPadreV2 struct {
	Id                int            `orm:"column(id);pk;auto"`
	PadreId           *DependenciaV2 `orm:"column(padre_id);rel(fk)"`
	HijaId            *DependenciaV2 `orm:"column(hija_id);rel(fk)"`
	Activo            bool           `orm:"column(activo)"`
	FechaCreacion     time.Time      `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion time.Time      `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}

//Estructura para construir el arbol de dependencia
type TreeDependencia struct {
	Id       int
	Nombre   string
	Opciones *[]TreeDependencia
}

var trDependenciaPadreV1 Diccionario

func (t *DependenciaPadreV2) TableName() string {
	return "dependencia_padre"
}

func init() {
	orm.RegisterModel(new(DependenciaPadreV2))

	trDependenciaPadreV1 = Diccionario{
		"Padre": "PadreId",
		"Hija":  "HijaId",
	}
}

// AddDependenciaPadre insert a new DependenciaPadre into database and returns
// last inserted Id on success.
func AddDependenciaPadre(m *DependenciaPadreV2) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDependenciaPadreById retrieves DependenciaPadre by Id. Returns error if
// Id doesn't exist
func GetDependenciaPadreById(id int) (v *DependenciaPadreV2, err error) {
	o := orm.NewOrm()
	v = &DependenciaPadreV2{Id: id}
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
	qs := o.QueryTable(new(DependenciaPadreV2)).RelatedSel(5)
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

	var l []DependenciaPadreV2
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

// Traduce los selectores (según se usen en query, filter, offset)
// según corresponda a la jerarquía actual
func (d *DependenciaPadreV2) SelectorsFromV1(in []string) (out []string) {
	out = make([]string, len(in))
	for k, v := range in { // Iterar parametros especificados
		// 1/3: Reemplazar "." por "__"
		temp := strings.Replace(v, ".", "__", -1)
		// 2/3: Trabajar sobre la parte inicial, correspondiente a esta entidad
		split := strings.SplitN(temp, "__", 2)
		if v, ok := trDependenciaPadreV1[split[0]]; ok {
			split[0] = v
			if len(split) > 1 { // Delegar la parte restante según la entidad (v2)
				switch v {
				case "PadreId", "HijaId":
					aux := DependenciaV2{}
					subqueryArr := aux.SelectorsFromV1([]string{split[1]})
					split[1] = subqueryArr[0]
				}
			}
		}
		// 3/3: Combinar el resultado
		temp = strings.Join(split, "__")
		out[k] = temp
	}
	return
}

// Ajusta los queries a la V2
func (d *DependenciaPadreV2) QueryFromV1(in map[string]string) (out map[string]string) {
	out = make(map[string]string)
	for k, v := range in { // Iterar cada criterio
		// 1/3: Reemplazar "." por "__"
		temp := strings.Replace(k, ".", "__", -1)
		value := v
		// 2/3: Trabajar sobre la parte inicial, correspondiente a esta entidad
		split := strings.SplitN(temp, "__", 2)
		if v2, ok := trDependenciaPadreV1[split[0]]; ok {
			split[0] = v2
			if len(split) > 1 { // Delegar la parte restante según la entidad (v2)
				switch v2 {
				case "PadreId", "HijaId":
					aux := DependenciaV2{}
					subqueryArr := aux.QueryFromV1(map[string]string{split[1]: value})
					if len(subqueryArr) == 1 {
						for k3, v3 := range subqueryArr {
							split[1] = k3
							value = v3
						}
					}
				}
			}
		}
		// 3/3: Combinar el resultado
		temp = strings.Join(split, "__")
		out[temp] = value
	}
	// logs.Debug("in:", in, "out:", out)
	return
}

// UpdateDependenciaPadre updates DependenciaPadre by Id and returns error if
// the record to be updated doesn't exist
func UpdateDependenciaPadreById(m *DependenciaPadreV2) (err error) {
	o := orm.NewOrm()
	v := DependenciaPadreV2{Id: m.Id}
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
	v := DependenciaPadreV2{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DependenciaPadreV2{Id: id}); err == nil {
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
	_, err := o.Raw(
		`SELECT dh.id AS id, dh.nombre AS nombre
		FROM ` + Esquema + `.dependencia d INNER JOIN ` + Esquema + `.dependencia_padre dp ON d.id = dp.padre_id
		INNER JOIN ` + Esquema + `.dependencia dh ON dh.id = dp.hija_id
		INNER JOIN ` + Esquema + `.dependencia_tipo_dependencia dtd ON dh.id = dtd.dependencia_id
		WHERE dtd.tipo_dependencia_id = 2`).
		QueryRows(&facultades)

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

	_, err := o.Raw(
		`SELECT DISTINCT de.id, de.nombre, dep.padre_id, dep.hija_id
		FROM `+Esquema+`.dependencia AS de
		LEFT JOIN `+Esquema+`.dependencia_padre AS dep ON de.id = dep.hija_id
		INNER JOIN `+Esquema+`.dependencia_tipo_dependencia dtd ON dep.hija_id = dtd.dependencia_id
		WHERE dep.padre_id = ? AND dtd.tipo_dependencia_id IN (1,14,15) ORDER BY de.id`,
		padre).
		QueryRows(&proyectos_curriculares)
	if err == nil {

		//Llena el elemento Opciones en la estructura del menú padre
		Facultad.Opciones = &proyectos_curriculares
	}
	return proyectos_curriculares
}

//Función que busca las dependencias que no tengan asignadas padre
func ConstruirDependenciasPadre() (dependencias []TreeDependencia) {
	o := orm.NewOrm()
	//Arreglo
	var dependenciaPadres []TreeDependencia
	num, err := o.Raw(
		`SELECT de.id AS id, de.nombre AS nombre, dep.padre_id AS padre
		FROM ` + Esquema + `.dependencia
		AS de left join ` + Esquema + `.dependencia_padre
		AS dep ON de.id = dep.hija_id
		WHERE padre_id IS NULL ORDER BY de.id`).
		QueryRows(&dependenciaPadres)

	if err == nil {
		fmt.Println("Dependencias padre encontradas: ", num)
		//For para que recorra los Ids en busca de hijos
		for i := 0; i < len(dependenciaPadres); i++ {
			//Me verifica que los Id tengan hijos
			ConstruirDependenciasHijas(&dependenciaPadres[i])
		}
	}
	return dependenciaPadres
}

//Función que busca los hijos de los padres encontrados en la función anterior
func ConstruirDependenciasHijas(Padre *TreeDependencia) (dependencias []TreeDependencia) {
	o := orm.NewOrm()
	//Conversión de entero a string
	padre := strconv.Itoa(Padre.Id)

	//Arreglo
	var dependenciaHijas []TreeDependencia

	num, err := o.Raw(
		`SELECT de.id, de.nombre, dep.padre_id, dep.hija_id
		FROM `+Esquema+`.dependencia AS de
		LEFT JOIN `+Esquema+`.dependencia_padre AS dep ON de.id = dep.hija_id
		WHERE dep.padre_id = ? ORDER BY de.id`,
		padre).
		QueryRows(&dependenciaHijas)

	//Condicional si el error es nulo
	if err == nil {
		fmt.Println("Dependencias Hijas encontradas: ", num)

		//Llena el elemento Opciones en la estructura del menú padre
		Padre.Opciones = &dependenciaHijas

		//For que recorre el subMenu en busca de hijos (Recursiva)
		for i := 0; i < len(dependenciaHijas); i++ {

			//Me verifica que los Id tengan hijos
			ConstruirDependenciasHijas(&dependenciaHijas[i])
		}
	}
	return dependenciaHijas
}

func TRDependenciaPadre(m *DependenciaPadreV2) (id int64, err error) {
	o := orm.NewOrm()
	Hija := m.HijaId
	Hija.Activo = true
	Hija.FechaCreacion = time.Now()
	Hija.FechaModificacion = time.Now()
	o.Begin()
	id, err = o.Insert(Hija)
	if err != nil {
		o.Rollback()
	} else {
		m.Activo = true
		m.FechaCreacion = time.Now()
		m.FechaModificacion = time.Now()
		_, err = o.Insert(m)
		if err != nil {
			o.Rollback()
		}
	}

	o.Commit()
	return
}
