package main

import (
	"encoding/json"
	"log"
	"os"
	_ "github.com/Inteli-College/2024-2A-T02-EC11-G01/api"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/configs"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/event/handler"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/rabbitmq"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/usecase/prediction_usecase"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/events"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/streadway/amqp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			App API
//	@version		1.0
//	@description	This is a.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	App API Support
//	@contact.url	https://github.com/Inteli-College/2024-2A-T02-EC11-G01
//	@contact.email	gomedicine@inteli.edu.br

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host	localhost:8080
//	@BasePath	/api/v1
func main() {
	/////////////////////// Configs /////////////////////////
	conn, isSet := os.LookupEnv("POSTGRES_URL")
	if !isSet {
		log.Fatalf("POSTGRES_URL is not set")
	}

	rabbitmqChannel, isSet := os.LookupEnv("RABBITMQ_CHANNEL")
	if !isSet {
		log.Fatalf("RABBITMQ_CHANNEL is not set")
	}

	db, err := configs.SetupPostgres(conn)
	if err != nil {
		panic(err)
	}

	ch, err := configs.SetupRabbitMQChannel(rabbitmqChannel)
	if err != nil {
		panic(err)
	}

	/////////////////////// Event Dispatcher /////////////////////////
	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("LocationCreated", &handler.LocationCreatedHandler{
		RabbitMQChannel: ch,
	})
	eventDispatcher.Register("PredictionCreated", &handler.PredictionCreatedHandler{
		RabbitMQChannel: ch,
	})

	/////////////////////// U'se Cases /////////////////////////
	pu := NewCreatePredictionUseCase(db, eventDispatcher)

	/////////////////////// Web Handlers /////////////////////////
	lh, err := NewLocationWebHandlers(db, eventDispatcher)
	if err != nil {
		panic(err)
	}

	ph, err := NewPredicitonWebHandlers(db, eventDispatcher)
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

	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	{
		predictionsGroup := api.Group("/predictions")
		{
			predictionsGroup.POST("", ph.PredictionWebHandlers.CreatePredictionHandler)
			predictionsGroup.GET("", ph.PredictionWebHandlers.FindAllPredictionsHandler)
			predictionsGroup.GET("/:id", ph.PredictionWebHandlers.FindPredictionByIdHandler)
			predictionsGroup.PUT("/:id", ph.PredictionWebHandlers.UpdatePredictionHandler)
			predictionsGroup.DELETE("/:id", ph.PredictionWebHandlers.DeletePredictionHandler)
		}
	}

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
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to start the web server: %v", err)
		}
	}()

	/////////////////////// Predictions Consumer /////////////////////////
	msgChan := make(chan amqp.Delivery)
	go func() {
		if err := rabbitmq.NewRabbitMQConsumer(ch).Consume(msgChan, "predictions"); err != nil {
			panic(err)
		}
	}()

	for msg := range msgChan {
		var prediction prediction_usecase.CreatePredictionInputDTO
		err := json.Unmarshal(msg.Body, &prediction)
		if err != nil {
			panic(err)
		}
		res, err := pu.Execute(prediction)
		if err != nil {
			panic(err)
		}
		log.Printf("Prediciton created: %v", res)
	}
}
