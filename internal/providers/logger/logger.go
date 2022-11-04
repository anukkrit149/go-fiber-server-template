package logger

import (
	"context"
)

//var log *logger.ZapLogger

func InitLogger(ctx context.Context) {
	//lgrConfig := logger.Config{
	//	LogLevel:      logger.Debug,
	//	ContextString: "go-foundation",
	//}
	//
	//Logger, err := logger.NewLogger(lgrConfig)
	//
	//if err != nil {
	//	panic("failed to initialize logger")
	//}
	//log = Logger
	//
	//context.WithValue(ctx, logger.LoggerCtxKey, Logger)

}

//func RootLogger() *logger.ZapLogger {
//	return log
//}
//
//func Logger(ctx context.Context) *logger.Entry {
//	ctxLogger, err := logger.Ctx(ctx)
//
//	if err == nil {
//		return ctxLogger
//	}
//
//	return nil
//}
