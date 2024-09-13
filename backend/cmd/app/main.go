package main

//
// import (
// 	"log"
//
// 	_ "github.com/Inteli-College/2024-1B-T02-EC10-G04/docs"
// 	"github.com/gin-gonic/gin"
// 	ginSwagger "github.com/swaggo/gin-swagger"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )
//
// // @title Location API
// // @version 1.0
// // @description API for managing locations.
// // @host localhost:8080
// // @BasePath /api/v1
// func main() {
// 	// Inicializa a conexão com o banco de dados
// 	db, err := gorm.Open(postgres.Open("host=localhost user=youruser dbname=yourdb password=yourpassword sslmode=disable"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}
//
// 	// Inicializa o handler usando Wire
// 	locationHandler := InitializeLocationHandler(db)
//
// 	// Inicializa o router Gin
// 	r := gin.Default()
//
// 	// Configura o Swagger
// 	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
//
// 	// Configura os endpoints da API
// 	apiGroup := r.Group("/api/v1")
// 	locationHandler.RegisterRoutes(apiGroup)
//
// 	// Inicia o servidor
// 	r.Run(":8080")
// }
