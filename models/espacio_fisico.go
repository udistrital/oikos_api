package models

import (
	"container/list"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/udistrital/utils_oas/formatdata"
)

// Traduce los selectores (según se usen en query, filter, offset)
// según corresponda a la jerarquía actual
func (d *EspacioFisicoV2) SelectorsFromV1(in []string) (out []string) {
	out = make([]string, len(in))
	for k, v := range in { // Iterar parametros especificados
		replace := ""
		for k2, v2 := range trEspacioFisicoV1 { // Iterar parametros a traducir
			if strings.HasPrefix(v, k2) {
				replace = strings.Replace(v, k2, v2, 1)
				// logs.Debug("k2:", k2, "v2:", v2)
				break
			}
		}
		if replace == "" {
			out[k] = v
		} else {
			out[k] = replace
		}
		// logs.Debug("v:", v, "replace:", replace)
	}
	return
}

// Ajusta los queries a la V2
func (d *EspacioFisicoV2) QueryFromV1(in map[string]string) (out map[string]string) {
	out = make(map[string]string)
	for k, v := range in { // Iterar cada criterio
		replaceK := ""
		replaceV := v
		for k2, v2 := range trEspacioFisicoV1 {
			if strings.HasPrefix(k, k2) {
				// Ajustar key, criterio
				replaceK = strings.Replace(k, k2, v2, 1)
				// Ajustar valores
				switch k2 {
				case "Estado":
					replaceV = fmt.Sprint(d.activoFromV1(v))
				}
				break
			}
		}
		if replaceK == "" {
			out[k] = v
		} else {
			out[replaceK] = replaceV
			// logs.Debug("k:", k, "replaceK:", replaceK, "replaceV:", replaceV)
		}
	}
	return
}

var elementMapEF = make(map[int]EspacioFisicoPadreHijo)
var lef = list.New()

const (
	EspacioFisicoEstadoActivo   = "Activo"
	EspacioFisicoEstadoInactivo = "Inactivo"
)

type EspacioFisico struct {
	Id          int                `orm:"column(id);pk;auto"`
	Estado      string             `orm:"column(estado)"`
	TipoEspacio *TipoEspacioFisico `orm:"column(tipo_espacio);rel(fk)"`
	Nombre      string             `orm:"column(nombre)"`
	Codigo      string             `orm:"column(codigo)"`
}

func (d *EspacioFisicoV2) FromV1(in EspacioFisico) error {
	if err := formatdata.FillStruct(in, &d); err != nil {
		return err
	}
	d.Activo = d.activoFromV1(in.Estado)
	d.CodigoAbreviacion = in.Codigo
	if in.TipoEspacio != nil {
		var esp TipoEspacioFisicoV2
		esp.FromV1(*in.TipoEspacio)
		d.TipoEspacioFisicoId = &esp
	}
	return nil
}
func (d *EspacioFisicoV2) activoFromV1(estado string) (activo bool) {
	return estado != EspacioFisicoEstadoInactivo
}
func (d *EspacioFisicoV2) estadoToV1(activo bool) (estado string) {
	if activo {
		return EspacioFisicoEstadoActivo
	} else {
		return EspacioFisicoEstadoInactivo
	}
}
func (d *EspacioFisicoV2) ToV1(out *EspacioFisico) error {
	if err := formatdata.FillStruct(d, &out); err != nil {
		return err
	}
	out.Codigo = d.CodigoAbreviacion
	out.Estado = d.estadoToV1(d.Activo)
	if d.TipoEspacioFisicoId != nil {
		var esp TipoEspacioFisico
		d.TipoEspacioFisicoId.ToV1(&esp)
		(*out).TipoEspacio = &esp
	}
	return nil
}

type EspacioFisicoV2 struct {
	Id                  int                  `orm:"column(id);pk;auto"`
	Nombre              string               `orm:"column(nombre)"`
	Descripcion         string               `orm:"column(descripcion);null"`
	CodigoAbreviacion   string               `orm:"column(codigo_abreviacion);null"`
	Activo              bool                 `orm:"column(activo)"`
	TipoTerrenoId       int                  `orm:"column(tipo_terreno_id)"`
	TipoEdificacionId   int                  `orm:"column(tipo_edificacion_id)"`
	TipoEspacioFisicoId *TipoEspacioFisicoV2 `orm:"column(tipo_espacio_fisico_id);rel(fk)"`
	FechaCreacion       time.Time            `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion   time.Time            `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}

func (t *EspacioFisicoV2) TableName() string {
	return "espacio_fisico"
}

var trEspacioFisicoV1 Diccionario

func init() {
	orm.RegisterModel(new(EspacioFisicoV2))
	trEspacioFisicoV1 = Diccionario{
		// V1   <---->     V2
		"Estado":      "Activo",
		"Codigo":      "CodigoAbreviacion",
		"TipoEspacio": "TipoEspacioFisicoId",
	}
}

// AddEspacioFisico insert a new EspacioFisico into database and returns
// last inserted Id on success.
func AddEspacioFisico(m *EspacioFisicoV2) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEspacioFisicoById retrieves EspacioFisico by Id. Returns error if
// Id doesn't exist
func GetEspacioFisicoById(id int) (v *EspacioFisicoV2, err error) {
	o := orm.NewOrm()
	v = &EspacioFisicoV2{Id: id}
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
	qs := o.QueryTable(new(EspacioFisicoV2)).RelatedSel(5)
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

	var l []EspacioFisicoV2
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
func UpdateEspacioFisicoById(m *EspacioFisicoV2) (err error) {
	o := orm.NewOrm()
	v := EspacioFisicoV2{Id: m.Id}
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
	v := EspacioFisicoV2{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EspacioFisicoV2{Id: id}); err == nil {
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
	num, err := o.Raw(
		`SELECT
			es.id, es.nombre AS nombre, es.codigo_abreviacion AS codigo, es.tipo_espacio_fisico_id AS tipo,
			CASE
				WHEN es.activo = TRUE
				THEN 'Activo'
				ELSE 'Inactivo'
			END estado
		FROM `+Esquema+`.espacio_fisico es WHERE es.tipo_espacio_fisico_id = ? AND es.id NOT IN
		(SELECT DISTINCT hijo_id FROM `+Esquema+`.espacio_fisico_padre)`,
		tipo_espacio).
		QueryRows(&espaciosHuerfanos)

	if err == nil {
		fmt.Println("Espacio físicos huerfanos encontrados: ", num)
	}
	return espaciosHuerfanos
}

func buscarEF(ef int) (padre EspacioFisicoPadreHijo) {

	var x EspacioFisicoPadreHijo
	for _, element := range elementMapEF {

		if ef == element.Id {
			x.Id = element.Id
			x.Nombre = element.Nombre
			x.Padre = element.Padre
			x.Hijo = element.Hijo

		}
	}

	return x
}

//Funcion recursiva que busca los espacios fisicos hijos a partir de un id del espacio físico padre
func getEspacioFisicoHijos(Padre *EspacioFisicoPadreHijo, padre int) (ef []EspacioFisicoPadreHijo) {

	for _, element := range elementMapEF {

		if padre == element.Padre {
			var x EspacioFisicoPadreHijo
			x.Id = element.Id
			x.Nombre = element.Nombre
			x.Padre = element.Padre
			j := &x
			x.Opciones = getEspacioFisicoHijos(j, element.Id)
			Padre.Opciones = append(Padre.Opciones, x)

		}

	}

	return Padre.Opciones

}

func GetEspaciosFisicosHijosById(espacioFisicoPadre int) (espaciosFisicos *EspacioFisicoPadreHijo, e error) {

	var espaciosFisicosHijos []EspacioFisicoPadreHijo

	qb, _ := orm.NewQueryBuilder("mysql")
	//buscar todos
	qb.Select("ef.id",
		"ef.nombre",
		"efp.padre_id",
		"efp.hijo_id").
		From(Esquema + ".espacio_fisico as ef").
		LeftJoin(Esquema + ".espacio_fisico_padre as efp").On("ef.id = efp.hijo_id").
		OrderBy("ef.id")

	sql := qb.String()

	o := orm.NewOrm()
	_, err := o.Raw(sql).QueryRows(&espaciosFisicosHijos)

	//TO MAP
	for _, s := range espaciosFisicosHijos {
		elementMapEF[s.Id] = s
	}

	c := buscarEF(espacioFisicoPadre)
	Cabeza := &c
	getEspacioFisicoHijos(Cabeza, espacioFisicoPadre)

	return Cabeza, err
}

func getEspaciosFisicosPadres(Hijo EspacioFisicoPadreHijo) (ef EspacioFisicoPadreHijo) {

	var x EspacioFisicoPadreHijo
	for _, element := range elementMapEF {

		if Hijo.Padre == element.Hijo {
			x.Id = element.Id
			x.Nombre = element.Nombre
			x.Padre = element.Padre
			x.Hijo = element.Hijo

			lef.PushFront(x)
			getEspaciosFisicosPadres(x)
		}
	}

	return x

}

func GetEspaciosFisicosPadresById(espacioFisicoHijo int) (espaciosFisicos []EspacioFisicoPadreHijo, e error) {

	var espacioFisicoPadres []EspacioFisicoPadreHijo
	var listaEspaciosFisicos []EspacioFisicoPadreHijo
	lef.Init()

	qb, _ := orm.NewQueryBuilder("mysql")
	//buscar todos
	qb.Select("ef.id",
		"ef.nombre",
		"efp.padre_id",
		"efp.hijo_id").
		From(Esquema + ".espacio_fisico as ef").
		LeftJoin(Esquema + ".espacio_fisico_padre as efp").On("ef.id = efp.hijo_id").
		OrderBy("ef.id")

	sql := qb.String()

	o := orm.NewOrm()
	_, err := o.Raw(sql).QueryRows(&espacioFisicoPadres)

	//TO MAP
	for _, s := range espacioFisicoPadres {
		elementMapEF[s.Id] = s
	}

	//Obtener informacion sobre espacio físico que se busca
	var Cola EspacioFisicoPadreHijo
	Cola.Id = elementMapEF[espacioFisicoHijo].Id
	Cola.Nombre = elementMapEF[espacioFisicoHijo].Nombre
	Cola.Padre = elementMapEF[espacioFisicoHijo].Padre
	Cola.Hijo = elementMapEF[espacioFisicoHijo].Hijo

	if Cola.Hijo != 0 {
		getEspaciosFisicosPadres(Cola)
		lef.PushBack(Cola)

		//Buscar cabeza de la lista
		p := buscarEF(lef.Front().Value.(EspacioFisicoPadreHijo).Padre)
		lef.PushFront(p)

		for temp := lef.Front(); temp != nil; temp = temp.Next() {
			listaEspaciosFisicos = append(listaEspaciosFisicos, temp.Value.(EspacioFisicoPadreHijo))
		}
	}

	return listaEspaciosFisicos, err
}
