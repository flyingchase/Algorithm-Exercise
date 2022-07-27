package exercises

import (
	"sync"
	"time"
)

var rwlock sync.RWMutex

func main() {
	rwlock.Lock()
	time.Sleep(10*time.Millisecond)
	rwlock.Unlock()


	rwlock.RLock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
}
