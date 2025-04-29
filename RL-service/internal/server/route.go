package server

import (
	"log/slog"

	"rl-service/internal/handler/clientHandlers"
	"rl-service/internal/middleware"
	"rl-service/internal/repository/bucketRepository"
	"rl-service/internal/repository/clientRepository"
	"rl-service/internal/usecase/bucketUsecase"
	"rl-service/internal/usecase/clientUsecase"
	"rl-service/internal/usecase/limiterUsecase"
)

// в методе создаются слои, и регистрируются обработчики
func (s *Server) MapHandlers() error {
	clientRepo := clientRepository.NewClientRepository(s.pool)
	clientUC := clientUsecase.NewClientUC(s.cfg, clientRepo)
	clientHandler := clientHandlers.NewClientHandlers()

	bucketRepo := bucketRepository.NewBucketRepository()
	bucketUC := bucketUsecase.NewBucketUC(s.cfg, bucketRepo)
	err := s.RecoverBuckets(bucketUC)
	if err != nil {
		slog.Error("error recover buckets", "error", err)
		return err
	}

	limiterUC := limiterUsecase.NewLimiterUC(s.cfg, clientUC, bucketUC)

	mw := middleware.NewMiddlewareManager(s.cfg, limiterUC)

	mux := clientHandlers.MapClientHandlers(clientHandler, mw)

	s.srv.Handler = mux

	return nil
}
