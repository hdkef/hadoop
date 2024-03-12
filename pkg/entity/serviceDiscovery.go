package entity

type ServiceDiscovery struct {
	services string
	id       string
	address  string
	port     uint32
}

// Setter method for services field
func (s *ServiceDiscovery) SetServices(services string) {
	s.services = services
}

// Getter method for services field
func (s *ServiceDiscovery) GetServices() string {
	return s.services
}

// Setter method for id field
func (s *ServiceDiscovery) SetID(id string) {
	s.id = id
}

// Getter method for id field
func (s *ServiceDiscovery) GetID() string {
	return s.id
}

// Setter method for address field
func (s *ServiceDiscovery) SetAddress(address string) {
	s.address = address
}

// Getter method for address field
func (s *ServiceDiscovery) GetAddress() string {
	return s.address
}

// Setter method for port field
func (s *ServiceDiscovery) SetPort(port uint32) {
	s.port = port
}

// Getter method for port field
func (s *ServiceDiscovery) GetPort() uint32 {
	return s.port
}
