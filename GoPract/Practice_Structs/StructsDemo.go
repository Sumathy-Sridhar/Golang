package main

import "fmt"

type Student struct {
	rollno      int
	studentname string
	course      string
}

func main() {
	stud1 := Student{ // fields of struct are mentioned exclusively
		rollno:      123,
		studentname: "Sumathy",
		course:      "MCA",
	}
	stud2 := Student{143, "Mathew", "B.E"} // without specifying the fields of struct values are declared
	fmt.Println("Student Details: ")
	fmt.Println("Student 1 ==> ", stud1)
	fmt.Println("Student 2 ==> ", stud2)
	book1 := struct { // anonymous struct creates only a new struct variable and does not define any new struct type like named structs
		bookname string

		bookid int
		author string
	}{
		bookname: "Can Love happen twice?",
		bookid:   1245,
		author:   "Ravindhar Singh",
	}
	fmt.Println("Book 1 ==> ", book1)
	//Accessing individual fields of all structs
	fmt.Println("Student Name: ", stud1.studentname)
	fmt.Println("Course of Student2 : ", stud2.course)
	fmt.Println("Book Name before modification: ", book1.bookname)
	book1.bookname = "Everyone has a story"
	fmt.Println("Book Name: ", book1.bookname)
	var stud3 Student                                // zero valued struct
	fmt.Println("Student Name: ", stud3.studentname) //returns empty string as no value is assigned
	fmt.Println("Roll No: ", stud3.rollno)           // reuturns zero
	fmt.Println("Course: ", stud3.course)
	// ignoring or specifying values for selected fields
	stud4 := Student{
		rollno:      143,
		studentname: "",
		course:      "MBA",
	}
	fmt.Println(stud4)
}
