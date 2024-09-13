package main

import (
	"log"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//	@title			Manager API
//	@version		1.0
//	@description	This is a.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Manager API Support
//	@contact.url	https://github.com/Inteli-College/2024-1B-T02-EC10-G04
//	@contact.email	gomedicine@inteli.edu.br

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host	localhost:8080
//	@BasePath	/api/v1

// @SecurityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description "Type: Bearer token"
// @scheme bearer
// @bearerFormat JWT
func main() {
	godotenv.Load()

	///////////////////////// Gin ///////////////////////////

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // TODO: change to false and make it for production
		AllowMethods:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	api := router.Group("/api/v1")

	/////////////////////// Handlers /////////////////////////

	var locationsHandler *web.LocationHandler
	var predictionsHandler *web.PredictionHandler
	var handlerError error

	locationsHandler, handlerError = InitializeLocationsHandler()
	predictionsHandler, handlerError = InitializePredictionsHandler()

	if handlerError != nil {
		log.Fatalf("Failing to initialize handlers: %+v\n", handlerError)
	}

	///////////////////////// Locations ///////////////////////////

	{
		locationsGroup := api.Group("/locations")
		if locationsHandler == nil {
			log.Fatal("Failing to load locationsHandler: nil pointer")
		}
		locationsHandler.RegisterRoutes(locationsGroup)
	}

	///////////////////////// Predictions ///////////////////////////

	{
		predictionsGroup := api.Group("/predictions")
		if predictionsHandler == nil {
			log.Fatal("Failing to load locationsHandler: nil pointer")
		}
		predictionsHandler.RegisterRoutes(predictionsGroup)
	}
}

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
