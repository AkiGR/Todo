package main

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        int
	Title     string
	Detail    string
	CreatedAt string
	Done      bool
}

var todos []Todo
var nextID = 1

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.SetFuncMap(template.FuncMap{
		"add": func(a, b int) int { return a + b },
	})
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todos,
		})
	})

	r.POST("/add", func(c *gin.Context) {
		title := c.PostForm("title")
		detail := c.PostForm("detail")

		if title != "" {
			todo := Todo{
				ID:        nextID,
				Title:     title,
				Detail:    detail,
				CreatedAt: time.Now().Format("2006/01/02 15:04"),
				Done:      false,
			}
			todos = append(todos, todo)
			nextID++
		}
		c.Redirect(http.StatusSeeOther, "/")
	})

	r.POST("/toggle/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err == nil {
			for i := range todos {
				if todos[i].ID == id {
					todos[i].Done = !todos[i].Done
					break
				}
			}
		}
		c.Redirect(http.StatusSeeOther, "/")
	})

	r.POST("/delete/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err == nil {
			for i, t := range todos {
				if t.ID == id {
					todos = append(todos[:i], todos[i+1:]...)
					break
				}
			}
		}
		c.Redirect(http.StatusSeeOther, "/")
	})

	r.Run(":8080")
}
