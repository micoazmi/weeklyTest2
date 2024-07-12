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

func TotalSalary(salary float64, additional float64) float64 {
	return salary + additional
}

func createEmployee(id int, ch chan<- interface{}) {
	salary := 5000 + rand.Float64()*(15000-5000)
	status := randomStatus()
	name := randomName()

	switch status {
	case models.Permanent:
		insurance := 500_000.00
		totalSalary := TotalSalary(salary, insurance)
		ch <- models.PermanentModel{
			Employee: models.Employee{
				EmpId:       id,
				FullName:    name,
				Salary:      salary,
				Status:      status,
				TotalSalary: totalSalary,
			},
			Insurance: insurance,
		}
	case models.Contract:
		overtime := 55_000.00
		totalSalary := TotalSalary(salary, overtime)
		ch <- models.ContractModel{
			Employee: models.Employee{
				EmpId:       id,
				FullName:    name,
				Salary:      salary,
				Status:      status,
				TotalSalary: totalSalary,
			},
			Overtime: overtime,
		}
	case models.Trainee:
		allowance := 500_000.00
		totalSalary := TotalSalary(salary, allowance)
		ch <- models.TraineeModel{
			Employee: models.Employee{
				EmpId:       id,
				FullName:    name,
				Salary:      salary,
				Status:      status,
				TotalSalary: totalSalary,
			},
			Allowance: allowance,
		}
	}

}

func main() {
	numEmployees := 100
	ch := make(chan interface{}, numEmployees)
	salaryCh := make(chan float64, numEmployees)
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

	var totalSalary float64
	for numEmployees > 0 {
		d := <-ch
		numEmployees--
		fmt.Println(d)
		switch emp := d.(type) {
		case models.PermanentModel:
			salaryCh <- emp.TotalSalary
		case models.ContractModel:
			salaryCh <- emp.TotalSalary
		case models.TraineeModel:
			salaryCh <- emp.TotalSalary

		}

	}
	close(salaryCh)

	for v := range salaryCh {
		totalSalary += v
	}

	fmt.Printf("Total salary all employee %.2f", totalSalary)

}
