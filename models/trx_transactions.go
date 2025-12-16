package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Trx_transactions struct {
	TransactionId          string `orm:"pk"`
	Amount                 float64
	TotalDebitAmount       float64
	Charge                 float64
	Commission             float64
	SenderAccountNumber    string    `orm:"size(255)"`
	RecipientAccountNumber string    `orm:"size(255)"`
	TransferCode           string    `orm:"size(150)"`
	Status                 *Status   `orm:"rel(fk)"`
	ResponseCode           string    `orm:"size(50)"`
	ResponseMessage        string    `orm:"size(255)"`
	Description            string    `orm:"size(255)"`
	DateCreated            time.Time `orm:"type(datetime)"`
	DateModified           time.Time `orm:"type(datetime)"`
	CreatedBy              int
	ModifiedBy             int
	Active                 int
}

func init() {
	orm.RegisterModel(new(Trx_transactions))
}

// AddTrx_transactions insert a new Trx_transactions into database and returns
// last inserted Id on success.
func AddTrx_transactions(m *Trx_transactions) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTrx_transactionsById retrieves Trx_transactions by Id. Returns error if
// Id doesn't exist
func GetTrx_transactionsById(id string) (v *Trx_transactions, err error) {
	o := orm.NewOrm()
	v = &Trx_transactions{TransactionId: id}
	if err = o.QueryTable(new(Trx_transactions)).Filter("TransactionId", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTrx_transactions retrieves all Trx_transactions matches certain condition. Returns empty list if
// no records exist
func GetAllTrx_transactions(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Trx_transactions))
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

	var l []Trx_transactions
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

// UpdateTrx_transactions updates Trx_transactions by Id and returns error if
// the record to be updated doesn't exist
func UpdateTrx_transactionsById(m *Trx_transactions) (err error) {
	o := orm.NewOrm()
	v := Trx_transactions{TransactionId: m.TransactionId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTrx_transactions deletes Trx_transactions by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTrx_transactions(id string) (err error) {
	o := orm.NewOrm()
	v := Trx_transactions{TransactionId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Trx_transactions{TransactionId: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
