package memorylimiterprocessor

import (
	"context"
	"runtime"
	"sync/atomic"
	"time"
)

// ... existing imports and struct definitions ...

func (ml *memoryLimiter) monitorMemory(ctx context.Context) {
	ticker := time.NewTicker(ml.checkInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			var mem runtime.MemStats
			runtime.ReadMemStats(&mem)
			currentUsage := mem.Alloc

			// Check if we are above the limit
			if currentUsage > ml.limit {
				atomic.StoreInt32(&ml.refusing, 1)
			} else {
				// Reset refusal state immediately when memory is safe
				atomic.StoreInt32(&ml.refusing, 0)
			}
		case <-ctx.Done():
			return
		}
	}
}