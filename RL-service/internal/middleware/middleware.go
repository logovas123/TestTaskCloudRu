package middleware

import (
	"context"

	"rl-service/internal/config"
	"rl-service/internal/usecase/limiterUsecase"
)

type LimiterUC interface {
	RateLimit(
		ctx context.Context,
		req limiterUsecase.RateLimitRequest,
	) error
}

type MDWManager struct {
	cfg       *config.Config
	limiterUC LimiterUC
}

func NewMiddlewareManager(
	cfg *config.Config,
	limiterUC LimiterUC,
) *MDWManager {
	return &MDWManager{
		cfg:       cfg,
		limiterUC: limiterUC,
	}
}
