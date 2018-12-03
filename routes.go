package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ConsultaPacientePorZonaEPatologia",
		"GET",
		"/consultaPacientePorZonaEPatologia",
		consultaPacientePorZonaEPatologiaIndex,
	},
	Route{
		"consultaPacientePorZonaEUF",
		"GET",
		"/consultaPacientePorZonaEUF",
		consultaPacientePorZonaEUFIndex,
	},
	Route{
		"consultaPacientePorZonaEClassificacaoEtaria",
		"GET",
		"/consultaPacientePorZonaEClassificacaoEtaria",
		consultaPacientePorZonaEClassificacaoEtariaIndex,
	},
	Route{
		"consultaPacientePorPatologiaEZona",
		"GET",
		"/consultaPacientePorPatologiaEZona",
		consultaPacientePorPatologiaEZonaIndex,
	},
	Route{
		"consultaPacientePorPatologiaEUF",
		"GET",
		"/consultaPacientePorPatologiaEUF",
		consultaPacientePorPatologiaEUFIndex,
	},
	Route{
		"consultaPacientePorPatologiaEClassificacaoEtaria",
		"GET",
		"/consultaPacientePorPatologiaEClassificacaoEtaria",
		consultaPacientePorPatologiaEClassificacaoEtariaIndex,
	},
	Route{
		"consultaPacientePorUFEZona",
		"GET",
		"/consultaPacientePorUFEZona",
		consultaPacientePorUFEZonaIndex,
	},
	Route{
		"consultaPacientePorUFEPatologia",
		"GET",
		"/consultaPacientePorUFEPatologia",
		consultaPacientePorUFEPatologiaIndex,
	},
	Route{
		"consultaPacientePorUFEClassificacaoEtaria",
		"GET",
		"/consultaPacientePorUFEClassificacaoEtaria",
		consultaPacientePorUFEClassificacaoEtariaIndex,
	},
	Route{
		"consultaPacientePorClassificacaoEtariaEZona",
		"GET",
		"/consultaPacientePorClassificacaoEtariaEZona",
		consultaPacientePorClassificacaoEtariaEZonaIndex,
	},
	Route{
		"consultaPacientePorClassificacaoEtariaEPatologia",
		"GET",
		"/consultaPacientePorClassificacaoEtariaEPatologia",
		consultaPacientePorClassificacaoEtariaEPatologiaIndex,
	},
	Route{
		"consultaPacientePorClassificacaoEtariaEUF",
		"GET",
		"/consultaPacientePorClassificacaoEtariaEUF",
		consultaPacientePorClassificacaoEtariaEUFIndex,
	},
	Route{
		"consultaPacienteZonaPorClassificacao",
		"GET",
		"/consultaPacienteZonaPorClassificacao",
		consultaPacienteZonaPorClassificacaoIndex,
	},
	Route{
		"pacienteZonaPorClassificacao",
		"GET",
		"/pacienteZonaPorClassificacao",
		pacienteZonaPorClassificacaoIndex,
	},
	Route{
		"pacienteZonaPorPatologia",
		"GET",
		"/pacienteZonaPorPatologia",
		pacienteZonaPorPatologiaIndex,
	},
	Route{
		"pacienteZonaPorUF",
		"GET",
		"/pacienteZonaPorUF",
		pacienteZonaPorUFIndex,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
		"GET",
		"/consultaPacientePorPatologiaEUF",
		consultaPacientePorPatologiaEUFIndex,
	},
	Route{
		"consultaPacientePorPatologiaEClassificacaoEtaria",
		"GET",
		"/consultaPacientePorPatologiaEClassificacaoEtaria",
		consultaPacientePorPatologiaEClassificacaoEtariaIndex,
	},
	Route{
		"consultaPacientePorUFEZona",
		"GET",
		"/consultaPacientePorUFEZona",
		consultaPacientePorUFEZonaIndex,
	},
	Route{
		"consultaPacientePorUFEPatologia",
		"GET",
		"/consultaPacientePorUFEPatologia",
		consultaPacientePorUFEPatologiaIndex,
	},
	Route{
		"consultaPacientePorUFEClassificacaoEtaria",
		"GET",
		"/consultaPacientePorUFEClassificacaoEtaria",
		consultaPacientePorUFEClassificacaoEtariaIndex,
	},
	Route{
		"consultaPacientePorClassificacaoEtariaEZona",
		"GET",
		"/consultaPacientePorClassificacaoEtariaEZona",
		consultaPacientePorClassificacaoEtariaEZonaIndex,
	},
	Route{
		"consultaPacientePorClassificacaoEtariaEPatologia",
		"GET",
		"/consultaPacientePorClassificacaoEtariaEPatologia",
		consultaPacientePorClassificacaoEtariaEPatologiaIndex,
	},
	Route{
		"consultaPacientePorClassificacaoEtariaEUF",
		"GET",
		"/consultaPacientePorClassificacaoEtariaEUF",
		consultaPacientePorClassificacaoEtariaEUFIndex,
	},
	Route{
		"consultaPacienteZonaPorClassificacao",
		"GET",
		"/consultaPacienteZonaPorClassificacao",
		consultaPacienteZonaPorClassificacaoIndex,
	},
	Route{
		"pacienteZonaPorClassificacao",
		"GET",
		"/pacienteZonaPorClassificacao",
		pacienteZonaPorClassificacaoIndex,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
