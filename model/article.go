package model

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"time"
)

const CreateArticleHeader = "INSERT INTO `articles` (`title`, `abstruct`) VALUE (:title, :abstruct)"
const CreateArticleBody = "INSERT INTO `article_bodies` (`id`, `body`) VALUE (:id, :body)"
const SelectArticleHeaders = "SELECT `id`, `title`, `abstruct`, `posting_date` FROM `articles` ORDER BY `id` DESC LIMIT :offset,:limit"
const SelectArticle = "SELECT `article_bodies`.`body` as `body`, `articles`.`id` as `header.id`, `articles`.`title` as `header.title`, `articles`.`abstruct` as `header.abstruct`, `articles`.`posting_date` as `header.posting_date` FROM `articles`, `article_bodies` WHERE `articles`.`id` = :id and `articles`.`id` = `article_bodies`.`id`"
const AbstructLength = 50

// ArticleHeader is header of article
type ArticleHeader struct {
	ID          uint32    `db:"id"`
	Title       string    `db:"title"`
	Abstruct    string    `db:"abstruct"`
	PostingDate time.Time `db:"posting_date"`
}

type Article struct {
	Header *ArticleHeader `db:"header"`
	Body   string         `db:"body"`
}

func (article *Article) createArticleHeader(tx *sqlx.Tx) (id uint32, error error) {
	stmt, err := tx.PrepareNamed(CreateArticleHeader)
	if err != nil {
		return 0, err
	}
	defer func() {
		err := stmt.Close()
		if error == nil && err != nil {
			error = err
		}
	}()

	args := map[string]interface{}{
		"title":    article.Header.Title,
		"abstruct": article.Header.Abstruct,
	}

	result, err := stmt.Exec(args)

	if err != nil {
		return 0, err
	}

	if c, err := result.LastInsertId(); err == nil {
		id = uint32(c)
	}

	return id, err
}

func (article *Article) createArticleBody(tx *sqlx.Tx) (error error) {
	stmt, err := tx.PrepareNamed(CreateArticleBody)
	if err != nil {
		return err
	}
	defer func() {
		err := stmt.Close()
		if error == nil && err != nil {
			error = err
		}
	}()

	args := map[string]interface{}{
		"id":   article.Header.ID,
		"body": article.Body,
	}
	_, err = stmt.Exec(args)

	return err
}

func (article *Article) Create(db *sqlx.DB) (error error) {
	if article == nil {
		return errors.New("nil object")
	}
	if article.Header == nil {
		return errors.New("header is nil")
	}

	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		} else if error != nil {
			tx.Rollback()
		} else {
			error = tx.Commit()
		}
	}()

	id, err := article.createArticleHeader(tx)

	if err != nil {
		return err
	}

	article.Header.ID = id
	return article.createArticleBody(tx)
}

func GetArticleHeaders(db *sqlx.DB, offset uint32) (_ []ArticleHeader, error error) {
	stmt, err := db.PrepareNamed(SelectArticleHeaders)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := stmt.Close()
		if error == nil && err != nil {
			error = err
		}
	}()
	articles := make([]ArticleHeader, 0, MaxQty)
	args := map[string]interface{}{
		"offset": offset * MaxQty,
		"limit":  MaxQty,
	}
	err = stmt.Select(&articles, args)

	return articles, error
}

func GetArticle(db *sqlx.DB, id uint32) (_ *Article, error error) {
	stmt, err := db.PrepareNamed(SelectArticle)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := stmt.Close()
		if error == nil && err != nil {
			error = err
		}
	}()

	article := &Article{}
	args := map[string]interface{}{
		"id": id,
	}
	err = stmt.Get(article, args)

	return article, err
}
