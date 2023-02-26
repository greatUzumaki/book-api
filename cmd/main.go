package main

import (
	"log"
	"test-api/pkg/common/config"
	"test-api/pkg/common/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	c := config.GetConfig() // Подлкючение конфига для получения переменных из .env

	r := gin.Default()
	db.Init(c.DbUrl)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  c.Port,
			"dbUrl": c.DbUrl,
		})
	})

	r.Run(c.Port)
}
