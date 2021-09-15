package ddb

type where struct {
	andParams  []paramPair  	// 参数
	orParams   []paramPair  	// or参数
}


func (w *where) and(query string, args ...interface{}) *where {
	w.andParams = append(w.andParams, paramPair{query, args})
	return w
}

func (w *where) or(query string, args ...interface{}) *where {
	w.orParams = append(w.orParams, paramPair{query, args})
	return w
}

