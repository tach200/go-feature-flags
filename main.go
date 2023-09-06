package main

import (
	"context"
	"log"

	"github.com/open-feature/go-sdk/pkg/openfeature"
)

func main() {
	openfeature.SetProvider(openfeature.NoopProvider{})

	client := openfeature.NewClient("app")

	v2Enabled, err := client.BooleanValue(context.Background(), "v2_enabled", true, openfeature.EvaluationContext{})
	if err != nil {
		log.Fatalf("error: couldn't query boolean flag: %v", err)
	}

	if v2Enabled {
		log.Print("v2 is enabled")
	}
}
