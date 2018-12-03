package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
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

	var contagem1 = consultaPacientePorZonaEPatologia(1, 510)
	var contagem2 = consultaPacientePorZonaEUF(0, 0)
	var contagem3 = consultaPacientePorZonaEClassificacaoEtaria(0, 0)

	var contagem4 = consultaPacientePorPatologiaEZona(0, 0)
	var contagem5 = consultaPacientePorPatologiaEUF(0, 0)
	var contagem6 = consultaPacientePorPatologiaEClassificacaoEtaria(0, 0)

	var contagem7 = consultaPacientePorUFEZona(0, 0)
	var contagem8 = consultaPacientePorUFEPatologia(0, 0)
	var contagem9 = consultaPacientePorUFEClassificacaoEtaria(0, 0)

	var contagem10 = consultaPacientePorClassificacaoEtariaEZona(0, 0)
	var contagem11 = consultaPacientePorClassificacaoEtariaEPatologia(0, 0)
	var contagem12 = consultaPacientePorClassificacaoEtariaEUF(0, 0)

	log.Println(contagem1)
	log.Println(contagem2)
	log.Println(contagem3)
	log.Println(contagem4)
	log.Println(contagem5)
	log.Println(contagem6)
	log.Println(contagem7)
	log.Println(contagem8)
	log.Println(contagem9)
	log.Println(contagem10)
	log.Println(contagem11)
	log.Println(contagem12)
}
