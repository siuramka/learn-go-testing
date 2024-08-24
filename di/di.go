package di

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Printf("Hello, %s", name)
}

type SumRepository interface {
	Add(n int)
	GetAll() []int
}

type sumRepository struct {
	numbers []int
}

func NewSumRepo() SumRepository {
	return &sumRepository{
		numbers: make([]int, 0),
	}
}

func (r *sumRepository) Add(n int) {
	r.numbers = append(r.numbers, n)
}

func (r *sumRepository) GetAll() []int {
	numbers := make([]int, len(r.numbers))
	copy(numbers, r.numbers)
	return numbers
}

func Calculate(r SumRepository) int {
	sum := 0
	for _, n := range r.GetAll() {
		sum += n
	}
	return sum
}
