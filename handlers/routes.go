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

func CreateRouter(base *db.DataBase, cache *service.CacheMap) *Router {
	router := gin.Default()
	return &Router{app: router, db: base, cache: cache}
}

func (r *Router) StartHandlers() {

	r.app.LoadHTMLGlob("web/*")

	r.app.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	r.app.GET("/order", func(c *gin.Context) {
		id := c.Query("id")
		data := r.cache.GetOrderByUID(id)
		if data == nil {
			c.Status(http.StatusNoContent)
			return
		}
		c.JSON(http.StatusOK, string(data))

	})

	err := r.app.Run("main:8080")
	if err != nil {
		log.Fatal(err)
	}
}
