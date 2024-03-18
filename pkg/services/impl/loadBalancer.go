package impl

import (
	"errors"
	"sync"

	pkgSvc "github.com/hdkef/hadoop/pkg/services"
)

type LoadBalancer struct {
	Mtx          *sync.Mutex
	CurrentPtr   int
	CurrentTotal int
}

// GetNextPtr implements service.RoundRobinLoadBalancer.
func (l *LoadBalancer) GetNextPtr(newTotal int) (int, error) {
	l.Mtx.Lock()
	defer l.Mtx.Unlock()

	if newTotal <= 0 {
		return 0, errors.New("no targets")
	}

	l.CurrentTotal = newTotal
	l.CurrentPtr++

	if l.CurrentPtr >= l.CurrentTotal {
		l.CurrentPtr = 0
	}

	return l.CurrentPtr, nil
}

func NewLoadBalancer(mtx *sync.Mutex) pkgSvc.RoundRobinLoadBalancer {
	return &LoadBalancer{
		Mtx: mtx,
	}
}
