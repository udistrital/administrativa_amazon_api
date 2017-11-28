package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ProveedorContratoPersona struct {
	Id int	`orm:"column(num_documento)"`
	Tipo_persona string `orm:"column(tipopersona)"`
	Direccion string `orm:"column(direccion)"`
	Correo string `orm:"column(correo)"`
	Nombre_completo string `orm:"column(nom_proveedor)"`
	Primer_Apellido string `orm:"column(primer_apellido)"`
	Segundo_Apellido string `orm:"column(segundo_apellido)"`
	Primer_Nombre string `orm:"column(primer_nombre)"`
	Segundo_Nombre string `orm:"column(segundo_nombre)"`
	Tipo_documento string `orm:"column(tipo_documento)"`
	Numero_contrato string `orm:"column(numero_contrato)"`
	Vigencia_contrato int `orm:"column(vigencia)"`
	Objeto_contrato string `orm:"column(objeto_contrato)"`
	Valor_contrato string `orm:"column(valor_contrato)"`
	Fecha_registro time.Time `orm:"column(fecha_registro);type(date)"`
	Observaciones string `orm:"column(observaciones)"`
	Justificacion string `orm:"column(justificacion)"`
	Unidad_ejecutora int `orm:"column(unidad_ejecutora)"`
	Ordenador_gasto string `orm:"column(ordenador_gasto)"`
	Numero_solicitud_necesidad int `orm:"column(numero_solicitud_necesidad)"`
	Numero_cdp int `orm:"column(numero_cdp)"`
	Unidad_ejecucion int `orm:"column(unidad_ejecucion)"`
	Id_proveedor string `orm:"column(id_proveedor)"`
	Consecutivo_contrato string `orm:"column(consecutivo_contrato_suscrito)"`
	Numero_suscrito string `orm:"column(numero_contrato_suscrito)"`
}

func init() {
	orm.RegisterModel(new(ProveedorContratoPersona))
}

func VigenciaProveedorContratoPersona(vigencia string) (proveedor_contrato_persona []ProveedorContratoPersona, err error) {
	o := orm.NewOrm()
	var temp []ProveedorContratoPersona
	_,err = o.Raw("SELECT * FROM agora.informacion_persona_natural INNER JOIN agora.informacion_proveedor ON agora.informacion_proveedor.num_documento = agora.informacion_persona_natural.num_documento_persona INNER JOIN argo.contrato_general ON agora.informacion_proveedor.id_proveedor = argo.contrato_general.contratista INNER JOIN argo.contrato_suscrito ON argo.contrato_suscrito.numero_contrato=argo.contrato_general.numero_contrato AND argo.contrato_suscrito.vigencia=argo.contrato_general.vigencia where argo.contrato_general.vigencia="+vigencia+";").QueryRows(&temp)
	if err == nil {
		return temp,nil
	}
	return nil, err
}

func ProveedorVigenciaContrato(contrato string,vigencia string) (proveedor_contrato_persona []ProveedorContratoPersona, err error) {
	o := orm.NewOrm()
	var temp []ProveedorContratoPersona
	_,err = o.Raw("SELECT * FROM agora.informacion_persona_natural INNER JOIN agora.informacion_proveedor ON agora.informacion_proveedor.num_documento = agora.informacion_persona_natural.num_documento_persona INNER JOIN argo.contrato_general ON agora.informacion_proveedor.id_proveedor = argo.contrato_general.contratista INNER JOIN argo.contrato_suscrito ON argo.contrato_suscrito.numero_contrato=argo.contrato_general.numero_contrato AND argo.contrato_suscrito.vigencia=argo.contrato_general.vigencia where argo.contrato_general.vigencia="+vigencia+" and argo.contrato_general.numero_contrato='"+contrato+"';").QueryRows(&temp)
	if err == nil {
		return temp,nil
	}
	return nil, err
}
