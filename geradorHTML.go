package main

import "fmt"

func GerarIndex() string {
	return `
<html>
	<head>
		<title>BDDATASUS</title>
	</head>
	<body>
		<ol>
			<li><a href="/pacienteZonaPorClassificacao">Classificacao Etaria Por Zona</a></li>
		</ol>
	</body>
</html>
`
}

func GeradorHTMLClassificacaoEtariaPorZona() string {
	var dados = ConstroiClassificacaoEtariaPorZona()
	var html = ""

	var classificacoes = dados[0].Classificacoes

	html += `<html>
	<head>
		<title>Classificação Etária Por Zona</title>
		<style>
			.cabecalho {
    			border-top: 1px solid white;
    			border-bottom: 1px solid white;
    			border-right: 1px solid white;
    			border-left: 1px solid white;
    			padding-top: 0.5em;
    			padding-bottom: 0.5em;
    			text-align: center;
			}
			.tabela {
    			border-collapse: collapse;
    			border-spacing: 0;
    			font-size: 10pt;
    			font-weight: normal;
    			color: #000;
    			background-color: transparent;
			}
			.tabela th {
		    	background-color: #7e919e;
		    	border: 2px solid #fff;
		    	color: #fff;
		    	padding: 3px 7px;
			}
			.tabela td {
		    	background-color: #e5f6fc;
		    	border: 2px solid #fff;
		    	padding: 3px 7px;
			}
		</style>
	</head>
    <body>
	`

	// region Construir HTML
	html += `<table width="100%" class="tabela">`
	// region Construir Cabeçalho
	html += `
	<tr>
		<th class="cabecalho">Zona</th>
	`
	for _, v := range classificacoes {
		html += fmt.Sprintf(`<th class="cabecalho">%s</th>`, v.ClassificacaoEtaria.Classificacao)
		html += "\n"
	}
	html += `
	</tr>`
	// endregion
	// region Construir Linhas
	for _, v := range dados {
		html += fmt.Sprintf(`
	<tr>
		<td>%s</td>
		`, v.Zona.Nome)

		for _, c := range v.Classificacoes {
			html += fmt.Sprintf(`<td>%d</td>`, c.Contagem)
			html += "\n"
		}
		html += `
	</tr>
		`
	}
	// endregion
	html += `</table>`
	// endregion

	html += `
	<BR>
	<a href="/">Voltar</a>
</body>
</html>
`

	return html
}