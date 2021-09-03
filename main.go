package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/dmaizel/my-gin/gin"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	stu1 := &student{Name: "Daniel", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}

	r.GET("/students", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title":  "gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gin.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gin.H{
			"title": "gin",
			"now":   time.Date(2020, 9, 18, 0, 0, 0, 0, time.UTC),
		})
	})

	r2 := gin.Default()
	r2.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World\n")
	})
	// index out of range for testing Recovery()
	r2.GET("/panic", func(c *gin.Context) {
		names := []string{"Daniel"}
		c.String(http.StatusOK, names[100])
	})

	r2.Run(":9998")
}
