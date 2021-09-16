package ddb

import (
	"github.com/by-zxy/dbcustom_plus"
	"gorm.io/gorm/utils"
)

type other struct {
	orders     		[]orderPair 	// 排序
	paging     		*paging      	// 分页
	having	   		[]paramPair
	groupBy			[]string
}

func (o *other) asc(column string) *other {
	o.orders = append(o.orders, orderPair{column: column, by: " ASC"})
	return o
}

func (o *other) desc(column string) *other {
	o.orders = append(o.orders, orderPair{column: column, by: " DESC"})
	return o
}

func (o *other) limit(limit int) *other {
	if o.paging == nil {
		o.page(1, limit)
	}else {
		o.page(o.paging.page,limit)
	}

	return o
}

func (o *other) page(page, limit int) *other {
	if o.paging == nil {
		o.paging = &paging{page: page, limit: limit}
	} else {
		o.paging.page = page
		o.paging.limit = limit
	}
	return o
}

func (o *other) condition(cdn cdn, args interface{}) *other {
	if dbcustom_plus.IsNon(args) {
		switch cdn {
			case asc:		o.asc(utils.ToString(args))
			case desc:		o.desc(utils.ToString(args))
			case limit:
				o.limit(dbcustom_plus.ToInt(args))
			case page:
				if o.paging == nil { //当前分页空
					o.page(dbcustom_plus.ToInt(args),1)
				} else { //当前存在分页
					o.page(dbcustom_plus.ToInt(args), o.paging.limit)
				}
		}
	}
	return o
}

//聚合函数
func (o *other) groupBys(column ...string) *other {
	o.groupBy = append(o.groupBy, column...)
	return o
}

func (o *other) havings(column string, args ...interface{}) *other {
	o.having = append(o.having, paramPair{column, args})
	return o
}
