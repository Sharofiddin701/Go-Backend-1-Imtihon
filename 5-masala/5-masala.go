package main

import (
	"fmt"
	"os"
)

type Student struct {
	name   string
	grades []int
	course string
}


func (s *Student) calculateAverage() float64 {
	sum := 0
	for _, grade := range s.grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.grades))
}


func (s *Student) writeResultToFile(fileName string) error {
	var result string
	if s.calculateAverage() < 60 {
		result = "You are failed"
	} else {
		result = "You are passed"
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(result)
	if err != nil {
		return err
	}

	return nil
}
func main() {
	student := Student{
		name:   "John",
		grades: []int{70, 80, 90},
		course: "Math",
	}

	resultFileName := "result.txt"
	err := student.writeResultToFile(resultFileName)
	if err != nil {
		fmt.Println("Error writing result to file:", err)
		return
	}

	fmt.Println("Result successfully written to", resultFileName)
}
