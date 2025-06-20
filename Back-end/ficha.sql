CREATE TABLE IF NOT EXISTS public.informacoes_exame
(
    uf_exame text COLLATE pg_catalog."default",
    cnes_unidade text COLLATE pg_catalog."default",
    unidade_saude text COLLATE pg_catalog."default",
    municipio_exame text COLLATE pg_catalog."default",
    prontuario text COLLATE pg_catalog."default",
    protocolo text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT pk_protocolo PRIMARY KEY (protocolo),
    CONSTRAINT informacoes_exame_protocolo_key UNIQUE (protocolo),
    CONSTRAINT protocolo_unique UNIQUE (protocolo)
)

CREATE TABLE IF NOT EXISTS public.paciente
(
    id integer NOT NULL DEFAULT nextval('paciente_id_seq'::regclass),
    cartao_sus text COLLATE pg_catalog."default",
    nome_mulher text COLLATE pg_catalog."default",
    nome_mae text COLLATE pg_catalog."default",
    apelido text COLLATE pg_catalog."default",
    cpf text COLLATE pg_catalog."default",
    nacionalidade text COLLATE pg_catalog."default",
    data_nascimento text COLLATE pg_catalog."default",
    idade text COLLATE pg_catalog."default",
    raca text COLLATE pg_catalog."default",
    logradouro text COLLATE pg_catalog."default",
    numero text COLLATE pg_catalog."default",
    complemento text COLLATE pg_catalog."default",
    bairro text COLLATE pg_catalog."default",
    uf_paciente text COLLATE pg_catalog."default",
    codigo_municipio text COLLATE pg_catalog."default",
    cep text COLLATE pg_catalog."default",
    municipio_paciente text COLLATE pg_catalog."default",
    ddd text COLLATE pg_catalog."default",
    telefone text COLLATE pg_catalog."default",
    referencia text COLLATE pg_catalog."default",
    escolaridade text COLLATE pg_catalog."default",
    responsavel_preenchimento text COLLATE pg_catalog."default",
    protocolo text COLLATE pg_catalog."default",
    CONSTRAINT paciente_pkey PRIMARY KEY (id),
    CONSTRAINT fk_paciente_protocolo FOREIGN KEY (protocolo)
        REFERENCES public.informacoes_exame (protocolo) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT fk_protocolo_paciente FOREIGN KEY (protocolo)
        REFERENCES public.informacoes_exame (protocolo) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
)

CREATE TABLE IF NOT EXISTS public.anamnese
(
    id integer NOT NULL DEFAULT nextval('anamnese_id_seq'::regclass),
    motivo text COLLATE pg_catalog."default",
    realizacao text COLLATE pg_catalog."default",
    gravida text COLLATE pg_catalog."default",
    anticoncepcional text COLLATE pg_catalog."default",
    tratamento_hormonal text COLLATE pg_catalog."default",
    data_ultima_mestruacao text COLLATE pg_catalog."default",
    sangramentors text COLLATE pg_catalog."default",
    sangramentom text COLLATE pg_catalog."default",
    responsavel_anamnese text COLLATE pg_catalog."default",
    protocolo text COLLATE pg_catalog."default",
    data_ultimo_exame text COLLATE pg_catalog."default",
    diu text COLLATE pg_catalog."default",
    usa_hormonio_menopausa text COLLATE pg_catalog."default",
    CONSTRAINT anamnese_pkey PRIMARY KEY (id),
    CONSTRAINT fk_anamnese_protocolo FOREIGN KEY (protocolo)
        REFERENCES public.informacoes_exame (protocolo) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT fk_protocolo_anamnese FOREIGN KEY (protocolo)
        REFERENCES public.informacoes_exame (protocolo) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
)

CREATE TABLE IF NOT EXISTS public.exame_clinico
(
    id integer NOT NULL DEFAULT nextval('exame_clinico_id_seq'::regclass),
    inspecao text COLLATE pg_catalog."default",
    dst text COLLATE pg_catalog."default",
    data_resultado text COLLATE pg_catalog."default",
    responsavel_coleta text COLLATE pg_catalog."default",
    protocolo text COLLATE pg_catalog."default",
    CONSTRAINT exame_clinico_pkey PRIMARY KEY (id),
    CONSTRAINT fk_exame_clinico_protocolo FOREIGN KEY (protocolo)
        REFERENCES public.informacoes_exame (protocolo) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT fk_protocolo_exame_clinico FOREIGN KEY (protocolo)
        REFERENCES public.informacoes_exame (protocolo) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
)

CREATE TABLE IF NOT EXISTS public.exame_laboratorial
(
    id integer NOT NULL DEFAULT nextval('exame_laboratorial_id_seq'::regclass),
    amostrarejeitada text COLLATE pg_catalog."default",
    epitelio text COLLATE pg_catalog."default",
    adequabilidade text COLLATE pg_catalog."default",
    dentro_limites text COLLATE pg_catalog."default",
    alteracoes text COLLATE pg_catalog."default",
    microbiologia text COLLATE pg_catalog."default",
    celulas_atipicas text COLLATE pg_catalog."default",
    atipias_escamosas text COLLATE pg_catalog."default",
    atipias_glandulares text COLLATE pg_catalog."default",
    neoplasias text COLLATE pg_catalog."default",
    endometriais text COLLATE pg_catalog."default",
    observacoes text COLLATE pg_catalog."default",
    responsavel_laboratorial text COLLATE pg_catalog."default",
    protocolo text COLLATE pg_catalog."default",
    data_resultado_coleta text COLLATE pg_catalog."default",
    CONSTRAINT exame_laboratorial_pkey PRIMARY KEY (id),
    CONSTRAINT fk_exame_laboratorial_protocolo FOREIGN KEY (protocolo)
        REFERENCES public.informacoes_exame (protocolo) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT fk_protocolo_exame_laboratorial FOREIGN KEY (protocolo)
        REFERENCES public.informacoes_exame (protocolo) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);
