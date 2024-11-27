package model

import (
	"strings"
)

type Person struct {
	id       int
	name     string
	cpf      string
	password string
	sex      string
}

func NewPerson(name string, cpf string, password string, sex string) Person {
	person := Person{
		name:     name,
		cpf:      cpf,
		password: password,
		sex:      sex,
	}
	return person
}

func (person *Person) GetId() int {
	return person.id
}

func (person *Person) SetId(id int) {
	person.id = id
}

func (person *Person) GetName() string {
	return person.name
}

func (person *Person) SetName(name string) {
	person.name = strings.TrimSpace(name)
}

func (person *Person) GetCpf() string {
	return person.cpf
}

func (person *Person) SetCpf(cpf string) {
	person.cpf = strings.TrimSpace(cpf)
}

func (person *Person) GetPassword() string {
	return person.password
}

func (person *Person) SetPassword(password string) {
	person.password = strings.TrimSpace(password)
}

func (person *Person) GetSex() string {
	return person.sex
}

func (person *Person) SetSex(sex string) {
	person.sex = strings.TrimSpace(sex)
}
