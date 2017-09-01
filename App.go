package main

import (
	"fmt"
	"net/http"
)

type App struct {
	taxaCrossover float64
	taxaMutacao float64
	elitismo bool
	tamanhoPopulacao int
}

func (a *App) Run(respWr http.ResponseWriter, req *http.Request, d *Data) {
	algoritmo := &Algoritmo{}
	//Define a solução
	algoritmo.setSolucao(d.toString())
	// Setar quantidade de formas
	algoritmo.setNumeroFormas(d.nCircles, d.nRectangles, d.nTriangles)
	//Define os caracteres existentes
	algoritmo.setCaracteres("abcdef1234567890")
	//taxa de crossover de 60%
	algoritmo.setTaxaDeCrossover(a.taxaCrossover)
	//taxa de mutação de 3%
	algoritmo.setTaxaDeMutacao(a.taxaMutacao)
	//elitismo
	elitismo := a.elitismo
	//tamanho da população
	tamPop := a.tamanhoPopulacao
	//numero mÃ¡ximo de gerações
	numMaxGeracoes := 10000

	//define o nÃºmero de genes do indivÃ­duo baseado na solução
	numGenes := len([]rune(algoritmo.getSolucao()))

	//cria a primeira população aleatÃ©rioa
	populacao := &Populacao{}
	populacao.InitRandom(numGenes, tamPop, algoritmo)

	temSolucao := false
	geracao := 0

	fmt.Printf("Iniciando... Aptidão da solução: %d\n", numGenes)

	//loop atÃ© o critÃ©rio de parada
	solucaoAnt := ""
	for !temSolucao && geracao < numMaxGeracoes {
		geracao++

		//cria nova populacao
		populacao = algoritmo.novaGeracao(populacao, elitismo)
		individuo0 := populacao.getIndividuo(0)

		fmt.Printf("Geração %d | Aptidão: %d | Melhor: %s\n", geracao, individuo0.getAptidao(), individuo0.getGenes())

		if solucaoAnt != individuo0.getGenes() {
			data := &Data{}
			data.fromString(individuo0.getGenes(), d.nCircles, d.nRectangles, d.nTriangles)
			label := fmt.Sprintf("Geração %d", geracao)
			PrintImage(respWr, req, data, label)
			solucaoAnt = individuo0.getGenes()
		}

		//verifica se tem a solucao
		temSolucao = populacao.temSolucao(algoritmo.getSolucao())
	}

	if geracao == numMaxGeracoes {
		individuo0 := populacao.getIndividuo(0)
		fmt.Printf("Numero Maximo de Gerações | %s %d\n", individuo0.getGenes(), individuo0.getAptidao())
	}

	if temSolucao {
		individuo0 := populacao.getIndividuo(0)
		fmt.Printf("Encontrado resultado na geração %d | %s (Aptidao: %d)\n", geracao, individuo0.getGenes(), individuo0.getAptidao())
	}
}
