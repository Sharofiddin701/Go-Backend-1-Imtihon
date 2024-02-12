package main

import (
	"fmt"
	"os"
)

type reader struct {
	Name   string
	Grades []int
	Course int
}

func main() {
	var arif, s, k float64

	student := reader{
		Name:   "John",
		Grades: []int{60, 70, 80, 90},
		Course: 7,
	}

	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Fprintln(file, "Name:", student.Name)
	fmt.Fprintln(file, "Grades:", student.Grades)
	fmt.Fprintln(file, "Course:", student.Course)

	for _, item := range student.Grades {
		s++
		k += float64(item)
		arif = k / s
	}
	fmt.Fprintln(file, "O'rtacha bahosi:", arif)

	var result string
	if arif > 60 {
		result = "You are passed"
	} else {
		result = "You are failed"
	}
	fmt.Fprintln(file, result)

	fmt.Println("Result has been written to result.txt")
}
