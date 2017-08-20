package main

import (
	"time"
	"math/rand"
)

type Individuo struct {
	algoritmo *Algoritmo
	genes string
	aptidao int
}

func (iv *Individuo) obterCaractereUTF8(caracteres string, index int) string {
	i := 0
    for _, x := range caracteres {
		if i == index {
			return string(x)
		}
		i++
	}
	return ""
}

func (iv *Individuo) InitRandom(numGenes int, alg *Algoritmo) {
	iv.genes = ""
	iv.algoritmo = alg

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	caracteres := alg.getCaracteres()
	tam := len([]rune(caracteres))
	for i := 0; i < numGenes; i++ {
		iv.genes += iv.obterCaractereUTF8(caracteres, r.Intn(tam))
	}

	iv.geraAptidao()
}

func (iv *Individuo) InitGenes(genes string, alg *Algoritmo) {
	iv.genes = genes
	iv.algoritmo = alg

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	//se for mutar, cria um gene aleatório
	if r.Float64() <= alg.getTaxaDeMutacao() {
		caracteres := alg.getCaracteres()
		geneNovo := ""
		posAleatoria := r.Intn(len([]rune(genes)))
		tamCaracteres := len([]rune(caracteres))
		i := 0
		for _, c := range genes {
			if i == posAleatoria {
				geneNovo += iv.obterCaractereUTF8(caracteres, r.Intn(tamCaracteres))
			} else {
				geneNovo += string(c)
			}
			i++
		}

		iv.genes = geneNovo
	}

	iv.geraAptidao()
}

//gera o valor de aptidão, será calculada pelo número de bits do gene iguais ao da solução
func (iv *Individuo) geraAptidao() {
	solucao := iv.algoritmo.getSolucao()
	for i := 0; i < len([]rune(solucao)); i++ {
		if iv.obterCaractereUTF8(solucao, i) == iv.obterCaractereUTF8(iv.genes, i) {
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
