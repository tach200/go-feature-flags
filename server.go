package main

import (
	"context"
	provider "go-feature-flags/my_provider"
	"log"
	"time"

	"github.com/open-feature/go-sdk/pkg/openfeature"
)

func main() {

	openfeature.SetProvider(provider.FeatureProvider{})

	client := openfeature.NewClient("app")

	evalCtx := openfeature.NewEvaluationContext(
		"", map[string]interface{}{
			"Location": "UK",
		},
	)

	for {

		x, err := client.BooleanValue(context.Background(), "feature-x", false, evalCtx)
		if err != nil {
			log.Fatal(err)
		}

		y, err := client.BooleanValue(context.Background(), "feature-y", false, evalCtx)
		if err != nil {
			log.Fatal(err)
		}

		z, err := client.BooleanValue(context.Background(), "feature-z", true, evalCtx)
		if err != nil {
			log.Fatal(err)
		}

		tickerInterval, err := client.IntValue(context.Background(), "timer-interval", 1, evalCtx)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second * time.Duration(tickerInterval))

		if x {
			featureX()
		}
		if y {
			featureY()
		}
		if z {
			featureZ()
		}
	}
}

func featureX() {
	log.Print("Info: feature X is active")
}

func featureY() {
	log.Print("Info: feature Y is active")
}

func featureZ() {
	log.Print("Info: feature Z is active")
}
