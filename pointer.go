package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	student := &Student{
		Name: "nguyen van A",
		Age:  20,
	}
	modifyStudent(student)
	fmt.Println(student)
}

func modifyStudent(student *Student) {
	student2 := Student{
		Name: "nguyen van B",
		Age:  30,
	}
	*student = student2
}
