package main

import "classes/polymorphism/income"

func main() {
	project1 := income.NewFixedBilling("Project 1", 5000)
	project2 := income.NewFixedBilling("Project 2",  10000)
	project3 := income.NewTimeAndMaterial("Project 3",  160,  25)
	incomeStreams := []income.Income{project1, project2, project3}
	income.CalculateNetIncome(incomeStreams)
}
