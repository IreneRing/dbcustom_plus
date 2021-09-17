package ddb

import (
	"github.com/by-zxy/dbcustom_plus/ddb_utils"
)

type processor struct {
	*Dbmp
	//*DB
}

// get a new processor
func (p *processor) getInstance(dp *Dbmp, model interface{}) (pp *processor) {

	pp =  &processor{
		Dbmp:	dp.clone(), // clone一份值的地址
	}

	pp.Dbmp.dDB.dest = model
	pp.Dbmp.processor = pp

	return
}


// finisherd
// new dbmp,原dbmp指针不变，会累积
func (p *processor) find() *processor {
	result := p.Dbmp.statement.build(p.dDB).db.Find(p.dDB.dest)

	p.Dbmp.dDB.db = result
	return p
}

func (p *processor) first() *processor {
	result := p.Dbmp.statement.build(p.dDB).db.First(p.dDB.dest)

	p.Dbmp.dDB.db = result
	return p
}
// isRest 是否重置db
func (p *processor) count(isRest bool) *processor {
	var count int64
	result := p.Dbmp.statement.build(p.dDB).db.Count(&count)
	if isRest {
		p.Dbmp.dDB.db = result
	}
	p.Result.Count = count
	return p
}

func (p *processor) create() *processor {
	of := ddb_utils.NewAllOfReflect(p.dDB.dest).NewOfDic(false)
	for k, _ := range of.ValueOfName {
		p.Dbmp.statement.columns(k)
	}

	result := p.Dbmp.statement.build(p.dDB).db.Create(p.dDB.dest)
	p.Dbmp.dDB.db = result
	return p
}

func (p *processor) update() *processor {
	result := p.Dbmp.statement.build(p.dDB).db.Updates(p.dDB.dest)
	p.Dbmp.dDB.db = result
	return p
}

func (p *processor) delete() *processor {
	result := p.Dbmp.statement.build(p.dDB).db.Delete(p.dDB.dest)
	p.Dbmp.dDB.db = result
	return p
}

// 获取新Dbmp,里面statement已经清空
func (p *processor) callMapper() *Dbmp {
	dataMapper := initDbmp()
	dataMapper.Result = Result{
		Error:        p.Dbmp.dDB.db.Error,
		RowsAffected: p.Dbmp.dDB.db.RowsAffected,
		Count:        p.Result.Count,
	}
	return dataMapper
}