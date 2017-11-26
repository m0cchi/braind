package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/m0cchi/braind/model"
	"github.com/utrack/gin-csrf"
	"net/http"
	"strconv"
	//	"time"
)

/*
func mockarticles() []model.ArticleHeader {
	ret := make([]model.ArticleHeader, 3, 3)

	ret[0] = model.ArticleHeader{
		ID:          0,
		Title:       "super title",
		Abstruct:    "hogehogehoge....",
		PostingDate: time.Now(),
	}
	ret[1] = model.ArticleHeader{
		ID:          1,
		Title:       "hyper title",
		Abstruct:    "title te te te...",
		PostingDate: time.Now(),
	}
	ret[2] = model.ArticleHeader{
		ID:          2,
		Title:       "middle title",
		Abstruct:    "sususususus",
		PostingDate: time.Now(),
	}
	return ret
}*/

// Index is Handler of gin
func Index(c *gin.Context) {
	currentNoStr := c.Query("page")
	currentNo, err := strconv.ParseUint(currentNoStr, 10, 32)
	if err != nil {
		currentNo = 0
	}
	canPrev := uint32(currentNo) > 0
	canNext := uint32(currentNo) < page.MaxPage
	var prevNum uint32
	var nextNum uint32
	if canPrev {
		prevNum = uint32(currentNo) - uint32(1)
	}

	if canNext {
		nextNum = uint32(currentNo) + uint32(1)
	}

	articles, err := model.GetArticleHeaders(datasource, uint32(currentNo))
	if err != nil {
		articles = make([]model.ArticleHeader, 0, 0)
	}

	c.HTML(http.StatusOK, "index.html.tmpl",
		gin.H{
			"canPrev":  canPrev,
			"canNext":  canNext,
			"prevNum":  prevNum,
			"nextNum":  nextNum,
			"postable": true,
			"csrf":     csrf.GetToken(c),
			"articles": articles,
			"page":     page,
		})

}
