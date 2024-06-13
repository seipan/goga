package goga

type GA struct {
	GAConfig
	Population
	Selector       Selector
	BestIndividual Individual
	PrintCallBack  func()
}

type GAConfig struct {
	PopulationSize uint
	NGenerations   uint
	CrossoverRate  float64
	MutationRate   float64
	ParallelEval   bool
}

type Population struct {
	Individuals Individuals
	Generations uint
}

func NewGA(gaConfig GAConfig, selector Selector) *GA {
	return &GA{
		GAConfig: gaConfig,
		Selector: selector,
	}
}

func (ga *GA) initPopulation(g Genome) {
	indis := make(Individuals, ga.PopulationSize)
	for i := range indis {
		indis[i].Chromosome = g.Initialization()
	}
	indis.Evaluate(ga.ParallelEval)
	ga.Population.Generations = 0
	ga.Population.Individuals = indis
	ga.Population.Individuals.SortByFitness()
	ga.BestIndividual = ga.Population.Individuals[0]
}
