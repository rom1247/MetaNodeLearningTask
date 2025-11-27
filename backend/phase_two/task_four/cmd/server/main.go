package main

import (
	"backend/backend/phase_two/task_four/internal/bootstrap"
	"log"
)

// @title 博客系统 API 文档
// @version 1.0
// @description 后端接口文档，包含用户、登录、文章、多端鉴权等内容。
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer {token}" in the header
func main() {
	app, err := bootstrap.InitializeApp()
	if err != nil {
		log.Fatal(err)
	}
	app.Run()
}
