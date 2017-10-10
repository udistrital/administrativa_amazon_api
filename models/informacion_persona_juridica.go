package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type InformacionPersonaJuridica struct {
	Id                           int                        `orm:"column(num_nit_empresa);pk"`
	DigitoVerificacion           float64                    `orm:"column(digito_verificacion)"`
	ProcedenciaEmpresa           string                     `orm:"column(procedencia_empresa)"`
	IdCiudadOrigen               float64                    `orm:"column(id_ciudad_origen);null"`
	CodigoPaisDian               float64                    `orm:"column(codigo_pais_dian);null"`
	CodigoPostal                 float64                    `orm:"column(codigo_postal);null"`
	TipoIdentificacionExtranjera string                     `orm:"column(tipo_identificacion_extranjera);null"`
	NumCedulaExtranjeria         float64                    `orm:"column(num_cedula_extranjeria);null"`
	NumPasaporte                 float64                    `orm:"column(num_pasaporte);null"`
	IdTipoConformacion           *ParametroTipoConformacion `orm:"column(id_tipo_conformacion);rel(fk)"`
	MontoCapitalAutorizado       float64                    `orm:"column(monto_capital_autorizado)"`
	ExclusividadProducto         bool                       `orm:"column(exclusividad_producto)"`
	RegimenContributivo          string                     `orm:"column(regimen_contributivo)"`
	Pyme                         bool                       `orm:"column(pyme)"`
	RegistroMercantil            bool                       `orm:"column(registro_mercantil)"`
	SujetoRetencion              bool                       `orm:"column(sujeto_retencion)"`
	AgenteRetenedor              bool                       `orm:"column(agente_retenedor)"`
	ResponsableICA               bool                       `orm:"column(responsable_ICA)"`
	ResponsableIVA               bool                       `orm:"column(responsable_IVA)"`
	Genero                       string                     `orm:"column(genero)"`
	NomProveedor                 string                     `orm:"column(nom_proveedor)"`
}

func (t *InformacionPersonaJuridica) TableName() string {
	return "informacion_persona_juridica"
}

func init() {
	orm.RegisterModel(new(InformacionPersonaJuridica))
}

// AddInformacionPersonaJuridica insert a new InformacionPersonaJuridica into database and returns
// last inserted Id on success.
func AddInformacionPersonaJuridica(m *InformacionPersonaJuridica) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetInformacionPersonaJuridicaById retrieves InformacionPersonaJuridica by Id. Returns error if
// Id doesn't exist
func GetInformacionPersonaJuridicaById(id int) (v *InformacionPersonaJuridica, err error) {
	o := orm.NewOrm()
	v = &InformacionPersonaJuridica{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllInformacionPersonaJuridica retrieves all InformacionPersonaJuridica matches certain condition. Returns empty list if
// no records exist
func GetAllInformacionPersonaJuridica(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(InformacionPersonaJuridica))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
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

	var l []InformacionPersonaJuridica
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

// UpdateInformacionPersonaJuridica updates InformacionPersonaJuridica by Id and returns error if
// the record to be updated doesn't exist
func UpdateInformacionPersonaJuridicaById(m *InformacionPersonaJuridica) (err error) {
	o := orm.NewOrm()
	v := InformacionPersonaJuridica{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteInformacionPersonaJuridica deletes InformacionPersonaJuridica by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInformacionPersonaJuridica(id int) (err error) {
	o := orm.NewOrm()
	v := InformacionPersonaJuridica{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&InformacionPersonaJuridica{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
