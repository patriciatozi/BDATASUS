package main


type Zona struct {
	Id_zona int
	Nome string
}

type ZonaClassificacaoContagem struct {
	ClassificacaoEtaria ClassificacaoEtaria
	Contagem int
}

type ZonaClassificacoes struct {
	Zona Zona
	Classificacoes []ZonaClassificacaoContagem
}

type ClassificacaoEtaria struct{
	Id_classificacao int
	Faixa_etaria string
	Classificacao string
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