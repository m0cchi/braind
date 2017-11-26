package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/m0cchi/braind/model"
	"net/http"
	"regexp"
)

var reAbstruct = regexp.MustCompile(fmt.Sprintf(".{0,%d}", model.AbstructLength))

func newAbstruct(body string) string {
	re := reAbstruct.Copy()
	return re.FindString(body)
}

// Post is Handler of gin
func Post(c *gin.Context) {
	title := c.PostForm("articleTitle")
	body := c.PostForm("articleBody")
	abstruct := newAbstruct(body)
	header := &model.ArticleHeader{
		Title:    title,
		Abstruct: abstruct,
	}
	article := &model.Article{
		Header: header,
		Body:   body,
	}
	err := article.Create(datasource)
	res := gin.H{}

	if err == nil {
		err = updatePage()
	}

	if err != nil {
		res["ok"] = false
		res["message"] = fmt.Sprintf("%v", err)
	} else {
		res["ok"] = true
	}

	c.JSON(http.StatusOK, res)
}
