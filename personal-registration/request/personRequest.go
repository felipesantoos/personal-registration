package request

type PersonRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Cpf      string `json:"cpf"`
	Sex      string `json:"sex"`
	Password string `json:"password"`
}
