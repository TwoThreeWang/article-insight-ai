package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Config struct {
	// Google Gemini API配置
	GeminiAPIKey string `json:"gemini_api_key"`
}

var (
	config *Config
	once   sync.Once
)

// LoadConfig 加载配置文件
func LoadConfig(configPath string) (*Config, error) {
	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	once.Do(func() {
		// 初始化默认配置
		config = &Config{}

		// 解析JSON配置文件
		if err := json.Unmarshal(data, config); err != nil {
			fmt.Printf("解析配置文件失败: %v，将使用默认配置", err)
		}
	})

	return config, nil
}

// GetConfig 获取配置实例
func GetConfig() *Config {
	return config
}
