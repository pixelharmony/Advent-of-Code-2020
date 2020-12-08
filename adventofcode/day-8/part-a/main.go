package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type operation struct {
	Operation          string
	NextOperationIndex int
	Completed          bool
}

func (o *operation) complete() {
	o.Completed = true
}

func main() {
	commands := load()
	operations := parse(commands)

	index := 0
	nextOperation := &operations[index]
	accumulator := 0

	for !nextOperation.Completed {

		op := nextOperation
		shift := advance(op, &accumulator)
		index = index + shift
		nextOperation = &operations[index]

		op.complete()
	}

	fmt.Printf("Acc: %v", accumulator)
}

func advance(op *operation, acc *int) (shift int) {
	shift = 1

	switch op.Operation {
	case "acc":
		*acc = *acc + op.NextOperationIndex
	case "jmp":
		shift = op.NextOperationIndex
	}

	return shift
}

func load() (commands []string) {
	content, _ := ioutil.ReadFile("./input.txt")
	commands = strings.Split(string(content), "\n")

	return commands
}

func parse(commands []string) (operations []operation) {
	for _, command := range commands {
		split := strings.Split(command, " ")
		op := split[0]
		nextOpIndex, _ := strconv.Atoi(split[1])

		operations = append(operations, operation{op, nextOpIndex, false})
	}

	return operations
}
