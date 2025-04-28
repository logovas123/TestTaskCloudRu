package bucket

import (
	"context"
	"log/slog"
)

// в отдельной горутине каждый тик будет добавляться токен в бакет
func (tb *TokenBucket) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-tb.ticker.C:
				tb.mutex.Lock()
				if tb.currentTokens < tb.maxTokens {
					tb.currentTokens++
				}
				tb.mutex.Unlock()
			case <-ctx.Done():
				tb.ticker.Stop()
				slog.Info(ctx.Err().Error())
				return
			}
		}
	}()
}
