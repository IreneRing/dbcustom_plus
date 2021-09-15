package ddb


type statement struct {
	sel     		// 要查询的字段，如果为空，表示查询所有字段
	where			// 查询条件
	pldWhere		// 连接
	joinWhere		// 连接
	other			// 排序，页码，分组...
	*Dbmp
}


// 条件构建
func (s *statement) build(db *dDB) *dDB {
	if s == nil{  // todo nil 也可以调用
		return db
	}

	//ret := s.where.build(db)
	ret := db
	// where
	if len(s.where.andParams) > 0 {
		for _, param := range s.where.andParams {
			ret.db = ret.db.Where(param.query, param.args...)
		}
	}

	if len(s.where.orParams) > 0 {
		for _, param := range s.where.orParams {
			ret.db = ret.db.Or(param.query, param.args...)
		}
	}


	if len(s.sel.column) > 0 {
		if s.sel.distinct {
			ret.db = ret.db.Distinct(s.sel.column)
		}else {
			ret.db = ret.db.Select(s.sel.column)
		}
	}

	// 聚合
	if len(s.other.groupBy) > 0 {
		for _, grp := range s.other.groupBy {
			ret.db = ret.db.Group(grp)
		}
	}

	if len(s.other.having) > 0 {
		for _, param := range s.other.having {
			ret.db = ret.db.Having(param.query, param.args...)
		}
	}



	// preload (关联 uuid)
	// 相同表明重置条件
	//"TTestBp","content = '1'" --> "t_test_bp"."uuid" = 1 AND content = '1'
	//"TTestBp","content = '1'","content = '1'" --> "t_test_bp"."uuid" = 1 AND "content = '1'" = 'content = \'1\''
	//"TTestBp","content = '1'","content = '1'","content = '1'" --> "t_test_bp"."uuid" = 1 AND "t_test_bp"."id" IN ('content = \'1\'','content = \'1\'','content = \'1\'')
	if s.pldWhere.preload != nil {
		for model, args := range s.pldWhere.preload {
			ret.db = ret.db.Preload(model,args)
		}
	}


	// order
	if len(s.other.orders) > 0 {
		for _, order := range s.other.orders {
			ret.db = ret.db.Order(order.column + order.by)
		}
	}

	// limit
	if s.other.paging != nil && s.other.paging.limit > 0 {
		ret.db = ret.db.Limit(s.other.paging.limit)
	}

	// offset
	if s.other.paging != nil && s.other.paging.offset() > 0 {
		ret.db = ret.db.Offset(s.other.paging.offset())
	}

	return ret

}
