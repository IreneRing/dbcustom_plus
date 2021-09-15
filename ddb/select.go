package ddb

type sel struct {
	column   		[]string
	distinct		bool
}

func (s *sel) columns(column ...string) *sel {
	s.column = append(s.column, column...)
	return s
}

func (s *sel) isDistinct() *sel {
	s.distinct = true
	return s
}

func (s *sel) alias(column, name string) *sel {
	s.column = append(s.column, column+" as "+name)
	return s
}

// 聚合函数
func (s *sel) max(column string, name ...string) *sel {
	maxStr := "Max("+column+")"
	if len(name) > 0 {
		maxStr += " as " + name[0]
	}
	s.column = append(s.column, maxStr)
	return s
}

func (s *sel) min(column string, name ...string) *sel {
	maxStr := "Min("+column+")"
	if len(name) > 0 {
		maxStr += " as " + name[0]
	}
	s.column = append(s.column, maxStr)
	return s
}

func (s *sel) sum(column string, name ...string) *sel {
	maxStr := "Sum("+column+")"
	if len(name) > 0 {
		maxStr += " as " + name[0]
	}
	s.column = append(s.column, maxStr)
	return s
}

func (s *sel) avg(column string, name ...string) *sel {
	maxStr := "Avg("+column+")"
	if len(name) > 0 {
		maxStr += " as " + name[0]
	}
	s.column = append(s.column, maxStr)
	return s
}

func (s *sel) count(column string, name ...string) *sel {
	maxStr := "count("+column+")"
	if len(name) > 0 {
		maxStr += " as " + name[0]
	}
	s.column = append(s.column, maxStr)
	return s
}
