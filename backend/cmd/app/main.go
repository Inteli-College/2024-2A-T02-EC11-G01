package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	ourSwagDocs "github.com/Inteli-College/2024-2A-T02-EC11-G01/swagger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/penglongli/gin-metrics/ginmetrics"
	amqp "github.com/rabbitmq/amqp091-go"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			App API
//	@version		1.0
//	@description	This is a.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	App API Support
//	@contact.url	https://github.com/Inteli-College/2024-2A-T02-EC11-G01
//	@contact.email	artemis@inteli.edu.br

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host	localhost:8080
// @BasePath	/api/v1
func main() {
	godotenv.Load()
	/////////////////////// Event Dispatcher /////////////////////////
	eventDispatcher, _ := NewEventDispatcher()

	locationCreatedHandler, _ := NewLocationCreatedHandler()
	eventDispatcher.Register("LocationCreated", locationCreatedHandler)

	predicitonCreatedHandler, _ := NewPredictionCreatedHandler()
	eventDispatcher.Register("PredictionCreated", predicitonCreatedHandler)

	/////////////////////// U'se Cases /////////////////////////
	pu, _ := NewCreatePredictionUseCase()

	/////////////////////// Web Handlers /////////////////////////
	lh, err := NewLocationWebHandlers()
	if err != nil {
		panic(err)
	}

	ph, err := NewPredicitonWebHandlers()
	if err != nil {
		panic(err)
	}

	/////////////////////// Web Server /////////////////////////

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // TODO: change to false and make it for production
		AllowMethods:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/api/v1/metrics")
	m.Use(router)

	router.GET("/api/v1/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	api := router.Group("/api/v1")

	///////////////////// Swagger //////////////////////

	if swaggerHost, ok := os.LookupEnv("SWAGGER_HOST"); ok {

		ourSwagDocs.SwaggerInfo.Host = swaggerHost
	}

	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	///////////////////////// Predictions ///////////////////////////

	{
		predictionsGroup := api.Group("/predictions")
		{
			predictionsGroup.POST("", ph.PredictionWebHandlers.CreatePrediction)
			predictionsGroup.GET("", ph.PredictionWebHandlers.FindAllPredictions)
			predictionsGroup.GET("/location/:location_id", ph.PredictionWebHandlers.FindAllPredictionsByLocationId)
			predictionsGroup.GET("/:id", ph.PredictionWebHandlers.FindPredictionByPredictionId)
		}
	}

	///////////////////////// Locations ///////////////////////////

	{
		locationsGroup := api.Group("/locations")
		{
			locationsGroup.POST("", lh.LocationWebHandlers.CreateLocation)
			locationsGroup.GET("", lh.LocationWebHandlers.FindAllLocations)
			locationsGroup.GET("/:id", lh.LocationWebHandlers.FindLocationById)
			locationsGroup.PUT("/:id", lh.LocationWebHandlers.UpdateLocation)
			locationsGroup.DELETE("/:id", lh.LocationWebHandlers.DeleteLocation)
		}
	}

	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to start the web server: %v", err)
		}
	}()

	/////////////////////// Predictions Consumer /////////////////////////

	msgChan := make(chan amqp.Delivery)
	rabbitmqConsumer, _ := NewRabbitMQConsumer()
	go func() {
		if err := rabbitmqConsumer.Consume(msgChan, "predictions"); err != nil {
			panic(err)
		}
	}()

	for msg := range msgChan {
		var prediction dto.CreatePredictionInputDTO
		ctx := context.Background()
		err := json.Unmarshal(msg.Body, &prediction)
		if err != nil {
			panic(err)
		}
		res, err := pu.Execute(ctx, &prediction)
		if err != nil {
			panic(err)
		}
		log.Printf("Prediciton created: %v", res)
	}
}

