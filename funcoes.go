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
		err = v.Scan(&zona.id_zona, &zona.nome)
		if err != nil {
			panic(err)
		}
		zonas = append(zonas, zona)
	}


	return zonas
}

func InsereClassificacaoEtaria(id int, faixa_etaria string, classificacao string) {
	_, err := insereClassificacaoEtariaStmt.Exec(id, faixa_etaria, classificacao)
	if err != nil {
		panic(err)
	}
}

func ListarClassificacaoEtaria() []ClassificacaoEtaria {
	var classificacao_etaria = make([]ClassificacaoEtaria, 0)
	v, err := listarClassificacaoEtariaStmt.Query()
	if err != nil {
		panic(err)
	}

	for v.Next() {
		var classificacaoEtaria ClassificacaoEtaria
		err = v.Scan(&classificacaoEtaria.id_classificacao, &classificacaoEtaria.faixa_etaria, &classificacaoEtaria.classificacao)
		if err != nil {
			panic(err)
		}
		classificacao_etaria = append(classificacao_etaria, classificacaoEtaria)
	}


	return classificacao_etaria
}

func InserePatologia(id int, nome string) {
	_, err := inserePatologiaStmt.Exec(id, nome)
	if err != nil {
		panic(err)
	}
}

	func ListarPatologia() []Patologia {
	var patologia = make([]Patologia, 0)
	v, err := listarPatologiaStmt.Query()
	if err != nil {
		panic(err)
	}

	for v.Next() {
		var patologias Patologia
		err = v.Scan(&patologias.id_patologia, &patologias.nome)
		if err != nil {
			panic(err)
		}
		patologia = append(patologia, patologias)
	}


	return patologia
}

func InsereUF(id int, sigla string, nome string) {
	_, err := inserePatologiaStmt.Exec(id, sigla, nome)
	if err != nil {
		panic(err)
	}
}

func ListarUF() []UF {
	var uf = make([]UF, 0)
	v, err := listarUFStmt.Query()
	if err != nil {
		panic(err)
	}

	for v.Next() {
		var estado UF
		err = v.Scan(&estado.id_uf, &estado.sigla, &estado.nome)
		if err != nil {
			panic(err)
		}
		uf = append(uf, estado)
	}


	return uf
}

func InserePaciente(id int, nome string, id_classificacao int, sexo string, endereco string, id_patologia int, presenca_patologia bool, id_uf int, id_zona int) {
	_, err := inserePatologiaStmt.Exec(id, nome, id_classificacao, sexo, endereco, id_patologia, presenca_patologia, id_uf, id_zona)
	if err != nil {
		panic(err)
	}
}

func ListarPaciente() []Paciente {
	var pacientes = make([]Paciente, 0)
	v, err := listarPacienteStmt.Query()
	if err != nil {
		panic(err)
	}

	for v.Next() {
		var paciente Paciente
		err = v.Scan(&paciente.id_paciente, &paciente.nome, &paciente.id_classificacao, &paciente.sexo, &paciente.endereco,&paciente.id_patologia, &paciente.presenca_patologia, &paciente.id_uf, &paciente.id_zona)
		if err != nil {
			panic(err)
		}
		pacientes = append(pacientes, paciente)
	}


	return pacientes
}