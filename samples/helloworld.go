package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/pkg/errors"

	"github.com/PandaXGO/device-sdk-go/client"
	"github.com/PandaXGO/device-sdk-go/util/wait"
)

const (
	_brokerAddr = "tcp://127.0.0.1:1883"
	_username   = "ZjI1M2IyNGMtNjNjZi0zMzM5LWFlMDMtYjBkOWVlYTQ4ZDNh"
	_pwd        = ""
)

func main() {
	log.SetFlags(log.Lshortfile)

	cli := client.NewClient(_brokerAddr, _username, _pwd)()

	err := cli.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	_ = cli.SubscribeRpcReq(context.TODO(), commandsTopicHandler)

	tm := time.Second * 1

	wait.EveryWithContext(context.Background(), func(ctx context.Context) {
		v, _ := deviceValue()
		// telemetry.
		_ = cli.PublishTelemetry(ctx, v)
	}, tm)

	select {}
}

func commandsTopicHandler(message client.Message) (interface{}, error) {
	fmt.Printf("commands=%s\n", string(message.Payload()))
	return "success", nil
}

func deviceValue() ([]byte, error) {
	mv := map[string]interface{}{
		"humidity":    rand.Intn(20),
		"temperature": rand.Intn(20),
	}

	bys, err := json.Marshal(mv)
	if err != nil {
		err = errors.Wrap(err, "deviceValue")
	}

	return bys, err
}
