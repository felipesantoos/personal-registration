package routers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"personal-registration/request"
	"personal-registration/src/control"
	"strconv"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {

	// Decodifica os dados JSON do corpo da solicitação
	var person request.PersonRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	registered := control.PersonRegister(person)

	if registered == false {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Cadastro realizado com sucesso")
	}
}

func EditPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var personEditRequest request.PersonRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&personEditRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	personIdRequest, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//personEditRequest := dataRequest.PatientRequest

	success := control.PersonEdit(personIdRequest, personEditRequest)

	if success == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Cadastro do paciente alterado com sucesso!")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("Não foi possível editar o cadastro do paciente!")
	}
}

func GetPersonById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "Id não foi informado", http.StatusBadRequest)
		return
	}

	personId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	person, err := control.PersonSelectByIdRegister(personId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
	w.WriteHeader(http.StatusOK)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	// Decodifica os dados JSON do corpo da solicitação
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "Id não foi informado", http.StatusBadRequest)
		return
	}

	personId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Chama a função que lê o paciente do banco de dados
	success := control.PersonDelete(personId)

	if success == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Registro da pessoa deletado com sucesso!")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("Não foi possível deletar o registro da pessoa!")
	}
}
