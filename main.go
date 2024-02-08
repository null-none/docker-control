package main

import (
	"net/http"

	Controller "example.com/example/controllers"

	"github.com/gin-gonic/gin"
)

type Validate struct {
	Image     string `form:"image" json:"image" xml:"image"  binding:"required"`
	Container string `form:"container" json:"container" xml:"container" binding:"required"`
}

func main() {
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
	r.POST("/images/pull", func(c *gin.Context) {
		var json Validate
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": Controller.PullImage(json.Image)})
	})
	r.POST("/container/run", func(c *gin.Context) {
		var json Validate
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": Controller.RunContainerBackground(json.Container)})
	})
	r.POST("/container/log", func(c *gin.Context) {
		var json Validate
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": Controller.LogContainer(json.Container)})
	})

	r.Run()
}
