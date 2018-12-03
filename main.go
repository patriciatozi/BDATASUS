package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@/formularioEletronico")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	db := ConnectDatabase()
	defer db.Close()

	InitStatements(db)
	defer CleanStatements()

	//log.Println("Inserindo Zona")
	//InsereZona(0, 1200, "HUEBR")
	//log.Println("Zona Inserida")

	//var zonas = ListarZonas()

	//log.Println(zonas)

	var contagem1 = consultaPacientePorZonaEPatologiaStmt(0, 0)
	var contagem2 = consultaPacientePorZonaEUFStmt(0, 0)
	var contagem3 = consultaPacientePorZonaEClassificacaoEtariaStmt(0, 0)

	var contagem4 = consultaPacientePorPatologiaEZonaStmt(0, 0)
	var contagem5 = consultaPacientePorPatologiaEUFStmt(0, 0)
	var contagem6 = consultaPacientePorPatologiaEClassificacaoEtariaStmt(0, 0)

	var contagem7 = consultaPacientePorUFEZonaStmt(0, 0)
	var contagem8 = consultaPacientePorUFEPatologiaStmt(0, 0)
	var contagem9 = consultaPacientePorUFEClassificacaoEtariaStmt(0, 0)

	var contagem10 = consultaPacientePorClassificacaoEtariaEZonaStmt(0, 0)
	var contagem11 = consultaPacientePorClassificacaoEtariaEPatologiaStmt(0, 0)
	var contagem12 = consultaPacientePorClassificacaoEtariaEUFStmt(0, 0)

	log.Println(contagem1)
}