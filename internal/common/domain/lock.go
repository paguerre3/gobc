package domain

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/paguerre3/blockchain/configs"
)

var (
	config = configs.Instance()
)

// General Lock Service. TODO: use Redis instead of Cache MAP in case of multiple worker nodes instances.
type Lock interface {
	Acquire(key string, timeout time.Duration) error
	Release(key string)
}

type lockImpl struct {
	locks sync.Map
}

func NewLock() Lock {
	return &lockImpl{}
}

func (l *lockImpl) Acquire(key string, timeout time.Duration) error {
	if len(key) == 0 {
		return errors.New("idempotencyKey is missing")
	}
	deadline := time.Now().Add(timeout)
	backoff := time.Duration(config.Lock().InitialBackoffInMillis()) * time.Millisecond
	for {
		if _, loaded := l.locks.LoadOrStore(key, struct{}{}); !loaded {
			return nil // Successfully acquired lock
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("failed to acquire lock for key %s within timeout", key)
		}
		time.Sleep(backoff)
		backoff *= time.Duration(config.Lock().BackoffMultiplier())

		// Avoid too long sleeps, i.g. 100*2*3*4=2.4 secs => max = timeout secs/2,
		// ie. 1.5 secs max backoff until deadline:
		if backoff > timeout/2 {
			backoff = timeout / 2
		}
	}
}

func (l *lockImpl) Release(key string) {
	l.locks.Delete(key)
}
