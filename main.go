package main

import (
	"database/sql"
	"log"
	"net/http"

	"herguardian/handlers"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=herguardian sslmode=disable password=apolo777")
	if err != nil {
		log.Fatal("Erro ao abrir conexão com o banco:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Erro ao testar conexão com o banco:", err)
	}

	http.HandleFunc("/ficha/informacoes-exame", handlers.InformacoesExame(db))
	http.HandleFunc("/ficha/informacoes-pessoais", handlers.InformacoesPessoais(db))
	http.HandleFunc("/ficha/dados-anamnese", handlers.Anamnese(db))
	http.HandleFunc("/ficha/exame-clinico", handlers.Exame_clinico(db))
	http.HandleFunc("/ficha/exame-laboratorial", handlers.Exame_laboratorial(db))
	http.HandleFunc("/api/pacientes", handlers.BuscarPacienteHandler(db))
	http.HandleFunc("/criar-conta", handlers.CriarContaHandler(db))
	http.HandleFunc("/login", handlers.LoginHandler(db))
	http.HandleFunc("/perfil", handlers.PerfilHandler(db))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/escolha-login.html", http.StatusSeeOther)
	})

	porta := ":8080"
	log.Println("Servidor rodando em http://localhost" + porta)
	log.Fatal(http.ListenAndServe(porta, nil))
}
