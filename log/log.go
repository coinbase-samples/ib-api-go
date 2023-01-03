package log

// TODO move me to shared util package

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/coinbase-samples/ib-api-go/config"
	"github.com/coinbase-samples/ib-api-go/model"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
)

type LogCtxKeyType string

const LogCtxKey LogCtxKeyType = "log"

type Entry struct {
	l *logrus.Entry
}

type IFields interface {
	map[string]interface{}
}

type Fields map[string]interface{}

func Init(app config.AppConfig) {
	logLevel, _ := logrus.ParseLevel(app.LogLevel)
	logger.SetLevel(logLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(true)
	if app.LogToFile == "true" {
		// open a file
		f, err := os.OpenFile("testing.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			fmt.Printf("error opening file: %v", err)
		}
		logger.SetOutput(f)
	} else {
		logger.SetOutput(os.Stdout)
	}
}

func NewEntry() *Entry {
	return &Entry{l: logrus.NewEntry(logger)}
}

func MakeContextLogger(l *logrus.Entry) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := uuid.New()
			ctx := ctxlogrus.ToContext(
				context.WithValue(r.Context(), model.RequestCtxKey, requestId.String()),
				l.WithField("requestId", requestId.String()),
			)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ToContext(ctx context.Context, entry *Entry) context.Context {
	return context.WithValue(ctx, LogCtxKey, entry)
}

func Extract(ctx context.Context) *Entry {
	l, ok := ctx.Value(LogCtxKey).(*Entry)
	if !ok || l == nil {
		return NewEntry()
	}
	return l
}

func (e *Entry) Tracef(format string, args ...interface{}) {
	e.l.Tracef(format, args...)
}

func (e *Entry) Trace(args ...interface{}) {
	e.l.Trace(args...)
}

func (e *Entry) Debugf(format string, args ...interface{}) {
	e.l.Debugf(format, args...)
}

func (e *Entry) Debug(args ...interface{}) {
	e.l.Debug(args...)
}

func (e *Entry) Infof(format string, args ...interface{}) {
	e.l.Infof(format, args...)
}

func (e *Entry) Info(args ...interface{}) {
	e.l.Info(args...)
}

func (e *Entry) Warnf(format string, args ...interface{}) {
	e.l.Warnf(format, args...)
}

func (e *Entry) Warn(args ...interface{}) {
	e.l.Warn(args...)
}

func (e *Entry) Errorf(format string, args ...interface{}) {
	e.l.Errorf(format, args...)
}

func (e *Entry) Error(args ...interface{}) {
	e.l.Error(args...)
}

func (e *Entry) Fatalf(format string, args ...interface{}) {
	e.l.Fatalf(format, args...)
}

func (e *Entry) Fatal(args ...interface{}) {
	e.l.Fatal(args...)
}

func (e *Entry) Panicf(format string, args ...interface{}) {
	e.l.Panicf(format, args...)
}

func (e *Entry) Panic(args ...interface{}) {
	e.l.Panic(args...)
}

func (e *Entry) WithField(key string, value interface{}) *Entry {
	e.l = e.l.WithField(key, value)
	return e
}

func (e *Entry) WithFields(fields map[string]interface{}) *Entry {
	e.l = e.l.WithFields(fields)
	return e
}

func Tracef(format string, args ...interface{}) {
	logger.Tracef(format, args...)
}

func Trace(args ...interface{}) {
	logger.Trace(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func CtxTracef(ctx context.Context, format string, args ...interface{}) {
	Extract(ctx).Tracef(format, args...)
}

func CtxTrace(ctx context.Context, args ...interface{}) {
	Extract(ctx).Trace(args...)
}

func CtxDebugf(ctx context.Context, format string, args ...interface{}) {
	Extract(ctx).Debugf(format, args...)
}

func CtxDebug(ctx context.Context, args ...interface{}) {
	Extract(ctx).Debug(args...)
}

func CtxInfof(ctx context.Context, format string, args ...interface{}) {
	Extract(ctx).Infof(format, args...)
}

func CtxInfo(ctx context.Context, args ...interface{}) {
	Extract(ctx).Info(args...)
}

func CtxWarnf(ctx context.Context, format string, args ...interface{}) {
	Extract(ctx).Warnf(format, args...)
}

func CtxWarn(ctx context.Context, format string, args ...interface{}) {
	Extract(ctx).Warn(args...)
}

func CtxFatal(ctx context.Context, args ...interface{}) {
	Extract(ctx).Fatal(args...)
}

func CtxFatalf(ctx context.Context, format string, args ...interface{}) {
	Extract(ctx).Fatalf(format, args...)
}

func CtxPanicf(ctx context.Context, format string, args ...interface{}) {
	Extract(ctx).Panicf(format, args...)
}

func CtxPanic(ctx context.Context, args ...interface{}) {
	Extract(ctx).Panic(args...)
}
