package serverpool

import "sync/atomic"

// метод атомарно увеличивает индекс текущего сервера в пуле, и возвращает этот индекс
func (s *ServerPool) NextIndex() int {
	return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.listOfSrvs)))
}
