package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name:%s, age:%d", "Luong Duong", 25) // like fmf.Printf(format, a)

	// //logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("Name", "Luong Duong"), zap.Int("Age", 25))
	// 2
	// logger := zap.NewExample()
	// logger.Info("Hello NewExample")
	// // Development
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment")
	// // Production
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	//3
	encoder := getEncoderLog()
	sync := getWriteSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core,zap.AddCaller())
	logger.Info("Info Log", zap.Int("line", 1))
	logger.Error("Error Log", zap.Int("line", 2))
}

// format log a message
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	// 1718211230.4289014 -> 2024-06-12T23:53:50.427+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> time
	encodeConfig.TimeKey = "time"
	// from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// "caller":"cli/main.log.go:19"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// return zapcore.NewConsoleEncoder()
	return zapcore.NewJSONEncoder(encodeConfig) // co ban

}

func getWriteSync() zapcore.WriteSyncer {
	// file, _ := os.OpenFile(name, flag, perm)
	file, _ := os.OpenFile("./log/log.txt",os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
