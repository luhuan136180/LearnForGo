package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"runtime"
	"time"
)

/*
我们先预定义了应用日志的 Level 和 Fields 的具体类型，
并且分为了 Debug、Info、Warn、Error、Fatal、Panic 六个日志等级，
便于在不同的使用场景中记录不同级别的日志。
*/
type Level int8 //自定义日至登记

type Fields map[string]interface{}

//设置常量
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

//重写string函数
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	default:
		return ""
	}

}

type Logger struct {
	newLogger *log.Logger //原生log
	ctx       context.Context
	fields    Fields
	callers   []string
}

//w:表示写入的地点
func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag) //
	return &Logger{newLogger: l}
}

//克隆
func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

//WithFields：设置日志公共字段。
func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}
	for k, v := range f {
		ll.fields[k] = v
	}
	return ll
}

//WithContext：设置日志上下文属性。
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

//WithCaller：设置当前某一层调用栈的信息（程序计数器、文件信息、行号）。
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	//Caller:报告当前go程调用栈所执行的函数的文件和行号信息。
	//实参skip为上溯的栈帧数，0表示Caller的调用者（Caller所在的调用栈）。
	//（由于历史原因，skip的意思在Caller和Callers中并不相同。）函数的返回值为调用栈标识符、文件名、
	//该调用在文件中的行号。如果无法获得信息，ok会被设为false。
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}

	}
	return ll
}

//WithCallersFrames：设置当前的整个调用栈信息。
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	//
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callers = append(callers, fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	ll := l.clone()
	ll.callers = callers
	return ll
}

//添加实现了jaeger，数据追踪的
func (l *Logger) WithTrace() *Logger {
	ginCTX, ok := l.ctx.(*gin.Context)
	if ok {
		return l.WithFields(Fields{
			"trace_id": ginCTX.MustGet("X-Trace-ID"),
			"span_id":  ginCTX.MustGet("X-Span-ID"),
		})
	}
	return l
}

//编写日志内容的格式化和日志输出动作的相关方法
//json格式
func (l *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	//额外四项信息
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	//将logger结构体的fields属性中的内容的数据转化
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok { //不存在的话加入
				data[k] = v
			}
		}
	}
	return data
}

//输出控制
func (l *Logger) Output(level Level, message string) {
	//转化成json
	body, _ := json.Marshal(l.JSONFormat(level, message))
	content := string(body)
	switch level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

//根据先前定义的日志分级，编写对应的日志输出的外部方法，继续写入如下代码：
//info
func (l *Logger) Info(v ...interface{}) {

	l.Output(LevelInfo, fmt.Sprint(v...))
}
func (l *Logger) Infof(format string, v ...interface{}) {
	//fmt.Sprintf 返回生成的格式化字符串
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

//fatal
func (l *Logger) Fatal(v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprint(v...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprintf(format, v...))
}

//debug
func (l *Logger) Debug(v ...interface{}) {
	l.Output(LevelDebug, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Output(LevelDebug, fmt.Sprintf(format, v...))
}

//Warn
func (l *Logger) Warn(v ...interface{}) {
	l.Output(LevelWarn, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Output(LevelWarn, fmt.Sprintf(format, v...))
}

//Error
func (l *Logger) Error(v ...interface{}) {
	l.Output(LevelError, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Output(LevelError, fmt.Sprintf(format, v...))
}

//Panic
func (l *Logger) Panic(v ...interface{}) {
	l.Output(LevelPanic, fmt.Sprint(v...))
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.Output(LevelPanic, fmt.Sprintf(format, v...))
}
