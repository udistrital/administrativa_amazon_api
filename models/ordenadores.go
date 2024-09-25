package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Ordenadores struct {
	Id              int       `orm:"column(id);pk"`
	IdOrdenador     int       `orm:"column(id_ordenador);null"`
	InfoResolucion  string    `orm:"column(info_resolucion);null"`
	IdCiudad        int       `orm:"column(id_ciudad);null"`
	FechaInicio     time.Time `orm:"column(fecha_inicio);type(date);null"`
	FechaFin        time.Time `orm:"column(fecha_fin);type(date);null"`
	Estado          bool      `orm:"column(estado);null"`
	Documento       int       `orm:"column(documento);null"`
	NombreOrdenador string    `orm:"column(nombre_ordenador);null"`
	RolOrdenador    string    `orm:"column(rol_ordenador);null"`
	RolId           int       `orm:"column(rol_id);null"`
}

func (t *Ordenadores) TableName() string {
	return "ordenadores"
}

func init() {
	orm.RegisterModel(new(Ordenadores))
}

// AddOrdenadores insert a new Ordenadores into database and returns
// last inserted Id on success.
func AddOrdenadores(m *Ordenadores) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrdenadoresById retrieves Ordenadores by Id. Returns error if
// Id doesn't exist
func GetOrdenadoresById(id int) (v *Ordenadores, err error) {
	o := orm.NewOrm()
	v = &Ordenadores{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOrdenadores retrieves all Ordenadores matches certain condition. Returns empty list if
// no records exist
func GetAllOrdenadores(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Ordenadores))
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

	var l []Ordenadores
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

// UpdateOrdenadores updates Ordenadores by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrdenadoresById(m *Ordenadores) (err error) {
	o := orm.NewOrm()
	v := Ordenadores{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrdenadores deletes Ordenadores by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrdenadores(id int) (err error) {
	o := orm.NewOrm()
	v := Ordenadores{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Ordenadores{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
