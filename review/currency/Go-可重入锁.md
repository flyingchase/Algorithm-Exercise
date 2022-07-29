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