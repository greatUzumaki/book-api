package main

import (
	"fmt"
	"log"
	"net/http"
	"test-api/pkg/books"
	"test-api/pkg/common/config"
	"test-api/pkg/common/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	docs "test-api/cmd/server/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

//	@title			Book API
//	@version		1.0
//	@description	Simple API GO + Gin

// @host		localhost:8001
// @BasePath	/api/v1
func main() {
	c := config.GetConfig() // Подлкючение конфига для получения переменных из .env

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DbLogin, c.DbPassword, c.DbHost, c.DbPort, c.DbName)
	fmt.Println(dbUrl)
	h := db.Init(dbUrl)

	r.GET("/", HealthCheck)

	v1 := r.Group("/api/v1")
	{
		books.RegisterRoutes(v1, h)
	}

	port := fmt.Sprintf(":%s", c.Port)
	fmt.Println("Listen on " + port)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(port)
}

// @Summary     Show the status of server.
// @Description get the status of server.
// @Tags        root
// @Accept      */*
// @Produce     json
// @Success     200 {object} map[string]interface{}
// @Router      / [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}
