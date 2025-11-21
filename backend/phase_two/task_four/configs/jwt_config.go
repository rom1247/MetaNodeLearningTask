package configs

import "time"

type JwtConfig struct {
	Secret    string        `yaml:"secret"`
	ExpiresIn time.Duration `yaml:"expires_in"`
}

func NewJwtConfig() *JwtConfig {
	return &JwtConfig{
		Secret:    "defaultSecret123",
		ExpiresIn: time.Hour * 24, // 默认 24 小时
	}
}
