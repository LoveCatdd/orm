package orm

import (
	"fmt"

	zlog "github.com/LoveCatdd/util/pkg/lib/core/log"
	"xorm.io/xorm/log"
)

var ormLog = zlog.OutZapLog(1)

// xormLogger 实现 xorm 的 ILogger 接口
type xormLogger struct {
	level   log.LogLevel
	showSQL bool
}

func NewXormLogger(level log.LogLevel) *xormLogger {
	return &xormLogger{level: level}
}

func (l *xormLogger) Debug(v ...interface{}) {
	if l.level <= log.LOG_DEBUG {
		ormLog.Debug(fmt.Sprint(v...))
	}
}

func (l *xormLogger) Info(v ...interface{}) {
	if l.level <= log.LOG_INFO {
		ormLog.Info(fmt.Sprint(v...))
	}
}

func (l *xormLogger) Warn(v ...interface{}) {
	if l.level <= log.LOG_WARNING {
		ormLog.Warn(fmt.Sprint(v...))
	}
}

func (l *xormLogger) Error(v ...interface{}) {
	if l.level <= log.LOG_ERR {
		ormLog.Error(fmt.Sprint(v...))
	}
}

func (l *xormLogger) Level() log.LogLevel {
	return l.level
}

func (l *xormLogger) SetLevel(level log.LogLevel) {
	l.level = level
}

func (l *xormLogger) ShowSQL(show ...bool) {
	if len(show) > 0 {
		l.showSQL = show[0]
		ormLog.Info(fmt.Sprintf("orm.xorm.ShowSQL set to %v", show[0]))
	}
}

func (l *xormLogger) IsShowSQL() bool {
	return l.showSQL
}

func withLevel(level string) log.LogLevel {
	switch level {
	case LOG_DEBUG:
		return log.LOG_DEBUG
	case LOG_INFO:
		return log.LOG_INFO
	case LOG_WARNING:
		return log.LOG_WARNING
	}
	ormLog.Error(fmt.Sprintf("unknown orm.xorm level: %v", level))
	return log.LOG_UNKNOWN
}

func (l *xormLogger) Debugf(format string, v ...interface{}) {
	if l.level <= log.LOG_DEBUG {
		ormLog.Debug(fmt.Sprintf(format, v...))
	}
}

func (l *xormLogger) Infof(format string, v ...interface{}) {
	if l.level <= log.LOG_INFO {
		ormLog.Info(fmt.Sprintf(format, v...))
	}
}

func (l *xormLogger) Warnf(format string, v ...interface{}) {
	if l.level <= log.LOG_WARNING {
		ormLog.Warn(fmt.Sprintf(format, v...))
	}
}

func (l *xormLogger) Errorf(format string, v ...interface{}) {
	if l.level <= log.LOG_ERR {
		ormLog.Error(fmt.Sprintf(format, v...))
	}
}
