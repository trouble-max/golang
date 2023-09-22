package employee

type Employee interface {
	GetSalary() float64
	SetSalary(float64)
	GetPosition() string
	SetPosition(string)
	GetAddress() string
	SetAddress(string)
}
