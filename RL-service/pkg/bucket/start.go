package bucket

import (
	"context"
	"log/slog"
)

//

/*
	В отдельной горутине каждый тик будет добавляться токен в бакет.

	Проблема!!!
	Горутины крутятся постоянно, и нужно как то красиво их завершать при остановке сервиса.
	Проще всего это сделать через отмену контекста, но я не придумал, как передать сюда такой контекст.
	Контекст, который везде используется - это контекст запроса и он отменяется, когда запрос выполнится,
	поэтому такой контекст нельзя сюда передавать.
*/

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
