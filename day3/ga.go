package day3

import (
	"math/rand/v2"
	"strconv"
)

type Individual struct {
	Genome  []bool
	Fitness int64
}

type Population struct {
	Individuals []Individual
}

type GAConfig struct {
	PopulationSize int
	GenomeLength   int
	MutationRate   float64
	CrossoverRate  float64
	Generations    int
}

func SolveGA(bank string, amountOfBatteries int, cfg GAConfig) Individual {
	n := len(bank)

	pop := make([]Individual, cfg.PopulationSize)
	for i := range pop {
		pop[i] = randomIndividual(n, amountOfBatteries)
		pop[i].Fitness = fitness(pop[i], bank)
	}

	best := pop[0]

	for g := 0; g < cfg.Generations; g++ {
		newPop := make([]Individual, cfg.PopulationSize)

		for i := 0; i < cfg.PopulationSize; i++ {
			p1 := tournament(pop)
			p2 := tournament(pop)

			var child Individual

			// FIX: crossover rate applied HERE
			if rand.Float64() < cfg.CrossoverRate {
				child = crossover(p1, p2, amountOfBatteries)
			} else {
				child = copyIndividual(p1)
			}

			// mutation
			if rand.Float64() < cfg.MutationRate {
				child = mutate(child)
			}

			child.Fitness = fitness(child, bank)
			newPop[i] = child

			if child.Fitness > best.Fitness {
				best = child
			}
		}

		pop = newPop
	}

	return best
}

func fitness(ind Individual, bank string) int64 {
	value := ""
	for i, take := range ind.Genome {
		if take {
			value = value + string(bank[i])
		}
	}

	result, _ := strconv.ParseInt(value, 10, 64)

	return result
}

func randomIndividual(n, k int) Individual {
	genome := make([]bool, n)
	indices := rand.Perm(n)[:k]

	for _, i := range indices {
		genome[i] = true
	}

	return Individual{Genome: genome}
}

func tournament(pop []Individual) Individual {
	a := pop[rand.IntN(len(pop))]
	b := pop[rand.IntN(len(pop))]
	if a.Fitness > b.Fitness {
		return a
	}
	return b
}

func crossover(p1, p2 Individual, k int) Individual {
	n := len(p1.Genome)
	child := make([]bool, n)

	for i := 0; i < n; i++ {
		if rand.Float64() < 0.5 {
			child[i] = p1.Genome[i]
		} else {
			child[i] = p2.Genome[i]
		}
	}

	// Repair to ensure exactly K bits set
	count := 0
	for _, b := range child {
		if b {
			count++
		}
	}

	for count > k {
		i := rand.IntN(n)
		if child[i] {
			child[i] = false
			count--
		}
	}

	for count < k {
		i := rand.IntN(n)
		if !child[i] {
			child[i] = true
			count++
		}
	}

	return Individual{Genome: child}
}

func mutate(ind Individual) Individual {

	var ones, zeros []int
	for i, b := range ind.Genome {
		if b {
			ones = append(ones, i)
		} else {
			zeros = append(zeros, i)
		}
	}

	if len(ones) > 0 && len(zeros) > 0 {
		ind.Genome[ones[rand.IntN(len(ones))]] = false
		ind.Genome[zeros[rand.IntN(len(zeros))]] = true
	}

	return ind
}

func copyIndividual(ind Individual) Individual {
	genomeCopy := make([]bool, len(ind.Genome))
	copy(genomeCopy, ind.Genome)

	return Individual{
		Genome:  genomeCopy,
		Fitness: ind.Fitness,
	}
}
