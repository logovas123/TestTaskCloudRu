package serverpool

import (
	"sync/atomic"
)

// Метод проходится по списку серверов, ищет живой сервер, сохраняет его текущим и возвращает его
func (s *ServerPool) GetNextActiveConn() ServerNodeHandler {
	next := s.NextIndex()
	l := len(s.listOfSrvs) + next
	for i := next; i < l; i++ {
		idx := i % len(s.listOfSrvs)
		if s.listOfSrvs[idx].IsAlive() {
			if i != next {
				atomic.StoreUint64(&s.current, uint64(idx))
			}
			return s.listOfSrvs[idx]
		}
	}
	return nil
}
