package cache

import (
	"sync"
	"sync/atomic"
	"time"
)

var globalMap sync.Map
var globalMapLen int64

func SetTimeout(key string, data interface{}, timeout time.Duration) {
	globalMap.Store(key, data)
	atomic.AddInt64(&globalMapLen, 1)
	time.AfterFunc(timeout, func() {
		atomic.AddInt64(&globalMapLen, -1)
		globalMap.Delete(key)
	})
}

func Set(key string, data interface{}) {
	globalMap.Store(key, data)
	atomic.AddInt64(&globalMapLen, 1)
}

func Get(key string) (value interface{}, ok bool) {
	return globalMap.Load(key)
}

func Delete(key string) {
	atomic.AddInt64(&globalMapLen, -1)
	globalMap.Delete(key)
}

func Len() int {
	return int(globalMapLen)
}
