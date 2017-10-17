package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type ResolucionVinculacion struct {
	Id              int       `orm:"column(id);pk;auto"`
	Estado          string    `orm:"column(estado)"`
	Numero          string    `orm:"column(numero)"`
	Vigencia        int       `orm:"column(vigencia)"`
	Facultad        string    `orm:"column(facultad)"`
	NivelAcademico  string    `orm:"column(nivel_academico)"`
	Dedicacion      string    `orm:"column(dedicacion)"`
	FechaExpedicion time.Time `orm:"column(fecha_expedicion);type(date)"`
}

func init() {
	orm.RegisterModel(new(ResolucionVinculacion))
}

func GetAllResolucionVinculacion() (arregloIDs []ResolucionVinculacion) {
	o := orm.NewOrm()
	var temp []ResolucionVinculacion
	//_, err := o.Raw("SELECT r.id_resolucion id, e.nombre_estado estado, r.numero_resolucion numero, r.vigencia vigencia, d.nombre facultad, rv.nivel_academico nivel_academico, rv.dedicacion dedicacion, r.fecha_expedicion fecha_expedicion FROM administrativa.resolucion r, administrativa.resolucion_vinculacion_docente rv, oikos.dependencia d, administrativa.resolucion_estado re, administrativa.estado_resolucion e WHERE rv.id_facultad=d.id AND r.id_resolucion=rv.id_resolucion AND re.resolucion=r.id_resolucion AND re.estado=e.id AND re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM administrativa.resolucion_estado re_aux WHERE re_aux.resolucion=r.id_resolucion) AND r.id_tipo_resolucion=1 ORDER BY id desc;").QueryRows(&temp)
	_, err := o.Raw("SELECT r.id_resolucion id, e.nombre_estado estado, r.numero_resolucion numero, r.vigencia vigencia, rv.nivel_academico nivel_academico, rv.dedicacion dedicacion, r.fecha_expedicion fecha_expedicion FROM administrativa.resolucion r, administrativa.resolucion_vinculacion_docente rv, administrativa.resolucion_estado re, administrativa.estado_resolucion e WHERE r.id_resolucion=rv.id_resolucion AND re.resolucion=r.id_resolucion AND re.estado=e.id AND re.fecha_registro=(SELECT MAX(re_aux.fecha_registro) FROM administrativa.resolucion_estado re_aux WHERE re_aux.resolucion=r.id_resolucion) AND r.id_tipo_resolucion=1 ORDER BY id desc;").QueryRows(&temp)
	
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp
}
