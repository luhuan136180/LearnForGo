package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

//定义本项目限流器需要的方法
type LimiterIface interface {
	//获取对应的限流器的键值对名称
	Key(c *gin.Context) string
	//获取令牌桶
	GetBucket(key string) (*ratelimit.Bucket, bool)
	//新增多个令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

//存储令牌桶与键值对名称的映射关系，
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

//用于存储令牌桶的一些相应规则属性
type LimiterBucketRule struct {
	//自定义键值对名称
	Key string
	//间隔对酒时间放N个令牌
	FillInterval time.Duration
	//令牌桶的容量
	Capacity int64
	//每次到达间隔时间锁放的具体令牌数量
	Quantum int64
}
