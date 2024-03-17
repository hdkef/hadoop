package impl

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	consul "github.com/hashicorp/consul/api"
	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	pkgSvc "github.com/hdkef/hadoop/pkg/services"
)

const (
	SERVICE_REGISTRY_HOST = "SERVICE_REGISTRY_HOST"
	SERVICE_REGISTRY_PORT = "SERVICE_REGISTRY_PORT"
	HEALTH_CHECK_INTERVAL = "HEALTH_CHECK_INTERVAL"
)

type ServiceRegistryConfig struct {
	Host           string
	Port           uint32
	HealthInterval time.Duration
}

type ServiceRegistry struct {
	c   *consul.Client
	cfg *ServiceRegistryConfig
}

// GetAll implements service.ServiceRegistry.
func (s *ServiceRegistry) GetAll(ctx context.Context, servicesName string, tag string) ([]*pkgEt.ServiceDiscovery, error) {
	passingOnly := true
	addrs, _, err := s.c.Health().Service(servicesName, tag, passingOnly, nil)
	if len(addrs) == 0 && err == nil {
		return nil, fmt.Errorf("service ( %s ) was not found", servicesName)
	}
	if err != nil {
		return nil, err
	}

	svd := []*pkgEt.ServiceDiscovery{}

	for _, v := range addrs {

		newEntry := &pkgEt.ServiceDiscovery{}
		newEntry.SetAddress(v.Node.Address)
		newEntry.SetID(v.Node.ID)
		newEntry.SetPort(uint32(v.Service.Port))
		newEntry.SetServices(v.Service.Service)

		svd = append(svd, newEntry)
	}

	return svd, nil
}

func (s *ServiceRegistry) RegisterNode(id string, serviceName string, grpcport int, address string) {

	registeration := &consul.AgentServiceRegistration{
		ID:      id,
		Name:    serviceName,
		Port:    grpcport,
		Address: address,
		Check: &consul.AgentServiceCheck{
			GRPC:     fmt.Sprintf("%s:%d", address, grpcport),
			Interval: s.cfg.HealthInterval.String(),
			Timeout:  "30s",
		},
	}

	regiErr := s.c.Agent().ServiceRegister(registeration)

	// retry mechanism

	for regiErr != nil {
		log.Printf("Failed to register service: %s:%v %s", address, grpcport, regiErr.Error())
		time.Sleep(5 * time.Second)
		regiErr = s.c.Agent().ServiceRegister(registeration)
	}

	log.Printf("successfully register service: %s:%v", address, grpcport)
}

func NewServiceRegistryConfig() *ServiceRegistryConfig {

	svcRegHost := os.Getenv(SERVICE_REGISTRY_HOST)

	if svcRegHost == "" {
		svcRegHost = "http://consul"
	}

	svcRegPort := os.Getenv(SERVICE_REGISTRY_PORT)

	if svcRegPort == "" {
		svcRegPort = "8500"
	}

	svcRegPortVal, err := strconv.Atoi(svcRegPort)
	if err != nil {
		panic(fmt.Sprintf("%s %s", SERVICE_REGISTRY_PORT, err.Error()))
	}

	healthInterval := os.Getenv(HEALTH_CHECK_INTERVAL)

	if healthInterval == "" {
		healthInterval = "2s"
	}

	healthIntervalVal, err := time.ParseDuration(healthInterval)
	if err != nil {
		panic(fmt.Sprintf("%s %s", HEALTH_CHECK_INTERVAL, err.Error()))
	}

	return &ServiceRegistryConfig{
		Host:           svcRegHost,
		Port:           uint32(svcRegPortVal),
		HealthInterval: healthIntervalVal,
	}
}

func NewServiceRegistry(cfg *ServiceRegistryConfig) pkgSvc.ServiceRegistry {

	if cfg == nil {
		panic("config is nil")
	}

	config := consul.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	c, err := consul.NewClient(config)
	if err != nil {
		panic(err)
	}

	return &ServiceRegistry{
		c:   c,
		cfg: cfg,
	}
}
