package ddb

import (
	"github.com/by-zxy/dbcustom_plus"
	"strings"
)

type argSet struct {
	column	string
	query	string
	arg  	interface{}
}

func argSetArg(cdn cdn, column string) argSet {
	apr := argSet{column:column}
	switch cdn {
	case eq:		apr.query = column+" = ?"
	case neq:		apr.query = column+" <> ?"
	case gt:		apr.query = column+" > ?"
	case gte:		apr.query = column+" >= ?"
	case lt:		apr.query = column+" < ?"
	case lte:		apr.query = column+" <= ?"
	case like:		apr.query = column+" LIKE concat('%',?,'%')"
	case starting:	apr.query = column+" LIKE concat(?,'%')"
	case ending:	apr.query = column+" LIKE concat('%',?)"
	case in:		apr.query = column+" in (?)"
	case nin:		apr.query = column+" not in (?)"
	}
	return apr
}

func argSetWithStr(cdn cdn, column string, arg string) argSet {
	apr := argSetArg(cdn,column)
	switch cdn {
	case in,nin:
		// 1,2,3 -> 1','2','3 -> '1','2','3'
		//apr.query = strings.ReplaceAll(apr.query,"?",strings.ReplaceAll("'?'","?",strings.ReplaceAll(arg,",","','")))
		apr.arg = strings.ReplaceAll("'?'","?",strings.ReplaceAll(arg,",","','"))
	default:
		//apr.query = strings.ReplaceAll(apr.query,"?","'"+arg+"'")
		apr.query = arg
	}
	return apr
}

func argSetWithIf(cdn cdn, column string, arg interface{}) argSet {
	apr := argSetArg(cdn,column)
	apr.arg = arg
	return apr
}

// 组合query
func buildQuery(cdn cdn, column string, arg string) string {
	argRt := buildQueryCol(cdn,column)
	switch cdn {
	case in,nin:
		// 1,2,3 -> 1','2','3 -> '1','2','3'
		argRt = strings.ReplaceAll(argRt,"?",strings.ReplaceAll("'?'","?",strings.ReplaceAll(arg,",","','")))
	default:
		argRt = strings.ReplaceAll(argRt,"?","'"+arg+"'")
	}
	return argRt
}
// 组合query and param
func buildQueryAndArg(cdn cdn, column string, arg interface{}) (string,interface{}) {
	argRt := buildQueryCol(cdn,column)
	switch cdn {
	// gorm 模糊通配符只能写入参数里面
	case like:
		argRt = column+" LIKE ?"
		arg = "%"+dbcustom_plus.ToString(arg)+"%"
	case starting:
		argRt = column+" LIKE ?"
		arg = dbcustom_plus.ToString(arg)+"%"
	case ending:
		argRt = column+" LIKE ?"
		arg = "%"+dbcustom_plus.ToString(arg)
	}
	return argRt,arg
}

func buildQueryCol(cdn cdn, column string) string {
	switch cdn {
	case eq:		return column+" = ?"
	case neq:		return column+" <> ?"
	case gt:		return column+" > ?"
	case gte:		return column+" >= ?"
	case lt:		return column+" < ?"
	case lte:		return column+" <= ?"
	case like:		return column+" LIKE concat('%',?,'%')"
	case starting:	return column+" LIKE concat(?,'%')"
	case ending:	return column+" LIKE concat('%',?)"
	case in:		return column+" in (?)"
	case nin:		return column+" not in (?)"
	}
	return ""
}