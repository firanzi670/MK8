package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Paciente struct {
	CartaoSUS       string `json:"cartao_sus"`
	NomeMulher      string `json:"nome_mulher"`
	NomeMae         string `json:"nome_mae"`
	Apelido         string `json:"apelido"`
	CPF             string `json:"cpf"`
	Nacionalidade   string `json:"nacionalidade"`
	DataNascimento  string `json:"data_nascimento"`
	Idade           string `json:"idade"`
	Raca            string `json:"raca"`
	Logradouro      string `json:"logradouro"`
	Numero          string `json:"numero"`
	Complemento     string `json:"complemento"`
	Bairro          string `json:"bairro"`
	UFPaciente      string `json:"uf_paciente"`
	CodigoMunicipio string `json:"codigo_municipio"`
	CEP             string `json:"cep"`
	Municipio       string `json:"municipio_paciente"`
	DDD             string `json:"ddd"`
	Telefone        string `json:"telefone"`
	Referencia      string `json:"referencia"`
	Escolaridade    string `json:"escolaridade"`
}

func BuscarPacienteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		termo := r.URL.Query().Get("termo")
		if termo == "" {
			http.Error(w, "Termo de busca obrigatório", http.StatusBadRequest)
			return
		}

		var paciente Paciente

		query := `
            SELECT cartao_sus, nome_mulher, nome_mae, apelido, cpf, nacionalidade,
                   data_nascimento, idade, raca, logradouro, numero, complemento,
                   bairro, uf_paciente, codigo_municipio, cep, municipio_paciente,
                   ddd, telefone, referencia, escolaridade
            FROM paciente
            WHERE nome_mulher ILIKE $1 OR cpf = $2 OR cartao_sus = $3
            LIMIT 1
        `

		termoNome := "%" + termo + "%"

		err := db.QueryRow(query, termoNome, termo, termo).Scan(
			&paciente.CartaoSUS,
			&paciente.NomeMulher,
			&paciente.NomeMae,
			&paciente.Apelido,
			&paciente.CPF,
			&paciente.Nacionalidade,
			&paciente.DataNascimento,
			&paciente.Idade,
			&paciente.Raca,
			&paciente.Logradouro,
			&paciente.Numero,
			&paciente.Complemento,
			&paciente.Bairro,
			&paciente.UFPaciente,
			&paciente.CodigoMunicipio,
			&paciente.CEP,
			&paciente.Municipio,
			&paciente.DDD,
			&paciente.Telefone,
			&paciente.Referencia,
			&paciente.Escolaridade,
		)

		if err != nil {
			http.Error(w, "Paciente não encontrado", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(paciente)
	}
}
