package hooks

import (
	"context"

	"github.com/open-feature/go-sdk/pkg/openfeature"
)

type Hooks struct{}

func (h Hooks) Before(ctx context.Context, hookContext openfeature.HookContext, hookHints openfeature.HookHints) (*openfeature.EvaluationContext, error) {
	// code to run before flag evaluation

	return nil, nil
}

func (h Hooks) After(ctx context.Context, hookContext openfeature.HookContext, flagEvaluationDetails openfeature.EvaluationDetails, hookHints openfeature.HookHints) error {
	// code to run after successful flag evaluation
	return nil
}

func (h Hooks) Error(ctx context.Context, hookContext openfeature.HookContext, err error, hookHints openfeature.HookHints) {
	// code to run if there's an error during before Hooks or during flag evaluation
}

func (h Hooks) Finally(ctx context.Context, hookContext openfeature.HookContext, hookHints openfeature.HookHints) {
	// code to run after all other stages, regardless of success/failure
}
