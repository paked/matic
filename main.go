package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/paked/configeur"
)

var (
	conf = configeur.New()

	min = conf.Int("min", 0, "The minimum number you want to be tested on")
	max = conf.Int("max", 12, "The maximum number you want to be tested on")

	addition       = conf.Bool("addition", true, "Whether you want to be tested on addition or not")
	subtraction    = conf.Bool("subtraction", true, "Whether you want to be tested on subtraction or not")
	multiplication = conf.Bool("multiplication", false, "Whether you want to be tested on multiplication or not")
	division       = conf.Bool("division", false, "Whether you want to be tested on division or not")
)

func init() {
	conf.Use(configeur.NewFlag())
	conf.Use(configeur.NewEnvironment())

	rand.Seed(time.Now().UnixNano())
}

func main() {
	conf.Parse()

	m := NewMatic(*min, *max, *addition, *subtraction, *multiplication, *division)

	m.Run()
}

type Matic struct {
	generator  func() int
	operations []Operator
}

func (m *Matic) Run() {
	for {
		a := m.generator()
		b := m.generator()

		op := m.operations[rand.Intn(len(m.operations))]

		answer := op.process(a, b)

		fmt.Printf("%d %s %d: ", a, op.name, b)

		var value int
		fmt.Scanf("%d", &value)

		fmt.Print("\n")

		if value == answer {
			fmt.Println("Correct!")
			break
		} else {
			fmt.Printf("Incorrect! The answer was %d\n", answer)
		}
	}
}

func NewMatic(min, max int, add, sub, mult, div bool) *Matic {
	gen := func() int {
		return rand.Intn(max) + min
	}

	var opers []Operator
	if add {
		addOp := Operator{
			process: adder,
			name:    "+",
		}

		opers = append(opers, addOp)
	}

	if sub {
		subOp := Operator{
			process: subber,
			name:    "-",
		}

		opers = append(opers, subOp)
	}

	if mult {
		multOp := Operator{
			process: multer,
			name:    "*",
		}

		opers = append(opers, multOp)
	}

	if div {
		divOp := Operator{
			process: divver,
			name:    "/",
		}

		opers = append(opers, divOp)
	}

	return &Matic{
		generator:  gen,
		operations: opers,
	}
}

func adder(a, b int) int {
	return a + b
}

func subber(a, b int) int {
	return a - b
}

func multer(a, b int) int {
	return a * b
}

func divver(a, b int) int {
	return a / b
}

type Operator struct {
	process func(a, b int) int
	name    string
}
