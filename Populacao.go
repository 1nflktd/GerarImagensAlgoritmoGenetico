package main

type Populacao struct {
	individuos []Individuo
	tamPopulacao int
}

//cria uma população com indivíduos aleatória
func (p *Populacao) InitRandom(numGenes, tamPop int) {
    p.tamPopulacao = tamPop
    p.individuos = make([]Individuo, tamPop)
    for i := 0; i < tamPop; i++ {
        p.individuos[i] = Individuo{}
        p.individuos[i].InitRandom(numGenes)
    }
}

//cria uma população com indivíduos sem valor, será composto posteriormente
func (p *Populacao) InitEmpty(tamPop int) {
    p.tamPopulacao = tamPop
    p.individuos = make([]Individuo, tamPop)
}

//coloca um indivíduo em uma certa posição da população
func (p *Populacao) setIndividuoPos(individuo Individuo, posicao int) {
	p.individuos[posicao] = individuo
}

//coloca um indivíduo na próxima posição disponível da população
func (p *Populacao) setIndividuo(individuo Individuo) {
    for i := 0; i < p.tamPopulacao; i++ {
        if p.individuos[i].getGenes() == "" {
            p.individuos[i] = individuo
            return
        }
    }
}

//verifica se algum indivíduo da população possui a solução
func (p *Populacao) temSolucao(solucao string) bool {
    var i Individuo
    for j := 0; j < p.tamPopulacao; j++ {
        if p.individuos[j].getGenes() == solucao {
            i = p.individuos[j]
            break;
        }
    }

    if i.getGenes() == "" { // == nil
        return false
    }

    return true
}

//ordena a população pelo valor de aptidão de cada indivíduo, do maior valor para o menor, assim se eu quiser obter o melhor indivíduo desta população, acesso a posição 0 do array de indivíduos
func (p *Populacao) ordenaPopulacao() {
    trocou := true
    for trocou {
        trocou = false
        for i := 0; i < (p.tamPopulacao - 1); i++ {
            if p.individuos[i].getAptidao() < p.individuos[i + 1].getAptidao() {
                temp := p.individuos[i]
                p.individuos[i] = p.individuos[i + 1]
                p.individuos[i + 1] = temp
                trocou = true
            }
        }
    }
}

//número de indivíduos existentes na população
func (p *Populacao) getNumIndividuos() int {
    num := 0
    for i := 0; i < p.tamPopulacao; i++ {
        if p.individuos[i].getGenes() != "" { //  != nil
            num++
        }
    }
    return num
}

func (p *Populacao) getTamPopulacao() int {
    return p.tamPopulacao
}

func (p *Populacao) getIndividuo(pos int) Individuo {
    return p.individuos[pos]
}
