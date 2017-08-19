package main

import (
	"time"
	"math/rand"
)

type Algoritmo struct {
	solucao string
	taxaDeCrossover float64
	taxaDeMutacao float64
	caracteres string
}

func (a *Algoritmo) novaGeracao(populacao Populacao, elitismo bool) (Populacao) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//nova população do mesmo tamanho da antiga
	novaPopulacao := Populacao{}
	novaPopulacao.InitEmpty(populacao.getTamPopulacao())

	//se tiver elitismo, mantém o melhor indivíduo da geração atual
	if elitismo {
		novaPopulacao.setIndividuo(populacao.getIndividuo(0))
	}

	//insere novos indivíduos na nova população, até atingir o tamanho máximo
	for novaPopulacao.getNumIndividuos() < novaPopulacao.getTamPopulacao() {
		//seleciona os 2 pais por torneio
		pais := a.selecaoTorneio(populacao)
		filhos := [2]Individuo{}

		//verifica a taxa de crossover, se sim realiza o crossover, se não, mantém os pais selecionados para a próxima geração
		if r.Float64() <= a.taxaDeCrossover {
			filhos = a.crossover(pais[1], pais[0])
		} else {
			filhos[0].InitGenes(pais[0].getGenes())
			filhos[1].InitGenes(pais[1].getGenes())
		}

		//adiciona os filhos na nova geração
		novaPopulacao.setIndividuo(filhos[0])
		novaPopulacao.setIndividuo(filhos[1])
	}

	//ordena a nova população
	novaPopulacao.ordenaPopulacao()

	return novaPopulacao
}

func (a *Algoritmo) crossover(individuo1, individuo2 Individuo) ([2]Individuo) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	//sorteia o ponto de corte
	pontoCorte1 := r.Intn((len(individuo1.getGenes())/2) - 2) + 1
	pontoCorte2 := r.Intn((len(individuo1.getGenes())/2) - 2) + len(individuo1.getGenes()) / 2

	//pega os genes dos pais
	genePai1 := individuo1.getGenes()
	genePai2 := individuo2.getGenes()

	//realiza o corte,
	/*
	geneFilho1 := genePai1.substring(0, pontoCorte1)
	geneFilho1 += genePai2.substring(pontoCorte1, pontoCorte2)
	geneFilho1 += genePai1.substring(pontoCorte2, len(genePai1))
	*/

	geneFilho1 := genePai1[0:pontoCorte1]
	geneFilho1 += genePai2[pontoCorte1:pontoCorte2]
	geneFilho1 += genePai1[pontoCorte2:len(genePai1)]

	/*
	geneFilho2 = genePai2.substring(0, pontoCorte1)
	geneFilho2 += genePai1.substring(pontoCorte1, pontoCorte2)
	geneFilho2 += genePai2.substring(pontoCorte2, len(genePai2))
	*/

	geneFilho2 := genePai2[0:pontoCorte1]
	geneFilho2 += genePai1[pontoCorte1:pontoCorte2]
	geneFilho2 += genePai2[pontoCorte2:len(genePai2)]

	//cria o novo indivíduo com os genes dos pais
	filhos := [2]Individuo{}
	filhos[0].InitGenes(geneFilho1)
	filhos[1].InitGenes(geneFilho2)

	return filhos
}

func (a *Algoritmo) selecaoTorneio(populacao Populacao) ([2]Individuo) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	populacaoIntermediaria := &Populacao{}
	populacaoIntermediaria.InitEmpty(3)

	//seleciona 3 indivíduos aleatóriamente na população
	populacaoIntermediaria.setIndividuo(populacao.getIndividuo(r.Intn(populacao.getTamPopulacao())))
	populacaoIntermediaria.setIndividuo(populacao.getIndividuo(r.Intn(populacao.getTamPopulacao())))
	populacaoIntermediaria.setIndividuo(populacao.getIndividuo(r.Intn(populacao.getTamPopulacao())))

	//ordena a população
	populacaoIntermediaria.ordenaPopulacao()

	//seleciona os 2 melhores deste população
	pais := [2]Individuo{}
	pais[0] = populacaoIntermediaria.getIndividuo(0)
	pais[1] = populacaoIntermediaria.getIndividuo(1)

	return pais
}

func (a *Algoritmo) getSolucao() string {
	return a.solucao
}

func (a *Algoritmo) setSolucao(solucao string) {
	a.solucao = solucao
}

func (a *Algoritmo) getTaxaDeCrossover() float64 {
	return a.taxaDeCrossover
}

func (a *Algoritmo) setTaxaDeCrossover(taxaDeCrossover float64) {
	a.taxaDeCrossover = taxaDeCrossover
}

func (a *Algoritmo) getTaxaDeMutacao() float64 {
	return a.taxaDeMutacao
}

func (a *Algoritmo) setTaxaDeMutacao(taxaDeMutacao float64) {
	a.taxaDeMutacao = taxaDeMutacao
}

func (a *Algoritmo) getCaracteres() string {
	return a.caracteres
}

func (a *Algoritmo) setCaracteres(caracteres string) {
	a.caracteres = caracteres
}
