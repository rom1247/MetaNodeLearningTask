package app

import (
	"backend/backend/phase_two/task_four/internal/domain/service"
	"backend/backend/phase_two/task_four/internal/interfaces/http/controller/response"
)

type PostAppService struct {
	// 依赖注入 指针类型 ，这里有点像java spring里的bean，单例模式，整个程序复用一个结构体
	postService *service.PostService
}

func NewPostAppService(postService *service.PostService) *PostAppService {
	return &PostAppService{
		postService: postService,
	}
}

func (s *PostAppService) CreatePost(userId uint, title string, content string) error {
	return s.postService.CreatePost(userId, title, content)
}

func (s *PostAppService) UpdatePost(id uint, title string, content string) error {
	return s.postService.UpdatePost(id, title, content)
}

func (s *PostAppService) DeletePost(id uint) error {
	return s.postService.DeletePost(id)
}

func (s *PostAppService) FindAllPosts() ([]response.PostResponse, error) {
	posts, err := s.postService.FindAllPosts()
	res := make([]response.PostResponse, len(posts))
	for i, post := range posts {
		res[i] = response.PostResponse{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
			UserID:  post.UserID,
		}
	}
	return res, err
}

func (s *PostAppService) FindOnePost(id uint) (response.PostResponse, error) {
	post, err := s.postService.FindOnePost(id)
	return response.PostResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.UserID,
	}, err
}
