package logman

import (
	"bytes"
	"runtime"
	"strings"
	"time"
)

type Entry struct {
	logger *logger
	Buffer *bytes.Buffer
	Map    map[string]interface{}
	Level  Level
	Time   time.Time
	File   string
	Line   int
	Func   string
	Format string
	Args   []interface{}
}

func entry(l *logger) *Entry {
	return &Entry{logger: l, Buffer: new(bytes.Buffer), Map: make(map[string]interface{}, 5)}
}

func (e *Entry) write(level Level, format string, args ...interface{}) {
	//传入的日志级别小于参数设置的级别，直接返回
	if e.logger.opt.level > level {
		return
	}
	e.Time = time.Now()
	e.Level = level
	e.Format = format
	e.Args = args
	//开启文件和行号输出
	if !e.logger.opt.disableCaller {
		if pc, file, line, ok := runtime.Caller(2); !ok {
			e.File = "???"
			e.Func = "???"
		} else {
			e.File, e.Line, e.Func = file, line, runtime.FuncForPC(pc).Name()
			e.Func = e.Func[strings.LastIndex(e.Func, "/")+1:]
		}
	}
	e.format()
	e.writer()
	e.release()
}

//日志格式化
func (e *Entry) format() {
	_ = e.logger.opt.formatter.Format(e)
}

//日志写入
func (e *Entry) writer() {
	e.logger.mu.Lock()
	_, _ = e.logger.opt.output.Write(e.Buffer.Bytes())
	e.logger.mu.Unlock()
}

//释放entry
func (e *Entry) release() {
	//属性重置
	e.Args, e.Line, e.File, e.Format, e.Func = nil, 0, "", "", ""
	//重置缓冲区
	e.Buffer.Reset()
	//放回临时对象池
	e.logger.entryPool.Put(e)
}
