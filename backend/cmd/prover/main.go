package main

/*
#cgo LDFLAGS: -L./lib -lprover -lpthread -ldl -lm -lstdc++ -static
#cgo CFLAGS: -I./include

#include <stdlib.h>

const char* notarize_request();
void free_string(char* s);
*/
import "C"
import (
	"log"
	"os"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/rollups_contracts"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/configs"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/rabbitmq"
	"github.com/ethereum/go-ethereum/common"
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

	rabbitmqChannel, isSet := os.LookupEnv("RABBITMQ_CHANNEL")
	if !isSet {
		log.Fatalf("RABBITMQ_CHANNEL is not set")
	}

	input_box_address, isSet := os.LookupEnv("INPUT_BOX_ADDRESS")
	if !isSet {
		log.Fatalf("INPUT_BOX_ADDRESS is not set")
	}

	application_address, isSet := os.LookupEnv("APPLICATION_ADDRESS")
	if !isSet {
		log.Fatalf("APPLICATION_ADDRESS is not set")
	}

	ch, err := configs.SetupRabbitMQChannel(rabbitmqChannel)
	if err != nil {
		panic(err)
	}

	client, opts, err := configs.SetupTransactor(rpcUrl, pk)
	if err != nil {
		panic(err)
	}

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
