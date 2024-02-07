package main

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	utils "htmx_example/utils"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.StaticFile("/static", "./static/index.html")
	r.StaticFile("/static/temp", "./static/temp.html")
	r.StaticFile("/static/loader.gif", "./static/loader.gif")

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "hello world"})
		})
		v1.GET("/users", func(c *gin.Context) {
			limit := c.Query("limit")
			limitInt, err := strconv.Atoi(limit)
			if err != nil {
				limitInt = 10
			}
			users := utils.GenerateRandomUsers(limitInt)
			c.HTML(http.StatusOK, "users.html", gin.H{
				"Users": users,
			})
		})

		// Handle POST request for temp conversion
		v1.POST("/convert", func(c *gin.Context) {
			r := c.Request
			err := r.ParseForm()
			if err != nil {
				slog.Error(err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
				return
			}
			temp := r.FormValue("fahrenheit")
			tempFloat, err := strconv.ParseFloat(temp, 64)
			if err != nil {
				slog.Error(err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid temperature"})
				return
			}
			convertedTemp := tempFloat*9/5 + 32
			temperature := strconv.FormatFloat(convertedTemp, 'f', 2, 64)
			time.Sleep(2 * time.Second)
			c.String(http.StatusOK, temperature)
		})
	}

	r.Run()
}
