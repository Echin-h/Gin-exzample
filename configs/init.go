package configs

// 1.从yaml中读取数据NewSettings 2.从viper中获取能用的数据readSection
import (
	"LearningGo/internal/global/log"
)

func Init() {
	err := SetUpSettings()
	if err != nil {
		log.SugarLogger.Error(err)
		return
	}
}
