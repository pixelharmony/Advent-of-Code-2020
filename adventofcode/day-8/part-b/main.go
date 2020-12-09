package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	Op                  string
	Value               int
	Executed            bool
	Line                int
	PreviousInstruction *instruction
	NextInstruction     *instruction
	Accumulator         int
}

func (o *instruction) complete(prev *instruction, next *instruction, accumulator int) {
	o.Executed = true
	o.PreviousInstruction = prev
	o.NextInstruction = next
	o.Accumulator = accumulator
}

func (o *instruction) shift(line int) int {
	if o.Op == "jmp" {
		return line + o.Value
	}
	return line + 1
}

func main() {
	commands := load()
	instructions := build(commands)
	_, brokenInstruction := run(instructions, 0)
	finalInstruction := flip(instructions, brokenInstruction)

	fmt.Println(finalInstruction.Accumulator)
}

func flip(instructions map[int]*instruction, finalInstruction *instruction) *instruction {

	i := finalInstruction.PreviousInstruction

	for {
		if i.Op == "nop" {
			i.Op = "jmp"
		}

		if i.Op == "jmp" {
			i.Op = "nop"
		}

		success, instruction := run(instructions, i.Line)

		if success {
			return instruction
		}

		i = i.PreviousInstruction
	}
}

func run(instructions map[int]*instruction, startIndex int) (success bool, finalInstruction *instruction) {

	index := startIndex
	instruction := instructions[index]
	prevInstruction := instruction.PreviousInstruction

	for {

		index = instruction.shift(index)
		nextInstruction := instructions[index]
		accumulator := 0

		if prevInstruction != nil {
			accumulator = accumulator + prevInstruction.Accumulator
		}

		if instruction.Op == "acc" {
			accumulator = accumulator + instruction.Value
		}

		if nextInstruction != nil && nextInstruction.Executed {
			instruction.complete(prevInstruction, nil, accumulator)
			success = false
			break
		} else {
			instruction.complete(prevInstruction, nextInstruction, accumulator)
		}

		if nextInstruction == nil {
			success = true
			break
		}

		prevInstruction = instruction
		instruction = nextInstruction
	}

	return success, instruction
}

func load() (commands []string) {
	content, _ := ioutil.ReadFile("./input.txt")
	commands = strings.Split(string(content), "\n")

	return commands
}

func build(commands []string) (instructions map[int]*instruction) {

	instructions = make(map[int]*instruction)

	for line, command := range commands {
		instructions[line] = parse(command, line)
	}

	return instructions
}

func parse(command string, line int) *instruction {
	split := strings.Split(command, " ")
	op := split[0]
	value, _ := strconv.Atoi(split[1])

	return &instruction{op, value, false, line, nil, nil, 0}
}
