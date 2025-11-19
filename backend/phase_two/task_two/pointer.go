package task_two

// AddTen 定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10,
func plusTen(num *int) {
	*num += 10
}

func plusPointerValue(num *int, num2 int) {
	*num += num2
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
func multiply(num *[]int, multiple int) {

	for i := 0; i < len(*num); i++ {
		(*num)[i] *= multiple
	}

}
