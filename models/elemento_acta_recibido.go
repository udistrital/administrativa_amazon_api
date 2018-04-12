package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ElementoActaRecibido struct {
	Id                int       `orm:"column(id_elemento_ac);pk"`
	FechaRegistro     time.Time `orm:"column(fecha_registro);type(date);null;auto_now"`
	Nivel             int       `orm:"column(nivel);null"`
	TipoBien          int       `orm:"column(tipo_bien);null"`
	Descripcion       string    `orm:"column(descripcion);null"`
	Cantidad          float64   `orm:"column(cantidad);null"`
	Unidad            string    `orm:"column(unidad);null"`
	Valor             float64   `orm:"column(valor);null"`
	Iva               int       `orm:"column(iva);null"`
	SubtotalSinIva    float64   `orm:"column(subtotal_sin_iva);null"`
	TotalIva          float64   `orm:"column(total_iva);null"`
	TotalIvaCon       float64   `orm:"column(total_iva_con);null"`
	TipoPoliza        int       `orm:"column(tipo_poliza);null"`
	FechaInicioPol    time.Time `orm:"column(fecha_inicio_pol);type(date);null;auto_now"`
	FechaFinalPol     time.Time `orm:"column(fecha_final_pol);type(date);null;auto_now"`
	Marca             string    `orm:"column(marca);null"`
	Serie             string    `orm:"column(serie);null"`
	IdActa            int       `orm:"column(id_acta);null"`
	Estado            bool      `orm:"column(estado);null"`
	Referencia        string    `orm:"column(referencia);null"`
	Placa             string    `orm:"column(placa);null"`
	Observacion       string    `orm:"column(observacion);null"`
	CodigoDependencia string    `orm:"column(codigo_dependencia);null"`
	Funcionario       string    `orm:"column(funcionario);null"`
	NumeroContrato    string    `orm:"column(numero_contrato);null"`
	Vigencia          int       `orm:"column(vigencia);null"`
}

func (t *ElementoActaRecibido) TableName() string {
	return "elemento_acta_recibido"
}

func init() {
	orm.RegisterModel(new(ElementoActaRecibido))
}

// AddElementoActaRecibido insert a new ElementoActaRecibido into database and returns
// last inserted Id on success.
func AddElementoActaRecibido(m *ElementoActaRecibido) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetElementoActaRecibidoById retrieves ElementoActaRecibido by Id. Returns error if
// Id doesn't exist
func GetElementoActaRecibidoById(id int) (v *ElementoActaRecibido, err error) {
	o := orm.NewOrm()
	v = &ElementoActaRecibido{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllElementoActaRecibido retrieves all ElementoActaRecibido matches certain condition. Returns empty list if
// no records exist
func GetAllElementoActaRecibido(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ElementoActaRecibido))
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

	var l []ElementoActaRecibido
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

// UpdateElementoActaRecibido updates ElementoActaRecibido by Id and returns error if
// the record to be updated doesn't exist
func UpdateElementoActaRecibidoById(m *ElementoActaRecibido) (err error) {
	o := orm.NewOrm()
	v := ElementoActaRecibido{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteElementoActaRecibido deletes ElementoActaRecibido by Id and returns error if
// the record to be deleted doesn't exist
func DeleteElementoActaRecibido(id int) (err error) {
	o := orm.NewOrm()
	v := ElementoActaRecibido{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ElementoActaRecibido{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
