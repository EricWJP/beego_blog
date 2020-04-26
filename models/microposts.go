package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"time"
)

type Microposts struct {
	Id            int64  `orm:"auto"`
	Title         string `orm:"size(256)"`
	Content       string `orm:"type(longtext)"`
	UserId        int64
	Status        bool
	CommentsCount int64
	FollowedCount int64
	UpdatedAt     time.Time `orm:"type(datetime)"`
	CreatedAt     time.Time `orm:"type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Microposts))
}

// AddMicroposts insert a new Microposts into database and returns
// last inserted Id on success.
func AddMicroposts(m *Microposts) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMicropostsById retrieves Microposts by Id. Returns error if
// Id doesn't exist
func GetMicropostsById(id int64) (v *Microposts, err error) {
	o := orm.NewOrm()
	v = &Microposts{Id: id}
	if err = o.QueryTable(new(Microposts)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMicroposts retrieves all Microposts matches certain condition. Returns empty list if
// no records exist
func GetAllMicroposts(query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Microposts))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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

	var l []Microposts
	qs = qs.OrderBy(sortFields...).RelatedSel()
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

// UpdateMicroposts updates Microposts by Id and returns error if
// the record to be updated doesn't exist
func UpdateMicropostsById(m *Microposts) (err error) {
	o := orm.NewOrm()
	v := Microposts{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMicroposts deletes Microposts by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMicroposts(id int64) (err error) {
	o := orm.NewOrm()
	v := Microposts{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Microposts{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
