package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/m0cchi/braind/model"
	"github.com/russross/blackfriday"
	"github.com/utrack/gin-csrf"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

// Article is Handler of gin
func Article(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.HTML(http.StatusOK, "article.html.tmpl",
			gin.H{
				"postable":   false,
				"csrf":       csrf.GetToken(c),
				"title":      "notfound",
				"body":       err,
				"posteddate": time.Now(),
			})
		return
	}
	article, err := model.GetArticle(datasource, uint32(id))
	if err != nil {
		c.HTML(http.StatusOK, "article.html.tmpl",
			gin.H{
				"postable":   false,
				"csrf":       csrf.GetToken(c),
				"title":      "notfound",
				"body":       err,
				"posteddate": time.Now(),
			})
		return
	}

	title := article.Header.Title
	body := article.Body
	html := blackfriday.Run([]byte(body))

	c.HTML(http.StatusOK, "article.html.tmpl",
		gin.H{
			"postable":   false,
			"csrf":       csrf.GetToken(c),
			"title":      title,
			"body":       template.HTML(string(html)),
			"posteddate": article.Header.PostingDate,
		})
}
