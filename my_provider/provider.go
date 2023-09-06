package provider

import (
	"context"
	"log"
	"os"

	"github.com/open-feature/go-sdk/pkg/openfeature"
	"gopkg.in/yaml.v3"
)

type FeatureProvider struct{}

func lookupYamlVal() (map[interface{}]interface{}, error) {
	featuresYmlFilePath := "my_provider/flags.yml"

	data, err := os.ReadFile(featuresYmlFilePath)
	if err != nil {
		log.Fatal(err)
	}

	featureMap := make(map[interface{}]interface{})

	err = yaml.Unmarshal(data, &featureMap)
	if err != nil {
		return nil, err
	}

	return featureMap, nil
}

func (fp FeatureProvider) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool, evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {

	featureMap, err := lookupYamlVal()
	if err != nil {
		return openfeature.BoolResolutionDetail{
			Value: defaultValue,
		}
	}

	boolValue := featureMap[flag].(bool)

	if flag == "feature-z" {
		//check for any blacklisted values
		blacklisted := convertToStrings(featureMap["feature-z-blacklist"].([]interface{}))

		if lookupValue(blacklisted, evalCtx["Location"].(string)) {
			boolValue = false
		}
	}

	return openfeature.BoolResolutionDetail{Value: boolValue}
}

func (fp FeatureProvider) IntEvaluation(ctx context.Context, flag string, defaultValue int64, evalCtx openfeature.FlattenedContext) openfeature.IntResolutionDetail {

	featureMap, err := lookupYamlVal()
	if err != nil {
		return openfeature.IntResolutionDetail{
			Value: defaultValue,
		}
	}

	// weird conversion here
	intValue := featureMap[flag].(int)
	int64Value := int64(intValue)

	return openfeature.IntResolutionDetail{Value: int64Value}
}

func (fp FeatureProvider) FloatEvaluation(ctx context.Context, flag string, defaultValue float64, evalCtx openfeature.FlattenedContext) openfeature.FloatResolutionDetail {
	return openfeature.FloatResolutionDetail{}
}

func (fp FeatureProvider) Hooks() []openfeature.Hook {
	return []openfeature.Hook{}
}

func (fp FeatureProvider) Metadata() openfeature.Metadata {
	return openfeature.Metadata{}
}

func (fp FeatureProvider) ObjectEvaluation(context.Context, string, interface{}, openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {
	return openfeature.InterfaceResolutionDetail{}
}

func (fp FeatureProvider) StringEvaluation(context.Context, string, string, openfeature.FlattenedContext) openfeature.StringResolutionDetail {
	return openfeature.StringResolutionDetail{}
}

func lookupValue(list []string, target string) bool {
	for _, value := range list {
		if value == target {
			return true
		}
	}
	return false
}

func convertToStrings(values []interface{}) []string {
	var result []string
	for _, v := range values {
		if str, ok := v.(string); ok {
			result = append(result, str)
		}
	}
	return result
}
