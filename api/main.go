package main

import (
	"net/http"
	"strconv"
	"yakshop/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/yak-shop/stock/:dayT", func(c *gin.Context) {
		var dayT int
		var err error
		if dayT, err = strconv.Atoi(c.Param("dayT")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		output, err := controller.CaculateYak(dayT)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"milk": output.MilkTotal, "woolTotal": output.WoolTotal})
	})

	router.GET("/yak-shop/herd/:dayT", func(c *gin.Context) {
		var dayT int
		var err error
		if dayT, err = strconv.Atoi(c.Param("dayT")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		output, err := controller.CaculateYak(dayT)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, output.Yak)
	})

	router.Run(":8080")
}
