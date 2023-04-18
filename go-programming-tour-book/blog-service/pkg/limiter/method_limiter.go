package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

/*
第一个编写的简单限流器针对路由器限流，
*/

//针对LimiterIface接口 实现了限流器：MethosLimiter；
//
type MethosLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimiterIface {
	return MethosLimiter{
		Limiter: &Limiter{
			limiterBuckets: make(map[string]*ratelimit.Bucket), //建一个新的map
		},
	}
}

//限流器需要的三个方法：

func (l MethosLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?") //在URL中使用strings.Index函数查找第一个'?'的位置
	if index == -1 {
		//如果返回值是-1，代表URL中没有'?'，直接返回完整的URL作为关键字
		return uri
	}

	//如果返回值不是-1，代表URL中有参数，去掉参数后面的部分，返回剩余部分作为关键字
	return uri[:index]
}

func (l MethosLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limiterBuckets[key]
	return bucket, ok
}

func (l MethosLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterIface {
	for _, rule := range rules {
		if _, ok := l.limiterBuckets[rule.Key]; !ok {

			//ratelimit.NewBucketWithQuantum()
			//这个函数是用来创建一个令牌桶限流器的，其参数是规则(RateLimitRule)中定义的决定限流策略的三个属性：填充的时间间隔、令牌桶的容量和每次填充的令牌数。
			//每个时间段（fillInterval）填充的令牌数量为quantum，令牌桶的容量为capacity
			//这个函数的返回值是一个针对函数参数规则确定的Bucket类型的限流器实例，通过调用该实例的Take（）方法就可以对请求进行限流了。
			l.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)

		}
	}

	return l
}
