package handlers

import (
	"database/sql"
	"log"
	"net/http"
)

func InformacoesExame(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()

		log.Println("Valores recebidos do formulário:")
		for key, value := range r.Form {
			log.Printf("  %s: %v\n", key, value)
		}

		if err != nil {
			log.Println("Erro ao processar formulário:", err)
			http.Error(w, "Erro ao processar dados", http.StatusBadRequest)
			return
		}

		_, err = db.Exec(`INSERT INTO informacoes_exame (
			uf_exame, cnes_unidade, unidade_saude, municipio_exame, prontuario, protocolo
		) VALUES ($1, $2, $3, $4, $5, $6)`,
			r.FormValue("uf_exame"),
			r.FormValue("cnes_unidade"),
			r.FormValue("unidade_saude"),
			r.FormValue("municipio_exame"),
			r.FormValue("prontuario"),
			r.FormValue("protocolo"),
		)

		if err != nil {
			log.Println("Erro ao salvar dados no banco:", err)
			http.Error(w, "Erro ao salvar os dados", http.StatusInternalServerError)
			return
		}

		log.Println("Dados salvos com sucesso!")
		http.Redirect(w, r, "/static/LoginProfissional/PaginaInicial/PreenchimentoFichaCitopatologica/informacoes-pessoais.html", http.StatusSeeOther)
	}
}

func InformacoesPessoais(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()

		log.Println("Valores recebidos do formulário:")
		for key, value := range r.Form {
			log.Printf("  %s: %v\n", key, value)
		}

		if err != nil {
			log.Println("Erro ao salvar dados:", err)
			http.Error(w, "Erro ao salvar os dados", http.StatusInternalServerError)
			return
		}

		protocolo := r.FormValue("protocolo")

		_, err = db.Exec(`INSERT INTO paciente  (
            protocolo, cartao_sus, nome_mulher, nome_mae, apelido, cpf, nacionalidade, data_nascimento,
            idade, raca, logradouro, numero, complemento, bairro, uf_paciente,
            codigo_municipio, cep, municipio_paciente, ddd, telefone, referencia,
            escolaridade, responsavel_preenchimento
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
                  $11, $12, $13, $14, $15, $16, $17, $18, $19,
                  $20, $21, $22, $23)`,
			protocolo,
			r.FormValue("cartao_sus"),
			r.FormValue("nome_mulher"),
			r.FormValue("nome_mae"),
			r.FormValue("apelido"),
			r.FormValue("cpf"),
			r.FormValue("nacionalidade"),
			r.FormValue("data_nascimento"),
			r.FormValue("idade"),
			r.FormValue("raca"),
			r.FormValue("logradouro"),
			r.FormValue("numero"),
			r.FormValue("complemento"),
			r.FormValue("bairro"),
			r.FormValue("uf_paciente"),
			r.FormValue("codigo_municipio"),
			r.FormValue("cep"),
			r.FormValue("municipio_paciente"),
			r.FormValue("ddd"),
			r.FormValue("telefone"),
			r.FormValue("referencia"),
			r.FormValue("escolaridade"),
			r.FormValue("responsavel_preenchimento"),
		)

		if err != nil {
			log.Println("Erro ao salvar dados:", err)
			http.Error(w, "Erro ao salvar os dados", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/static/LoginProfissional/PaginaInicial/PreenchimentoFichaCitopatologica/dados-anamnese.html", http.StatusSeeOther)
	}
}

func Anamnese(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()

		log.Println("Valores recebidos do formulário:")
		for key, value := range r.Form {
			log.Printf("  %s: %v\n", key, value)
		}

		if err != nil {
			log.Println("Erro ao salvar dados:", err)
			http.Error(w, "Erro ao salvar os dados", http.StatusInternalServerError)
			return
		}

		protocolo := r.FormValue("protocolo")

		_, err = db.Exec(`INSERT INTO anamnese (
            protocolo, motivo, realizacao, data_ultimo_exame, diu, gravida, anticoncepcional, usa_hormonio_menopausa,
			tratamento_hormonal, data_ultima_mestruacao, sangramentors, sangramentom, responsavel_anamnese
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,
			protocolo,
			r.FormValue("motivo"),
			r.FormValue("realizacao"),
			r.FormValue("data_ultimo_exame"),
			r.FormValue("diu"),
			r.FormValue("gravida"),
			r.FormValue("anticoncepcional"),
			r.FormValue("usa_hormonio_menopausa"),
			r.FormValue("tratamento_hormonal"),
			r.FormValue("data_ultima_mestruacao"),
			r.FormValue("sangramentors"),
			r.FormValue("sangramentom"),
			r.FormValue("responsavel_anamnese"),
		)

		if err != nil {
			log.Println("Erro ao salvar dados:", err)
			http.Error(w, "Erro ao salvar os dados", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/static/LoginProfissional/PaginaInicial/PreenchimentoFichaCitopatologica/exame-clinico.html", http.StatusSeeOther)
	}
}

func Exame_clinico(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()

		log.Println("Valores recebidos do formulário:")
		for key, value := range r.Form {
			log.Printf("  %s: %v\n", key, value)
		}

		if err != nil {
			log.Println("Erro ao salvar dados:", err)
			http.Error(w, "Erro ao salvar os dados", http.StatusInternalServerError)
			return
		}

		protocolo := r.FormValue("protocolo")

		_, err = db.Exec(`INSERT INTO exame_clinico (
            protocolo, inspecao, DST, data_resultado, responsavel_coleta
        ) VALUES ($1, $2, $3, $4, $5)`,
			protocolo,
			r.FormValue("inspecao"),
			r.FormValue("DST"),
			r.FormValue("data_resultado"),
			r.FormValue("responsavel_coleta"),
		)

		if err != nil {
			log.Println("Erro ao salvar dados:", err)
			http.Error(w, "Erro ao salvar os dados", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/static/LoginProfissional/PaginaInicial/PreenchimentoFichaCitopatologica/exame-laboratorio.html", http.StatusSeeOther)
	}
}

func Exame_laboratorial(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()

		log.Println("Valores recebidos do formulário:")
		for key, value := range r.Form {
			log.Printf("  %s: %v\n", key, value)
		}

		if err != nil {
			log.Println("Erro ao salvar dados:", err)
			http.Error(w, "Erro ao salvar os dados", http.StatusInternalServerError)
			return
		}

		protocolo := r.FormValue("protocolo")

		_, err = db.Exec(`INSERT INTO exame_laboratorial (
            protocolo, amostraRejeitada, epitelio, adequabilidade, dentro_limites, alteracoes, microbiologia, 
            celulas_atipicas, atipias_escamosas, atipias_glandulares, neoplasias, endometriais,
            observacoes, data_resultado_coleta, responsavel_laboratorial
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`,
			protocolo,
			r.FormValue("amostraRejeitada"),
			r.FormValue("epitelio"),
			r.FormValue("adequabilidade"),
			r.FormValue("dentro_limites"),
			r.FormValue("alteracoes"),
			r.FormValue("microbiologia"),
			r.FormValue("celulas_atipicas"),
			r.FormValue("atipias_escamosas"),
			r.FormValue("atipias_glandulares"),
			r.FormValue("neoplasias"),
			r.FormValue("endometriais"),
			r.FormValue("observacoes"),
			r.FormValue("data_resultado_coleta"),
			r.FormValue("responsavel_laboratorial"),
		)

		if err != nil {
			log.Println("Erro ao salvar dados:", err)
			http.Error(w, "Erro ao salvar os dados", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/static/LoginProfissional/PaginaInicial/pagina-inicial.html", http.StatusSeeOther)
	}
}
