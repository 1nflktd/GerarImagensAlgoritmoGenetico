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
	nCircles, nRectangles, nTriangles int
}

func (a *Algoritmo) novaGeracao(populacao * Populacao, elitismo bool) (*Populacao) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//nova população do mesmo tamanho da antiga
	novaPopulacao := &Populacao{}
	novaPopulacao.InitEmpty(populacao.getTamPopulacao())

	//se tiver elitismo, mantém o melhor indivíduo da geração atual
	if elitismo {
		novaPopulacao.setIndividuo(populacao.getIndividuo(0))
	}

	//insere novos indivíduos na nova população, até atingir o tamanho máximo
	for novaPopulacao.getNumIndividuos() < novaPopulacao.getTamPopulacao() {
		//seleciona os 2 pais por torneio
		pais := a.selecaoTorneio(populacao)
		filhos := [2]*Individuo{}

		//verifica a taxa de crossover, se sim realiza o crossover, se não, mantém os pais selecionados para a próxima geração
		if r.Float64() <= a.taxaDeCrossover {
			filhos = a.crossover(pais[1], pais[0])
		} else {
			filhos[0] = &Individuo{}
			filhos[0].InitGenes(pais[0].getGenes(), a)
			filhos[1] = &Individuo{}
			filhos[1].InitGenes(pais[1].getGenes(), a)
		}

		//adiciona os filhos na nova geração
		novaPopulacao.setIndividuo(filhos[0])
		novaPopulacao.setIndividuo(filhos[1])
	}

	//ordena a nova população
	novaPopulacao.ordenaPopulacao()

	return novaPopulacao
}

func (a *Algoritmo) obterPontosCorte(pos1, pos2, pos3, pos4 int, genes string) (pontoCorte1, pontoCorte2, pontoCorte3, pontoCorte4 int) {
	i := 0
	for p, _ := range genes {
		if i == pos1 {
			pontoCorte1 = p
		} else if i == pos2 {
			pontoCorte2 = p
		} else if i == pos3 {
			pontoCorte3 = p
		} else if i == pos4 {
			pontoCorte4 = p
			break
		}
		i++
	}

	return
}

func (a *Algoritmo) crossover(individuo1, individuo2 *Individuo) ([2]*Individuo) {
	//pega os genes dos pais
	genePai1 := individuo1.getGenes()
	genePai2 := individuo2.getGenes()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	tamanho := len([]rune(genePai1))
	pos1 := r.Intn((tamanho/4) - 2) + 1
	pos2 := r.Intn((tamanho/4) - 2) + tamanho / 4
	pos3 := r.Intn((tamanho/4) - 2) + (tamanho*2) / 4
	pos4 := r.Intn((tamanho/4) - 2) + (tamanho*3) / 4

	//sorteia o ponto de corte
	pontoCorte1Pai1, pontoCorte2Pai1, pontoCorte3Pai1, pontoCorte4Pai1 := a.obterPontosCorte(pos1, pos2, pos3, pos4, genePai1)
	pontoCorte1Pai2, pontoCorte2Pai2, pontoCorte3Pai2, pontoCorte4Pai2 := a.obterPontosCorte(pos1, pos2, pos3, pos4, genePai2)

	//realiza o corte,
	geneFilho1 := string(genePai1[0:pontoCorte1Pai1])
	geneFilho1 += string(genePai2[pontoCorte1Pai2:pontoCorte2Pai2])
	geneFilho1 += string(genePai1[pontoCorte2Pai1:pontoCorte3Pai1])
	geneFilho1 += string(genePai2[pontoCorte3Pai2:pontoCorte4Pai2])
	geneFilho1 += string(genePai1[pontoCorte4Pai1:])

	geneFilho2 := string(genePai2[0:pontoCorte1Pai2])
	geneFilho2 += string(genePai1[pontoCorte1Pai1:pontoCorte2Pai1])
	geneFilho2 += string(genePai2[pontoCorte2Pai2:pontoCorte3Pai2])
	geneFilho2 += string(genePai1[pontoCorte3Pai1:pontoCorte4Pai1])
	geneFilho2 += string(genePai2[pontoCorte4Pai2:])

	//cria o novo indivíduo com os genes dos pais
	filhos := [2]*Individuo{}
	filhos[0] = &Individuo{}
	filhos[0].InitGenes(geneFilho1, a)
	filhos[1] = &Individuo{}
	filhos[1].InitGenes(geneFilho2, a)

	return filhos
}

func (a *Algoritmo) selecaoTorneio(populacao * Populacao) ([2]*Individuo) {
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
	pais := [2]*Individuo{}
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

func (a *Algoritmo) setNumeroFormas(nCircles, nRectangles, nTriangles int) {
	a.nCircles, a.nRectangles, a.nTriangles = nCircles, nRectangles, nTriangles
}

func (a *Algoritmo) getNumeroFormas() (int, int, int) {
	return a.nCircles, a.nRectangles, a.nTriangles
}
