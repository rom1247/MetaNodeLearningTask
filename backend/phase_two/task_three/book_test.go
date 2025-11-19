package task_three

import (
	"backend/backend/phase_two/task_three/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initBook() {
	Init()
	DB.Migrator().DropTable(&models.Book{})
	DB.AutoMigrate(&models.Book{})
}

// 使用Sqlx执行一个复杂的查询 查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全
func TestQueryBook(t *testing.T) {
	initBook()
	_, e := Xdb.Exec("INSERT INTO books (id,title, author, price) VALUES " +
		" (1,'Go 语言基础', 'Go 语言中文网', 49.9)," +
		" (2,'Go 语言实战', 'Go 语言中文网', 69.9)," +
		" (3,'Go 语言微服务', 'Go 语言中文网', 59.9)")

	assert.Nil(t, e)
	var books []models.Book
	err := Xdb.Select(&books, "SELECT * FROM books WHERE price > ?", 50)
	assert.Nil(t, err)

}
