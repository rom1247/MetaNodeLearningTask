package response

type CommentResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	PostID  uint   `json:"post_id"`
}
