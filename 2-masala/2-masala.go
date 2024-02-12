package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

type Man struct {
	First_name  string `json:"first_name"`
	Second_name string `json:"second_name"`
	Salary      int    `json:"salary"`
	Experience  int    `json:"experience"`
	Age         int    `json:"age"`
}

func main() {
	data, err := os.ReadFile("employees.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	var mans []Man
	if err := json.Unmarshal(data, &mans); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	file, err := os.Create("orderByAge.txt")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	var maxAge int
	for _, man := range mans {
		if man.Age > maxAge {
			maxAge = man.Age

			_, err := fmt.Fprintf(file, "First_name: %v, Second_name: %v, Salary : %v, Experience: %v, Age: %v\n", man.First_name, man.Second_name, man.Salary, man.Experience, man.Age)
			if err != nil {
				log.Fatalf("Error writing to file: %v", err)
			}
		}
	}

	fmt.Println("Results have been written to orderByAge.txt")

	file1, err := os.Create("topSalaryEmployees.txt")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file1.Close()

	offSet := 0

	sort.Slice(mans, func(i, j int) bool {
		return mans[i].Salary > mans[j].Salary
	})
	for i := 0; i <= 2; i++ {

		text := fmt.Sprintf("firstname %v , secondname %v , salary: %v,expp: %v , age:%v \n", mans[i].First_name, mans[i].Second_name, mans[i].Salary, mans[i].Experience, mans[i].Age)
		written, err := file1.WriteAt([]byte(text), int64(offSet))

		if err != nil {
			panic(err)
		}
		offSet += written
	}

}
