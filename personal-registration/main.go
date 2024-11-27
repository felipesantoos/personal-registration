package main

import "personal-registration/server"

//func main() {
//	var person model.Person
//
//	person = model.NewPerson("Martinho Lutero", "12345678911", "12345678", "M")
//
//	registered := control.PersonRegister(person)
//	fmt.Println("Valor booleano resultante: ", registered)
//	if registered == true {
//		fmt.Println("Pessoa cadastrada com sucesso!")
//	} else {
//		fmt.Println("Não foi possível realizar o cadastro")
//	}
//}

//func main() {
//	reader := bufio.NewReader(os.Stdin)
//
//	for {
//		fmt.Println("\nMenu de Opções:")
//		fmt.Println("1 - Cadastrar nova pessoa")
//		fmt.Println("0 - Sair")
//		fmt.Print("Escolha uma opção: ")
//
//		op, _ := reader.ReadString('\n')
//		op = strings.TrimSpace(op)
//
//		switch op {
//		case "1":
//			fmt.Println("\n--- Cadastro de Pessoa ---")
//
//			fmt.Print("Informe o nome: ")
//			name, _ := reader.ReadString('\n')
//			name = strings.TrimSpace(name)
//
//			fmt.Print("Informe o CPF: ")
//			cpf, _ := reader.ReadString('\n')
//			cpf = strings.TrimSpace(cpf)
//
//			fmt.Print("Informe o password: ")
//			password, _ := reader.ReadString('\n')
//			password = strings.TrimSpace(password)
//
//			fmt.Print("Informe o sexo (M/F): ")
//			sex, _ := reader.ReadString('\n')
//			sex = strings.TrimSpace(sex)
//
//			person := model.NewPerson(name, cpf, password, sex)
//
//			registered := control.PersonRegister(person)
//			if registered {
//				fmt.Println("Pessoa cadastrada com sucesso!")
//			} else {
//				fmt.Println("Não foi possível realizar o cadastro.")
//			}
//
//		case "0":
//			fmt.Println("Encerrando o programa.")
//			return
//
//		default:
//			fmt.Println("Opção inválida. Tente novamente.")
//		}
//	}
//}

func main() {
	server.OpenServerTest()
}
