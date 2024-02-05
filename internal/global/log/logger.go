package log

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// https://www.topgoer.com/%E9%A1%B9%E7%9B%AE/log/ZapLogger.html
var (
	SugarLogger *zap.SugaredLogger
)

// 做一个闭包
func Init() gin.HandlerFunc {
	return func(c *gin.Context) {
		initLogger()
		c.Next()
	}
}

func initLogger() {
	Encoder := getEncoder()
	WriterSyncer := getWriterSyncer()
	core := zapcore.NewCore(Encoder, WriterSyncer, zapcore.DebugLevel)
	//  zap.AddCaller()可以实现记录函数信息
	logger := zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()
}

// 编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 写入位置
func getWriterSyncer() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./text.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
