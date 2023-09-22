package accountant

type Accountant struct {
	salary   float64
	position string
	address  string
}

func (a *Accountant) GetSalary() float64 {
	return a.salary
}

func (a *Accountant) SetSalary(salary float64) {
	a.salary = salary
}

func (a *Accountant) GetPosition() string {
	return a.position
}

func (a *Accountant) SetPosition(position string) {
	a.position = position
}

func (a *Accountant) GetAddress() string {
	return a.address
}

func (a *Accountant) SetAddress(address string) {
	a.address = address
}
