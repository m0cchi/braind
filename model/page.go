package model

import (
	"github.com/jmoiron/sqlx"
)

const CountPage = "SELECT count(*) as `size` FROM `articles`"

// number of articles that page can include
const MaxQty = uint32(5)

type Page struct {
	Size    uint32 `db:"size"`
	MaxPage uint32
}

func GetPage(db *sqlx.DB) (page *Page, error error) {
	stmt, err := db.PrepareNamed(CountPage)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := stmt.Close()
		if error == nil && err != nil {
			error = err
		}
	}()
	page = &Page{}
	err = stmt.Get(page, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	page.MaxPage = page.Size / MaxQty

	return page, nil
}
