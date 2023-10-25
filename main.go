package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

// Funcao para autenticação na página
func Secret(user, realm string) string {
	if user == "admin" { // usuário
		return "$1$23EFFqa7$1q6g9py/8onZzZ9THbxt6/" // password: senhasecreta
	}
	return ""
}

func main() {
	// Tratamento de entrada dos argumentos de entrada
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <diretorio> <porta>")
		os.Exit(1)
	}

	// Variáveis de argumentos
	dirPath := os.Args[1]
	portServer := os.Args[2]

	authenticator := auth.NewBasicAuthenticator("meuserver.com", Secret)
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(dirPath)).ServeHTTP(w, &r.Request)
	}))
	fmt.Printf("Subindo servidor na porta %s\n", portServer)        // Subindo servidor na porta argumentada
	fmt.Printf("Acesse através da url: localhost:%s\n", portServer) // Método de acesso
	log.Fatal(http.ListenAndServe(":"+portServer, nil))
}
