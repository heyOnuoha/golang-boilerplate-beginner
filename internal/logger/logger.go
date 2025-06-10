package logger

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogger(env string) {

	var config zap.Config

	if env == "development" {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	config.OutputPaths = []string{"stdout"}

	var err error
	Logger, err = config.Build()

	if err != nil {
		panic(err)
	}

	// Don't defer Logger.Sync() here as it would only execute when this function returns
	// Instead, we should call Logger.Sync() in the main function before the program exits

	// Replace the global zap logger
	zap.ReplaceGlobals(Logger)
}
