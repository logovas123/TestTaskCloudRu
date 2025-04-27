package srvnode

// возвращает статус сервера
func (b *SrvNode) IsAlive() bool {
	b.mux.RLock()
	defer b.mux.RUnlock()

	return b.Alive
}
