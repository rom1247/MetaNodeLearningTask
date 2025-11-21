package app

import "backend/backend/phase_two/task_four/internal/domain/service"

type PostAppService struct {
	// 依赖注入 指针类型 ，这里有点像java spring里的bean，单例模式，整个程序复用一个结构体
	postService *service.PostService
}

func NewPostAppService(postService *service.PostService) *PostAppService {
	return &PostAppService{
		postService: postService,
	}
}
