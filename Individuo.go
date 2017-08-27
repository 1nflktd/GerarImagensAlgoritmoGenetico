package main

import (
	"time"
	"math/rand"
	"math"
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
		posAleatoria1 := r.Intn(len([]rune(genes)))
		posAleatoria2 := r.Intn(len([]rune(genes)))
		posAleatoria3 := r.Intn(len([]rune(genes)))
		tamCaracteres := len([]rune(caracteres))
		i := 0
		for _, c := range genes {
			if i == posAleatoria1 || i == posAleatoria2 || i == posAleatoria3  {
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

func (iv *Individuo) geraAptidao2() {
	nCircles, nRectangles, nTriangles := iv.algoritmo.getNumeroFormas()

	dataSolucao := &Data{}
	dataSolucao.fromString(iv.algoritmo.getSolucao(), nCircles, nRectangles, nTriangles)

	data := &Data{}
	data.fromString(iv.genes, nCircles, nRectangles, nTriangles)

	// comparar a diferenca entre os valores
	iv.aptidao = 0
	for i, c := range dataSolucao.circles {
		/*
		iv.aptidao -= int(math.Abs(float64(c.x - data.circles[i].x)))
		iv.aptidao -= int(math.Abs(float64(c.y - data.circles[i].y)))
		iv.aptidao -= int(math.Abs(float64(c.r - data.circles[i].r)))
		iv.aptidao -= int(math.Abs(float64(c.red - data.circles[i].red)))
		iv.aptidao -= int(math.Abs(float64(c.green - data.circles[i].green)))
		iv.aptidao -= int(math.Abs(float64(c.blue - data.circles[i].blue)))
		*/
		iv.aptidao -= int(math.Abs(float64(c.x + c.y + c.r + c.red + c.green + c.blue - data.circles[i].x - data.circles[i].y - data.circles[i].r - data.circles[i].red - data.circles[i].green - data.circles[i].blue)))
	}

	for i, c := range dataSolucao.rectangles {
		/*
		iv.aptidao -= int(math.Abs(float64(c.x - data.rectangles[i].x)))
		iv.aptidao -= int(math.Abs(float64(c.y - data.rectangles[i].y)))
		iv.aptidao -= int(math.Abs(float64(c.w - data.rectangles[i].w)))
		iv.aptidao -= int(math.Abs(float64(c.h - data.rectangles[i].h)))
		iv.aptidao -= int(math.Abs(float64(c.red - data.rectangles[i].red)))
		iv.aptidao -= int(math.Abs(float64(c.green - data.rectangles[i].green)))
		iv.aptidao -= int(math.Abs(float64(c.blue - data.rectangles[i].blue)))
		*/
		iv.aptidao -= int(math.Abs(float64(c.x + c.y + c.w + c.h + c.red + c.green + c.blue - data.rectangles[i].x - data.rectangles[i].y - data.rectangles[i].w - data.rectangles[i].h - data.rectangles[i].red - data.rectangles[i].green - data.rectangles[i].blue)))
	}

	for i, c := range dataSolucao.triangles {
		/*
		iv.aptidao -= int(math.Abs(float64(c.p1 - data.triangles[i].p1)))
		iv.aptidao -= int(math.Abs(float64(c.p2 - data.triangles[i].p2)))
		iv.aptidao -= int(math.Abs(float64(c.p3 - data.triangles[i].p3)))
		iv.aptidao -= int(math.Abs(float64(c.red - data.triangles[i].red)))
		iv.aptidao -= int(math.Abs(float64(c.green - data.triangles[i].green)))
		iv.aptidao -= int(math.Abs(float64(c.blue - data.triangles[i].blue)))
		*/
		iv.aptidao -= int(math.Abs(float64(c.p1 + c.p2 + c.p3 + c.red + c.green + c.blue - data.triangles[i].p1 - data.triangles[i].p2 - data.triangles[i].p3 - data.triangles[i].red - data.triangles[i].green - data.triangles[i].blue)))
	}
}

func (iv *Individuo) getAptidao() int {
	return iv.aptidao
}

func (iv *Individuo) getGenes() string {
	return iv.genes
}
