package models

import (
	"github.com/astaxie/beego/orm"
)

var proyectosCurricularesMap = make(map[int]DependenciaPadreHijo)

func getProyectosPorFacultad(Padre *DependenciaPadreHijo, padre int) (dep []DependenciaPadreHijo) {

	for _, element := range proyectosCurricularesMap {

		if padre == element.Padre {
			var x DependenciaPadreHijo
			x.Id = element.Id
			x.Nombre = element.Nombre
			x.Padre = element.Padre
			j := &x
			x.Opciones = getProyectosPorFacultad(j, element.Id)
			Padre.Opciones = append(Padre.Opciones, x)

		}

	}

	return Padre.Opciones

}

func GetAllProyectosByFacultades() (facultad []DependenciaPadreHijo, e error) {
	//Declaración objeto ORM
	o := orm.NewOrm()

	var facultades []DependenciaPadreHijo

	//Se buscan todas las facultades de la Universidad.
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("d.id AS id",
		"d.nombre AS nombre").
		From(Esquema + ".dependencia d").
		InnerJoin(Esquema + ".dependencia_tipo_dependencia dp").On(" d.id = dp.dependencia_id").
		Where("dp.tipo_dependencia_id = 2")

	sql := qb.String()
	_, err := o.Raw(sql).QueryRows(&facultades)

	if err == nil {

		var pc []DependenciaPadreHijo
		qb, _ := orm.NewQueryBuilder("mysql")

		//buscar todos los proyectos curriculares
		qb.Select("de.id",
			"de.nombre",
			"dep.padre_id",
			"dep.hija_id").
			From(Esquema + ".dependencia as de").
			LeftJoin(Esquema + ".dependencia_padre as dep").On("de.id = dep.hija_id").
			InnerJoin(Esquema + ".dependencia_tipo_dependencia dtd").On("dep.hija_id = dtd.dependencia_id").
			Where("dtd.tipo_dependencia_id IN (1,14,15)").
			OrderBy("de.id")

		sql := qb.String()

		o := orm.NewOrm()
		_, err := o.Raw(sql).QueryRows(&pc)

		if err == nil {
			//TO MAP
			for _, s := range pc {
				proyectosCurricularesMap[s.Id] = s
			}

			//Se busca por cada facultad sus proyectos curriculares
			for i := 0; i < len(facultades); i++ {
				getProyectosPorFacultad(&facultades[i], facultades[i].Id)
			}

		}

	}

	return facultades, err

}

func GetAllProyectosByFacultadId(idFacultad int) (facultad []DependenciaPadreHijo, e error) {
	//Declaración objeto ORM
	o := orm.NewOrm()

	var facultades []DependenciaPadreHijo

	//Se buscan todas las facultades de la Universidad.
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("d.id AS id",
		"d.nombre AS nombre").
		From(Esquema + ".dependencia d").
		InnerJoin(Esquema + ".dependencia_tipo_dependencia dp").On(" d.id = dp.dependencia_id").
		Where("dp.tipo_dependencia_id = 2").
		And("d.id = ?")

	sql := qb.String()
	_, err := o.Raw(sql, idFacultad).QueryRows(&facultades)

	if err == nil {

		var pc []DependenciaPadreHijo
		qb, _ := orm.NewQueryBuilder("mysql")

		//buscar todos los proyectos curriculares
		qb.Select("de.id",
			"de.nombre",
			"dep.padre_id",
			"dep.hija_id").
			From(Esquema + ".dependencia as de").
			LeftJoin(Esquema + ".dependencia_padre as dep").On("de.id = dep.hija_id").
			InnerJoin(Esquema + ".dependencia_tipo_dependencia dtd").On("dep.hija_id = dtd.dependencia_id").
			Where("dtd.tipo_dependencia_id IN (1,14,15)").
			OrderBy("de.id")

		sql := qb.String()

		o := orm.NewOrm()
		_, err := o.Raw(sql).QueryRows(&pc)

		if err == nil {
			//TO MAP
			for _, s := range pc {
				proyectosCurricularesMap[s.Id] = s
			}

			//Se busca por cada facultad sus proyectos curriculares
			for i := 0; i < len(facultades); i++ {
				getProyectosPorFacultad(&facultades[i], facultades[i].Id)
			}

		}

	}

	return facultades, err

}
