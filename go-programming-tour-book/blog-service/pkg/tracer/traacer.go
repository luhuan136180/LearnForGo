package tracer

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

// 定义一个函数：创建并返回 Jaeger Tracer 实例
// serviceName：服务名称，agentHostPort：Jaeger Agent 的地址和端口（格式为“host:port”）
func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	// 初始化 Jaeger 配置，使用常量采样器，并将日志记录开启
	//config.Configuration ： jaeger client 配置项，设置应用的基本信息
	cfg := &config.Configuration{
		ServiceName: serviceName, // 服务名称
		Sampler: &config.SamplerConfig{ //采样配置
			Type:  "const", //采样类型
			Param: 1,       //参数
		},
		Reporter: &config.ReporterConfig{ // 报告配置
			LogSpans:            true,            // 记录所有采样的 span
			BufferFlushInterval: 1 * time.Second, // 报告刷新缓存的时间间隔
			LocalAgentHostPort:  agentHostPort,   // Agent 的地址和端口
		},
	}

	// 创建 Jaeger Tracer 实例
	//fg.NewTracer(): 根据配置初始化tracer，返回 opentracing.Tracer 类型
	tracer, closer, err := cfg.NewTracer() // 通过配置创建一个新的 Jaeger Tracer 实例
	if err != nil {
		return nil, nil, err
	}

	opentracing.SetGlobalTracer(tracer) // 设置全局变量，将该 Tracer 实例设置为全局可用的
	return tracer, closer, nil
}
