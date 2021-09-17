package ddb

import (
	"errors"
	"gorm.io/gorm"
)

/**
	不同框架随意亲自搭建db
	需要用上封装方法，可以自己重写当前go
*/

// if no modal ,DB func must return
type dDB struct {
	db *gorm.DB
	dest interface{}  //保存目标modal对象
}

// get a db`transaction
func (db *dDB) transaction(fc func(tx *dDB) error) error {
	return db.db.Transaction(func(tx *gorm.DB) error{
		db.db = tx
		return fc(db) //只需返回error作判断即可
	})
}

// new DB and copy origin DB.DB and reset origin DB.DB
func (db *dDB) clone() *dDB {
	tmpDB := *db.db
	tx := &dDB{
		db: &tmpDB,
	}
	return tx
}

var gdb *gorm.DB
// you can update this func
// get a new DB
func db(defDb *gorm.DB) {
	gdb = defDb
}

func newDB() *dDB {
	if gdb == nil {
		panic(errors.New("newDB: 参数gorm.DB为指针"))
	}
	tmpDB := *gdb
	return &dDB{
		db: &tmpDB,
	}
}
