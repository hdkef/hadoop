package service

type RoundRobinLoadBalancer interface {
	GetNextPtr(newSize int) (int, error)
}
