package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("123456")

type Profissional struct {
	ID                  int    `json:"id"`
	NomeCompleto        string `json:"nome_completo"`
	CPF                 string `json:"cpf"`
	CodigoAutorizacao   string `json:"codigo_autorizacao"`
	Telefone            string `json:"telefone"`
	Email               string `json:"email"`
	CEP                 string `json:"cep"`
	CRMCoreN            string `json:"crm_coren"`
	FuncaoEspecialidade string `json:"funcao_especialidade"`
	Senha               string `json:"senha,omitempty"`
}

type Claims struct {
	CPF string `json:"cpf"`
	jwt.RegisteredClaims
}

func CriarContaHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		var p Profissional
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, "Dados inválidos", http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Senha), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Erro ao gerar senha", http.StatusInternalServerError)
			return
		}

		query := `INSERT INTO profissionais 
			(nome_completo, cpf, codigo_autorizacao, telefone, email, cep, crm_coren, funcao_especialidade, senha_hash)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`
		err = db.QueryRow(query,
			p.NomeCompleto, p.CPF, p.CodigoAutorizacao, p.Telefone, p.Email, p.CEP, p.CRMCoreN, p.FuncaoEspecialidade, string(hashedPassword),
		).Scan(&p.ID)
		if err != nil {
			http.Error(w, "Erro ao salvar usuário: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Conta criada com sucesso"})
	}
}

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		var cred struct {
			CPF   string `json:"cpf"`
			Senha string `json:"senha"`
		}
		err := json.NewDecoder(r.Body).Decode(&cred)
		if err != nil {
			http.Error(w, "Dados inválidos", http.StatusBadRequest)
			return
		}

		var storedHash string
		err = db.QueryRow("SELECT senha_hash FROM profissionais WHERE cpf=$1", cred.CPF).Scan(&storedHash)
		if err != nil {
			http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(cred.Senha))
		if err != nil {
			http.Error(w, "Senha incorreta", http.StatusUnauthorized)
			return
		}

		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &Claims{
			CPF: cred.CPF,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}

func PerfilHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token não fornecido", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		var p Profissional
		err = db.QueryRow(`SELECT id, nome_completo, cpf, codigo_autorizacao, telefone, email, cep, crm_coren, funcao_especialidade
							FROM profissionais WHERE cpf=$1`, claims.CPF).
			Scan(&p.ID, &p.NomeCompleto, &p.CPF, &p.CodigoAutorizacao, &p.Telefone, &p.Email, &p.CEP, &p.CRMCoreN, &p.FuncaoEspecialidade)
		if err != nil {
			http.Error(w, "Usuário não encontrado", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
	}
}
