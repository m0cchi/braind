package main

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	//	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/m0cchi/braind/handler"
	"github.com/m0cchi/gfalcon/util"
	"github.com/utrack/gin-csrf"

	"os"
)

func init() {

}

func main() {
	datasource, err := sqlx.Connect("mysql", os.Getenv("DATASOURCE"))
	if err != nil {
		fmt.Printf("failed to open datasource: %v\n", err)
		os.Exit(1)
	}
	err = handler.InitHandler(datasource)
	if err != nil {
		fmt.Printf("failed to InitHandler: %v\n", err)
		os.Exit(1)
	}
	defer handler.Final()

	r := gin.Default()

	templates := multitemplate.New()
	templates.AddFromFiles("index.html.tmpl",
		"./resources/templates/layout.html.tmpl",
		"./resources/templates/navbar.html.tmpl",
		"./resources/templates/index.html.tmpl")
	templates.AddFromFiles("article.html.tmpl",
		"./resources/templates/layout.html.tmpl",
		"./resources/templates/navbar.html.tmpl",
		"./resources/templates/article.html.tmpl")
	r.HTMLRender = templates

	r.Static("/statics", "./resources/public/")

	secret := util.GenerateSessionID(255)
	store := sessions.NewCookieStore([]byte(secret))
	r.Use(sessions.Sessions("session", store))

	r.Use(csrf.Middleware(csrf.Options{
		Secret: secret,
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	r.GET("/", handler.Index)

	r.GET("/article", handler.Article)

	r.POST("/post", handler.Post)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	r.Run(":" + port)

}
