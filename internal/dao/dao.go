package dao

import (
	"sync"

	"github.com/astaxie/beego/orm"
)

var globalOrm orm.Ormer
var once sync.Once

// GetOrmer :set ormer singleton
func GetOrmer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}

// Transactional invoke lambda function within transaction
func Transactional(ormer orm.Ormer, handle func() error) (err error) {
	err = ormer.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			ormer.Rollback()
			panic(p)
		} else if err != nil {
			ormer.Rollback()
		} else {
			err = ormer.Commit()
		}
	}()
	err = handle()
	return
}
