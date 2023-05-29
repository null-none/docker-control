package main

import (
	Controller "./controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/*
		Controller.RunContainerBackground("hello-world")
		Controller.PullImage("hello-world")
	*/
	r := gin.Default()
	r.GET("/containers", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"containers": Controller.ListContainers(),
		})
	})
	r.GET("/images", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"images": Controller.ListImages(),
		})
	})
	r.GET("/containers/stop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"containers": Controller.StopRunningContainers(),
		})
	})
	r.Run()
}
