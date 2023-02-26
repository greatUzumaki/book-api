package main

import (
	"fmt"
	"log"
	"test-api/pkg/common/config"
	"test-api/pkg/common/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	c := config.GetConfig() // Подлкючение конфига для получения переменных из .env

	r := gin.Default()

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DbLogin, c.DbPassword, c.DbHost, c.DbPort, c.DbName)
	db.Init(dbUrl)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  c.Port,
			"dbUrl": c.DbHost,
		})
	})

	r.Run(c.Port)
}
