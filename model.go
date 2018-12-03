package main


type Zona struct {
	id_zona int
	nome string
}

type ClassificacaoEtaria struct{
	id_classificacao int
	faixa_etaria string
	classificacao string
}

type Patologia struct {
	id_patologia int
	nome string
}

type UF	struct {
	id_uf int
	sigla string
	nome string
}

type Paciente struct {
	id_paciente int
	nome string
	id_classificacao int
	sexo string
	endereco string
	id_patologia int
	presenca_patologia bool
	id_uf int
	id_zona int
}