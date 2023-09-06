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

	for {

		x, err := client.BooleanValue(context.Background(), "feature-x", false, openfeature.EvaluationContext{})
		if err != nil {
			log.Fatal(err)
		}

		y, err := client.BooleanValue(context.Background(), "feature-y", false, openfeature.EvaluationContext{})
		if err != nil {
			log.Fatal(err)
		}

		z, err := client.BooleanValue(context.Background(), "feature-z", true, openfeature.EvaluationContext{})
		if err != nil {
			log.Fatal(err)
		}

		tickerInterval, err := client.IntValue(context.Background(), "timer-interval", 1, openfeature.EvaluationContext{})
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
