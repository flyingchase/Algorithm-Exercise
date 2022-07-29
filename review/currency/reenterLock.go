package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

/*
可重入锁：不断续期的锁
	不同服务器上的多个线程中，只有一个线程可以抢到一个锁来执行一个任务
	- Lock——>DoJob——>Unlock
锁一般不会放在某个进程中，会借用第三方存储，比如 Redis 来分布式锁
保证Unlock一定运行：
	为锁提供超时时间作为兜底：LockByExpire——>DoJob——>Unlock
出现异常情况时Unlock 没有运行，可以在 duration 时间锁后被释放，redis 中使用`sex ex`设定锁的超时设定
超时 Expire时间设定：
	比DoJob时间时间长，否则出现并发执行任务，无法准确得知 DoJob 时间
	DoJob设定超时时间，最多执行多久，让分布式锁比最多执行时长久一点即可——>TimeoutContext
	为锁设定较小的expire，不断续期锁即可，可视为不断重新开始加锁）
		除主线程外，可重入锁必须有另外一个线程/协程对锁进行续期，即为watchDog
*/

type DistributeLockRedis struct {
	key        string             // mutex key
	expire     int64              // mutex expireDuration
	status     bool               // mutex running flag
	cancelFunc context.CancelFunc //cancel renew
	redis      redis.Client       // redis句柄
}

func NewDistributeLockRedis(key string, expire int64) *DistributeLockRedis {
	return &DistributeLockRedis{
		key:    key,
		expire: expire,
	}
}

// tryLock 上锁
func (dl *DistributeLockRedis) TryLock() (err error) {

	if err == dl.lock(); err != nil {
		return err
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	dl.cancelFunc = cancelFunc
	dl.startWatchDog(ctx)
	dl.status = true
	return nil
}

// competition 竞争锁
func (dl *DistributeLockRedis) lock() error {
	if res, err := redis.NewStringCmd(dl.redis.Do(context.Background(), "SET", dl.key, 1, "NX", "EX", dl.expire)); err != nil {
		return nil
	}
	return nil
}

// guard 守护协程 自动续期

func (dl *DistributeLockRedis) startWatchDog(ctx context.Context) {
	safeGo(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			default:
				// 续租
				if dl.status {
					if res, err := redis.NewIntResult(dl.redis.Do(context.Background(), "EXPIRE", dl.key, dl.expire)); err != nil {
						return nil
					}
				}
				time.Sleep(time.Duration(dl.expire/2) * time.Second)
			}
		}
	})
}

// Unlock
func (dl *DistributeLockRedis) Unlock() (err error) {
	if dl.cancelFunc != nil {
		dl.cancelFunc()
	}
	var res int
	if dl.status {
		if res, err := redis.NewIntCmd(dl.redis.Do(context.Background(), "Del", dl.key)); err != nil {
			return fmt.Errorf("释放失败")
		}
		if res == 1 {
			dl.status = false
			return nil
		}
	}
	return fmt.Errorf("释放失败")
}
