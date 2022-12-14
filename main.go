package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Name : Mohammad Ali Ashraf
//ranaaliashraf12@gmail.com
func Hello(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		name = "Annonymous user"
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Hello %s !!", name))
}

func SetupRouter(v ...any) *gin.Engine {
	var router *gin.Engine
	if len(v) > 0 && v[0] == "test" {
		gin.SetMode(gin.TestMode)
	}

	router = gin.Default()
	router.Handle("GET", "/login", Hello)
	return router
}
func main() {
	r := SetupRouter()
	r.Run(":8000")
}
