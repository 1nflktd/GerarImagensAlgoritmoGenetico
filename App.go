package main

import (
	"fmt"
)

type App struct {}

func (a *App) Run() {
	algoritmo := Algoritmo{}
	//Define a solução
	algoritmo.setSolucao("A dúvida é o princípio da sabedoria")
	//Define os caracteres existentes
	algoritmo.setCaracteres("!,.:;àáãâúíóôõéêQWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890 ")
	//taxa de crossover de 60%
	algoritmo.setTaxaDeCrossover(0.6)
	//taxa de mutação de 3%
	algoritmo.setTaxaDeMutacao(0.3)
	//elitismo
	elitismo := true
	//tamanho da população
	tamPop := 100
	//numero mÃ¡ximo de gerações
	numMaxGeracoes := 10000

	//define o nÃºmero de genes do indivÃ­duo baseado na solução
	numGenes := len(algoritmo.getSolucao())

	//cria a primeira população aleatÃ©rioa
	populacao := Populacao{}
	populacao.InitRandom(numGenes, tamPop)

	temSolucao := false
	geracao := 0

	fmt.Printf("Iniciando... Aptidão da solução: %d\n", numGenes)

	//loop atÃ© o critÃ©rio de parada
	for !temSolucao && geracao < numMaxGeracoes {
		geracao++

		//cria nova populacao
		populacao = algoritmo.novaGeracao(populacao, elitismo)
		individuo0 := populacao.getIndividuo(0)

		fmt.Printf("Geração %d | Aptidão: %d | Melhor: %d\n", geracao, individuo0.getAptidao(), individuo0.getGenes())

		//verifica se tem a solucao
		temSolucao = populacao.temSolucao(algoritmo.getSolucao())
	}

	if geracao == numMaxGeracoes {
		individuo0 := populacao.getIndividuo(0)
		fmt.Printf("NÃºmero Maximo de Gerações | %d %d\n", individuo0.getGenes(), individuo0.getAptidao())
	}

	if temSolucao {
		individuo0 := populacao.getIndividuo(0)
		fmt.Printf("Encontrado resultado na geração %d | %d  (AptidÃ£o: %d)\n", geracao, individuo0.getGenes(), individuo0.getAptidao())
	}
}