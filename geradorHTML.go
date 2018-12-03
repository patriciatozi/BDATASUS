package main

import "fmt"

func GerarIndex() string {
	return `
<html>
	<head>
		<title>MedConsulta</title>
	</head>
	<body>
		<ol>
			<li>Zona Por Patologia <a href="/pacienteZonaPorPatologia">Tabela</a></li>
			<li>Zona Por UF <a href="/pacienteZonaPorUF">Tabela</a></li>
			<li>Zona Por Classificacao Etaria <a href="/pacienteZonaPorClassificacao">Tabela</a></li>
			
		</ol>
	</body>
</html>
`
}
func GeradorHTMLZonaPorPatologia() string {
	var dados = ConstroiZonaPorPatologia()
	var html = ""

	var patologias = dados[0].Patologias

	html += `<html>
	<head>
		<title>Zona por Patologia</title>
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
		<th class="cabecalho">Patologia</th>
	`
	for _, v := range patologias {
		html += fmt.Sprintf(`<th class="cabecalho">%s</th>`, v.Patologia.Nome)
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

		for _, c := range v.Patologias {
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

func GeradorHTMLZonaPorUF() string {
	var dados = ConstroiZonaPorUF()
	var html = ""

	var ufs = dados[0].UFS

	html += `<html>
	<head>
		<title>Zona por UF</title>
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
		<th class="cabecalho">UF</th>
	`
	for _, v := range ufs {
		html += fmt.Sprintf(`<th class="cabecalho">%s</th>`, v.UF.Nome)
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

		for _, c := range v.UFS {
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

func GeradorHTMLZonaPorClassificacaoEtaria() string {
	var dados = ConstroiZonaPorClassificacaoEtaria()
	var html = ""

	var classificacoes = dados[0].Classificacoes

	html += `<html>
	<head>
		<title>Zona Por Classificação Etária</title>
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
