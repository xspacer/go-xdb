package xdb

import (
	"github.com/jinzhu/gorm"
)

type DB struct {
	*gorm.DB
}

func New(driver, dataSource string, opts ...Option) (*DB, func(), error) {
	options := options{singularTable: _defaultSingularTable}
	for _, o := range opts {
		o.apply(&options)
	}

	db, err := gorm.Open(driver, dataSource)
	if err != nil {
		return nil, nil, err
	}

	db.SingularTable(options.singularTable)
	return &DB{db}, func() { db.Close() }, nil
}

func (db *DB) RunInTransaction(fn func(*DB) error) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := fn(&DB{tx}); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
