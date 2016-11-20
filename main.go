package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/snickers54/dynamic-json-binding/binding"
	"github.com/snickers54/dynamic-json-binding/middlewares"
	"github.com/snickers54/dynamic-json-binding/models"
)

func main() {
	binding.RegisterTypes(models.User{})
	r := gin.Default()
	r.Use(middlewares.Bind)
	r.POST("/ping", func(c *gin.Context) {
		value, exists := c.Get("JSON_BINDING")
		if exists == true {
			my := value.(*models.User)
			fmt.Println(my)
		}
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and server on 0.0.0.0:8080
}
