package main

import "database/sql"

// region String Queries
const queryConsultaPorZonaPatologia = "SELECT COUNT(*) FROM Paciente WHERE id_zona = ? AND id_patologia = ?"
const queryConsultaPorZonaUF = "SELECT COUNT(*) FROM Paciente WHERE id_zona = ? AND id_uf = ?"
const queryConsultaPorZonaClassificacao = "SELECT COUNT(*) FROM Paciente WHERE id_zona = ? AND id_classificacao = ?"

const queryConsultaPorPatologiaZona = "SELECT COUNT(*) FROM Paciente WHERE id_patologia = ? AND id_zona = ?"
const queryConsultaPorPatologiaUF = "SELECT COUNT(*) FROM Paciente WHERE id_patologia = ? AND id_uf = ?"
const queryConsultaPorPatologiaClassificacao = "SELECT COUNT(*) FROM Paciente WHERE id_patologia = ? AND id_classificacao = ?"

const queryConsultaPorUFZona = "SELECT COUNT(*) FROM Paciente WHERE id_uf = ? AND id_zona = ?"
const queryConsultaPorUFPatologia = "SELECT COUNT(*) FROM Paciente WHERE id_uf = ? AND id_patologia = ?"
const queryConsultaPorUFClassificacao = "SELECT COUNT(*) FROM Paciente WHERE id_uf = ? AND id_classificacao = ?"

const queryConsultaPorClassificacaoZona = "SELECT COUNT(*) FROM Paciente WHERE id_classificacao = ? AND id_zona = ?"
const queryConsultaPorClassificacaoPatologia = "SELECT COUNT(*) FROM Paciente WHERE id_classificacao = ? AND id_patologia = ?"
const queryConsultaPorClassificacaoUF = "SELECT COUNT(*) FROM Paciente WHERE id_classificacao = ? AND id_uf = ?"
// endregion

var insereZonaStmt *sql.Stmt
var listarZonaStmt *sql.Stmt

//var inserePatologiaStmt *sql.Stmt
var listarPatologiaStmt *sql.Stmt

//var insereUFStmt *sql.Stmt
var listarUFStmt *sql.Stmt

//var insereClassificacaoEtariaStmt *sql.Stmt
var listarClassificacaoEtariaStmt *sql.Stmt

//var inserePacienteStmt *sql.Stmt
//var listarPacienteStmt *sql.Stmt

var consultaPacientePorZonaEPatologiaStmt *sql.Stmt
var consultaPacientePorZonaEUFStmt *sql.Stmt
var consultaPacientePorZonaEClassificacaoEtariaStmt *sql.Stmt

var consultaPacientePorPatologiaEZonaStmt *sql.Stmt
var consultaPacientePorPatologiaEUFStmt *sql.Stmt
var consultaPacientePorPatologiaEClassificacaoEtariaStmt *sql.Stmt

var consultaPacientePorUFEZonaStmt *sql.Stmt
var consultaPacientePorUFEPatologiaStmt *sql.Stmt
var consultaPacientePorUFEClassificacaoEtariaStmt *sql.Stmt

var consultaPacientePorClassificacaoEtariaEZonaStmt *sql.Stmt
var consultaPacientePorClassificacaoEtariaEPatologiaStmt *sql.Stmt
var consultaPacientePorClassificacaoEtariaEUFStmt *sql.Stmt

func InitStatements(db *sql.DB) {
	// Prepare statement for inserting data in Zona
	stmtIns, err := db.Prepare("INSERT INTO Zona VALUES( ?, ? )")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	insereZonaStmt = stmtIns

	stmtQuery, err := db.Prepare("SELECT * FROM Zona")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	listarZonaStmt = stmtQuery

	stmtQuery, err = db.Prepare("SELECT * FROM UF")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	listarUFStmt = stmtQuery

	stmtQuery, err = db.Prepare("SELECT * FROM ClassificacaoEtaria")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	listarClassificacaoEtariaStmt = stmtQuery

	stmtQuery, err = db.Prepare("SELECT * FROM Patologia")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	listarPatologiaStmt = stmtQuery

	//Zona e Patologia
	stmtQuery, err = db.Prepare(queryConsultaPorZonaPatologia)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorZonaEPatologiaStmt = stmtQuery

	//Zona e UF
	stmtQuery, err = db.Prepare(queryConsultaPorZonaUF)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorZonaEUFStmt = stmtQuery

	//Zona e Classificacao Etaria
	stmtQuery, err = db.Prepare(queryConsultaPorZonaClassificacao)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorZonaEClassificacaoEtariaStmt = stmtQuery

	//Patologia e Zona
	stmtQuery, err = db.Prepare(queryConsultaPorPatologiaZona)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorPatologiaEZonaStmt = stmtQuery

	//Patologia e UF
	stmtQuery, err = db.Prepare(queryConsultaPorPatologiaUF)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorPatologiaEUFStmt = stmtQuery

	//Patologia e Classificacao
	stmtQuery, err = db.Prepare(queryConsultaPorPatologiaClassificacao)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorPatologiaEClassificacaoEtariaStmt= stmtQuery

	//UF e Zona
	stmtQuery, err = db.Prepare(queryConsultaPorUFZona)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorUFEZonaStmt = stmtQuery

	//UF e Patologia
	stmtQuery, err = db.Prepare(queryConsultaPorUFPatologia)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorUFEPatologiaStmt = stmtQuery

	//UF e Classificacao Etaria
	stmtQuery, err = db.Prepare(queryConsultaPorUFClassificacao)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorUFEClassificacaoEtariaStmt = stmtQuery

	//Classificacao Etaria e Zona
	stmtQuery, err = db.Prepare(queryConsultaPorClassificacaoZona)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorClassificacaoEtariaEZonaStmt = stmtQuery

	//Classificacao Etaria e Patologia
	stmtQuery, err = db.Prepare(queryConsultaPorClassificacaoPatologia)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorClassificacaoEtariaEPatologiaStmt = stmtQuery

	//Classificacao Etaria e UF
	stmtQuery, err = db.Prepare(queryConsultaPorClassificacaoUF)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorClassificacaoEtariaEUFStmt = stmtQuery
}


func CleanStatements() {
	insereZonaStmt.Close()
	listarZonaStmt.Close()

	//inserePatologiaStmt.Close()
	listarPatologiaStmt.Close()

	//insereUFStmt.Close()
	//listarUFStmt.Close()

	//insereClassificacaoEtariaStmt.Close()
	listarClassificacaoEtariaStmt.Close()

	// inserePacienteStmt.Close()
	//listarPacienteStmt.Close()

	consultaPacientePorZonaEPatologiaStmt.Close()
	consultaPacientePorZonaEUFStmt.Close()
	consultaPacientePorUFEClassificacaoEtariaStmt.Close()

	consultaPacientePorPatologiaEZonaStmt.Close()
	consultaPacientePorPatologiaEUFStmt.Close()
	consultaPacientePorPatologiaEClassificacaoEtariaStmt.Close()

	consultaPacientePorUFEZonaStmt.Close()
	consultaPacientePorUFEPatologiaStmt.Close()
	consultaPacientePorUFEClassificacaoEtariaStmt.Close()

	consultaPacientePorClassificacaoEtariaEZonaStmt.Close()
	consultaPacientePorClassificacaoEtariaEPatologiaStmt.Close()
	consultaPacientePorClassificacaoEtariaEUFStmt.Close()

}

var consultaPacientePorClassificacaoEtariaEZonaStmt *sql.Stmt
var consultaPacientePorClassificacaoEtariaEPatologiaStmt *sql.Stmt
var consultaPacientePorClassificacaoEtariaEUFStmt *sql.Stmt

func InitStatements(db *sql.DB) {
	// Prepare statement for inserting data in Zona
	stmtIns, err := db.Prepare("INSERT INTO Zona VALUES( ?, ? )")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	insereZonaStmt = stmtIns

	stmtQuery, err := db.Prepare("SELECT * FROM Zona")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	listarZonaStmt = stmtQuery

	stmtQuery, err = db.Prepare("SELECT * FROM ClassificacaoEtaria")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	listarClassificacaoEtariaStmt = stmtQuery

	//Zona e Patologia
	stmtQuery, err = db.Prepare(queryConsultaPorZonaPatologia)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorZonaEPatologiaStmt = stmtQuery

	//Zona e UF
	stmtQuery, err = db.Prepare(queryConsultaPorZonaUF)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorZonaEUFStmt = stmtQuery

	//Zona e Classificacao Etaria
	stmtQuery, err = db.Prepare(queryConsultaPorZonaClassificacao)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorZonaEClassificacaoEtariaStmt = stmtQuery

	//Patologia e Zona
	stmtQuery, err = db.Prepare(queryConsultaPorPatologiaZona)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorPatologiaEZonaStmt = stmtQuery

	//Patologia e UF
	stmtQuery, err = db.Prepare(queryConsultaPorPatologiaUF)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorPatologiaEUFStmt = stmtQuery

	//Patologia e Classificacao
	stmtQuery, err = db.Prepare(queryConsultaPorPatologiaClassificacao)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorPatologiaEClassificacaoEtariaStmt= stmtQuery

	//UF e Zona
	stmtQuery, err = db.Prepare(queryConsultaPorUFZona)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorUFEZonaStmt = stmtQuery

	//UF e Patologia
	stmtQuery, err = db.Prepare(queryConsultaPorUFPatologia)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorUFEPatologiaStmt = stmtQuery

	//UF e Classificacao Etaria
	stmtQuery, err = db.Prepare(queryConsultaPorUFClassificacao)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorUFEClassificacaoEtariaStmt = stmtQuery

	//Classificacao Etaria e Zona
	stmtQuery, err = db.Prepare(queryConsultaPorClassificacaoZona)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorClassificacaoEtariaEZonaStmt = stmtQuery

	//Classificacao Etaria e Patologia
	stmtQuery, err = db.Prepare(queryConsultaPorClassificacaoPatologia)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorClassificacaoEtariaEPatologiaStmt = stmtQuery

	//Classificacao Etaria e UF
	stmtQuery, err = db.Prepare(queryConsultaPorClassificacaoUF)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	consultaPacientePorClassificacaoEtariaEUFStmt = stmtQuery
}


func CleanStatements() {
	insereZonaStmt.Close()
	listarZonaStmt.Close()

	//inserePatologiaStmt.Close()
	//listarPatologiaStmt.Close()

	//insereUFStmt.Close()
	//listarUFStmt.Close()

	//insereClassificacaoEtariaStmt.Close()
	listarClassificacaoEtariaStmt.Close()

	// inserePacienteStmt.Close()
	//listarPacienteStmt.Close()

	consultaPacientePorZonaEPatologiaStmt.Close()
	consultaPacientePorZonaEUFStmt.Close()
	consultaPacientePorUFEClassificacaoEtariaStmt.Close()

	consultaPacientePorPatologiaEZonaStmt.Close()
	consultaPacientePorPatologiaEUFStmt.Close()
	consultaPacientePorPatologiaEClassificacaoEtariaStmt.Close()

	consultaPacientePorUFEZonaStmt.Close()
	consultaPacientePorUFEPatologiaStmt.Close()
	consultaPacientePorUFEClassificacaoEtariaStmt.Close()

	consultaPacientePorClassificacaoEtariaEZonaStmt.Close()
	consultaPacientePorClassificacaoEtariaEPatologiaStmt.Close()
	consultaPacientePorClassificacaoEtariaEUFStmt.Close()

}

