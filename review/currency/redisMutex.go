package currency

import (
	"fmt"
	"os"
	"time"

	redis "code.byted.org/kv/goredis"
)

type RedisMutex struct {
	client      *redis.Client
	key         string
	value       string
	timeout     time.Duration
	tryInterval time.Duration
	maxRetries  int
	locked      bool
}

func NewRedisMutex(client *redis.Client, key string, timeout int, tryInterval int, maxRetries int) *RedisMutex {
	var redisMutex = new(RedisMutex)
	redisMutex.client = client
	redisMutex.key = "mutex:" + key

	hostName, err := os.Hostname()
	if err != nil {
		hostName = "unknown"
	}
	pid := os.Getpid()
	redisMutex.value = fmt.Sprintf("%v|%v", hostName, pid)

	redisMutex.timeout = time.Duration(timeout) * time.Second
	redisMutex.tryInterval = time.Duration(tryInterval) * time.Millisecond
	redisMutex.maxRetries = maxRetries
	redisMutex.locked = false

	return redisMutex
}

func (m *RedisMutex) TouchLock() bool {
	m.locked, _ = m.client.SetNX(m.key, m.value, m.timeout).Result()
	return m.locked
}

func (m *RedisMutex) Lock() bool {
	retry := 0
	for !m.locked {
		m.locked, _ = m.client.SetNX(m.key, m.value, m.timeout).Result()
		if m.locked {
			break
		}

		retry += 1
		if retry > m.maxRetries {
			// 查看是否设置了过期时间，未设置则进行del
			ttl, _ := m.client.TTL(m.key).Result()
			if ttl.String() == "-1s" {
				m.client.Del(m.key)
			}
			break
		}

		time.Sleep(m.tryInterval)
	}

	return m.locked
}

func (m *RedisMutex) Unlock() {
	if !m.locked {
		return
	}
	m.client.Del(m.key)
}

func (m *RedisMutex) SafeRelease() error {
	luaScript := "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end"
	return m.client.Eval(luaScript, []string{m.key}, m.value).Err()
}
