package handler

import (
	"github.com/jmoiron/sqlx"
	"github.com/m0cchi/braind/model"

	"sync"
)

var datasource *sqlx.DB
var page *model.Page

var pageLock *sync.Mutex

// InitHandler is initializer of handler
func InitHandler(_datasource *sqlx.DB) error {
	var err error
	datasource = _datasource
	page, err = model.GetPage(datasource)
	if err != nil {
		return err
	}
	pageLock = new(sync.Mutex)
	return nil
}

func updatePage() error {
	pageLock.Lock()
	defer pageLock.Unlock()
	var err error
	page, err = model.GetPage(datasource)
	return err
}
