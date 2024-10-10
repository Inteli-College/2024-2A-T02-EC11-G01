package main

/*
#cgo LDFLAGS: -L./lib -lprover -lpthread -ldl -lm -lstdc++
#cgo CFLAGS: -I./include

#include <stdlib.h>

const char* notarize_request();
void free_string(char* s);
*/
import "C"
import (
	"log"
	"os"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/configs"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/rabbitmq"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/rollups_contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	/////////////////////// Configs /////////////////////////
	pk, isSet := os.LookupEnv("TESTNET_PRIVATE_KEY")
	if !isSet {
		log.Fatalf("TESTNET_PRIVATE_KEY is not set")
	}

	rpcUrl, isSet := os.LookupEnv("TESTNET_RPC_URL")
	if !isSet {
		log.Fatalf("TESTNET_RPC_URL is not set")
	}

	input_box_address, isSet := os.LookupEnv("INPUT_BOX_ADDRESS")
	if !isSet {
		log.Fatalf("INPUT_BOX_ADDRESS is not set")
	}

	application_address, isSet := os.LookupEnv("APPLICATION_ADDRESS")
	if !isSet {
		log.Fatalf("APPLICATION_ADDRESS is not set")
	}

	ch, err := configs.SetupRabbitMQChannel()
	if err != nil {
		panic(err)
	}

	client, opts, err := configs.SetupTransactor(rpcUrl, pk)
	if err != nil {
		panic(err)
	}

	/////////////////////// Webserver /////////////////////////
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

	/////////////////////// Predictions Consumer /////////////////////////
	msgChan := make(chan amqp.Delivery)
	go func() {
		if err := rabbitmq.NewRabbitMQConsumer(ch).Consume(msgChan, "prediction.created"); err != nil {
			panic(err)
		}
	}()

	for msg := range msgChan {
		log.Printf("Event received: %v", string(msg.Body))
		result := C.notarize_request()

		instance, err := rollups_contracts.NewInputBox(common.HexToAddress(input_box_address), client)
		if err != nil {
			log.Fatal(err)
		}

		tx, err := instance.AddInput(opts, common.HexToAddress(application_address), []byte(C.GoString(result)))
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Transaction sent: %s", tx.Hash().Hex())

		C.free_string(result)
	}
}
