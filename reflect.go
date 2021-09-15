package dbcustom_plus

import (
	"reflect"
)

type OfReflect struct {
	Vof []reflect.Value
	Tof []reflect.StructField
}

type OfDic struct {
	TagsOfName 	map[string]map[string]string  	//Column所有tags
	TagsOfJson 	map[string]map[string]string  	//Json所有tags
	ValueOfName map[string]interface{}  		//Column所有value
	ValueOfJson map[string]interface{}  		//Json所有value
}

type OfReflectName 		map[string]int
//type OfReflectValue 	map[string]interface{}
type OfReflectTag 		map[string]string
type OfReflectTags 		map[string]map[string]string

type Kind = reflect.Kind
const (
	MAP			= reflect.Map
	PTR			= reflect.Ptr
	STRING		= reflect.String
	SLICE		= reflect.Slice
	STRUCT		= reflect.Struct
	NIL			= reflect.UnsafePointer + 1
)

const (
	JSON 	=	"json"
	FORM 	=	"form"
	QUERY 	=	"query"
	JOIN 	=	"join"
)
var Tags = []string{JSON, FORM, QUERY, JOIN}

// 反射字段构造器，不反射字段里struct，struct变量
//反射获取所有字段(包含镶嵌) []reflect.StructField,[]reflect.Value
// struct中 &struct,[]struct 无法反射字段
func NewOfReflect(i interface{}) *OfReflect {
	structFields, values := toReflect(i,false)
	return &OfReflect{
		Vof: values,
		Tof: structFields,
	}
}

// 反射字段全参构造器，反射字段里struct，struct变量
func NewAllOfReflect(i interface{}) *OfReflect {
	structFields, values := toReflect(i,true)
	return &OfReflect{
		Vof: values,
		Tof: structFields,
	}
}

// 反射获取接口类型
func GetKind(i interface{}) Kind {
	if i == nil {
		return NIL
	}
	return reflect.TypeOf(i).Kind()
}

// 反射获取信息
func getReflect(i interface{}) (reflect.Type, reflect.Value) {
	tof := reflect.TypeOf(i)
	var vof reflect.Value
	switch tof.Kind() {
	case STRUCT: vof = reflect.ValueOf(i)
	case PTR:
		if tof.Elem().Kind() == STRUCT {
			vof = reflect.ValueOf(i).Elem()
		}

	}
	return tof,vof
}

// 反射字段,接口类型是指针或则实体
func toReflect(i interface{},isAll bool) (rts []reflect.StructField,rvs []reflect.Value) {
	_, vof := getReflect(i)
	for k := 0; k<vof.NumField(); k++ {
		if vof.Field(k).Kind() == STRUCT && isAll {
			rt,rv := nReflect(vof.Field(k).Interface())
			if rt == nil && rv == nil{
				rvs = append(rvs,vof.Field(k))
				rts = append(rts,vof.Type().Field(k))
			}else {
				rvs = append(rvs,rv...)  //数组添加数组
				rts = append(rts,rt...)  //数组添加数组
			}

		} else{
			rvs = append(rvs,vof.Field(k))
			rts = append(rts,vof.Type().Field(k))
		}
	}
	return
}

// 获取非第一层多次反射struct，有json tag的字段
func nReflect(i interface{}) (rts []reflect.StructField,rvs []reflect.Value) {
	tof, vof := getReflect(i)
	count := 0
	for k := 0; k<vof.NumField(); k++ {
		if vof.Field(k).Kind() == STRUCT {
			if rt, rv := nReflect(vof.Field(k).Interface()); rt != nil && rv != nil{
				rts = append(rts, rt...)
				rvs = append(rvs, rv...)
			}
		}
		if _, ok := tof.Field(k).Tag.Lookup("json"); ok{
			rvs = append(rvs,vof.Field(k))
			rts = append(rts,vof.Type().Field(k))
		}else {
			count++
		}
	}
	if count == vof.NumField(){
		return nil,nil
	}
	return
}

// 取出含有json注解字段
func (o *OfReflect) GetJsonReflect () *OfReflect {
	rts,rvs := make([]reflect.StructField,0), make([]reflect.Value,0)
	for i,field := range o.Tof {
		if _,ok := field.Tag.Lookup(JSON); ok{
			rts = append(rts, field)
			rvs = append(rvs, o.Vof[i])
		}
	}
	return &OfReflect{Tof: rts,Vof: rvs}
}

// 取出非空且含有json注解字段
func (o *OfReflect) GetNonJsonReflect () *OfReflect {
	rts,rvs := make([]reflect.StructField,0), make([]reflect.Value,0)
	for i, value := range o.Vof {
		if IsNon(value.Interface()){
			field := o.Tof[i]
			if _,ok := field.Tag.Lookup(JSON); ok{
				rts = append(rts, field)
				rvs = append(rvs, value)
			}
		}
	}
	return &OfReflect{Tof: rts,Vof: rvs}
}

// 反射Column/Json的tags/value
func (o *OfReflect) NewOfDic(hasNull bool) *OfDic{
	tagsOfName := make(map[string]map[string]string)
	tagsOfJson := make(map[string]map[string]string)
	valueOfName := make(map[string]interface{})
	valueOfJson := make(map[string]interface{})

	var of *OfReflect
	if hasNull{ // 获取含有json字段
		of = o.GetJsonReflect()
	}else { // 获取非空且含有json字段
		of = o.GetNonJsonReflect()
	}
	for i, field := range of.Tof {
		//Tags
		mtags := make(map[string]string)
		var json string
		for _,tag := range Tags {
			if val,ok := field.Tag.Lookup(tag); ok {
				mtags[tag] = val
				if tag == JSON {
					json = val
				}
			}
		}
		//TagsOfName
		tagsOfName[field.Name] = mtags
		//ValueOfName
		valueOfName[field.Name] = of.Vof[i].Interface()
		//TagsOfJson
		tagsOfJson[json] = mtags
		//ValueOfJson
		valueOfJson[json] = of.Vof[i].Interface()
	}
	return &OfDic{TagsOfName: tagsOfName,TagsOfJson: tagsOfJson,ValueOfName: valueOfName,ValueOfJson: valueOfJson}
}
