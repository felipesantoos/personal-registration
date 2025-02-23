package dao

import (
	"fmt"
	"personal-registration/src/connection"
	"personal-registration/src/model"
)

func PersonInsert(person model.Person) (bool, error) {

	db, err := connection.NewConnection()
	var registered bool
	if err != nil {
		fmt.Println("Error0: ", err.Error())
		return registered, err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error1: ", err.Error())
		return registered, err
	}

	sql := "select exists(select id from person where cpf=$1) as exist"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		return registered, err
		fmt.Println("Error2: ", err.Error())
	}

	rows, err := tx.Query(sql,
		person.GetCpf())
	if err != nil {
		tx.Rollback()
		fmt.Println("Error3: ", err.Error())
		return registered, err
	}

	var exist bool
	for rows.Next() {
		err = rows.Scan(&exist)
		if err != nil {
			tx.Rollback()
			return registered, err
			fmt.Println("Error5: ", err.Error())
		}
	}

	if exist == true {
		tx.Rollback()
		return registered, err
	}

	for rows.Next() {
		err = rows.Scan(&exist)
		if err != nil {
			tx.Rollback()
			fmt.Println("Error6: ", err.Error())
			return registered, err
		}
	}

	sql = "insert into person(name, cpf, password, sex)" +
		" values ($1, $2, $3, $4)"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error1: ", err.Error())
		return registered, err
	}

	_, err = tx.Exec(sql,
		person.GetName(),
		person.GetCpf(),
		person.GetPassword(),
		person.GetSex(),
	)

	if err != nil {
		tx.Rollback()
		fmt.Println("Error7: ", err.Error())
		return registered, err
	}

	tx.Commit()
	registered = true
	return registered, err
}

func PersonEdit(person model.Person) (bool, error) {

	db, err := connection.NewConnection()
	var registered bool
	if err != nil {
		fmt.Println("Error0: ", err.Error())
		return registered, err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error1: ", err.Error())
		return registered, err
	}

	sql := "update person set name=$1, cpf=$2, password=$3, sex=$4 where id=$5"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error2: ", err.Error())
		return registered, err
	}

	_, err = tx.Exec(sql,
		person.GetName(),
		person.GetCpf(),
		person.GetPassword(),
		person.GetSex(),
		person.GetId(),
	)

	if err != nil {
		tx.Rollback()
		fmt.Println("Error3: ", err.Error())
		return registered, err
	}

	tx.Commit()
	registered = true
	return registered, err
}

func PersonSelectById(personInt int) (model.Person, error) {
	db, err := connection.NewConnection()
	if err != nil {
		println("Error1: ", err.Error())
	}
	defer db.Close()

	sql := "select name, cpf, password, sex from person where id=$1"
	_, err = db.Prepare(sql)
	if err != nil {
		println("Error2: ", err.Error())
	}

	var personNameDB, personCpfDB, personPasswordDB, personSexDB string
	var person model.Person
	rows, err := db.Query(sql, personInt)
	for rows.Next() {
		err = rows.Scan(&personNameDB, &personCpfDB, &personPasswordDB, &personSexDB)
		if err != nil {
			println("Error nos dados retornados: ", err.Error())
			return person, err
		}
	}
	person = model.NewPerson(personNameDB, personCpfDB, personPasswordDB, personSexDB)

	return person, nil

}

func PersonDelete(personId int) (bool, error) {
	var success bool
	db, err := connection.NewConnection()
	if err != nil {
		println("ErroDelecao0: ", err.Error())
		return false, err
	}

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("ErrorDelecao1: ", err.Error())
		return success, err
	}

	sql := "delete from person where id = $1"
	_, err = db.Prepare(sql)
	if err != nil {
		println("ErroDelecao2: ", err.Error())
		return success, err
	}

	_, err = db.Exec(sql, personId)
	if err != nil {
		println("ErroDelecao3: ", err.Error())
		return success, err
	}

	err = tx.Commit()
	if err != nil {
		println("ErroDelecao4: ", err.Error())
		return false, err
	}
	success = true
	return success, err
}
