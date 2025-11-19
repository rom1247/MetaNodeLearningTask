package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string
	PostCount int
}

type Post struct {
	gorm.Model
	UserID       uint
	Title        string
	CommentCount int
	HasComment   bool
}

type Comment struct {
	gorm.Model
	UserId uint
	PostID uint
	Text   string
}

type CommentCount struct {
	PostId uint
	Count  int
}

// AfterCreate 创建一个钩子，文章发布后，用户的文章数加1
func (post *Post) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(&User{}).Where("id = ?", post.UserID).
		Update("post_count", gorm.Expr("post_count + ?", 1)).Error
}

func (comment *Comment) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(&Post{}).Where("id = ?", comment.PostID).
		Updates(map[string]interface{}{
			"has_comment":   gorm.Expr("EXISTS(SELECT 1 FROM comments WHERE post_id = ?)", comment.PostID),
			"comment_count": gorm.Expr("comment_count + ?", 1),
		}).Error
}

// AfterDelete 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (comment *Comment) AfterDelete(tx *gorm.DB) (err error) {

	return tx.Model(&Post{}).Where("id = ?", comment.PostID).
		Updates(map[string]interface{}{
			"has_comment":   gorm.Expr("EXISTS(SELECT 1 FROM comments WHERE post_id = ?)", comment.PostID),
			"comment_count": gorm.Expr("comment_count - ?", 1),
		}).Error
}
