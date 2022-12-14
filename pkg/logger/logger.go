package logger

import (
	"io"
	"log"
	"os"
	"sync/atomic"
)

// log level
const (
	LDEBUG = iota + 1 // 1
	LWARN             // 2
	LINFO             // 3
	LERROR            // 4
	LFATAL            // 5
)

var Logger *MyLogger

type MyLogger struct {
	allwriter      io.Writer
	allwriterlevel int64
	stdlevel       int64
	debugLogger    *log.Logger
	warnLogger     *log.Logger
	infoLogger     *log.Logger
	errLogger      *log.Logger
	fatalLogger    *log.Logger
}

func init() {
	Logger = &MyLogger{
		debugLogger: log.New(os.Stdout, "[DEBUG]", log.LstdFlags|log.Lmsgprefix),
		warnLogger:  log.New(os.Stdout, "[WARN]", log.LstdFlags|log.Lmsgprefix),
		infoLogger:  log.New(os.Stdout, "[INFO]", log.LstdFlags|log.Lmsgprefix),
		errLogger:   log.New(os.Stdout, "[ERROR]", log.LstdFlags|log.Lmsgprefix),
		fatalLogger: log.New(os.Stdout, "[FATAL]", log.LstdFlags|log.Lmsgprefix),
	}
}

func SetStdoutLv(level int64) {
	if level < LDEBUG || level > LFATAL {
		return
	}
	atomic.StoreInt64(&Logger.stdlevel, level)
}
func SetAllwriter(w io.Writer) {
	Logger.allwriter = w
}
func Setallwriterlevel(level int64) {
	if level < LDEBUG || level > LFATAL {
		return
	}
	atomic.StoreInt64(&Logger.allwriterlevel, level)
}
func PF(lv int, format string, v ...interface{}) {
	var ios []io.Writer
	switch lv {
	case LERROR:
		if atomic.LoadInt64(&Logger.stdlevel) <= LERROR {
			ios = append(ios, os.Stdout)
		}
		if atomic.LoadInt64(&Logger.allwriterlevel) <= LERROR {
			ios = append(ios, Logger.allwriter)
		}
		if len(ios) != 0 {
			Logger.errLogger.SetOutput(io.MultiWriter(ios...))
			Logger.errLogger.Printf(format, v...)
		}
	case LDEBUG:
		if atomic.LoadInt64(&Logger.stdlevel) <= LDEBUG {
			ios = append(ios, os.Stdout)
		}
		if atomic.LoadInt64(&Logger.allwriterlevel) <= LDEBUG {
			ios = append(ios, Logger.allwriter)
		}
		if len(ios) != 0 {
			Logger.debugLogger.SetOutput(io.MultiWriter(ios...))
			Logger.debugLogger.Printf(format, v...)
		}
	case LFATAL:
		if atomic.LoadInt64(&Logger.stdlevel) <= LFATAL {
			ios = append(ios, os.Stdout)
		}
		if atomic.LoadInt64(&Logger.allwriterlevel) <= LFATAL {
			ios = append(ios, Logger.allwriter)
		}
		if len(ios) != 0 {
			Logger.fatalLogger.SetOutput(io.MultiWriter(ios...))
			Logger.fatalLogger.Printf(format, v...)
		}
	case LWARN:
		if atomic.LoadInt64(&Logger.stdlevel) <= LWARN {
			ios = append(ios, os.Stdout)
		}
		if atomic.LoadInt64(&Logger.allwriterlevel) <= LWARN {
			ios = append(ios, Logger.allwriter)
		}
		if len(ios) != 0 {
			Logger.warnLogger.SetOutput(io.MultiWriter(ios...))
			Logger.warnLogger.Printf(format, v...)
		}
	case LINFO:
		if atomic.LoadInt64(&Logger.stdlevel) <= LINFO {
			ios = append(ios, os.Stdout)
		}
		if atomic.LoadInt64(&Logger.allwriterlevel) <= LINFO {
			ios = append(ios, Logger.allwriter)
		}
		if len(ios) != 0 {
			Logger.infoLogger.SetOutput(io.MultiWriter(ios...))
			Logger.infoLogger.Printf(format, v...)
		}
	}

}
