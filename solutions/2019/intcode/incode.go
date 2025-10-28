package intcode

import (
	"advent-of-go/utils"
	"fmt"
	"math"
	"slices"
	"strings"
)

// types

const HALT_CODE = math.MinInt

type opfunc func()
type outFunc func(int)
type inFunc func() int

type instruction struct {
	op     int
	params []param
}

type Op struct {
	nArgs int
	run   opfunc
}

type param struct {
	val, mode int
}

type program struct {
	mem []int
	inst instruction
	pos int
	base int
	inputFunc inFunc
	outputFunc outFunc
	halt bool
}

var N_ARGS = [10]int{-1, 3, 3, 1, 1, 2, 2, 3, 3, 1}

func RunString(instructionString string, input inFunc, output outFunc) error {
	ops, e := utils.StringsToInts(strings.Split(instructionString, ","))
	if e != nil {
		return fmt.Errorf("Error parsing intcode input: %w", e)
	}
	return Run(ops, input, output)
}

func RunBasicString(instructionString string) error {
	ops, e := utils.StringsToInts(strings.Split(instructionString, ","))
	if e != nil {
		return fmt.Errorf("Error parsing incode input: %w", e)
	}
	return RunBasic(ops)
}

func Run(rawInstructions []int, input inFunc, output outFunc) error {
	program := program{mem: rawInstructions, inputFunc: input, outputFunc: output}
	for {
		if halt := program.parseInstruction(); halt {
			return nil
		}
		if e := program.runInst(); e != nil {
			return fmt.Errorf("Error running intocode program: %w", e)
		}
		if program.halt {
			return nil
		}
	}
}

func RunAt(rawInstructions []int, input inFunc, output outFunc, pos int) error {
	program := program{mem: rawInstructions, inputFunc: input, outputFunc: output, pos: pos}
	for {
		if halt := program.parseInstruction(); halt {
			return nil
		}
		if e := program.runInst(); e != nil {
			return fmt.Errorf("Error running intocode program: %w", e)
		}
		if program.halt {
			return nil
		}
	}
}

func RunBasic(ops []int) error {
	noopInput := func() int { return -1 }
	noopOutput := func(int) {}
	Run(ops, noopInput, noopOutput)
	return nil
}

func (prog *program) runInst() error {
	switch prog.inst.op {
		case 1: prog.add()
		case 2: prog.mult()
		case 3: prog.in()
		case 4: prog.out()
		case 5: prog.jmpT()
		case 6: prog.jmpF()
		case 7: prog.lt()
		case 8: prog.eq()
		case 9: prog.crb()
		default: return fmt.Errorf("Unknown operation %d detected!", prog.inst.op)
	}
	return nil
}

func (p *program) at(i int) int {
	if i >= len(p.mem) {
		return 0
	}
	return p.mem[i]
}


func (prog *program) get(i int) int {
	p := prog.inst.params[i]
	switch p.mode {
		case 0: return prog.at(p.val)
		case 1: return p.val
		case 2: return prog.at(p.val + prog.base)
		default: return 0
	}
}

func (p *program) set(paramIdx, v int) {
	param := p.inst.params[paramIdx]
	addr := param.val
	if param.mode == 2 {
		addr += p.base
	}
	if addr >= len(p.mem) {
		p.mem = slices.Concat(p.mem, make([]int, (addr + 1) - len(p.mem)))
	}
	p.mem[addr] = v
}

func (p *program) parseInstruction() bool {
	opCode := p.at(p.pos) % 100

	if opCode == 99 {
		return true
	}

	nArgs := N_ARGS[opCode]
	modes := p.at(p.pos) / 100
	params := make([]param, nArgs)

	for i := 0; i < nArgs; i++ {
		params[i] = param{p.at(p.pos + 1 + i), modes % 10}
		modes /= 10
	}

	p.inst = instruction{opCode, params}
	p.pos += len(params) + 1
	return false
}