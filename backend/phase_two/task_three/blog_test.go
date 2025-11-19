package task_three

import (
	"backend/backend/phase_two/task_three/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initBlog() {
	Init()
	DB.Migrator().DropTable(&models.User{}, &models.Post{}, &models.Comment{})
	DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}

// 使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
func TestQueryUserPostAndComment(t *testing.T) {
	initBlog()
	//初始化用户 、文章、两个评论
	user := models.User{Name: "张三"}
	DB.Create(&user)
	post1 := models.Post{Title: "第一篇博客", UserID: user.ID}
	DB.Create(&post1)
	comment1 := models.Comment{Text: "第一篇博客的评论1", UserId: user.ID, PostID: post1.ID}
	DB.Create(&comment1)
	comment2 := models.Comment{Text: "第一篇博客的评论2", UserId: user.ID, PostID: post1.ID}
	DB.Create(&comment2)

	var posts []models.Post
	DB.Find(&posts, "user_id = ?", user.ID)

	assert.Equal(t, 1, len(posts))

	var comments []models.Comment
	DB.Find(&comments, "user_id = ?", user.ID)
	assert.Equal(t, 2, len(comments))

	//再初始化第二篇文章 和 三个评论
	post2 := models.Post{Title: "第二篇博客", UserID: user.ID}
	DB.Create(&post2)
	comment3 := models.Comment{Text: "第二篇博客的评论1", UserId: user.ID, PostID: post2.ID}
	DB.Create(&comment3)
	comment4 := models.Comment{Text: "第二篇博客的评论2", UserId: user.ID, PostID: post2.ID}
	DB.Create(&comment4)
	comment5 := models.Comment{Text: "第二篇博客的评论3", UserId: user.ID, PostID: post2.ID}
	DB.Create(&comment5)

	//使用Gorm查询评论数量最多的文章信息
	var commentCount models.CommentCount
	DB.Model(&models.Comment{}).Select("post_id, count(*) as count").
		Group("post_id").
		Order("count desc").
		First(&commentCount)
	assert.Equal(t, post2.ID, commentCount.PostId)
	assert.Equal(t, 3, commentCount.Count)
}

// 为 Post 模型添加一个钩子函数
func TestHook(t *testing.T) {
	initBlog()

	//初始化用户 、文章、两个评论
	user := models.User{Name: "张三"}
	DB.Create(&user)
	post1 := models.Post{Title: "第一篇博客", UserID: user.ID}
	DB.Create(&post1)
	comment1 := models.Comment{Text: "第一篇博客的评论1", UserId: user.ID, PostID: post1.ID}
	DB.Create(&comment1)
	comment2 := models.Comment{Text: "第一篇博客的评论2", UserId: user.ID, PostID: post1.ID}
	DB.Create(&comment2)

	var user1 models.User
	DB.First(&user1, user.ID)
	assert.Equal(t, 1, user1.PostCount)
	var post models.Post
	DB.First(&post, post1.ID)
	assert.True(t, post.HasComment)
	assert.Equal(t, 2, post.CommentCount)

	DB.Unscoped().Delete(&comment1) //真删除 ，不然统计的时候要加上deleted_at IS NULL

	DB.First(&post, post1.ID)
	assert.True(t, post.HasComment)

	DB.Unscoped().Delete(&comment2)
	DB.First(&post, post1.ID)
	assert.False(t, post.HasComment)

}
