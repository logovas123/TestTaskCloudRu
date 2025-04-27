package bucket

import (
	"context"
	"time"
)

// метод удаляет токен из бакета
// если в бакете нет токенов, то идёт sleep, и снова повторяется попытка удаления
func (tb *TokenBucket) Wait(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			tb.mutex.Lock()
			if tb.currentTokens > 0 {
				tb.currentTokens--
				tb.mutex.Unlock()
				return nil
			}
			tb.mutex.Unlock()
			time.Sleep(10 * time.Millisecond)
		}
	}
}
