package control

import (
	"fmt"
	"personal-registration/request"
	"personal-registration/response"
	"personal-registration/src/dao"
	"personal-registration/src/model"
)

func PersonRegister(personRequest request.PersonRequest) bool {
	var registered bool
	var err error

	person := model.NewPerson(
		personRequest.Name,
		personRequest.Cpf,
		personRequest.Password,
		personRequest.Sex,
	)

	if person != (model.Person{}) {
		registered, err = dao.PersonInsert(person)
		if err != nil {
			fmt.Println("Error cadastro: ", err.Error())
			return false
		}
	}

	return registered
}

func PersonEdit(personIdRequest int, personRequest request.PersonRequest) bool {
	var changed bool
	var err error

	person := model.NewPerson(
		personRequest.Name,
		personRequest.Cpf,
		personRequest.Password,
		personRequest.Sex,
	)
	person.SetId(personIdRequest)

	if person != (model.Person{}) {
		changed, err = dao.PersonEdit(person)
		if err != nil {
			fmt.Println("Error Alteração: ", err.Error())
			return false
		}
	}

	return changed
}

func PersonSelectByIdRegister(personId int) (response.PersonResponse, error) {
	var p model.Person
	var personResponse response.PersonResponse
	var err error
	if personId != 0 {
		p, err = dao.PersonSelectById(personId)
		if err != nil {
			println("Error na busca das informações da pessoa - camada controller: ", err.Error())
			return personResponse, err
		}
		personResponse = response.PersonResponse{
			Name:     p.GetName(),
			Cpf:      p.GetCpf(),
			Password: p.GetPassword(),
			Sex:      p.GetSex(),
		}

	} else {
		println("O id informado é inválido!")
	}
	return personResponse, err
}

func PersonDelete(personId int) bool {
	var success bool
	var err error
	if personId != 0 {
		success, err = dao.PersonDelete(personId)
		if err != nil {
			println("Error durante a deleção do registro da pessoa: ", err.Error())
			return success
		}
	} else {
		println("Informe um id de pessoa válido!")
		success = false
		return success
	}

	return success
}
