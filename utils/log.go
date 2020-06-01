package utils

import "go.uber.org/zap"

var ZlLoggor *zap.SugaredLogger

func init (){
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	ZlLoggor = logger.Sugar()
}