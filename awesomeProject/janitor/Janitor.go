package janitor

type Janitor struct {
	salary   float64
	position string
	address  string
}

func (j *Janitor) GetSalary() float64 {
	return j.salary
}

func (j *Janitor) SetSalary(salary float64) {
	j.salary = salary
}

func (j *Janitor) GetPosition() string {
	return j.position
}

func (j *Janitor) SetPosition(position string) {
	j.position = position
}

func (j *Janitor) GetAddress() string {
	return j.address
}

func (j *Janitor) SetAddress(address string) {
	j.address = address
}
