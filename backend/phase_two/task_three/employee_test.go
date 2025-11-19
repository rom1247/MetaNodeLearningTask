package task_three

import (
	"backend/backend/phase_two/task_three/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initEmployee() {
	Init()
	DB.Migrator().DropTable(&models.Employee{})
	DB.AutoMigrate(&models.Employee{})
}

// 查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片
func TestQueryEmployee(t *testing.T) {
	initEmployee()
	Xdb.Exec("insert into employees (id, name, department, salary) values (1, '张三', '技术部', 5000), (2, '李四', '销售部', 4000), (3, '王五', '财务部', 3000), (4, '赵六', '技术部', 6000), (5, '孙七', '销售部', 5000)")
	var employees []models.Employee
	Xdb.Select(employees, "select * from employees where department = ?", "技术部")
	for _, employee := range employees {
		assert.Contains(t, []int{1, 4}, employee.Id)
	}

}

// 查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中
func TestQueryMaxSalaryEmployee(t *testing.T) {
	initEmployee()
	Xdb.Exec("insert into employees (id, name, department, salary) values (1, '张三', '技术部', 5000), (2, '李四', '销售部', 4000), (3, '王五', '财务部', 3000), (4, '赵六', '技术部', 6000), (5, '孙七', '销售部', 5000)")
	employee := models.Employee{}
	err := Xdb.Get(&employee, "select * from employees where salary = (select max(salary) from employees)")
	assert.Nil(t, err)
	assert.Equal(t, "赵六", employee.Name)
}
