package main

import (
	"time"
	"math/rand"
)

type Individuo struct {
	genes string
	aptidao int
}

func (iv *Individuo) InitRandom(numGenes int) {
	iv.genes = ""
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	alg := Algoritmo{}
	caracteres := alg.getCaracteres()

	for i := 0; i < numGenes; i++ {
		iv.genes += string(caracteres[r.Intn(len(caracteres))])
	}

	iv.geraAptidao()
}

func (iv *Individuo) InitGenes(genes string) {
	iv.genes = genes

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	//se for mutar, cria um gene aleatório
	alg := Algoritmo{}
	if r.Float64() <= alg.getTaxaDeMutacao() {
		caracteres := alg.getCaracteres()
		geneNovo := ""
		posAleatoria := r.Intn(len(genes))
		for i := 0; i < len(genes); i++ {
			if i == posAleatoria {
				geneNovo += string(caracteres[r.Intn(len(caracteres))])
			} else {
				geneNovo += string(genes[i])
			}
		}
		iv.genes = geneNovo
	}
	iv.geraAptidao()
}

//gera o valor de aptidão, será calculada pelo número de bits do gene iguais ao da solução
func (iv *Individuo) geraAptidao() {
	alg := Algoritmo{}
	solucao := alg.getSolucao()
	for i := 0; i < len(solucao); i++ {
		if solucao[i] == iv.genes[i] {
			iv.aptidao++
		}
	}
}

func (iv *Individuo) getAptidao() int {
	return iv.aptidao
}

func (iv *Individuo) getGenes() string {
	return iv.genes
}
