package middleware

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"

	"rl-service/internal/usecase/limiterUsecase"
	"rl-service/pkg/errlst"
)

/*
	Главный мидлвар, который реализует логику Rate Limiter.
	Мидлвар не пропускает запрос дальше, пока не выполнятся необходимые условия.
	Логика этого слоя (handler):
	1) Клиент отправляет запрос с query параметрами (rate и count), которые определяют настройки для этого клиента
	2) Если параметры придут пустыми, то настройки заполняются значениями по умолчанию,
	которые задаются в env-файле.
	3) дальнейшая работа передаётся в слой usecase
	4) по ответу из usecase определяем дальнейшие действия для запроса клиента
*/

type RateLimitRequest struct {
	IP    string
	Count string
	Rate  string
}

type RateLimitResponse struct {
	IP    string
	Count uint64
	Rate  uint64
}

func (m *MDWManager) RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count := r.URL.Query().Get("count")
		if count == "" {
			count = os.Getenv("DEFAULT_COUNT")
		}

		rate := r.URL.Query().Get("rate")
		if rate == "" {
			rate = os.Getenv("DEFAULT_RATE")
		}

		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			slog.Error("error parse remote addr", "error", err, "remoteAddr", r.RemoteAddr)
			http.Error(w, errlst.ErrMsgInternalError, http.StatusInternalServerError)
			return
		}
		ip := host

		// переходим в usecase
		err = m.limiterUC.RateLimit(r.Context(), toRateLimitRequest(RateLimitRequest{
			IP:    ip,
			Count: count,
			Rate:  rate,
		}))
		if err != nil {
			switch {
			case errors.Is(err, errlst.ErrBucketExist):
				http.Error(w, fmt.Sprintf("\"error\": %s; \"ip\": %s", errlst.ErrMsgBucketExist, ip),
					http.StatusInternalServerError,
				)

			case errors.Is(err, errlst.ErrBucketNotExist):
				http.Error(w, fmt.Sprintf("\"error\": %s; \"ip\": %s", errlst.ErrMsgBucketNotExist, ip),
					http.StatusNotFound,
				)

			case errors.Is(err, errlst.ErrCountOrRateNoValid):
				http.Error(w, fmt.Sprintf("\"error\": %s; \"ip\": %s", errlst.ErrMsgCountOrRateNoValid, ip),
					http.StatusBadRequest,
				)

			default:
				http.Error(w, fmt.Sprintf("\"error\": %s; \"ip\": %s", errlst.ErrMsgInternalError, ip),
					http.StatusInternalServerError,
				)
			}
			slog.Error("error rate limiter", "error", err, "IP", ip)
			return
		}

		slog.Info("request redirect success", "IP", ip)

		next.ServeHTTP(w, r) // запрос идёт дальше
	})
}

func toRateLimitRequest(req RateLimitRequest,
) limiterUsecase.RateLimitRequest {
	return limiterUsecase.RateLimitRequest{
		IP:    req.IP,
		Count: req.Count,
		Rate:  req.Rate,
	}
}
