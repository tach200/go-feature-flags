package main

import (
	"context"
	"log"
	"net/http"
	"time"

	gofeatureflag "github.com/open-feature/go-sdk-contrib/providers/go-feature-flag/pkg"
	"github.com/open-feature/go-sdk/pkg/openfeature"
)

func main() {

	opts := gofeatureflag.ProviderOptions{
		Endpoint: "http://0.0.0.0:1031",
		HTTPClient: &http.Client{
			Timeout: 3 * time.Second,
		},
	}

	provider, err := gofeatureflag.NewProviderWithContext(context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	}

	err = openfeature.SetProvider(provider)
	if err != nil {
		log.Fatal(err)
	}

	client := openfeature.NewClient("app")

	evalCtx := openfeature.NewEvaluationContext(
		"metadata", map[string]interface{}{
			"user-agent": "OLD-CLI",
			"location":   "UK",
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

		z, err := client.BooleanValue(context.Background(), "feature-z", false, evalCtx)
		if err != nil {
			log.Fatal(err)
		}

		ticker, err := client.IntValue(context.Background(), "ticker", 1, evalCtx)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second * time.Duration(ticker))

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
