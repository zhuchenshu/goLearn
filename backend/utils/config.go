package utils

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
}

var CONFIG = new(Config)

func init() {
	viper.SetConfigName("goLearn")

	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		Infof("error", err.Error())
	}

	viper.AutomaticEnv()
}

// GetFloat64 获取浮点数配置
func (Config) GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

// GetString 获取字符串配置
func (Config) GetString(key string) string {
	return viper.GetString(key)
}

// GetInt 获取整数配置
func (Config) GetInt(key string) int {
	return viper.GetInt(key)
}

// GetInt32 获取 int32 配置
func (Config) GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

// GetInt64 获取 int64 配置
func (Config) GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

// GetDuration 获取时间配置
func (Config) GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

// GetBool 获取配置布尔配置
func (Config) GetBool(key string) bool {
	return viper.GetBool(key)
}

// Set 设置配置，仅用于测试
func (Config) Set(key string, value string) {
	viper.Set(key, value)
}
