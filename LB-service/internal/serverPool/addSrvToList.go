package serverpool

// Добавляем сервер в список серверов(пул)
func (s *ServerPool) AddSrvToList(srv ServerNodeHandler) {
	s.listOfSrvs = append(s.listOfSrvs, srv)
}
