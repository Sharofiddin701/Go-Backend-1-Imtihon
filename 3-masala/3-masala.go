package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Sex  string
	Job  string
	Age  int
}

func main() {
	var n int

	fmt.Printf("Enter the number of animals:")
	fmt.Scan(&n)
	animals := make([]Animal, 0)

	for i := 0; i < n; i++ {
		var name, sex, job string
		var age int

		fmt.Print("Name: ")
		fmt.Scan(&name)

		fmt.Print("Sex: ")
		fmt.Scan(&sex)

		fmt.Print("Job: ")
		fmt.Scan(&job)

		fmt.Print("Age: ")
		fmt.Scan(&age)

		animals = append(animals, Animal{Name: name, Sex: sex, Job: job, Age: age})
	}

	fmt.Println("Which do you want to see? (males | females | all)")
	var choice string
	fmt.Scan(&choice)

	switch choice {
	case "males":
		fmt.Println("Male Animals:")
		for _, animal := range animals {
			if animal.Sex == "males" {
				fmt.Println(animal)
			}
		}
	case "females":
		fmt.Println("Female Animals:")
		for _, animal := range animals {
			if animal.Sex == "females" {
				fmt.Println(animal)
			}
		}
	case "all":
		fmt.Println("All Animals:")
		for _, animal := range animals {
			fmt.Println(animal)
		}
	default:
		fmt.Println("Invalid choice")
	}
}
