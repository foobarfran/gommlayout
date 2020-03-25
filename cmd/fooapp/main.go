package main

import (
	"github.com/foobarfran/gommlayout/pkg/fooapi/foo"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		flight := &foo.Flight{Origin: "BUE", Destination: "MIA"}
		c.JSON(200, flight)

	})

	r.Run()

}
