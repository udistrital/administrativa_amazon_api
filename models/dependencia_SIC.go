package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type DependenciaSIC struct {
	ESFIDESPACIO    string    `orm:"column(ESF_ID_ESPACIO)"`
	ESFCODIGODEP    string    `orm:"column(ESF_CODIGO_DEP)"`
	ESFDEPENCARGADA string    `orm:"column(ESF_DEP_ENCARGADA);null"`
	ESFESTADO       string    `orm:"column(ESF_ESTADO);null"`
	Ciudad          string    `orm:"column(ciudad);null"`
	Localidad       string    `orm:"column(localidad);null"`
	Barrio          string    `orm:"column(barrio);null"`
	Direccion       string    `orm:"column(direccion);null"`
	EstadoRegistro  bool      `orm:"column(estado_registro)"`
	FechaRegistro   time.Time `orm:"column(fecha_registro);type(date);null"`
	Id              int       `orm:"column(id_dependencia);pk"`
}

func (t *DependenciaSIC) TableName() string {
	return "dependencia_SIC"
}

func init() {
	orm.RegisterModel(new(DependenciaSIC))
}

// AddDependenciaSIC insert a new DependenciaSIC into database and returns
// last inserted Id on success.
func AddDependenciaSIC(m *DependenciaSIC) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDependenciaSICById retrieves DependenciaSIC by Id. Returns error if
// Id doesn't exist
func GetDependenciaSICById(id int) (v *DependenciaSIC, err error) {
	o := orm.NewOrm()
	v = &DependenciaSIC{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDependenciaSIC retrieves all DependenciaSIC matches certain condition. Returns empty list if
// no records exist
func GetAllDependenciaSIC(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DependenciaSIC))
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

	var l []DependenciaSIC
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

// UpdateDependenciaSIC updates DependenciaSIC by Id and returns error if
// the record to be updated doesn't exist
func UpdateDependenciaSICById(m *DependenciaSIC) (err error) {
	o := orm.NewOrm()
	v := DependenciaSIC{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDependenciaSIC deletes DependenciaSIC by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDependenciaSIC(id int) (err error) {
	o := orm.NewOrm()
	v := DependenciaSIC{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DependenciaSIC{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
