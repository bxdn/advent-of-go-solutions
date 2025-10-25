package utils

import "sync"

type TaskGroup struct {
	funcs []func()
	wg *sync.WaitGroup
}

func (tg *TaskGroup) Add(fn func()) {
	tg.funcs = append(tg.funcs, fn)
}

func (tg *TaskGroup) Start() {
	if tg.wg == nil {
		tg.wg = new(sync.WaitGroup)
	}
	tg.wg.Add(len(tg.funcs))
	for _, fn := range tg.funcs {
		fnWithCleanup := func() {
			fn()
			tg.wg.Done()
		}
		go fnWithCleanup()
	}
}

func (tg *TaskGroup) Wait() {
	if tg.wg == nil {
		return
	}
	tg.wg.Wait()
}