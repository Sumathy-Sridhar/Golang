package main

import "fmt"

type Student struct {
	name   string
	rollno int
	course string
}

func (stud Student) stuMethod() {
	fmt.Println("Student Name: ", stud.name, "Roll Number: ", stud.rollno, "Course: ", stud.course)
}

func main() {
	stud1 := Student{
		name:   "Sumathy",
		rollno: 123,
		course: "MCA",
	}
	stud1.stuMethod()
}
