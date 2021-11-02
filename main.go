package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

const PORT = "8080"

func SetupRouter() *gin.Engine {
	router := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	router.GET("/health", HealthCheck)

	return router
}

func main() {
	// Gin instance
	r := SetupRouter()

	// Start server
	if err := r.Run(":" + PORT); err != nil {
		log.Fatal(err)
	}
}

func HealthCheck(c *gin.Context) {
	log.WithFields(log.Fields{
		"user_agent": c.Request.Header.Get("User-Agent"),
	}).Info("received request")

	c.JSON(http.StatusOK, gin.H{
		"data": "Server is up and running",
	})
}
