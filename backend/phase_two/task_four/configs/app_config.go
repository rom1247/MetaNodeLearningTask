package configs

type AppConfig struct {
	AutoMigrate bool // true 表示启动时自动迁移
}

func NewConfig() *AppConfig {
	return &AppConfig{
		AutoMigrate: true, // 默认自动迁移，可修改为false
	}
}
