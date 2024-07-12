package models

type Status string

const (
	Permanent Status = "Permanent"
	Contract         = "Contract"
	Trainee          = "Trainee"
)

type Employee struct {
	EmpId       int
	FullName    string
	Salary      float64
	Status      Status
	TotalSalary float64
}
