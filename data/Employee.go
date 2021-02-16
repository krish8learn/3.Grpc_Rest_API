package data

import "fmt"

type Emp struct {
	Empid     string
	Empname   string
	Empmail   string
	Empmobile string
}

func Stored(str string) []Emp {
	fmt.Println(str)
	p1 := Emp{
		Empid:     "1",
		Empname:   "Kevin_ Konroy",
		Empmail:   "xyz@gmail.com",
		Empmobile: "+2144656565",
	}
	p2 := Emp{
		Empid:     "2",
		Empname:   "Harry Kane",
		Empmail:   "kane@gmail.com",
		Empmobile: "+19088766565",
	}
	p3 := Emp{
		Empid:     "3",
		Empname:   "Kai Havertz",
		Empmail:   "havertz@gmail.com",
		Empmobile: "+31557656565",
	}
	EmpSlice := []Emp{}
	EmpSlice = append(EmpSlice, p1, p2, p3)
	return EmpSlice
}
