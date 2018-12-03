package main

func InsereZona(id int, nome string) {
	_, err := insereZonaStmt.Exec(id, nome)
	if err != nil {
		panic(err)
	}
}

func ListarZona() []Zona {
	var zonas = make([]Zona, 0)
	v, err := listarZonaStmt.Query()
	if err != nil {
		panic(err)
	}

	for v.Next() {
		var zona Zona
		err = v.Scan(&zona.Id_zona, &zona.Nome)
		if err != nil {
			panic(err)
		}
		zonas = append(zonas, zona)
	}

	return zonas
}

//func InsereClassificacaoEtaria(id int, faixa_etaria string, classificacao string) {
//	_, err := insereClassificacaoEtariaStmt.Exec(id, faixa_etaria, classificacao)
//	if err != nil {
//		panic(err)
//	}
//}

func ListarClassificacaoEtaria() []ClassificacaoEtaria {
	var classificacao_etaria = make([]ClassificacaoEtaria, 0)
	v, err := listarClassificacaoEtariaStmt.Query()
	if err != nil {
		panic(err)
	}

	for v.Next() {
		var classificacaoEtaria ClassificacaoEtaria
		err = v.Scan(&classificacaoEtaria.Id_classificacao, &classificacaoEtaria.Faixa_etaria, &classificacaoEtaria.Classificacao)
		if err != nil {
			panic(err)
		}
		classificacao_etaria = append(classificacao_etaria, classificacaoEtaria)
	}

	return classificacao_etaria
}


func ConstroiClassificacaoEtariaPorZona() []ZonaClassificacoes {
	var zonas = ListarZona()
	var classificacoes = ListarClassificacaoEtaria()

	var dados = make([]ZonaClassificacoes, len(zonas) + 1)

	// region Criar Total
	var total = ZonaClassificacoes{
		Zona: Zona{
			Id_zona: -1,
			Nome: "TOTAL",
		},
		Classificacoes: make([]ZonaClassificacaoContagem, len(classificacoes)),
	}
	for j := 0; j < len(classificacoes); j++ {
		var classificacao = classificacoes[j]
		total.Classificacoes[j] = ZonaClassificacaoContagem{
			ClassificacaoEtaria: classificacao,
			Contagem: 0,
		}
	}
	// endregion
	// region Construir Classificações
	for i := 0; i < len(zonas); i++ {
		var zona = zonas[i]
		var dado = ZonaClassificacoes{
			Zona: zona,
			Classificacoes: make([]ZonaClassificacaoContagem, len(classificacoes)),
		}
		for j := 0; j < len(classificacoes); j++ {
			var classificacao = classificacoes[j]
			dado.Classificacoes[j] = ZonaClassificacaoContagem{
				ClassificacaoEtaria: classificacao,
				Contagem: consultaPacientePorClassificacaoEtariaEZona(classificacao.Id_classificacao, zona.Id_zona),
			}

			total.Classificacoes[j].Contagem += dado.Classificacoes[j].Contagem
		}
		dados[i+1] = dado
	}
	// endregion

	dados[0] = total

	return dados
}

//func InserePatologia(id int, nome string) {
//	_, err := inserePatologiaStmt.Exec(id, nome)
//	if err != nil {
//		panic(err)
//	}
//}

//func ListarPatologia() []Patologia {
//	var patologia = make([]Patologia, 0)
//	v, err := listarPatologiaStmt.Query()
//	if err != nil {
//		panic(err)
//	}
//
//	for v.Next() {
//		var patologias Patologia
//		err = v.Scan(&patologias.id_patologia, &patologias.nome)
//		if err != nil {
//			panic(err)
//		}
//		patologia = append(patologia, patologias)
//	}
//
//	return patologia
//}

//func InsereUF(id int, sigla string, nome string) {
//	_, err := inserePatologiaStmt.Exec(id, sigla, nome)
//	if err != nil {
//		panic(err)
//	}
//}

//func ListarUF() []UF {
//	var uf = make([]UF, 0)
//	v, err := listarUFStmt.Query()
//	if err != nil {
//		panic(err)
//	}
//
//	for v.Next() {
//		var estado UF
//		err = v.Scan(&estado.id_uf, &estado.sigla, &estado.nome)
//		if err != nil {
//			panic(err)
//		}
//		uf = append(uf, estado)
//	}
//
//	return uf
//}

//func InserePaciente(id int, nome string, id_classificacao int, sexo string, endereco string, id_patologia int, presenca_patologia bool, id_uf int, id_zona int) {
//	_, err := inserePatologiaStmt.Exec(id, nome, id_classificacao, sexo, endereco, id_patologia, presenca_patologia, id_uf, id_zona)
//	if err != nil {
//		panic(err)
//	}
//}

//func ListarPaciente() []Paciente {
//	var pacientes = make([]Paciente, 0)
//	v, err := listarPacienteStmt.Query()
//	if err != nil {
//		panic(err)
//	}
//
//	for v.Next() {
//		var paciente Paciente
//		err = v.Scan(&paciente.id_paciente, &paciente.nome, &paciente.id_classificacao, &paciente.sexo, &paciente.endereco, &paciente.id_patologia, &paciente.presenca_patologia, &paciente.id_uf, &paciente.id_zona)
//		if err != nil {
//			panic(err)
//		}
//		pacientes = append(pacientes, paciente)
//	}
//
//	return pacientes
//}


func consultaPacientePorZonaEPatologia(idZona, idPatologia int) int {
	var contagem int
	v, err := consultaPacientePorZonaEPatologiaStmt.Query(idZona, idPatologia)
	if err != nil {
		panic (err)
	}

	v.Next()

	err = v.Scan(&contagem)
	if err != nil{
		panic(err)
	}

	return contagem
}

func consultaPacientePorZonaEUF(idZona int, idUF int) int {
	var contagem int
	v, err := consultaPacientePorZonaEUFStmt.Query(idZona, idUF)
	if err != nil {
		panic (err)
	}

	v.Next()

	err = v.Scan(&contagem)
	if err != nil{
		panic(err)
	}

	return contagem
}

func consultaPacientePorZonaEClassificacaoEtaria(idZona int, idClassificacao int) int {
	var contagem int
	v, err := consultaPacientePorZonaEClassificacaoEtariaStmt.Query(idZona, idClassificacao)
	if err != nil {
		panic (err)
	}

	v.Next()

	err = v.Scan(&contagem)
	if err != nil{
		panic(err)
	}

	return contagem
}

func consultaPacientePorPatologiaEZona(idPatologia int, idZona int) int {
	var contagem int
	v, err := consultaPacientePorPatologiaEZonaStmt.Query(idPatologia, idZona)
	if err != nil {
		panic (err)
	}

	v.Next()

	err = v.Scan(&contagem)
	if err != nil{
		panic(err)
	}

	return contagem
}

func consultaPacientePorPatologiaEUF(idPatologia int, idUF int) int {
	var contagem int
	v, err := consultaPacientePorPatologiaEUFStmt.Query(idPatologia, idUF)
	if err != nil {
		panic (err)
	}

	v.Next()

	err = v.Scan(&contagem)
	if err != nil{
		panic(err)
	}

	return contagem
}

func consultaPacientePorPatologiaEClassificacaoEtaria(idPatologia int, idClassificacao int) int {
	var contagem int
	v, err := consultaPacientePorPatologiaEClassificacaoEtariaStmt.Query(idPatologia, idClassificacao)
	if err != nil {
		panic (err)
	}

	v.Next()

	err = v.Scan(&contagem)
	if err != nil{
		panic(err)
	}

	return contagem
}

func consultaPacientePorUFEZona(idUF int, idZona int) int {
	var contagem int
	v, err := consultaPacientePorUFEZonaStmt.Query(idUF, idZona)
	if err != nil {
		panic (err)
	}

	v.Next()

	err = v.Scan(&contagem)
	if err != nil{
		panic(err)
	}

	return contagem
}

func consultaPacientePorUFEPatologia(idUF int, idPatologia int) int {
	var contagem int
	v, err := consultaPacientePorUFEPatologiaStmt.Query(idUF, idPatologia)
	if err != nil {
		panic (err)
	}

	v.Next()

	err = v.Scan(&contagem)
	if err != nil{
		panic(err)
	}

	return contagem
}

func consultaPacientePorUFEClassificacaoEtaria(idUF int, idClassificacao int) int {
	var contagem int
	v, err := consultaPacientePorUFEClassificacaoEtariaStmt.Query(idUF, idClassificacao)
	if err != nil {
		panic (err)
	}

	v.Next()

	err = v.Scan(&contagem)
	if err != nil{
		panic(err)
	}

	return contagem
}

func consultaPacientePorClassificacaoEtariaEZona(idClassificacao int, idZona int) int {
	var contagem int
	v, err := consultaPacientePorClassificacaoEtariaEZonaStmt.Query(idClassificacao, idZona)
	if err != nil {
		panic (err)
	}

	v.Next()

	err = v.Scan(&contagem)
	if err != nil{
		panic(err)
	}

	return contagem
}

func consultaPacientePorClassificacaoEtariaEPatologia(idClassificacao int, idPatologia int) int {
	var contagem int
	v, err := consultaPacientePorClassificacaoEtariaEPatologiaStmt.Query(idClassificacao, idPatologia)
	if err != nil {
		panic (err)
	}

	v.Next()

	err = v.Scan(&contagem)
	if err != nil{
		panic(err)
	}

	return contagem
}

func consultaPacientePorClassificacaoEtariaEUF(idClassificacao int, idUF int) int {
	var contagem int
	v, err := consultaPacientePorClassificacaoEtariaEUFStmt.Query(idClassificacao, idUF)
	if err != nil {
		panic (err)
	}

	v.Next()

	err = v.Scan(&contagem)
	if err != nil{
		panic(err)
	}

	return contagem
}
