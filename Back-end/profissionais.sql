CREATE TABLE IF NOT EXISTS public-profissionais (
  id SERIAL PRIMARY KEY, 
  nome TEXT NOT NULL,
  cpf TEXT UNIQUE NOT NULL, 
  codigo_autorizacao TEXT NOT NULL, 
  telefone TEXT,
  email TEXT UNIQUE NOT NULL, 
  cep TEXT, 
  crm_coren TEXT, 
  funcao TEXT, 
  senha TEXT NOT NULL, 
  criado_em TIMESTAMP DEFAULT CURRENT_ TIMPESTAMP
)
