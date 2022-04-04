package snowflake

import (
	"sync"
)

type Snowflake struct {
	sync.Mutex
	// 时间戳 毫秒
	timestamp int64
	// 工作结点
	workerid int64
	// 数据中心机房 id
	datacenterid int64
	// 序列号
	sequence int64
}
const (
	epoch = int64(1577808000000)
	timestampBits =uint(41)
	datacenteridBits  = uint(2)
	workeridBits = 
)
