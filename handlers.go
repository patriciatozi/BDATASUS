package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func consultaPacientePorZonaEPatologiaIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	q := r.URL.Query()

	idZonaString := q.Get("idZona")
	idPatologiaString := q.Get("idPatologia")

	idZona, _ := strconv.ParseInt(idZonaString, 10,32)
	idPatologia, _ := strconv.ParseInt(idPatologiaString, 10,32)

	contagem := consultaPacientePorZonaEPatologia(int(idZona), int(idPatologia))

	retorno := make(map[string]int)

	retorno["contagem"] = contagem

	x, err := json.MarshalIndent(retorno," ","    ")
	if err != nil {
		panic(err)
	}

	w.Write([]byte(x))
}

func consultaPacientePorZonaEUFIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	q := r.URL.Query()

	idZonaString := q.Get("idZona")
	idUFString := q.Get("idUF")

	idZona, _ := strconv.ParseInt(idZonaString, 10,32)
	idUF, _ := strconv.ParseInt(idUFString, 10,32)

	contagem := consultaPacientePorZonaEUF(int(idZona), int(idUF))

	retorno := make(map[string]int)

	retorno["contagem"] = contagem

	x, err := json.MarshalIndent(retorno," ","    ")
	if err != nil {
		panic(err)
	}

	w.Write([]byte(x))
}

func consultaPacientePorZonaEClassificacaoEtariaIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	q := r.URL.Query()

	idZonaString := q.Get("idZona")
	idClassificacaoString := q.Get("idClassificacao")

	idZona, _ := strconv.ParseInt(idZonaString, 10,32)
	idClassificacao, _ := strconv.ParseInt(idClassificacaoString, 10,32)

	contagem := consultaPacientePorZonaEClassificacaoEtaria(int(idZona), int(idClassificacao))

	retorno := make(map[string]int)

	retorno["contagem"] = contagem

	x, err := json.MarshalIndent(retorno," ","    ")
	if err != nil {
		panic(err)
	}

	w.Write([]byte(x))
}

func consultaPacientePorPatologiaEZonaIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	q := r.URL.Query()

	idPatologiaString := q.Get("idPatologia")
	idZonaString := q.Get("idZona")

	idPatologia, _ := strconv.ParseInt(idPatologiaString, 10,32)
	idZona, _ := strconv.ParseInt(idZonaString, 10,32)

	contagem := consultaPacientePorPatologiaEZona(int(idPatologia), int(idZona))

	retorno := make(map[string]int)

	retorno["contagem"] = contagem

	x, err := json.MarshalIndent(retorno," ","    ")
	if err != nil {
		panic(err)
	}

	w.Write([]byte(x))
}

func consultaPacientePorPatologiaEUFIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	q := r.URL.Query()

	idPatologiaString := q.Get("idPatologia")
	idUFString := q.Get("idUF")

	idPatologia, _ := strconv.ParseInt(idPatologiaString, 10,32)
	idUF, _ := strconv.ParseInt(idUFString, 10,32)

	contagem := consultaPacientePorPatologiaEUF(int(idPatologia), int(idUF))

	retorno := make(map[string]int)

	retorno["contagem"] = contagem

	x, err := json.MarshalIndent(retorno," ","    ")
	if err != nil {
		panic(err)
	}

	w.Write([]byte(x))
}

func consultaPacientePorPatologiaEClassificacaoEtariaIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	q := r.URL.Query()

	idPatologiaString := q.Get("idPatologia")
	idClassificacaoString := q.Get("idClassificacao")

	idPatologia, _ := strconv.ParseInt(idPatologiaString, 10,32)
	idClassificacao, _ := strconv.ParseInt(idClassificacaoString, 10,32)

	contagem := consultaPacientePorPatologiaEClassificacaoEtaria(int(idPatologia), int(idClassificacao))

	retorno := make(map[string]int)

	retorno["contagem"] = contagem

	x, err := json.MarshalIndent(retorno," ","    ")
	if err != nil {
		panic(err)
	}

	w.Write([]byte(x))
}

func consultaPacientePorUFEZonaIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	q := r.URL.Query()

	idUFString := q.Get("idUF")
	idZonaString := q.Get("idZona")

	idUF, _ := strconv.ParseInt(idUFString, 10,32)
	idZona, _ := strconv.ParseInt(idZonaString, 10,32)

	contagem := consultaPacientePorUFEZona(int(idUF), int(idZona))

	retorno := make(map[string]int)

	retorno["contagem"] = contagem

	x, err := json.MarshalIndent(retorno," ","    ")
	if err != nil {
		panic(err)
	}

	w.Write([]byte(x))
}

func consultaPacientePorUFEPatologiaIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	q := r.URL.Query()

	idUFString := q.Get("idUF")
	idPatologiaString := q.Get("idPatologia")

	idUF, _ := strconv.ParseInt(idUFString, 10,32)
	idPatologia, _ := strconv.ParseInt(idPatologiaString, 10,32)

	contagem := consultaPacientePorUFEPatologia(int(idUF), int(idPatologia))

	retorno := make(map[string]int)

	retorno["contagem"] = contagem

	x, err := json.MarshalIndent(retorno," ","    ")
	if err != nil {
		panic(err)
	}

	w.Write([]byte(x))
}

func consultaPacientePorUFEClassificacaoEtariaIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	q := r.URL.Query()

	idUFString := q.Get("idUF")
	idClassificacaoString := q.Get("idClassificacao")

	idUF, _ := strconv.ParseInt(idUFString, 10,32)
	idClassificacao, _ := strconv.ParseInt(idClassificacaoString, 10,32)

	contagem := consultaPacientePorUFEClassificacaoEtaria(int(idUF), int(idClassificacao))

	retorno := make(map[string]int)

	retorno["contagem"] = contagem

	x, err := json.MarshalIndent(retorno," ","    ")
	if err != nil {
		panic(err)
	}

	w.Write([]byte(x))
}

func consultaPacientePorClassificacaoEtariaEZonaIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	q := r.URL.Query()

	idClassificacaoString := q.Get("idClassificacao")
	idZonaString := q.Get("idZona")

	idClassificacao, _ := strconv.ParseInt(idClassificacaoString, 10,32)
	idZona, _ := strconv.ParseInt(idZonaString, 10,32)

	contagem := consultaPacientePorClassificacaoEtariaEZona(int(idClassificacao), int(idZona))

	retorno := make(map[string]int)

	retorno["contagem"] = contagem

	x, err := json.MarshalIndent(retorno," ","    ")
	if err != nil {
		panic(err)
	}

	w.Write([]byte(x))
}

func consultaPacientePorClassificacaoEtariaEPatologiaIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	q := r.URL.Query()

	idClassificacaoString := q.Get("idClassificacao")
	idPatologiaString := q.Get("idPatologia")

	idClassificacao, _ := strconv.ParseInt(idClassificacaoString, 10,32)
	idPatologia, _ := strconv.ParseInt(idPatologiaString, 10,32)

	contagem := consultaPacientePorClassificacaoEtariaEPatologia(int(idClassificacao), int(idPatologia))

	retorno := make(map[string]int)

	retorno["contagem"] = contagem

	x, err := json.MarshalIndent(retorno," ","    ")
	if err != nil {
		panic(err)
	}

	w.Write([]byte(x))
}

func consultaPacientePorClassificacaoEtariaEUFIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	q := r.URL.Query()

	idClassificacaoString := q.Get("idClassificacao")
	idUFString := q.Get("idUF")

	idClassificacao, _ := strconv.ParseInt(idClassificacaoString, 10, 32)
	idUF, _ := strconv.ParseInt(idUFString, 10, 32)

	contagem := consultaPacientePorClassificacaoEtariaEUF(int(idClassificacao), int(idUF))

	retorno := make(map[string]int)

	retorno["contagem"] = contagem

	x, err := json.MarshalIndent(retorno, " ", "    ")
	if err != nil {
		panic(err)
	}

	w.Write([]byte(x))
}

func consultaPacienteZonaPorClassificacaoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	dados := ConstroiClassificacaoEtariaPorZona()

	x, err := json.MarshalIndent(dados, " ", "    ")
	if err != nil {
		panic(err)
	}

	w.Write([]byte(x))
}

func pacienteZonaPorClassificacaoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	dados := GeradorHTMLClassificacaoEtariaPorZona()
	w.Write([]byte(dados))
}