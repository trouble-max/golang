package salesman

type Salesman struct {
	salary   float64
	position string
	address  string
}

func (s *Salesman) GetSalary() float64 {
	return s.salary
}

func (s *Salesman) SetSalary(salary float64) {
	s.salary = salary
}

func (s *Salesman) GetPosition() string {
	return s.position
}

func (s *Salesman) SetPosition(position string) {
	s.position = position
}

func (s *Salesman) GetAddress() string {
	return s.address
}

func (s *Salesman) SetAddress(address string) {
	s.address = address
}
