package bucket

import (
	"context"
	"time"
)

func (tb *TokenBucket) Wait(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			tb.mutex.Lock()
			if tb.currentTokens > 0 {
				tb.currentTokens-- // Если токен есть, забираем его
				tb.mutex.Unlock()
				return nil
			}
			tb.mutex.Unlock()
			time.Sleep(10 * time.Millisecond) // Пауза перед новой попыткой
		}
	}
}
