package bucket

import (
	"context"
	"log/slog"
)

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
				slog.Info("context canceled")
				return
			}
		}
	}()
}
