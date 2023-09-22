package engineer

type Engineer struct {
	salary   float64
	position string
	address  string
}

func (e *Engineer) GetSalary() float64 {
	return e.salary
}

func (e *Engineer) SetSalary(salary float64) {
	e.salary = salary
}

func (e *Engineer) GetPosition() string {
	return e.position
}

func (e *Engineer) SetPosition(position string) {
	e.position = position
}

func (e *Engineer) GetAddress() string {
	return e.address
}

func (e *Engineer) SetAddress(address string) {
	e.address = address
}
