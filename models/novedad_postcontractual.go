package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type NovedadPostcontractual struct {
	Id              int       `orm:"column(id);pk;auto"`
	NumeroContrato  string    `orm:"column(numero_contrato)"`
	Vigencia        int       `orm:"column(vigencia)"`
	TipoNovedad     float64   `orm:"column(tipo_novedad)"`
	FechaInicio     time.Time `orm:"column(fecha_inicio);type(date);null"`
	FechaFin        time.Time `orm:"column(fecha_fin);type(date);null"`
	FechaRegistro   time.Time `orm:"column(fecha_registro);type(date)"`
	Contratista     float64   `orm:"column(contratista);null"`
	NumeroCdp       int       `orm:"column(numero_cdp);null"`
	VigenciaCdp     int       `orm:"column(vigencia_cdp);null"`
	PlazoEjecucion  int       `orm:"column(plazo_ejecucion);null"`
	UnidadEjecucion int       `orm:"column(unidad_ejecucion);null"`
	ValorNovedad    float64   `orm:"column(valor_novedad);null"`
}

func (t *NovedadPostcontractual) TableName() string {
	return "novedad_postcontractual"
}

func init() {
	orm.RegisterModel(new(NovedadPostcontractual))
}

// AddNovedadPostcontractual insert a new NovedadPostcontractual into database and returns
// last inserted Id on success.
func AddNovedadPostcontractual(m *NovedadPostcontractual) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNovedadPostcontractualById retrieves NovedadPostcontractual by Id. Returns error if
// Id doesn't exist
func GetNovedadPostcontractualById(id int) (v *NovedadPostcontractual, err error) {
	o := orm.NewOrm()
	v = &NovedadPostcontractual{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNovedadPostcontractual retrieves all NovedadPostcontractual matches certain condition. Returns empty list if
// no records exist
func GetAllNovedadPostcontractual(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(NovedadPostcontractual))
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

	var l []NovedadPostcontractual
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

// UpdateNovedadPostcontractual updates NovedadPostcontractual by Id and returns error if
// the record to be updated doesn't exist
func UpdateNovedadPostcontractualById(m *NovedadPostcontractual) (err error) {
	o := orm.NewOrm()
	v := NovedadPostcontractual{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNovedadPostcontractual deletes NovedadPostcontractual by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNovedadPostcontractual(id int) (err error) {
	o := orm.NewOrm()
	v := NovedadPostcontractual{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&NovedadPostcontractual{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
