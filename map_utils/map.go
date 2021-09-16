package map_dbcustom_plus

import (
	"errors"
	"fmt"
	"github.com/by-zxy/dbcustom_plus/ddb_utils"
	"github.com/by-zxy/dbcustom_plus/str_utils"
)

type iInt = map[int]int
type iStr = map[int]string
type iVal = map[int]interface{}
type sInt = map[string]int
type sStr = map[string]string
type sVal = map[string]interface{}
type vInt = map[interface{}]int
type vStr = map[interface{}]string
type vVal = map[interface{}]interface{}

type Map struct {
	iiMap  	iInt
	isMap	iStr
	ivMap	iVal
	siMap	sInt
	ssMap	sStr
	svMap	sVal
	viMap	vInt
	vsMap	vStr
	vvMap	vVal
	status 	string
}

// 实例化
func IiMap() *Map {
	return &Map{
		iiMap: make(iInt,0),
		status: "iiMap",
	}
}
func IsMap() *Map {
	return &Map{
		isMap: make(iStr,0),
		status: "isMap",
	}
}
func IvMap() *Map {
	return &Map{
		ivMap: make(iVal,0),
		status: "ivMap",
	}
}
func SiMap() *Map {
	return &Map{
		siMap: make(sInt,0),
		status: "siMap",
	}
}
func SsMap() *Map {
	return &Map{
		ssMap: make(sStr,0),
		status: "ssMap",
	}
}
func SvMap() *Map {
	return &Map{
		svMap: make(sVal,0),
		status: "svMap",
	}
}
func ViMap() *Map {
	return &Map{
		viMap: make(vInt,0),
		status: "viMap",
	}
}
func VsMap() *Map {
	return &Map{
		vsMap: make(vStr,0),
		status: "vsMap",
	}
}
func VvMap() *Map {
	return &Map{
		vvMap: make(vVal,0),
		status: "vvMap",
	}
}

// 获取当前使用类型
func (m *Map) getMap() string {
	if m.iiMap != nil{
		return "iiMap"
	}
	if m.isMap != nil{
		return "isMap"
	}
	if m.ivMap != nil{
		return "ivMap"
	}
	if m.siMap != nil{
		return "siMap"
	}
	if m.ssMap != nil{
		return "ssMap"
	}
	if m.svMap != nil{
		return "svMap"
	}
	if m.viMap != nil{
		return "viMap"
	}
	if m.vsMap != nil{
		return "vsMap"
	}
	if m.vvMap != nil{
		return "vvMap"
	}
	return ""
}

// map.put
func (m *Map) Put(key,value interface{}) *Map {
	if ddb_utils.GetKind(key) == ddb_utils.PTR || ddb_utils.GetKind(value) == ddb_utils.PTR {
		panic(errors.New("Map.Put: 参数只能是值，不能为指针"))
	}
	switch m.status {
	case "iiMap": m.iiMap[key.(int)] = value.(int)
	case "isMap": m.isMap[key.(int)] = value.(string)
	case "ivMap": m.ivMap[key.(int)] = value
	case "siMap": m.siMap[key.(string)] = value.(int)
	case "ssMap": m.ssMap[key.(string)] = value.(string)
	case "svMap": m.svMap[key.(string)] = value
	case "viMap": m.viMap[key] = value.(int)
	case "vsMap": m.vsMap[key] = value.(string)
	case "vvMap": m.vvMap[key] = value
	}
	return m
}

func (m *Map) PutAll(maps interface{}) *Map {
	if ddb_utils.GetKind(maps) == ddb_utils.PTR {
		panic(errors.New("Map.PutAll: 参数map的值"))
	}
	switch m.status {
	case "iiMap":
		for k, v := range maps.(iInt) {
			m.iiMap[k] = v
		}
	case "isMap":
		for k, v := range maps.(iStr) {
			m.isMap[k] = v
		}
	case "ivMap":
		for k, v := range maps.(iVal) {
			m.ivMap[k] = v
		}
	case "siMap":
		for k, v := range maps.(sInt) {
			m.siMap[k] = v
		}
	case "ssMap":
		for k, v := range maps.(sStr) {
			m.ssMap[k] = v
		}
	case "svMap":
		for k, v := range maps.(sVal) {
			m.svMap[k] = v
		}
	case "viMap":
		for k, v := range maps.(vInt) {
			m.viMap[k] = v
		}
	case "vsMap":
		for k, v := range maps.(vStr) {
			m.vsMap[k] = v
		}
	case "vvMap":
		for k, v := range maps.(vVal) {
			m.vvMap[k] = v
		}
	}
	return m
}

// map.get
func (m *Map) Get(key interface{}) interface{} {
	if ddb_utils.GetKind(key) == ddb_utils.PTR {
		panic(errors.New("Map.Get: 参数只能为值"))
	}
	switch m.status {
	case "iiMap": return m.iiMap[key.(int)]
	case "isMap": return m.isMap[key.(int)]
	case "ivMap": return m.ivMap[key.(int)]
	case "siMap": return m.siMap[key.(string)]
	case "ssMap": return m.ssMap[key.(string)]
	case "svMap": return m.svMap[key.(string)]
	case "viMap": return m.viMap[key]
	case "vsMap": return m.vsMap[key]
	case "vvMap": return m.vvMap[key]
	}
	return nil
}

// map.delete
func (m *Map) DelKey(key interface{}) *Map {
	if ddb_utils.GetKind(key) == ddb_utils.PTR {
		panic(errors.New("Map.DelKey: 参数只能为值"))
	}
	switch m.status {
	case "iiMap":
		delete(m.iiMap, key.(int))
	case "isMap":
		delete(m.isMap, key.(int))
	case "ivMap":
		delete(m.ivMap, key.(int))
	case "siMap":
		delete(m.siMap, key.(string))
	case "ssMap":
		delete(m.ssMap, key.(string))
	case "svMap":
		delete(m.svMap, key.(string))
	case "viMap":
		delete(m.viMap, key)
	case "vsMap":
		delete(m.vsMap, key)
	case "vvMap":
		delete(m.vvMap, key)
	}

	return m
}

func (m *Map) DelKeys(keys ...interface{}) *Map {
	for _, key := range keys {
		m.DelKey(key)
	}
	return m
}

func (m *Map) Clean() *Map {
	switch m.status {
	case "iiMap": m.iiMap = make(iInt,0)
	case "isMap": m.isMap = make(iStr,0)
	case "ivMap": m.ivMap = make(iVal,0)
	case "siMap": m.siMap = make(sInt,0)
	case "ssMap": m.ssMap = make(sStr,0)
	case "svMap": m.svMap = make(sVal,0)
	case "viMap": m.viMap = make(vInt,0)
	case "vsMap": m.vsMap = make(vStr,0)
	case "vvMap": m.vvMap = make(vVal,0)
	}
	return m
}

func (m *Map) AllKv() (interface{}, interface{}) {
	switch m.status {
	case "iiMap":
		ks := make([]int,0)
		vs := make([]int,0)
		for k, v := range m.iiMap {
			ks = append(ks, k)
			vs = append(vs, v)
		}
		return ks,vs
	case "isMap":
		ks := make([]int,0)
		vs := make([]string,0)
		for k, v := range m.isMap {
			ks = append(ks, k)
			vs = append(vs, v)
		}
		return ks,vs
	case "ivMap":
		ks := make([]int,0)
		vs := make([]interface{},0)
		for k, v := range m.ivMap {
			ks = append(ks, k)
			vs = append(vs, v)
		}
		return ks,vs
	case "siMap":
		ks := make([]string,0)
		vs := make([]int,0)
		for k, v := range m.siMap {
			ks = append(ks, k)
			vs = append(vs, v)
		}
		return ks,vs
	case "ssMap":
		ks := make([]string,0)
		vs := make([]string,0)
		for k, v := range m.ssMap {
			ks = append(ks, k)
			vs = append(vs, v)
		}
		return ks,vs
	case "svMap":
		ks := make([]string,0)
		vs := make([]interface{},0)
		for k, v := range m.svMap {
			ks = append(ks, k)
			vs = append(vs, v)
		}
		return ks,vs
	case "viMap":
		ks := make([]interface{},0)
		vs := make([]int,0)
		for k, v := range m.viMap {
			ks = append(ks, k)
			vs = append(vs, v)
		}
		return ks,vs
	case "vsMap":
		ks := make([]interface{},0)
		vs := make([]string,0)
		for k, v := range m.vsMap {
			ks = append(ks, k)
			vs = append(vs, v)
		}
		return ks,vs
	case "vvMap":
		ks := make([]interface{},0)
		vs := make([]interface{},0)
		for k, v := range m.vvMap {
			ks = append(ks, k)
			vs = append(vs, v)
		}
		return ks,vs
	}
	return nil,nil
}
// map 所有key
func (m *Map) AllOfIntKey() []int {
	sets := make([]int,0)
	if m.iiMap != nil {
		for i := range m.iiMap {
			sets = append(sets, i)
		}
	}
	if m.isMap != nil {
		for i := range m.isMap {
			sets = append(sets, i)
		}
	}
	if m.ivMap != nil {
		for i := range m.ivMap {
			sets = append(sets, i)
		}
	}
	return sets
}
func (m *Map) AllOfStrKey() []string {
	sets := make([]string,0)
	if m.siMap != nil {
		for i := range m.siMap {
			sets = append(sets, i)
		}
	}
	if m.ssMap != nil {
		for i := range m.ssMap {
			sets = append(sets, i)
		}
	}
	if m.svMap != nil {
		for i := range m.svMap {
			sets = append(sets, i)
		}
	}
	return sets
}
func (m *Map) AllOfIfKey() []interface{} {
	sets := make([]interface{},0)
	if m.viMap != nil {
		for i := range m.viMap {
			sets = append(sets, i)
		}
	}
	if m.vsMap != nil {
		for i := range m.vsMap {
			sets = append(sets, i)
		}
	}
	if m.vvMap != nil {
		for i := range m.vvMap {
			sets = append(sets, i)
		}
	}
	return sets
}

// map 所有val
func (m *Map) AllOfIntVal() []int {
	sets := make([]int,0)
	if m.iiMap != nil {
		for _,v := range m.iiMap {
			sets = append(sets, v)
		}
	}
	if m.siMap != nil {
		for _,v := range m.siMap {
			sets = append(sets, v)
		}
	}
	if m.viMap != nil {
		for _,v := range m.viMap {
			sets = append(sets, v)
		}
	}
	return sets
}
func (m *Map) AllOfStrVal() []string {
	sets := make([]string,0)
	if m.isMap != nil {
		for _,v := range m.isMap {
			sets = append(sets, v)
		}
	}
	if m.ssMap != nil {
		for _,v := range m.ssMap {
			sets = append(sets, v)
		}
	}
	if m.vsMap != nil {
		for _,v := range m.vsMap {
			sets = append(sets, v)
		}
	}
	return sets
}
func (m *Map) AllOfIfVal() []interface{} {
	sets := make([]interface{},0)
	if m.ivMap != nil {
		for _,v := range m.ivMap {
			sets = append(sets, v)
		}
	}
	if m.svMap != nil {
		for _,v := range m.svMap {
			sets = append(sets, v)
		}
	}
	if m.vvMap != nil {
		for _,v := range m.vvMap {
			sets = append(sets, v)
		}
	}
	return sets
}

func (m *Map) Contain(maps interface{}) bool {
	if ddb_utils.GetKind(maps) == ddb_utils.PTR {
		panic(errors.New("Map.Contain: 参数只能为map值"))
	}
	switch m.status {
	case "iiMap":
		for k, v := range maps.(iInt) {
			if vv, ok := m.iiMap[k]; !ok{
				return false
			}else {
				if vv != v {
					return false
				}
			}
		}
		return true
	case "isMap":
		for k, v := range maps.(iStr) {
			if vv, ok := m.isMap[k]; !ok{
				return false
			}else {
				if !str_utils.Equals(vv,v) {
					return false
				}
			}
		}
		return true
	case "ivMap":
		for k, v := range maps.(iVal) {
			if vv, ok := m.ivMap[k]; !ok{
				return false
			}else {
				if ddb_utils.GetKind(vv) != ddb_utils.GetKind(v) {
					return false
				}
				if !str_utils.Equals(fmt.Sprintf("%v", vv),fmt.Sprintf("%v", v)) {
					return false
				}
			}
		}
		return true
	case "siMap":
		for k, v := range maps.(sInt) {
			if vv, ok := m.siMap[k]; !ok{
				return false
			}else {
				if vv != v {
					return false
				}
			}
		}
		return true
	case "ssMap":
		for k, v := range maps.(sStr) {
			if vv, ok := m.ssMap[k]; !ok{
				return false
			}else {
				if !str_utils.Equals(vv,v) {
					return false
				}
			}
		}
		return true
	case "svMap":
		for k, v := range maps.(sVal) {
			if vv, ok := m.svMap[k]; !ok{
				return false
			}else {
				if ddb_utils.GetKind(vv) != ddb_utils.GetKind(v) {
					return false
				}
				if !str_utils.Equals(fmt.Sprintf("%v", vv),fmt.Sprintf("%v", v)) {
					return false
				}
			}
		}
		return true
	case "viMap":
		for k, v := range maps.(vInt) {
			if vv, ok := m.viMap[k]; !ok{
				return false
			}else {
				if vv != v {
					return false
				}
			}
		}
		return true
	case "vsMap":
		for k, v := range maps.(vStr) {
			if vv, ok := m.vsMap[k]; !ok{
				return false
			}else {
				if !str_utils.Equals(vv,v) {
					return false
				}
			}
		}
		return true
	case "vvMap":
		for k, v := range maps.(vVal) {
			if vv, ok := m.vvMap[k]; !ok{
				return false
			}else {
				if ddb_utils.GetKind(vv) != ddb_utils.GetKind(v) {
					return false
				}
				if !str_utils.Equals(fmt.Sprintf("%v", vv),fmt.Sprintf("%v", v)) {
					return false
				}
			}
		}
		return true
	}
	return false
}

func (m *Map) End() interface{} {
	if m.iiMap != nil{
		return m.iiMap
	}
	if m.isMap != nil{
		return m.isMap
	}
	if m.ivMap != nil{
		return m.ivMap
	}
	if m.siMap != nil{
		return m.siMap
	}
	if m.ssMap != nil{
		return m.ssMap
	}
	if m.svMap != nil{
		return m.svMap
	}
	if m.viMap != nil{
		return m.viMap
	}
	if m.vsMap != nil{
		return m.vsMap
	}
	if m.vvMap != nil{
		return m.vvMap
	}
	return nil
}
