package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type PersonaEscalafon struct {
	Id              int    `orm:"column(id);pk;auto"`
	PrimerNombre    string `orm:"column(primer_nombre)"`
	SegundoNombre   string `orm:"column(segundo_nombre)"`
	PrimerApellido  string `orm:"column(primer_apellido)"`
	SegundoApellido string `orm:"column(segundo_apellido)"`
	Escalafon       string `orm:"column(escalafon)"`
}

func init() {
	orm.RegisterModel(new(PersonaEscalafon))
}

func GetAllPersonaEscalafon() (arregloIDs []PersonaEscalafon) {
	o := orm.NewOrm()
	var temp []PersonaEscalafon
	_, err := o.Raw("SELECT ipn.num_documento_persona id, ipn.primer_nombre, ipn.segundo_nombre, ipn.primer_apellido, ipn.segundo_apellido, e.nombre_escalafon escalafon FROM agora.informacion_persona_natural ipn INNER JOIN agora.informacion_proveedor ip ON ipn.num_documento_persona = ip.num_documento INNER JOIN administrativa.escalafon_persona ep ON ip.id_proveedor = ep.id_persona_natural INNER JOIN administrativa.escalafon e ON ep.id_escalafon = e.id_escalafon;").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp
}