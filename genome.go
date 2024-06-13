package goga

type Genome interface {
	Initialization() Genome
	Fitness() int
	Mutation()
	Crossover(Genome) Genome
}
