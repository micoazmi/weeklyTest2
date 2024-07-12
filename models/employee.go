package models

type Status string

const (
	Permanent Status = "Permanent"
	Contract         = "Contract"
	Trainee          = "Trainee"
)

type Employee struct {
	empId    int
	fullName string
	salary   float64
	status   Status
}
