package main


type Zona struct {
	Id_zona int
	Nome string
}

type ZonaPatologiaContagem struct {
	Patologia Patologia
	Contagem int
}

type ZonaPatologias struct {
	Zona Zona
	Patologias []ZonaPatologiaContagem
}

type ZonaUFContagem struct {
	UF UF
	Contagem int
}

type ZonaUFS struct {
	Zona Zona
	UFS []ZonaUFContagem
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
	Id_patologia int
	Nome string
}

type UF	struct {
	Id_uf int
	Sigla string
	Nome string
}

type Paciente struct {
	Id_paciente int
	Nome string
	Id_classificacao int
	Sexo string
	Endereco string
	Id_patologia int
	Presenca_patologia bool
	Id_uf int
	Id_zona int
}
