package hl

import "fmt"

func main(){
	var name string
	var age int
	fmt.Printf("Name & Age: ")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf(hello(name, age))
}

func hello(name string, age int) string {
	var message string
	message = fmt.Sprintf("Hi, %s %d, welcome!", name, age)
	// message := fmt.Sprintf("Hi, %v, welcome!", name) 直接定义赋值
	return message
}