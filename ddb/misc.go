package ddb

// params[where]
type cdn string
const (
	eq       	cdn = "eq"
	neq      	cdn = "neq"
	gt       	cdn = "gt"
	gte      	cdn = "gte"
	lt       	cdn = "lt"
	lte      	cdn = "lte"
	like     	cdn = "like"
	starting 	cdn = "starting"
	ending   	cdn = "ending"
	in       	cdn = "in"
	nin      	cdn = "nin"

	asc      	cdn = "asc"
	desc     	cdn = "desc"
	limit    	cdn = "limit"
	page     	cdn = "page"
)


// 分页请求数据
type paging struct {
	page  int   // 页码
	limit int   // 每页条数
	total int64 // 总数据条数
}

func (p *paging) offset() int {
	offset := 0
	if p.page > 0 {
		offset = (p.page - 1) * p.limit
	}
	return offset
}

func (p *paging) totalPage() int {
	if p.total == 0 || p.limit == 0 {
		return 0
	}
	totalPage := int(p.total) / p.limit
	if int(p.total)%p.limit > 0 {
		totalPage = totalPage + 1
	}
	return totalPage
}

type paramPair struct {
	query string        // 查询
	args  []interface{} // 参数
}

// 排序信息
type orderPair struct {
	column string // 排序字段
	by    string  // 正序/反序
}
