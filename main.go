package main

import (
	"fmt"
	"math/rand"
	"sync"
	"weeklyTest2/models"
)

func randomName() string {
	names := []string{"Budi", "Ario", "Joni", "Loti", "Nasir", "Jarwo"}
	return names[rand.Intn(len(names))]
}

func randomStatus() models.Status {
	statuses := []models.Status{models.Permanent, models.Contract, models.Trainee}
	return statuses[rand.Intn(len(statuses))]
}

func createEmployee(id int, ch chan<- interface{}) {
	salary := 5000 + rand.Float64()*(15000-5000)
	status := randomStatus()
	name := randomName()

	switch status {
	case models.Permanent:
		ch <- models.PermanentModel{
			Employee: models.Employee{
				EmpId:    id,
				FullName: name,
				Salary:   salary,
				Status:   status,
			},
			Insurance: 500_000,
		}
	case models.Contract:
		ch <- models.ContractModel{
			Employee: models.Employee{
				EmpId:    id,
				FullName: name,
				Salary:   salary,
				Status:   status,
			},
			Overtime: 55_000,
		}
	case models.Trainee:
		ch <- models.TraineeModel{
			Employee: models.Employee{
				EmpId:    id,
				FullName: name,
				Salary:   salary,
				Status:   status,
			},
			Allowance: 100_000,
		}
	}

}

func main() {
	numEmployees := 100
	ch := make(chan interface{}, numEmployees)
	var wg sync.WaitGroup

	for i := 1; i <= numEmployees; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			createEmployee(id, ch)
		}(i)
	}

	wg.Wait()
	close(ch)

	var permanentEmployees []models.PermanentModel
	var contractEmployees []models.ContractModel
	var traineeEmployees []models.TraineeModel

	for e := range ch {
		switch emp := e.(type) {
		case models.PermanentModel:
			permanentEmployees = append(permanentEmployees, emp)
		case models.ContractModel:
			contractEmployees = append(contractEmployees, emp)
		case models.TraineeModel:
			traineeEmployees = append(traineeEmployees, emp)
		}

	}

	fmt.Println("Permanen employees")
	for _, v := range permanentEmployees {
		fmt.Println(v)
	}

	fmt.Println("Contracrt employees")
	for _, v := range contractEmployees {
		fmt.Println(v)
	}
	fmt.Println("Trainee employees")
	for _, v := range traineeEmployees {
		fmt.Println(v)
	}

}
