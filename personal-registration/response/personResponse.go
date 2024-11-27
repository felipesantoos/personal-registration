package response

type PersonResponse struct {
	Name     string `json:"name"`
	Cpf      string `json:"cpf"`
	Password string `json:"password"`
	Sex      string `json:"sex"`
}
