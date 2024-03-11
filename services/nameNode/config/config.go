package config

import "time"

type Config struct {
	ReplicationTarget uint32
	BlockSplitTarget  uint32
	MinLeaseTime      time.Duration
}
