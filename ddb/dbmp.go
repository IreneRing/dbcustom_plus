package ddb

import (
	"github.com/by-zxy/dbcustom_plus/ddb_utils"
	"github.com/by-zxy/dbcustom_plus/str_utils"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"log"
	"strings"
)

type Dbmp struct {
	*processor //单纯调用
	*statement
	*dDB
	*Tx
	Result
}

type Result struct {
	Error        	error //错误
	RowsAffected 	int64 //影响条数
	Count 			int64 //总数
}

func NewDbmp(defDb *gorm.DB) *Dbmp {
	db(defDb)
	return initDbmp()
}

func initDbmp() *Dbmp {
	dataMapper := &Dbmp{}
	dataMapper.dDB = newDB()

	dataMapper.statement = &statement{
		Dbmp: dataMapper,
	}
	return dataMapper
}

// new Mapper and a statement , 保持原来mapper不变
func (dp *Dbmp) clone() *Dbmp {
	dataMapper := &Dbmp{}
	dataMapper.dDB = dp.dDB.clone()  // todo

	if dp.statement != nil {
		dataMapper.statement = &statement{
			sel:   		dp.statement.sel,
			where: 		dp.statement.where,
			pldWhere:  	dp.statement.pldWhere,
			other: 		dp.statement.other,
			Dbmp:  		dataMapper,
		}
	}else {
		dataMapper.statement = &statement{
			Dbmp: dataMapper,
		}
	}
	return dataMapper
}

/**
	sel
*/
func (dp *Dbmp) SelectId() *Dbmp {
	dp.statement.sel.columns("id")
	return dp
}

func (dp *Dbmp) SelectAll() *Dbmp {
	dp.statement.sel.columns("*")
	return dp
}

func (dp *Dbmp) Select(column ...string) *Dbmp {
	dp.statement.sel.columns(column...)
	return dp
}


//是否开启查询字段distinct
func (dp *Dbmp) Distinct() *Dbmp {
	dp.statement.sel.isDistinct()
	return dp
}

func (dp *Dbmp) SelectByAlias(column, name string) *Dbmp {
	dp.statement.sel.alias(column, name)
	return dp
}

// 清空查询字段
func (dp *Dbmp) SelectReset() *Dbmp {
	dp.statement.sel = sel{}
	return dp
}

//聚合函数
// column - 表字段。name - 别名(只取第一位)
func (dp *Dbmp) Max(column string, name ...string) *Dbmp {
	dp.statement.sel.max(column,name...)
	return dp
}
func (dp *Dbmp) MaxDistinct(column string, name ...string) *Dbmp {
	dp.Max("Distinct "+column,name...)
	return dp
}

func (dp *Dbmp) Min(column string, name ...string) *Dbmp {
	dp.statement.sel.min(column,name...)
	return dp
}
func (dp *Dbmp) MinDistinct(column string, name ...string) *Dbmp {
	dp.Min("Distinct "+column,name...)
	return dp
}

func (dp *Dbmp) Sum(column string, name ...string) *Dbmp {
	dp.statement.sel.sum(column,name...)
	return dp
}
func (dp *Dbmp) SumDistinct(column string, name ...string) *Dbmp {
	dp.Sum("Distinct "+column,name...)
	return dp
}

func (dp *Dbmp) Avg(column string, name ...string) *Dbmp {
	dp.statement.sel.avg(column,name...)
	return dp
}
func (dp *Dbmp) AvgDistinct(column string, name ...string) *Dbmp {
	dp.Avg("Distinct "+column,name...)
	return dp
}

func (dp *Dbmp) CountSel(column string, name ...string) *Dbmp {
	dp.statement.sel.count(column,name...)
	return dp
}
func (dp *Dbmp) CountSelDistinct(column string, name ...string) *Dbmp {
	dp.CountSel("Distinct "+column,name...)
	return dp
}

/**
	where
 */
func (dp *Dbmp) Eq(column string, args interface{}) *Dbmp {
	dp.where.and(buildQueryAndArg(eq,column,args))
	return dp
}

func (dp *Dbmp) EqOr(column string, args interface{}) *Dbmp {
	dp.where.or(buildQueryAndArg(eq,column,args))
	return dp
}

func (dp *Dbmp) NotEq(column string, args interface{}) *Dbmp {
	dp.where.and(buildQueryAndArg(neq,column,args))
	return dp
}

func (dp *Dbmp) NotEqOr(column string, args interface{}) *Dbmp {
	dp.where.or(buildQueryAndArg(neq,column,args))
	return dp
}

func (dp *Dbmp) Gt(column string, args interface{}) *Dbmp {
	dp.where.and(buildQueryAndArg(gt,column,args))
	return dp
}

func (dp *Dbmp) GtOr(column string, args interface{}) *Dbmp {
	dp.where.or(buildQueryAndArg(gt,column,args))
	return dp
}

func (dp *Dbmp) Gte(column string, args interface{}) *Dbmp {
	dp.where.and(buildQueryAndArg(gte,column,args))
	return dp
}

func (dp *Dbmp) GteOr(column string, args interface{}) *Dbmp {
	dp.where.or(buildQueryAndArg(gte,column,args))
	return dp
}

func (dp *Dbmp) Lt(column string, args interface{}) *Dbmp {
	dp.where.and(buildQueryAndArg(lt,column,args))
	return dp
}

func (dp *Dbmp) LtOr(column string, args interface{}) *Dbmp {
	dp.where.or(buildQueryAndArg(lt,column,args))
	return dp
}

func (dp *Dbmp) Lte(column string, args interface{}) *Dbmp {
	dp.where.and(buildQueryAndArg(lte,column,args))
	return dp
}

func (dp *Dbmp) LteOr(column string, args interface{}) *Dbmp {
	dp.where.or(buildQueryAndArg(lte,column,args))
	return dp
}

func (dp *Dbmp) Like(column string, str string) *Dbmp {
	dp.where.and(buildQueryAndArg(like,column,str))
	return dp
}

func (dp *Dbmp) LikeOr(column string, str string) *Dbmp {
	dp.where.or(buildQueryAndArg(like,column,str))
	return dp
}

func (dp *Dbmp) Starting(column string, str string) *Dbmp {
	dp.where.and(buildQueryAndArg(starting,column,str))
	return dp
}

func (dp *Dbmp) StartingOr(column string, str string) *Dbmp {
	dp.where.or(buildQueryAndArg(starting,column,str))
	return dp
}

func (dp *Dbmp) Ending(column string, str string) *Dbmp {
	dp.where.and(buildQueryAndArg(ending,column,str))
	return dp
}

func (dp *Dbmp) EndingOr(column string, str string) *Dbmp {
	dp.where.or(buildQueryAndArg(ending,column,str))
	return dp
}

func (dp *Dbmp) In(column string, args interface{}) *Dbmp {
	dp.where.and(buildQueryAndArg(in,column,args))
	return dp
}

func (dp *Dbmp) InOr(column string, args interface{}) *Dbmp {
	dp.where.or(buildQueryAndArg(in,column,args))
	return dp
}

func (dp *Dbmp) NotIn(column string, args interface{}) *Dbmp {
	dp.where.and(buildQueryAndArg(nin,column,args))
	return dp
}

func (dp *Dbmp) NotInOr(column string, args interface{}) *Dbmp {
	dp.where.or(buildQueryAndArg(nin,column,args))
	return dp
}

/**
	join
*/
//func (dp *Dbmp) Join(raw ...string) *Dbmp {
//	dp.statement.joinWhere.joins(raw...)
//	return dp
//}

// 想要预加载，请调用此方法
// 存在此表的预加条件，使用次方法会重置清空
func (dp *Dbmp) Preload(model string) *Dbmp {
	dp.statement.pldWhere.preloads(model,"")
	return dp
}

func (dp *Dbmp) PreloadEq(model, column string, arg string) *Dbmp {
	dp.statement.pldWhere.preloads(model,buildQuery(eq,column,arg))
	return dp
}

func (dp *Dbmp) PreloadNeq(model, column string, arg string) *Dbmp {
	dp.statement.pldWhere.preloads(model,buildQuery(neq,column,arg))
	return dp
}

func (dp *Dbmp) PreloadGt(model, column string, arg string) *Dbmp {
	dp.statement.pldWhere.preloads(model,buildQuery(gt,column,arg))
	return dp
}

func (dp *Dbmp) PreloadGte(model, column string, arg string) *Dbmp {
	dp.statement.pldWhere.preloads(model,buildQuery(gte,column,arg))
	return dp
}

func (dp *Dbmp) PreloadLt(model, column string, arg string) *Dbmp {
	dp.statement.pldWhere.preloads(model,buildQuery(lt,column,arg))
	return dp
}

func (dp *Dbmp) PreloadLte(model, column string, arg string) *Dbmp {
	dp.statement.pldWhere.preloads(model,buildQuery(lte,column,arg))
	return dp
}

func (dp *Dbmp) PreloadLike(model, column string, arg string) *Dbmp {
	dp.statement.pldWhere.preloads(model,buildQuery(like,column,arg))
	return dp
}

func (dp *Dbmp) PreloadStarting(model, column string, arg string) *Dbmp {
	dp.statement.pldWhere.preloads(model,buildQuery(starting,column,arg))
	return dp
}

func (dp *Dbmp) PreloadEnding(model, column string, arg string) *Dbmp {
	dp.statement.pldWhere.preloads(model,buildQuery(ending,column,arg))
	return dp
}

func (dp *Dbmp) PreloadIn(model, column string, arg string) *Dbmp {
	dp.statement.pldWhere.preloads(model,buildQuery(in,column,arg))
	return dp
}

func (dp *Dbmp) PreloadNin(model, column string, arg string) *Dbmp {
	dp.statement.pldWhere.preloads(model,buildQuery(nin,column,arg))
	return dp
}

/**
	other
*/
func (dp *Dbmp) OrderBy(column string, isAsc bool) *Dbmp {
	if isAsc{
		dp.statement.other.asc(column)
	}else {
		dp.statement.other.desc(column)
	}

	return dp
}

func (dp *Dbmp) Page(page,limit int) *Dbmp {
	dp.statement.other.page(page,limit)
	return dp
}

func (dp *Dbmp) Limit(limit int) *Dbmp {
	dp.statement.other.limit(limit)
	return dp
}

func (dp *Dbmp) GroupBys(column ...string) *Dbmp {
	dp.statement.other.groupBys(column...)
	return dp
}

func (dp *Dbmp) Havings(column string, arg ...interface{}) *Dbmp {
	dp.statement.other.havings(column,arg...)
	return dp
}

func (dp *Dbmp) HavingEq(column string, arg interface{}) *Dbmp {
	dp.statement.other.havings(buildQueryAndArg(eq,column,arg))
	return dp
}

func (dp *Dbmp) HavingNeq(column string, arg interface{}) *Dbmp {
	dp.statement.other.havings(buildQueryAndArg(neq,column,arg))
	return dp
}

func (dp *Dbmp) HavingGt(column string, arg interface{}) *Dbmp {
	dp.statement.other.havings(buildQueryAndArg(gt,column,arg))
	return dp
}

func (dp *Dbmp) HavingGte(column string, arg interface{}) *Dbmp {
	dp.statement.other.havings(buildQueryAndArg(gte,column,arg))
	return dp
}

func (dp *Dbmp) HavingLt(column string, arg interface{}) *Dbmp {
	dp.statement.other.havings(buildQueryAndArg(lt,column,arg))
	return dp
}

func (dp *Dbmp) HavingLte(column string, arg interface{}) *Dbmp {
	dp.statement.other.havings(buildQueryAndArg(lte,column,arg))
	return dp
}

func (dp *Dbmp) HavingIn(column string, arg interface{}) *Dbmp {
	dp.statement.other.havings(buildQueryAndArg(in,column,arg))
	return dp
}

func (dp *Dbmp) HavingNin(column string, arg interface{}) *Dbmp {
	dp.statement.other.havings(buildQueryAndArg(nin,column,arg))
	return dp
}

// 初始化
// for init params to make a statement
func (dp *Dbmp) InitParams(model interface{}) *Dbmp {

	of := ddb_utils.NewAllOfReflect(model).NewOfDic(false)
	vj := of.ValueOfJson
	for json, tag := range of.TagsOfJson {
		if str_utils.IsNotBlank(tag[ddb_utils.QUERY]){
			// where
			querys := strings.Split(tag[ddb_utils.QUERY], ",")
			// join -  先判断是否为关联库的字段
			if str_utils.IsNotBlank(tag[ddb_utils.JOIN]){
				//dp.statement.join.condition(cdn(querys[0]),tag[utils.JOIN],querys[1],vj[json])
				dp.statement.pldWhere.preloads(tag[ddb_utils.JOIN],buildQuery(cdn(querys[0]),querys[1],utils.ToString(vj[json])) )
				continue
			}

			// (and / or)
			//dp.statement.where.condition(cdn(querys[0]),querys[1],vj[json])
			if strings.Contains(querys[0],"_") {
				dp.statement.where.or(buildQueryAndArg(cdn(querys[0]),querys[1],vj[json]))
			}else {
				dp.statement.where.and(buildQueryAndArg(cdn(querys[0]),querys[1],vj[json]))
			}

		}else {
			// order &limit
			dp.statement.other.condition(cdn(json),vj[json])
		}
	}
	return dp
}

//copy 一个当前状态的dbmp
func (dp *Dbmp) Copy() *Dbmp {
	return dp.clone()
}


// curd持久方法
// finisher,不会改变原本Mapper的地址
func (dp *Dbmp) FindAndCount(model interface{}) *Dbmp {
	if model == nil {
		log.Println("model must be a ptr")
		return dp
	}
	return dp.getInstance(dp,model).find().count(false).callMapper()
}

func (dp *Dbmp) Find(model interface{}) *Dbmp {
	if model == nil {
		log.Println("model must be a ptr")
		return dp
	}
	return dp.getInstance(dp,model).find().callMapper()
}

func (dp *Dbmp) First(model interface{}) *Dbmp {
	if model == nil {
		log.Println("model must be a ptr")
		return dp
	}
	return dp.getInstance(dp,model).first().callMapper()
}

func (dp *Dbmp) Count(model interface{}) *Dbmp {
	if model == nil {
		log.Println("model must be a ptr")
		return dp
	}
	return dp.getInstance(dp,model).count(true).callMapper()
}

func (dp *Dbmp) Create(model interface{}) *Dbmp {
	if model == nil {
		log.Println("model must be a ptr")
		return dp
	}
	return dp.getInstance(dp,model).create().callMapper()
}

func (dp *Dbmp) Update(model interface{}) *Dbmp {
	if model == nil {
		log.Println("model must be a ptr")
		return dp
	}
	return dp.getInstance(dp, model).update().callMapper()
}

func (dp *Dbmp) Delete(model interface{}) *Dbmp {
	if model == nil {
		log.Println("model must be a ptr")
		return dp
	}
	return dp.getInstance(dp,model).delete().callMapper()
}

// 事务
func (dp *Dbmp) transaction( fc func(m *Dbmp) error) error {
	return dp.dDB.transaction(func(tx *dDB) error {
		cloneDp := dp.clone()
		cloneDp.dDB = tx
		return fc(cloneDp) //必须返回fc的error
	})
}

// 手动原生sql
func Raw(ddb *dDB, modal interface{}, sql string, values ...interface{}) error {
	return ddb.db.Raw(sql,values).Scan(modal).Error
}
