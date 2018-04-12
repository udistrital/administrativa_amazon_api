package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SolicitudNecesidad struct {
	Id                     int                       `orm:"column(id);pk"`
	Numero                 int                       `orm:"column(numero)"`
	Vigencia               float64                   `orm:"column(vigencia)"`
	Dependencia            int                       `orm:"column(dependencia)"`
	JefeDependencia        *JefeDependenciaTemporal  `orm:"column(jefe_dependencia);rel(fk)"`
	DependenciaDestino     *DependenciaTemporal      `orm:"column(dependencia_destino);rel(fk)"`
	ObjetoContractual      string                    `orm:"column(objeto_contractual)"`
	Estado                 *EstadoSolicitudNecesidad `orm:"column(estado);rel(fk)"`
	FechaSolicitud         time.Time                 `orm:"column(fecha_solicitud);type(date);auto_now"`
	ValorContratacion      float64                   `orm:"column(valor_contratacion)"`
	Justificacion          string                    `orm:"column(justificacion)"`
	UnidadEjecutora        int                       `orm:"column(unidad_ejecutora)"`
	DiasDuracion           float64                   `orm:"column(dias_duracion)"`
	UnicoPago              bool                      `orm:"column(unico_pago)"`
	AgotarPresupuesto      bool                      `orm:"column(agotar_presupuesto)"`
	NovedadOtroSi          bool                      `orm:"column(novedad_otro_si)"`
	CodigoContratoOtroSi   int                       `orm:"column(codigo_contrato_otro_si);null"`
	ModalidadSeleccion     *ModalidadSeleccion       `orm:"column(modalidad_seleccion);rel(fk)"`
	Entidad                int                       `orm:"column(entidad)"`
	Servicio               *Servicio                 `orm:"column(servicio);rel(fk)"`
	PlanAnualAdquisiciones int                       `orm:"column(plan_anual_adquisiciones)"`
	EstudioMercado         string                    `orm:"column(estudio_mercado);null"`
	FechaEvaluacion        time.Time                 `orm:"column(fecha_evaluacion);type(date);null;auto_now"`
	OrdenadorGasto         *OrdenadorGastoTemporal   `orm:"column(ordenador_gasto);rel(fk)"`
	JustificacionRechazo   string                    `orm:"column(justificacion_rechazo);null"`
	TipoContratacion       *TipoContratacion         `orm:"column(tipo_contratacion);rel(fk)"`
	AnalisisRiesgo         string                    `orm:"column(analisis_riesgo);null"`
	TecnicasUniformes      bool                      `orm:"column(tecnicas_uniformes)"`
}

func (t *SolicitudNecesidad) TableName() string {
	return "solicitud_necesidad"
}

func init() {
	orm.RegisterModel(new(SolicitudNecesidad))
}

// AddSolicitudNecesidad insert a new SolicitudNecesidad into database and returns
// last inserted Id on success.
func AddSolicitudNecesidad(m *SolicitudNecesidad) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSolicitudNecesidadById retrieves SolicitudNecesidad by Id. Returns error if
// Id doesn't exist
func GetSolicitudNecesidadById(id int) (v *SolicitudNecesidad, err error) {
	o := orm.NewOrm()
	v = &SolicitudNecesidad{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSolicitudNecesidad retrieves all SolicitudNecesidad matches certain condition. Returns empty list if
// no records exist
func GetAllSolicitudNecesidad(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SolicitudNecesidad))
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

	var l []SolicitudNecesidad
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

// UpdateSolicitudNecesidad updates SolicitudNecesidad by Id and returns error if
// the record to be updated doesn't exist
func UpdateSolicitudNecesidadById(m *SolicitudNecesidad) (err error) {
	o := orm.NewOrm()
	v := SolicitudNecesidad{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSolicitudNecesidad deletes SolicitudNecesidad by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSolicitudNecesidad(id int) (err error) {
	o := orm.NewOrm()
	v := SolicitudNecesidad{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SolicitudNecesidad{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
