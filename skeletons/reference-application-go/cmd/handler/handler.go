package handler

import (
	"fmt"
	"github.com/coreeng/core-reference-application-go/cmd/database"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func Router() http.Handler {
	router := gin.Default()

	setupHelloRoutes(router)
	setupCounterRoutes(router)
	setupDownstreamRoutes(router)
	return router
}

func InternalRouter() http.Handler {
	router := gin.Default()
	routerGroup1 := router.Group("/internal")
	routerGroup1.GET("/status", func(c *gin.Context) { c.Status(http.StatusOK) })
	return router
}

func setupDownstreamRoutes(r *gin.Engine) {
	r.GET("/delay/:delayAmountSeconds", handleDelay)
	r.GET("/status/:status", handleStatus)
}

func setupCounterRoutes(r *gin.Engine) {
	r.GET("/counter/:counterName", handleGetCounter)
	r.PUT("/counter/:counterName", handleIncrementCounter)
}

func setupHelloRoutes(r *gin.Engine) gin.IRoutes {
	return r.GET("/hello", handleHello)
}

func handleHello(c *gin.Context) {
	nameOrDefault := c.DefaultQuery("name", "world")
	c.String(http.StatusOK, "Hello %s!", nameOrDefault)
}

func handleGetCounter(c *gin.Context) {
	counter := database.GetCounter(c.Param("counterName"))
	c.JSON(http.StatusOK, counter)
}

func handleIncrementCounter(c *gin.Context) {
	incrementedCounter := database.IncrementCounterValue(c.Param("counterName"))
	c.JSON(http.StatusOK, incrementedCounter)
}

func handleDelay(c *gin.Context) {
	host := os.Getenv("DOWNSTREAM_HOST")
	if host == "" {
		host = "localhost"
	}
	r, err := http.Get(fmt.Sprintf("http://%s:9898/delay/%s", host, c.Param("delayAmountSeconds")))
	if err != nil {
		log.Errorf("Client failed with the status code \"%d\" and error %s", r.StatusCode, err)
		return
	}

	c.String(http.StatusOK, "Delay amount "+c.Param("delayAmountSeconds"))
}

func handleStatus(c *gin.Context) {
	host := os.Getenv("DOWNSTREAM_HOST")
	if host == "" {
		host = "localhost"
	}
	r, err := http.Get(fmt.Sprintf("http://%s:9898/status/%s", host, c.Param("status")))
	if err != nil {
		log.Errorf("Client failed with the status code \"%d\" and error %s", r.StatusCode, err)
		return
	}
	c.String(http.StatusOK, "Status "+c.Param("status"))
}
