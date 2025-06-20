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
	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao testar conexão com o banco:", err)
	}

	http.HandleFunc("/ficha/informacoes-exame", handlers.InformacoesExame(db))
	http.HandleFunc("/ficha/informacoes-pessoais", handlers.InformacoesPessoais(db))
	http.HandleFunc("/ficha/dados-anamnese", handlers.Anamnese(db))
	http.HandleFunc("/ficha/exame-clinico", handlers.Exame_clinico(db))
	http.HandleFunc("/ficha/exame-laboratorial", handlers.Exame_laboratorial(db))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	porta := ":8080"
	log.Println("Servidor rodando em http://localhost" + porta)
	log.Fatal(http.ListenAndServe(porta, nil))

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/Front-end/escolha-login.html", http.StatusSeeOther)
	})

}
