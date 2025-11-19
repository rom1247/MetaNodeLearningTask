package task_three

import (
	"backend/backend/phase_two/task_three/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initStudent() {
	Init()
	DB.Migrator().DropTable(&models.Student{})
	DB.AutoMigrate(&models.Student{})
}

/*
*编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
 */
func TestCreateStudent(t *testing.T) {
	initStudent()
	student := models.Student{Name: "张三", Age: 20, Grade: "三年级"}
	create := DB.Create(&student)
	rowsAffected := create.RowsAffected
	assert.Equal(t, int64(1), rowsAffected)
}

/*
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息
*/
func TestQueryStudent(t *testing.T) {
	initStudent()
	//插入数据
	student := []models.Student{
		{Name: "张三", Age: 20, Grade: "三年级"},
		{Name: "李四", Age: 17, Grade: "三年级"},
	}
	DB.Save(&student)
	var students []models.Student
	DB.Where("age > ?", 18).Find(&students)
	assert.Equal(t, 1, len(students))
}

// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
func TestUpdateStudent(t *testing.T) {
	initStudent()
	student := models.Student{Name: "张三", Age: 20, Grade: "三年级"}
	DB.Create(&student)
	DB.Debug().Model(student).Where("name = ?", "张三").Update("grade", "四年级")
	assert.Equal(t, "三年级", student.Grade) // 不使用指针类型更新，原来的值不会变化，但是数据库已经更新

	student2 := models.Student{}
	DB.Where("name = ?", "张三").Find(&student2)
	assert.Equal(t, "四年级", student2.Grade)

	student3 := models.Student{Name: "李四", Age: 20, Grade: "三年级"}
	DB.Create(&student3)
	DB.Debug().Model(&student3).Where("name = ?", "李四").Update("grade", "四年级")
	assert.Equal(t, "四年级", student3.Grade) // 不使用指针类型更新，原来的值不会变化，但是数据库已经更新

}

// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
func TestDeleteStudent(t *testing.T) {
	initStudent()
	student := []models.Student{
		{Name: "张三", Age: 20, Grade: "三年级"},
		{Name: "李四", Age: 17, Grade: "三年级"},
		{Name: "王五", Age: 14, Grade: "三年级"},
	}
	tx := DB.Save(&student)
	assert.Equal(t, int64(3), tx.RowsAffected)

	DB.Where("age < ?", 15).Delete(&models.Student{})

	var students []models.Student
	DB.Find(&students)
	assert.Equal(t, 2, len(students))
}
