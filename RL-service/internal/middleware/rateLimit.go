package middleware

import (
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
		count := r.Header.Get("count")
		if count == "" {
			count = os.Getenv("DEFAULT_COUNT")
		}

		rate := r.Header.Get("rate")
		if rate == "" {
			rate = os.Getenv("DEFAULT_RATE")
		}

		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, errlst.ErrMsgInternalError, http.StatusInternalServerError)
			return
		}
		ip := host

		err = m.limiterUC.RateLimit(r.Context(), toRateLimitRequest(RateLimitRequest{
			IP:    ip,
			Count: count,
			Rate:  rate,
		}))
		if err != nil {
			http.Error(w, errlst.ErrMsgInternalError, http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
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
