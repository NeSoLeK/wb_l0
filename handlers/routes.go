package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"wb_l0/db"
	"wb_l0/service"
)

type Router struct {
	app   *gin.Engine
	db    *db.DataBase
	cache *service.CacheMap
}

func CreateRouter() *Router {
	router := gin.Default()
	return &Router{app: router}
}

func (r *Router) StartHandlers(m *service.CacheMap) {

	r.app.LoadHTMLGlob("web/*")

	r.app.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	r.app.GET("/order", func(c *gin.Context) {
		id := c.Query("id")
		data := m.GetOrderByUID(id)

		c.JSON(http.StatusOK, string(data))
		if data == nil {
			c.Status(http.StatusNoContent)
		}

	})

	err := r.app.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
}
