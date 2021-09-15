package ddb

type pldWhere struct {
	preload    map[string]string  	// 预加载
}

type joinWhere struct {
	join	[]string
}

func (jw *joinWhere) joins(raw ...string) *joinWhere {
	jw.join = append(jw.join, raw...)
	return jw
}

func (p *pldWhere) preloads(model, arg string) *pldWhere {
	if len(model) == 0 {
		return p
	}

	if p.preload == nil{
		p.preload = make(map[string]string)
	}

	if val,ok := p.preload[model]; ok {
		if len(arg) > 0 {
			p.preload[model] = val + " and " + arg
			return p
		}
	}
	p.preload[model] = arg
	return p
}
