package interseptors

import (
	"context"

	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/sirupsen/logrus"
)

// Maybe refactor grpc's in folder
// That there will be opts and interseptors create func with parametrs
// TODO interseptors
func InterceptorLogger(l *logrus.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, lvl grpclog.Level, msg string, fields ...any) {
		var logrusLevel logrus.Level
		switch lvl {
		case grpclog.LevelDebug:
			logrusLevel = logrus.DebugLevel
		case grpclog.LevelInfo:
			logrusLevel = logrus.InfoLevel
		case grpclog.LevelWarn:
			logrusLevel = logrus.WarnLevel
		case grpclog.LevelError:
			logrusLevel = logrus.ErrorLevel
		default:
			logrusLevel = logrus.InfoLevel
		}

		l.WithFields(logrus.Fields{"details": fields}).Log(logrusLevel, msg)
	})
}
