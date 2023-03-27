package pubsub

import (
	"sync"
	"time"
)

type (
	subscriber chan interface{}         // 订阅者为一个管道
	topicFunc  func(v interface{}) bool //主题为一个过滤器
)

//发布者对象
type Publisher struct {
	m           sync.RWMutex             //读写锁
	buffer      int                      //订阅队列的缓存大小
	timeout     time.Duration            //发布超时时间
	subscribers map[subscriber]topicFunc //订阅者信息：k：订阅者是谁；V：订阅者订阅的主题
}

//构造一个发布者对象，也是设置发布超市时间和缓存队列的长度
func NewPublisher(publisherTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publisherTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

//用于添加一个新的该发布者的订阅者，且订阅该发布者的所有主题咨询
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeToptic(nil)
}

//用于添加一个该发布者的订阅者信息，按照过滤器帅选后的主题订阅
func (p *Publisher) SubscribeToptic(toptic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer) //创建订阅者信息（订阅通道）
	p.m.Lock()                             //上锁
	p.subscribers[ch] = toptic
	p.m.Unlock()
	return ch
}

//退出订阅(输入需要退订的订阅者信息——通道)
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)

}

//发布一个主题
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendToptic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// 关闭发布者对象，同时关闭所有的订阅者管道。(删除该发布者)
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

//发送主题
func (p *Publisher) sendToptic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done() //
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):

	}
}
