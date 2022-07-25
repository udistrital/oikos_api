package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/udistrital/utils_oas/formatdata"
)

type TipoUsoEspacioFisico struct {
	Id              int            `orm:"column(id);pk;auto"`
	TipoUsoId       *TipoUso       `orm:"column(tipo_uso_id);rel(fk)"`
	EspacioFisicoId *EspacioFisico `orm:"column(espacio_fisico_id);rel(fk)"`
}

func (d *TipoUsoEspacioFisicoV2) FromV1(in TipoUsoEspacioFisico) (err error) {
	if err = formatdata.FillStruct(in, &d); err != nil {
		return
	}
	if in.TipoUsoId != nil {
		var tu TipoUsoV2
		tu.FromV1(*in.TipoUsoId)
		d.TipoUsoId = &tu
	}
	if in.EspacioFisicoId != nil {
		var ef EspacioFisicoV2
		ef.FromV1(*in.EspacioFisicoId)
		d.EspacioFisicoId = &ef
	}
	return
}

func (d *TipoUsoEspacioFisicoV2) ToV1(out *TipoUsoEspacioFisico) (err error) {
	formatdata.FillStruct(d, &out)
	if d.TipoUsoId != nil {
		var tu TipoUso
		d.TipoUsoId.ToV1(&tu)
		out.TipoUsoId = &tu
	}
	if d.EspacioFisicoId != nil {
		var ef EspacioFisico
		d.EspacioFisicoId.ToV1(&ef)
		out.EspacioFisicoId = &ef
	}
	return
}

type TipoUsoEspacioFisicoV2 struct {
	Id                int              `orm:"column(id);pk;auto"`
	TipoUsoId         *TipoUsoV2       `orm:"column(tipo_uso_id);rel(fk)"`
	EspacioFisicoId   *EspacioFisicoV2 `orm:"column(espacio_fisico_id);rel(fk)"`
	Activo            bool             `orm:"column(activo)"`
	FechaCreacion     time.Time        `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion time.Time        `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}

var trTipoUsoEspacioFisicoV1 Diccionario

func (t *TipoUsoEspacioFisicoV2) TableName() string {
	return "tipo_uso_espacio_fisico"
}

func init() {
	orm.RegisterModel(new(TipoUsoEspacioFisicoV2))

	trTipoUsoEspacioFisicoV1 = Diccionario{
		// v1        -->  v2
		"TipoUsoId":       "TipoUsoId",
		"EspacioFisicoId": "EspacioFisicoId",
	}
}

// AddTipoUsoEspacioFisico insert a new TipoUsoEspacioFisico into database and returns
// last inserted Id on success.
func AddTipoUsoEspacioFisico(m *TipoUsoEspacioFisicoV2) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTipoUsoEspacioFisicoById retrieves TipoUsoEspacioFisico by Id. Returns error if
// Id doesn't exist
func GetTipoUsoEspacioFisicoById(id int) (v *TipoUsoEspacioFisicoV2, err error) {
	o := orm.NewOrm()
	v = &TipoUsoEspacioFisicoV2{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTipoUsoEspacioFisico retrieves all TipoUsoEspacioFisico matches certain condition. Returns empty list if
// no records exist
func GetAllTipoUsoEspacioFisico(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TipoUsoEspacioFisicoV2)).RelatedSel(5)
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

	var l []TipoUsoEspacioFisicoV2
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
func (d *TipoUsoEspacioFisicoV2) SelectorsFromV1(in []string) (out []string) {
	out = make([]string, len(in))
	for k, v := range in { // Iterar parametros especificados
		// 1/3: Reemplazar "." por "__"
		temp := strings.Replace(v, ".", "__", -1)
		// 2/3: Trabajar sobre la parte inicial, correspondiente a esta entidad
		split := strings.SplitN(temp, "__", 2)
		if v, ok := trTipoUsoEspacioFisicoV1[split[0]]; ok {
			split[0] = v
			if len(split) > 1 { // Delegar la parte restante según la entidad (v2)
				switch v {
				case "TipoUsoId":
					aux := TipoUsoV2{}
					subqueryArr := aux.SelectorsFromV1([]string{split[1]})
					split[1] = subqueryArr[0]
				case "EspacioFisicoId":
					aux := EspacioFisicoV2{}
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
func (d *TipoUsoEspacioFisicoV2) QueryFromV1(in map[string]string) (out map[string]string) {
	out = make(map[string]string)
	for k, v := range in { // Iterar cada criterio
		// 1/3: Reemplazar "." por "__"
		temp := strings.Replace(k, ".", "__", -1)
		value := v
		// 2/3: Trabajar sobre la parte inicial, correspondiente a esta entidad
		split := strings.SplitN(temp, "__", 2)
		if v2, ok := trTipoUsoEspacioFisicoV1[split[0]]; ok {
			split[0] = v2
			if len(split) > 1 { // Delegar la parte restante según la entidad (v2)
				switch v2 {
				case "TipoUsoId":
					aux := TipoUsoV2{}
					subqueryArr := aux.QueryFromV1(map[string]string{split[1]: value})
					if len(subqueryArr) == 1 {
						for k3, v3 := range subqueryArr {
							split[1] = k3
							value = v3
						}
					}
				case "EspacioFisicoId":
					aux := EspacioFisicoV2{}
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

// UpdateTipoUsoEspacioFisico updates TipoUsoEspacioFisico by Id and returns error if
// the record to be updated doesn't exist
func UpdateTipoUsoEspacioFisicoById(m *TipoUsoEspacioFisicoV2) (err error) {
	o := orm.NewOrm()
	v := TipoUsoEspacioFisicoV2{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTipoUsoEspacioFisico deletes TipoUsoEspacioFisico by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTipoUsoEspacioFisico(id int) (err error) {
	o := orm.NewOrm()
	v := TipoUsoEspacioFisico{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TipoUsoEspacioFisicoV2{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
