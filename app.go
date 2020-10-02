package main

import "github.com/gin-gonic/gin"

import "net/http"

import "fmt"

type myForm struct {
	Colors []string `form:"colors[]"`
}


func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// For each matched request Context will hold the route definition
	router.POST("/user/:name/*action", func(c *gin.Context) {
		fmt.Printf(c.FullPath())
	})

	router.LoadHTMLGlob("views/*")
	router.GET("/", indexHandler)
	router.POST("/", formHandler)

	router.Run(":9080") // listen and serve on 0.0.0.0:9080
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.Bind(&fakeForm)
	c.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
}
