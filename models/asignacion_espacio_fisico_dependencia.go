package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	"github.com/udistrital/utils_oas/formatdata"
)

type AsignacionEspacioFisicoDependencia struct {
	Id               int            `orm:"column(id);pk;auto"`
	Estado           string         `orm:"column(estado)"`
	FechaInicio      time.Time      `orm:"column(fecha_inicio);type(date)"`
	FechaFin         time.Time      `orm:"column(fecha_fin);type(date);null"`
	DocumentoSoporte string         `orm:"column(documento_soporte)"`
	EspacioFisicoId  *EspacioFisico `orm:"column(espacio_fisico_id);rel(fk)"`
	DependenciaId    *Dependencia   `orm:"column(dependencia_id);rel(fk)"`
}

const (
	AsignacionEspacioFisicoDependenciaEstadoActivo   = "Activo"
	AsignacionEspacioFisicoDependenciaEstadoInactivo = "Inactivo"
)

func (d *AsignacionEspacioFisicoDependenciaV2) FromV1(in AsignacionEspacioFisicoDependencia) (err error) {
	if err = formatdata.FillStruct(in, &d); err != nil {
		return
	}
	d.Activo = in.Estado != AsignacionEspacioFisicoDependenciaEstadoInactivo
	d.DocumentoSoporte, _ = strconv.Atoi(in.DocumentoSoporte)
	if in.EspacioFisicoId != nil {
		var ef EspacioFisicoV2
		ef.FromV1(*in.EspacioFisicoId)
		d.EspacioFisicoId = &ef
	}
	if in.DependenciaId != nil {
		var dep DependenciaV2
		dep.FromV1(*in.DependenciaId)
		d.DependenciaId = &dep
	}
	return
}
func (d *AsignacionEspacioFisicoDependenciaV2) ToV1(out *AsignacionEspacioFisicoDependencia) (err error) {
	formatdata.FillStruct(d, &out)
	if d.Activo {
		out.Estado = AsignacionEspacioFisicoDependenciaEstadoActivo
	} else {
		out.Estado = AsignacionEspacioFisicoDependenciaEstadoInactivo
	}
	out.DocumentoSoporte = fmt.Sprint(d.DocumentoSoporte)
	if d.EspacioFisicoId != nil {
		var ef EspacioFisico
		d.EspacioFisicoId.ToV1(&ef)
		(*out).EspacioFisicoId = &ef
	}
	if d.DependenciaId != nil {
		var dep Dependencia
		d.DependenciaId.ToV1(&dep)
		(*out).DependenciaId = &dep
	}
	logs.Debug("convertido:")
	formatdata.JsonPrint(out)
	fmt.Println()
	return
}

type AsignacionEspacioFisicoDependenciaV2 struct {
	Id                int              `orm:"column(id);pk;auto"`
	EspacioFisicoId   *EspacioFisicoV2 `orm:"column(espacio_fisico_id);rel(fk)"`
	DependenciaId     *DependenciaV2   `orm:"column(dependencia_id);rel(fk)"`
	Activo            bool             `orm:"column(activo)"`
	FechaInicio       time.Time        `orm:"column(fecha_inicio);type(timestamp without time zone)"`
	FechaFin          time.Time        `orm:"column(fecha_fin);type(timestamp without time zone);null"`
	DocumentoSoporte  int              `orm:"column(documento_soporte);null"`
	FechaCreacion     time.Time        `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion time.Time        `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}

var trAsignacionEspacioFisicoDependenciaV1 Diccionario

func (t *AsignacionEspacioFisicoDependenciaV2) TableName() string {
	return "asignacion_espacio_fisico_dependencia"
}

func init() {
	orm.RegisterModel(new(AsignacionEspacioFisicoDependenciaV2))

	trAsignacionEspacioFisicoDependenciaV1 = Diccionario{
		"EspacioFisicoId": "EspacioFisicoId",
		"DependenciaId":   "DependenciaId",
	}
}

// AddAsignacionEspacioFisicoDependencia insert a new AsignacionEspacioFisicoDependencia into database and returns
// last inserted Id on success.
func AddAsignacionEspacioFisicoDependencia(m *AsignacionEspacioFisicoDependenciaV2) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAsignacionEspacioFisicoDependenciaById retrieves AsignacionEspacioFisicoDependencia by Id. Returns error if
// Id doesn't exist
func GetAsignacionEspacioFisicoDependenciaById(id int) (v *AsignacionEspacioFisicoDependenciaV2, err error) {
	o := orm.NewOrm()
	v = &AsignacionEspacioFisicoDependenciaV2{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAsignacionEspacioFisicoDependencia retrieves all AsignacionEspacioFisicoDependencia matches certain condition. Returns empty list if
// no records exist
func GetAllAsignacionEspacioFisicoDependencia(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AsignacionEspacioFisicoDependenciaV2)).RelatedSel(5)
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

	var l []AsignacionEspacioFisicoDependenciaV2
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
func (d *AsignacionEspacioFisicoDependenciaV2) SelectorsFromV1(in []string) (out []string) {
	out = make([]string, len(in))
	for k, v := range in { // Iterar parametros especificados
		// 1/3: Reemplazar "." por "__"
		temp := strings.Replace(v, ".", "__", -1)
		// 2/3: Trabajar sobre la parte inicial, correspondiente a esta entidad
		split := strings.SplitN(temp, "__", 2)
		if v, ok := trAsignacionEspacioFisicoDependenciaV1[split[0]]; ok {
			split[0] = v
			if len(split) > 1 { // Delegar la parte restante según la entidad (v2)
				switch v {
				case "EspacioFisicoId":
					aux := EspacioFisicoV2{}
					subqueryArr := aux.SelectorsFromV1([]string{split[1]})
					split[1] = subqueryArr[0]
				case "DependenciaId":
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
func (d *AsignacionEspacioFisicoDependenciaV2) QueryFromV1(in map[string]string) (out map[string]string) {
	out = make(map[string]string)
	for k, v := range in { // Iterar cada criterio
		// 1/3: Reemplazar "." por "__"
		temp := strings.Replace(k, ".", "__", -1)
		value := v
		// 2/3: Trabajar sobre la parte inicial, correspondiente a esta entidad
		split := strings.SplitN(temp, "__", 2)
		if v2, ok := trAsignacionEspacioFisicoDependenciaV1[split[0]]; ok {
			split[0] = v2
			if len(split) > 1 { // Delegar la parte restante según la entidad (v2)
				switch v2 {
				case "EspacioFisicoId":
					aux := EspacioFisicoV2{}
					subqueryArr := aux.QueryFromV1(map[string]string{split[1]: value})
					if len(subqueryArr) == 1 {
						for k3, v3 := range subqueryArr {
							split[1] = k3
							value = v3
						}
					}
				case "DependenciaId":
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

// UpdateAsignacionEspacioFisicoDependencia updates AsignacionEspacioFisicoDependencia by Id and returns error if
// the record to be updated doesn't exist
func UpdateAsignacionEspacioFisicoDependenciaById(m *AsignacionEspacioFisicoDependenciaV2) (err error) {
	o := orm.NewOrm()
	v := AsignacionEspacioFisicoDependenciaV2{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAsignacionEspacioFisicoDependencia deletes AsignacionEspacioFisicoDependencia by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAsignacionEspacioFisicoDependencia(id int) (err error) {
	o := orm.NewOrm()
	v := AsignacionEspacioFisicoDependenciaV2{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AsignacionEspacioFisicoDependenciaV2{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
