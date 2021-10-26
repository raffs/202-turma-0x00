package store

import (
	"sync"
)

type Service struct {
	Name         string
	ContainerId  string
	Image        string
	IPAddress    string
	Endpoint     string
	Hostname     string
	Labels       map[string]string
	Ports        []string
}

type ServiceStore struct {
	services map[string]Service

	sync.RWMutex // disponibilizar funcoes para locks/unlocks tanto de leitura e escrita.	
}

func (s *ServiceStore) Add(nome string, serv Service) {
	s.Lock()
	s.services[nome] = serv
	s.Unlock()
}

func (s *ServiceStore) Rem(nome string) {
	s.Lock()
	delete(s.services, nome)
	s.Unlock()
}

func (s *ServiceStore) List() (services []Service) {
	s.RLock()
	defer s.RUnlock()

	// var services []Services
	for _, s := range s.services {
		services = append(services, s)
	}

	return services
}

func New() ServiceStore {
	return ServiceStore{
		services: make(map[string]Service),
	}
}
