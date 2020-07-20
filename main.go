package main

import (
	"awesomeProject/config"
	"awesomeProject/uploadDownload"
	"awesomeProject/user"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func init()  {
	config.Connect()
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	setRoute(r)
	// set module
	user.GetRoute(r)
	uploadDownload.GetRoute(r)

	// run
	r.Run()
}

func setRoute(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
}
