package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	appSwagDocs "github.com/Inteli-College/2024-2A-T02-EC11-G01/api"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/rabbitmq"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/usecase/prediction_usecase"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/events"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

// @host	localhost:8081
// @BasePath	/api/v1
func main() {
	/////////////////////// Event Dispatcher /////////////////////////

	eventDispatcher := events.NewEventDispatcher()

	locationCreatedHandler, err := NewLocationCreatedHandler()
	if err != nil {
		panic(err)
	}
	eventDispatcher.Register("LocationCreated", locationCreatedHandler)

	predictionCreatedHandler, err := NewPredictionCreatedHandler()
	if err != nil {
		panic(err)
	}
	eventDispatcher.Register("PredictionCreated", predictionCreatedHandler)

	/////////////////////// Use Cases /////////////////////////
	pu, err := NewCreatePredictionUseCase(eventDispatcher)
	if err != nil {
		panic(err)
	}

	/////////////////////// Web Handlers /////////////////////////
	lh, err := NewLocationWebHandlers(eventDispatcher)
	if err != nil {
		panic(err)
	}

	ph, err := NewPredicitonWebHandlers(eventDispatcher)
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

		appSwagDocs.SwaggerInfo.Host = swaggerHost
	}

	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	///////////////////////// Predictions ///////////////////////////

	{
		predictionsGroup := api.Group("/predictions")
		{
			predictionsGroup.POST("", ph.PredictionWebHandlers.CreatePredictionHandler)
			predictionsGroup.GET("", ph.PredictionWebHandlers.FindAllPredictionsHandler)
			predictionsGroup.GET("/:id", ph.PredictionWebHandlers.FindPredictionByIdHandler)
			predictionsGroup.GET("/location/:id", ph.PredictionWebHandlers.FindAllPredictionsByLocationIdHandler)
			predictionsGroup.PUT("/:id", ph.PredictionWebHandlers.UpdatePredictionHandler)
			predictionsGroup.DELETE("/:id", ph.PredictionWebHandlers.DeletePredictionHandler)
		}
	}

	///////////////////////// Locations ///////////////////////////

	{
		locationsGroup := api.Group("/locations")
		{
			locationsGroup.POST("", lh.LocationWebHandlers.CreateLocationHandler)
			locationsGroup.GET("", lh.LocationWebHandlers.FindAllLocationsHandler)
			locationsGroup.GET("/:id", lh.LocationWebHandlers.FindLocationByIdHandler)
			locationsGroup.PUT("/:id", lh.LocationWebHandlers.UpdateLocationHandler)
			locationsGroup.DELETE("/:id", lh.LocationWebHandlers.DeleteLocationHandler)
		}
	}

	go func() {
		if err := router.Run(":8081"); err != nil {
			log.Fatalf("Failed to start the web server: %v", err)
		}
	}()

	/////////////////////// Predictions Consumer /////////////////////////
	ch, err := NewRabbitChannel()
	if err != nil {
		panic(err)
	}

	msgChan := make(chan amqp.Delivery)
	go func() {
		if err := rabbitmq.NewRabbitMQConsumer(ch).Consume(msgChan, "prediction"); err != nil {
			panic(err)
		}
	}()

	for msg := range msgChan {
		var prediction prediction_usecase.CreatePredictionInputDTO
		err := json.Unmarshal(msg.Body, &prediction)
		if err != nil {
			panic(err)
		}
		ctx := context.Background()
		res, err := pu.Execute(ctx, prediction)
		if err != nil {
			panic(err)
		}
		log.Printf("Prediciton created: %v", res)
	}
}
