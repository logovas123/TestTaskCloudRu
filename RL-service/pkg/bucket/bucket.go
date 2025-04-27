package bucket

import (
	"sync"
	"time"
)

type TokenBucket struct {
	rate          uint64
	maxTokens     uint64
	currentTokens uint64
	ticker        *time.Ticker
	mutex         sync.Mutex
}

// метод создаёт новый бакет
// MaxTokens - это размер бакета
// Rate - скорость добавления токенов, показывает колиство раз добавления в секунду
func NewTokenBucket(MaxTokens uint64, Rate uint64) *TokenBucket {
	everyMs := 1 / float64(Rate) * 1000

	return &TokenBucket{
		rate:          Rate,
		maxTokens:     MaxTokens,
		ticker:        time.NewTicker(time.Duration(int64(everyMs) * int64(time.Millisecond))),
		currentTokens: MaxTokens,
		mutex:         sync.Mutex{},
	}
}
