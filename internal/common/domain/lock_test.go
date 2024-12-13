package domain

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLock_ConcurrentAccess(t *testing.T) {
	lock := NewLock()
	key := "test-key"
	timeout := time.Duration(config.Lock().TimeOutInSeconds()) * time.Second
	//timeout := 3 * time.Second

	var wg sync.WaitGroup
	var mus, mue sync.Mutex
	successfulAcquires := 0
	errorsCount := 0

	// Simulate 10 goroutines trying to acquire the same lock
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := lock.Acquire(key, timeout)
			if err != nil {
				if err.Error() == fmt.Sprintf("failed to acquire lock for key %s within timeout", key) {
					mue.Lock()
					errorsCount++
					mue.Unlock()
				} else {
					t.Errorf("Unexpected error: %v", err)
				}
				return
			}
			mus.Lock()
			successfulAcquires++
			mus.Unlock()

			// TASK WORK SIMULATION:
			// simulate a litle more than TO work
			// for acquiring the lock otherwise success count will be greater than the expected.
			taskSimulation := time.Duration(config.Lock().TimeOutInSeconds()+2) * time.Second
			time.Sleep(taskSimulation)
			lock.Release(key)
		}()
	}

	wg.Wait()

	// Assert only one goroutine successfully acquires the lock
	assert.Equal(t, 1, successfulAcquires, "Expected 1 successful lock acquisition, but got %d", successfulAcquires)

	// Assert the remaining goroutines failed due to timeout
	assert.Equal(t, 9, errorsCount, "Expected 9 errors due to timeout, but got %d", errorsCount)
}
