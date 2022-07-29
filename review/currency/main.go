package main

import (
	"context"
	"fmt"
	"time"
	// "code.byted.org/gopkg/singleflight"
)

// const (
// 	funcKey = "key"
// 	times   = 5
// 	randNum = 100
// )
//
// var g singleflight.Group
//
// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }
//
// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(times)
// 	for i := 0; i < times; i++ {
// 		go func() {
// 			defer wg.Done()
// 			num, err := run(funcKey)
// 			if err != nil {
// 				log.Fatal(err)
// 				return
// 			}
// 			log.Println(num)
// 		}()
// 	}
// 	wg.Wait()
// }
// func run(key string) (num int, err error) {
// 	v, err, isShare := g.Do(key, func() (interface{}, error) {
// 		// time.Sleep(time.Second * 5)
// 		num = rand.Intn(randNum)
// 		return num, nil
// 	})
//
// 	if err != nil {
// 		log.Fatal(err)
// 		return 0, err
// 	}
// 	data := v.(int)
// 	log.Println(isShare)
// 	return data, nil
//
// }

// 构造函数、struct封装对象再 new 对象
type Cluster struct {
	opts options
}
type options struct {
	connectionTimeout time.Duration
	readTimeout       time.Duration
	writeTimeout      time.Duration
	logError          func(ctx context.Context, err error)
}

// Golang中函数指针
type Option func(c *options)

// 闭包 设置参数的具体实现
func LogError(f func(ctx context.Context, err error)) Option {
	return func(c *options) {
		c.logError = f
	}
}

// 关键数据变量赋值使用方法来实现
// 而非直接赋值
func ConnectionTimeout(d time.Duration) Option {
	return func(c *options) {
		c.connectionTimeout = d
	}
}
func WriteTimeout(d time.Duration) Option {
	return func(c *options) {
		c.writeTimeout = d
	}
}

func ReadTimeout(d time.Duration) Option {
	return func(c *options) {
		c.readTimeout = d
	}
}

// 函数具体实现，多参
func NewCluster(opts ...Option) *Cluster {
	clusterOpts := options{}
	for _, opt := range opts {
		// 函数指针赋值调用
		opt(&clusterOpts)
	}

	cluster := new(Cluster)
	cluster.opts = clusterOpts
	return cluster
}

func main() {
	// 初始化参数setting
	commonOpts := []Option{
		ConnectionTimeout(time.Second * 1),
		ReadTimeout(2 * time.Second),
		WriteTimeout(3 * time.Second),
		LogError(func(ctx context.Context, err error) {
		}),
	}

	// 构造函数
	cluster := NewCluster(commonOpts...)

	//测试
	fmt.Println(cluster.opts.connectionTimeout)
	fmt.Println(cluster.opts.writeTimeout)
}

// 组合与继承
type MsgModel struct {
	msgId   int
	msgTyep int
}

// 成员方法，设置msgId
func (msg *MsgModel) SetId(msgId int) {
	msg.msgId = msgId
}
