package log

import "go.uber.org/zap"

// New function return the sugared logger instance.
func New() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	return logger.Sugar()
}
