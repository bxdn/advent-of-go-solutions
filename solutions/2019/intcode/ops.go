package intcode

func (prog *program) add() {
	op1 := prog.get(0)
	op2 := prog.get(1)
	prog.set(2, op1+op2)
}

func (prog *program) mult() {
	op1 := prog.get(0)
	op2 := prog.get(1)
	prog.set(2, op1*op2)
}

func (prog *program) in() {
	prog.set(0, prog.inputFunc())
}

func (prog *program) out() {
	prog.outputFunc(prog.get(0))
}

func (prog *program) jmpT() {
	val := prog.get(0)
	if val != 0 {
		prog.pos = prog.get(1)
	}
}

func (prog *program) jmpF() {
	val := prog.get(0)
	if val == 0 {
		prog.pos = prog.get(1)
	}
}

func (prog *program) lt() {
	if prog.get(0) < prog.get(1) {
		prog.set(2, 1)
	} else {
		prog.set(2, 0)
	}
}

func (prog *program) eq() {
	if prog.get(0) == prog.get(1) {
		prog.set(2, 1)
	} else {
		prog.set(2, 0)
	}
}

func (prog *program) crb() {
	prog.base += prog.get(0)
}